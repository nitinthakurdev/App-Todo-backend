package models

type Image struct {
	ImageURL string `bson:"image_url,omitempty"`
	ImageId  string `bson:"image_Id,omitempty"`
}

type UserModel struct {
	Username string `bson:"username,omitempty" validate:"required,min=2"`
	Email    string `bson:"email,omitempty" validate:"required,email"`
	Password string `bson:"password,omitempty" validate:"required,min=4,max=16"`
	Image    Image  `bson:"image,omitempty"`
}
