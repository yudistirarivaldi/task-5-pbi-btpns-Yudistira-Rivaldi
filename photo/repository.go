package photo

import "gorm.io/gorm"

type Repository interface {
	Save(photo Photo) (Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(photo Photo) (Photo, error) {

	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}
