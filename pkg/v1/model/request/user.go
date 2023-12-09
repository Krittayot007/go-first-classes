package request

type UserRequestPost struct {
	Usernames []string `json:"usernames"`
	FirstName []string `json:"firs_name"`
}
