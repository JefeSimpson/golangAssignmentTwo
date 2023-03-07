package service

import "secondAssignment/model"

type Collection struct {
	userIterator, itemIterator int
	Users                      []model.User
	Items                      []model.Item
}

func NewCollection(userIterator int, itemIterator int, users []model.User, items []model.Item) *Collection {
	return &Collection{userIterator: userIterator, itemIterator: itemIterator, Users: users, Items: items}
}

var isAuthorized bool = false

func GetIsAuthorized() bool {
	return isAuthorized
}

func SetIsAuthorized(enter bool) {
	isAuthorized = enter
}

type UserManager interface {
	SignUp(username, password string) bool
	SignIn(username, password string) bool
	GetUser(username string) []model.User
	UserTakeData()
	UserSaveData()
}

type ItemManager interface {
	ItemPush(name string, price, rating float64)
	SearchItemsByName(name string) []model.Item
	FilterItemsByPrice(price float64) []model.Item
	FilterItemsByRating(rating float64) []model.Item
	SetRating(rating float64, id int)
	ItemTakeData()
	ItemSaveData()
}
