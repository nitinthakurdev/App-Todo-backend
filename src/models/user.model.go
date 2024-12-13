package models

type UserModel struct {
	Username string `bson:"username,omitempty" validate:"required,min=2"`
	Email    string `bson:"email,omitempty" validate:"required,email"`
	Password string `bson:"password,omitempty" validate:"required,min=4,max=16"`
}
