package dto

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Admin    bool `form:"admin"`
}
