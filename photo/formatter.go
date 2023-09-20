package photo

// fungsi struct ini agar mengubah menjadi format json

type PhotoFormatter struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PhotoURL   string `json:"photo_url"`
	UserID     string `json:"user_id"`
}

func FormatPhoto(photo Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		ID:         photo.ID,
		Title:      photo.Title,
		Caption: 	photo.Caption,
		PhotoURL:   photo.PhotoURL,
		UserID:   	photo.AvatarFileName,
	}

	return formatter

}