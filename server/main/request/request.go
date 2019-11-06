package request

type GoogleFormParam struct {
	Name string `json:"name"`
	WorkName string `json:"work"`
	Method string `json:"method"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}