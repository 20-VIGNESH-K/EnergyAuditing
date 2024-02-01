package models

type CreateUser struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Address  string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	Country  string `json:"country" bson:"country"`
	Zipcode  string `json:"zipcode" bson:"zipcode"`
}

type LogIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
