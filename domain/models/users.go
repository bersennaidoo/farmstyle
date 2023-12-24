package models

type User struct {
	Uuid     string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	FullName string `json:"fullname" bson:"fullname"`
}

type NewUser struct {
	Uuid     string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	FullName string `json:"fullname" bson:"fullname"`
	Password string `json:"password" bson:"password"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
