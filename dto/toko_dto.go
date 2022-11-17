package dto

type TokoCreateDTO struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokoUpdateDTO struct {
	ID       uint64 `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokoLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
