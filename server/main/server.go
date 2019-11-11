package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/inc-brick/Medici/server/main/entity"
	"github.com/inc-brick/Medici/server/main/request"
	"github.com/inc-brick/Medici/server/main/response"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	db *sqlx.DB
	err error
)


// TODO(takimura): 削除
type SignupParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// TODO(takimura): 削除
type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// TODO(takimura): 削除
type EventsParam struct {
	Language string `json:"language"`
	Area string `json:"area"`
}

// TODO(takimura): 削除
type EventApiResponse struct {
	ResultsReturned int `json:"results_returned"`
	ResultsAvailable int `json:"results_available"`
	ResultsStart int `json:"results_start"`
	Events []Event `json:"events"`
}

// TODO(takimura): 削除
type Event struct {
	EventId int `json:"event_id"`
	Title string `json:"title"`
	Catch string `json:"catch"`
	Description string `json:"description"`
	EventUrl string `json:"event_url"`
	HashTag string `json:"hash_tag"`
	StartedAt string `json:"started_at"`
	EndedAt string `json:"ended_at"`
	Limit int `json:"limit"`
	EventType string `json:"event_type"`
	Series Series `json:"series"`
	Address string `json:"address"`
	Place string `json:"place"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	OwnerId int `json:"owner_id"`
	OwnerNickname string `json:"owner_nickname"`
	OwnerDisplayName string `json:"owner_display_name"`
	Accepted int `json:"accepted"`
	Waiting int `json:"waiting"`
	UpdatedAt string `json:"updated_at"`
}

// TODO(takimura): 削除
type Series struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
}

type User struct {
	Id int64 `json:"user_id" db:"id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"-" db:"digested_password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func init() {
}

func signup(c echo.Context) error {
	log.Info("singup is called")
	req := new(SignupParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	// validation
	if req.Password != req.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, "password is not matched")
	}
	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, "password length less than 8")
	}
	log.Info("password validation: OK")
	// store
	digestedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Failed to encrypt password. %s", err.Error())
	}
	log.Info("password encryption: OK")
	tx, err := db.Begin()
	if err != nil {
		log.Errorf("Failed to begin transaction. %s", err.Error())
	}
	log.Info("transaction begin: OK")
	res, err := tx.Exec("INSERT INTO `users` (`id`, `username`, `digested_password`, `created_at`, `updated_at`) VALUES (null, ?,?,?,?)",
		req.UserName,
		digestedPassword,
		time.Now(),
		time.Now())
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	err = tx.Commit()
	if err != nil {
		log.Errorf("Failed to commit. %s", err.Error())
	}
	log.Info("db execution: OK %s", res)

	// session
	lastInsertedId, _ := res.LastInsertId()
	saveSession(lastInsertedId, req.UserName, true, c)
	// select user
	var user User
	if err := db.Get(&user,"SELECT * FROM `users` WHERE `id` = ?", lastInsertedId); err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func login(c echo.Context) error {
	log.Info("login is called")
	req := new(LoginParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	var user User
	if err := db.Get(&user, "SELECT * FROM `users` WHERE `username` = ?", req.UserName); err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
		return c.JSON(http.StatusUnauthorized, "Can not find user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil{
		log.Info(err)
		return c.JSON(http.StatusUnauthorized, "password doesn't match")
	}
	// session
	saveSession(user.Id, user.UserName, true, c)
	return c.JSON(http.StatusOK, user)
}

func logout(c echo.Context) error {
	authenticated, err := authenticated(c)
	if authenticated == 0 {
		return err
	}
	sess, _ := session.Get("EVENTWATCH_SESSION", c)
	saveSession(0, "", false, c)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Errorf("Can not save the session. %s", err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func test(c echo.Context) error {
	userId, err := authenticated(c)
	if userId == 0 {
		return err
	}
	req := new(EventsParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	url := "https://connpass.com/api/v1/event/"
	client := &http.Client{}
	apiReq, err := http.NewRequest("GET", url+buildQueryParam(req.Language, req.Area), nil)
	if err != nil {
		log.Error(err)
		return err
	}
	apiRes, err := client.Do(apiReq)
	if err != nil {
		log.Info(err)
		return err
	}
	defer apiRes.Body.Close()

	if apiRes.StatusCode != http.StatusOK {
		log.Info(apiRes)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	body, err := ioutil.ReadAll(apiRes.Body)
	if err != nil {
		log.Errorf("Can not read response body. %s", err.Error())
	}

	var eventApiResponse EventApiResponse
	err = json.Unmarshal(body, &eventApiResponse)
	if err != nil {
		log.Errorf("Could not unmarshal response body. %s", err.Error())
	}
	var events []Event
	for _, v := range eventApiResponse.Events {
		if strings.Contains(v.Address, req.Area) {
			events = append(events, v)
		}
	}
	return c.JSON(http.StatusOK, events)
}

func fetchArtistInfo(c echo.Context) error {
	log.Info("fetchArtistInfo is called")
	artistId, err := strconv.Atoi(c.Param("artistId"))
	if err != nil {
		log.Fatal("Could not parse artistId. artistId: %s ¥nCause is... %s", artistId, err.Error())
		return c.JSON(http.StatusBadRequest, "Can not find artist")
	}
	var res response.FetchArtistInfo
	res.ArtistId = artistId
	

	// ======== 処理内容 ========
	// ComArtistMstから情報を取得
	// ArtWorkMstから情報を取得
	// ArtArtistEventTrnから情報を取得
	// MediaTrnから情報を取得
	// response作って返す

	//if err := db.Get(&user, "SELECT * FROM `users` WHERE `username` = ?", req.UserName); err != nil {
	//	log.Errorf("Failed to execute query. %s", err.Error())
	//	return c.JSON(http.StatusUnauthorized, "Can not find user")
	//}
	return c.JSON(http.StatusOK, res)
}

func dbHealthCheck(c echo.Context) error {
	log.Info("login is called")
	var comArtistPrivateInfo entity.ComArtistPrivateInfo
	if err := db.Get(&comArtistPrivateInfo, "SELECT * FROM `COM_ARTIST_PIVATE_INFO_MST`"); err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
		return c.JSON(http.StatusUnauthorized, "Can not find user")
	}
	return c.JSON(http.StatusOK, comArtistPrivateInfo)
}

func postGoogleForm(c echo.Context) error {
	c.Echo().Logger.Info("post google form is called")
	req := new(request.GoogleFormParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	// request bodyの作成
	reqPram := url2.Values{}
	reqPram.Add("entry.1239827993", req.Name) // 名前
	reqPram.Add("entry.1372048078", req.WorkName) // 作品名称
	reqPram.Add("entry.1700798423", req.Method) // 連絡方法
	reqPram.Add("entry.1682303839", req.Email) // メールアドレス
	reqPram.Add("entry.1390168152", req.Phone) // 電話番号
	c.Echo().Logger.Fatal("entry.1239827993(name): %s", req.Name)
	c.Echo().Logger.Fatal("entry.1372048078(work): %s", req.WorkName)
	c.Echo().Logger.Fatal("entry.1700798423(method): %s", req.Method)
	c.Echo().Logger.Fatal("entry.1682303839(email): %s", req.Email)
	c.Echo().Logger.Fatal("entry.1390168152(phone): %s", req.Phone)
	// google form のURL
	url := "https://docs.google.com/forms/u/1/d/e/1FAIpQLSc10-M1uZi5jD2jmyK_ICom4KipjEWXv6O6xHqTQq6vyvO_hg/formResponse"
	// requestの生成
	apiReq, err := http.NewRequest("POST", url, strings.NewReader(reqPram.Encode()))
	if err != nil {
		log.Error(err)
		return err
	}
	// headerの追加
	apiReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	apiRes, err := client.Do(apiReq) // API実行
	if err != nil {
		log.Info(err)
		return err
	}
	defer apiRes.Body.Close()

	if apiRes.StatusCode != http.StatusOK {
		log.Info(apiRes)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	body, err := ioutil.ReadAll(apiRes.Body)
	if err != nil {
		log.Errorf("Can not read response body. %s", err.Error())
	}
	return c.JSON(http.StatusOK, body)
}

func main() {
	dbHost := os.Getenv("MEDICI_MYSQL_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPort := os.Getenv("MEDICI_MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	_, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("failed to read DB port number from an environment variable MYSQL_PORT.\nError: %s", err.Error())
	}
	dbUserName := os.Getenv("MEDICI_MYSQL_USERNAME")
	if dbUserName == "" {
		dbUserName = "brickdev01"
	}
	dbPassword := os.Getenv("MEDICI_MYSQL_PASSWORD")
	if dbPassword == "" {
		dbPassword = "brickdev012019"
	}
	dbName := os.Getenv("MEDICI_MYSQL_DBNAME")
	if dbName == "" {
		dbName = "medici"
	}
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbUserName, dbPassword, dbHost, dbPort, dbName)
	db, err = sqlx.Open("mysql", datasource)
	if err != nil {
		log.Fatalf("failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("MEDICI_COOKIE_STORE_KEY"))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://app-brick.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
	}))
	e.POST("/signup", signup)
	e.POST("/login", login)
	e.POST("/logout", logout)
	e.POST("/events", test)
	e.POST("/db/healthCheck", dbHealthCheck)
	e.POST("/artist/:artistId", fetchArtistInfo)
	e.POST("/post/contact", postGoogleForm)
	e.Logger.Fatal(e.Start(":8000"))
}

func authenticated(c echo.Context) (interface{}, error) {
	sess, _ := session.Get("MEDICI_SESSION", c)
	b, _ := sess.Values["authenticated"]
	if b != true {
		log.Info("User not logged in")
		return 0, c.JSON(http.StatusUnauthorized, "please sign up or login at first")
	}
	userId, ok := sess.Values["userId"]
	if !ok {
		log.Error("Can not get userId from session store.")
		return 0, c.JSON(http.StatusUnauthorized, "please sign up or login at first")
	}
	return userId, nil
}

func saveSession(userId int64, username string, authenticated bool, c echo.Context) {
	// session
	sess, _ := session.Get("MEDICI_SESSION", c)
	sess.Options = &sessions.Options{
		Path:	"/",
		MaxAge:	3600,
		HttpOnly: true,
	}
	sess.Values["userId"] = userId
	sess.Values["username"] = username
	sess.Values["authenticated"] = authenticated
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Errorf("Can not save the session. %s", err.Error())
	}
}

func buildQueryParam(language, area string) string {
	s := ""
	if len(language) != 0 {
		s = s + "?keyword=" + language
	}
	if len(area) != 0 {
		s = s + "&keyword=" + area
	}
	return s
}