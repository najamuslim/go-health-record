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

type RequestCreateNurse struct {
	Nip                string  `json:"nip"`
	Name               string `json:"name"`
	IdentityCardScanImg string `json:"identityCardScanImg"` // URL to the identity card scan image
}
