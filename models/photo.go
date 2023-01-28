package models

import "gorm.io/gorm"

type Photo struct {
	PhotoID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	PhotoTitle string `gorm:"type:varchar(255)" json:"photo_title"`
	PhotoCaption string `gorm:"type:text" json:"photo_caption"`
	PhotoURL string `gorm:"type:text"`
	UserID uint64 `gorm:"not null" json:"-"`
	User User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE"`
}

type PhotoRepo interface {
	InsertPhoto(photo Photo) Photo
	UpdatePhoto(photo Photo) Photo
	DeletePhoto(photo Photo)
	GetPhotoByID(photo Photo) Photo
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoConnection{
		connection: db,
	}
}

func (db *photoConnection) InsertPhoto(photo Photo) Photo {
	db.connection.Save(&photo)
	db.connection.Preload("User").Find(&photo)
	return photo
}

func (db *photoConnection) UpdatePhoto(photo Photo) Photo {
	db.connection.Save(&photo)
	db.connection.Preload("User").Find(&photo)
	return photo
}

func (db *photoConnection) DeletePhoto(photo Photo){
	db.connection.Delete(&photo)
}

func (db *photoConnection) GetPhotoByID(photoID uint64) Photo {
	var photos Photo
	db.connection.Preload("User").Find(&photos, photoID)
	return photos
}