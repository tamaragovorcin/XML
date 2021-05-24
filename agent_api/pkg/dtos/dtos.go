package dtos

type UserRequest struct {
	Name  string `bson:"name,omitempty"`
	Surname string `bson:"surname,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	RepeatedPassword string `bson:"repeatedPassword,omitempty"`
}
type LoginRequest struct {
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}