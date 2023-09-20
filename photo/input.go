package photo

import "crowdfunding/user"

type CreatePhotoInput struct {
	Title    string `form:"title" binding:"required"`
	Caption  string `form:"caption" binding:"required"`
	User     user.User
}
