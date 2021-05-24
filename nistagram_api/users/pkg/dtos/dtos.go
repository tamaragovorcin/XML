package dtos

type UserRequest struct {
	Name  string
	Surname string
	Email string
	Password string
	RepeatedPassword string
}

type LoginRequest struct {
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
type UserTokenState struct {
	AccessToken string
	ExpiresIn int64
	Roles []string
}



