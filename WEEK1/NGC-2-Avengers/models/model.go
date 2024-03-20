package models

type Heroes struct {
	ID       int64  `json:"heroid"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	Imageurl string `json:"imageurl"`
}

type Villain struct {
	ID       int64  `json:"villainid"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Imageurl string `json:"imageurl"`
}
