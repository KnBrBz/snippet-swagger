package auth

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthSuccess struct {
	Token string `json:"token"`
}
