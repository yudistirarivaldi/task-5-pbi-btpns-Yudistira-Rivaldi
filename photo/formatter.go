package photo



type PhotoFormatter struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PhotoURL   string `json:"photo_url"`
	UserID     int    `json:"user_id"`
}

func FormatPhoto(photo Photo) PhotoFormatter {
	formatter := PhotoFormatter{
		ID:         photo.ID,
		Title:      photo.Title,
		Caption: 	photo.Caption,
		PhotoURL:   photo.PhotoURL,
		UserID:   	photo.UserID,
	}

	return formatter

}