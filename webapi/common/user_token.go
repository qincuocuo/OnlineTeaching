package common

type UserToken struct {
	UserId int `json:"userId" bson:"user_id"`
	Role   int `json:"role" bson:"role"`
}
