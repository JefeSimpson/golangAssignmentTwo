package model

type User struct {
	Id       int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//func (u *User) Validator() error {
//	if u.Id == 0 {
//		return errors.New("not [pr")
//	}
//}

type Item struct {
	Id       int     `json:"itemId"`
	ItemName string  `json:"itemName"`
	Price    float64 `json:"price"`
	Rating   float64 `json:"rating"`
}
