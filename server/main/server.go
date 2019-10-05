package main

import (
"encoding/json"
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/gorilla/sessions"
"github.com/jmoiron/sqlx"
"github.com/labstack/echo"
"github.com/labstack/echo-contrib/session"
"github.com/labstack/gommon/log"
"golang.org/x/crypto/bcrypt"
"io/ioutil"
"net/http"
"os"
"strconv"
"strings"
"time"
)

var (
	db *sqlx.DB
	err error
)


// TODO(takimura): typeは将来的には別で管理する
type SignupParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type EventsParam struct {
	Language string `json:"language"`
	Area string `json:"area"`
}

type EventApiResponse struct {
	ResultsReturned int `json:"results_returned"`
	ResultsAvailable int `json:"results_available"`
	ResultsStart int `json:"results_start"`
	Events []Event `json:"events"`
}

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

type ComArtistPrivateInfo struct {
	ArtistPrivateId int `json:"artist_private_id" db:"ARTIST_PRIVATE_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	PhoneNumber int `json:"phone_number" db:"PHONE_NUMBER"`
	MailAddress string `json:"mail_address" db:"MAIL_ADDRESS"`
	PostalCode string `json:"postal_code" db:"POSTAL_CODE"`
	Address string `json:"address" db:"ADDRESS"`
	WorkPlacePostalCode string `json:"work_place_postal_code" db:"WORK_PLACE_POSTAL_CODE"`
	WorkPlaceAddress int `json:"work_place_address" db:"WORK_PLACE_ADDRESS"`
	FormalName int `json:"formal_name" db:"FORMAL_NAME"`
}

type ComArtistPrivateInfo struct {
	ArtistPrivateId int `json:"artist_private_id" db:"ARTIST_PRIVATE_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	PhoneNumber int `json:"phone_number" db:"PHONE_NUMBER"`
	MailAddress string `json:"mail_address" db:"MAIL_ADDRESS"`
	PostalCode string `json:"postal_code" db:"POSTAL_CODE"`
	Address string `json:"address" db:"ADDRESS"`
	WorkPlacePostalCode string `json:"work_place_postal_code" db:"WORK_PLACE_POSTAL_CODE"`
	WorkPlaceAddress string `json:"work_place_address" db:"WORK_PLACE_ADDRESS"`
	FormalName string `json:"formal_name" db:"FORMAL_NAME"`
}

type ComArtist struct {
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	Sequence int `json:"sequence" db:"SEQUENCE"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	ArtistNameJa string `json:"artist_name_ja" db:"ARTIST_NAME_JA"`
	ArtistNameEn string `json:"artist_name_en" db:"ARTIST_NAME_EN"`
	DescriptionJa string `json:"description_ja" db:"DESCRIPTION_JA"`
	DescriptionEn string `json:"description_en" db:"DESCRIPTION_EN"`
}

type ArtWork struct {
	WorkId int `json:"work_id" db:"WORK_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	WorkTitleJa string `json:"work_title_ja" db:"WORK_TITLE_JA"`
	WorkTitleEn string `json:"work_title_en" db:"WORK_TITLE_EN"`
	WorkDescriptionJa string `json:"work_description_ja" db:"WORK_DESCRIPTION_JA"`
	WorkDescriptionEn string `json:"work_description_en" db:"WORK_DESCRIPTION_EN"`
	WorkImagePath string `json:"work_image_path" db:"WORK_IMAGE_PATH"`
	CreateYear int `json:"create_year" db:"CREATE_YEAR"`
	Size string `json:"size" db:"SIZE"`
	SizeXPosition int `json:"size_x_position" db:"SIZE_X_POSITION"`
	SizeYPosition int `json:"size_y_position" db:"SIZE_Y_POSITION"`
	SizeZPosition int `json:"size_z_position" db:"SIZE_Z_POSITION"`
	Material string `json:"material" db:"MATERIAL"`
}

type ArtRelatedItem struct {
	ItemId int `json:"item_id" db:"ITEM_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	ItemTitleJa string `json:"item_title_ja" db:"ITEM_TITLE_JA"`
	ItemTitleEn string `json:"item_title_en" db:"ITEM_TITLE_EN"`
	ItemDescriptionJa string `json:"item_description_ja" db:"ITEM_DESCRIPTION_JA"`
	ItemDescriptionEn string `json:"item_description_en" db:"ITEM_DESCRIPTION_EN"`
	ItemImagePath string `json:"item_image_path" db:"ITEM_IMAGE_PATH"`
	CreateYear int `json:"create_year" db:"CREATE_YEAR"`
	Size string `json:"size" db:"SIZE"`
	SizeXPosition int `json:"size_x_position" db:"SIZE_X_POSITION"`
	SizeYPosition int `json:"size_y_position" db:"SIZE_Y_POSITION"`
	SizeZPosition int `json:"size_z_position" db:"SIZE_Z_POSITION"`
	Material string `json:"material" db:"MATERIAL"`
}

type ArtArtistEvent struct {
	EventId int `json:"event_id" db:"EVENT_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	EventNameJA string `json:"event_name_ja" db:"EVENT_NAME_JA"`
	EventNameEN string `json:"event_name_en" db:"EVENT_NAME_EN"`
	EventDescriptionJA string `json:"event_description_ja" db:"EVENT_DESCRIPTION_JA"`
	EventDescriptionEN string `json:"event_description_en" db:"EVENT_DESCRIPTION_EN"`
	EventImagePath string `json:"event_image_path" db:"EVENT_IMAGE_PATH"`
	EventStartAt time.Time `json:"event_start_at" db:"EVENT_START_AT"`
	EventEndAt time.Time `json:"event_end_at" db:"EVENT_END_AT"`
	EventPhone int `json:"event_phone" db:"EVENT_PHONE"`
	EventMailAddress string `json:"event_mail_address" db:"EVENT_MAIL_ADDRESS"`
	EventUrl string `json:"event_url" db:"EVENT_URL"`
}

type ArtMedia struct {
	MediaId int `json:"media_id" db:"MEDIA_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	MediaType int `json:"media_type" db:"MEDIA_TYPE"`
	MediaNameJa string `json:"media_name_ja" db:"MEDIA_NAME_JA"`
	MediaNameEn string `json:"media_name_en" db:"MEDIA_NAME_EN"`
	MediaDescriptionJa string `json:"media_description_ja" db:"MEDIA_DESCRIPTION_JA"`
	MediaDescriptionEn string `json:"media_description_en" db:"MEDIA_DESCRIPTION_EN"`
	MediaUrl int `json:"media_url" db:"MEDIA_URL"`
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
	log.Info("login is called")
	artistId := c.Param("artistId")
	var user User
	if err := db.Get(&user, "SELECT * FROM `users` WHERE `username` = ?", req.UserName); err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
		return c.JSON(http.StatusUnauthorized, "Can not find user")
	}
	return c.JSON(http.StatusOK, artistId)
}

func main() {
	dbHost := os.Getenv("MEDICI _MYSQL_HOST")
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
	e.POST("/signup", signup)
	e.POST("/login", login)
	e.POST("/logout", logout)
	e.POST("/events", test)
	e.POST("/artist/:artistId", fetchArtistInfo)
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