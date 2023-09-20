package photo

import "time"

type Photo struct {
	ID             int
	Title          string
	Caption  	   string
	PhotoURL       string
	PasswordHash   string
	AvatarFileName string
	UserID		   int
	CreatedAt      time.Time
	UpdatedAt      time.Time  
}