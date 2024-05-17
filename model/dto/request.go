package dto

type RequestAuth struct {
	Nip    int64 `json:"nip"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Nip    int64 `json:"nip"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestCreateNurse struct {
	Nip                int64  `json:"nip"`
	Name               string `json:"name"`
	IdentityCardScanImg string `json:"identityCardScanImg"` // URL to the identity card scan image
}
