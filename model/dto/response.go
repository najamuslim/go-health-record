package dto

type ResponseStatusAndMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserDTO struct {
	UserId      string `json:"userId"`
	NIP int64 `json:"nip"`
	Name        string `json:"name"`
	CreatedAt string `json:"createdAt"`
}