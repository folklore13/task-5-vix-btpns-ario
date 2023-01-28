package database

type PhotoTable struct {
	PhotoID string `json:"photo_id"`
	PhotoFile string `json:"photo_file"`
	PhotoURL string `json:"photo_url"`
	UserID string `json:"user_id"`
}