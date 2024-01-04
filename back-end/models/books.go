package models

import (
	"time"
)

type Books struct {
	ID        string `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Quantity  int    `json:"quantity"`
	Image     []byte `json:"image"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Borrowing struct {
	ID              string    `json:"id" gorm:"primary_key"`
	Cost            int       `json:"cost"`
	Fine            int       `json:"fine"`
	DateOfReturn    time.Time `json:"date_of_return"`
	ReturningAt     time.Time `json:"returning_at"`
	ReturningStatus bool      `json:"returning_status"`

	// foreign keys
	ID_Profile string  `json:"id_profile"`
	Profile    Profile `json:"profile" gorm:"foreignKey:ID_Profile`
	ID_Book    string  `json:"id_book"`
	Books      Books   `json:"books" gorm:"foreignKey:ID_Book"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
