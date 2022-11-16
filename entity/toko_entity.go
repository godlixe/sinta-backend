package entity

type Toko struct {
	ID       uint64 `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"-"`
	BaseModel
}
