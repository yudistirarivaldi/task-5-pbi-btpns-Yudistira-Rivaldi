package photo

type Service interface {
	CreatePhoto(input CreatePhotoInput, fileLocation string) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input CreatePhotoInput, fileLocation string) (Photo, error) {

	photo := Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = fileLocation
	photo.UserID = input.User.ID

	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}
