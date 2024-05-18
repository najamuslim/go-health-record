package dto

type ResponseStatusAndMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserDTO struct {
	UserId      string 	`json:"userId"`
	NIP			int64 	`json:"nip"`
	Name        string 	`json:"name"`
	CreatedAt 	string 	`json:"createdAt"`
}

type PatientDTO struct {
	IdentityNumber		int			`json:"identityNumber"`
	PhoneNumber			string		`json:"phoneNumber"`
	Name				string		`json:"name"`
	BirthDate			string		`json:"birthDate"`
	Gender				string		`json:"gender"`
	CreatedAt			string		`json:"createdAt"`
}