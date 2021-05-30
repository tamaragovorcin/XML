package dtos

type Gender int
const (
	Male Gender = iota
	Female
	Other
)
type UserRequest struct {
	Name  string
	LastName string
	Email string
	Username string
	Password string
	PhoneNumber string
	Gender string
	DateOfBirth string
	Private bool
	Biography string
}

type LoginRequest struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
type UserTokenState struct {
	AccessToken string
	ExpiresIn int64
	Roles string
}



