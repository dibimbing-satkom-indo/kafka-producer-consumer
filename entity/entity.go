package entity

type User struct {
	ID   uint `json:"id"`
	Name string
}

type Event struct {
	Name string `json:"name"`
	Data User   `json:"data"`
}
