package controller

type User struct {
	Id       int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var IsAuthorized bool = false

type Item struct {
	Id       int     `json:"itemId"`
	ItemName string  `json:"itemName"`
	Price    float64 `json:"price"`
	Rating   float64 `json:"rating"`
}

type Collection struct {
	userIterator, itemIterator int
	Users                      []User
	Items                      []Item
}

func NewCollection(userIterator int, itemIterator int, users []User, items []Item) *Collection {
	return &Collection{userIterator: userIterator, itemIterator: itemIterator, Users: users, Items: items}
}
