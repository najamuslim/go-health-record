package dto

type RequestAuth struct {
	Nip    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Nip    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
