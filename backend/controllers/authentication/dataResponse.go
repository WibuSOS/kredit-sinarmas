package authentication

type DataResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
