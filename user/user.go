package user

//easyjson:json
type User struct {
	Browsers []string `json:"browsers,intern"`
	Company  string
	Country  string
	Email    string `json:"email,intern"`
	Job      string
	Name     string `json:"name,intern"`
	Phone    string
}
