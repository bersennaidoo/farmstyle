package models

type Review struct {
	Message string `json:"message" bson:"message"`
	Rating  int    `json:"rating" bson:"rating"`
	Uuid    string `json:"uuid" bson:"uuid"`
	UserID  string `json:"userId,omitempty" bson:"userId"`
}
type ReviewFilters struct {
	MaxRating int `json:"maxRating"`
}

type DeletedReview struct {
	Uuid string `json:"uuid" bson:"uuid"`
}
