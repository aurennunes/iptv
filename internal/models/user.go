package models

type UserInfo struct {
	Auth                 int      `json:"auth"`
	Username             string   `json:"username"`
	Password             string   `json:"password"`
	Message              string   `json:"message"`
	Status               string   `json:"status"`
	MaxConnections       string   `json:"max_connections"`
	IsTrial              string   `json:"is_trial"`
	ActiveCons           string   `json:"active_cons"`
	ExpDate              string   `json:"exp_date"`
	CreatedAt            string   `json:"created_at"`
	AllowedOutputFormats []string `json:"allowed_output_formats"`
}

type Auth struct {
	UserInfo UserInfo `json:"user_info"`
}
