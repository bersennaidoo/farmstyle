package models

type Review struct {
	Message string `json:"message" bson:"message"`
	Rating  int    `json:"rating" bson:"rating"`
	Uuid    string `json:"uuid" bson:"uuid"`
	UserID  string `json:"-" bson:"-"`
}
type ReviewFilters struct {
	MaxRating int `json:"maxRating" bson:"maxRating"`
}

type DeletedReview struct {
	Uuid string `json:"uuid" bson:"uuid"`
}
