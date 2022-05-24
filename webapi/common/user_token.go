package common

type UserToken struct {
	UserId string `json:"userId" bson:"user_id"`
	Role   int `json:"role" bson:"role"`
}
