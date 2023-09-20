package photo

type Service interface {
	CreatePhoto(input CreatePhotoInput) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input CreatePhotoInput) (Photo, error) {

	photo := Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL

	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}
