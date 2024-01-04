package models

import "time"

type Users struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     bool   `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Profile struct {
	ID        string    `json:"id" gorm:"primary_key"`
	NIK       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Birthdate time.Time `json:"birthdate"`

	// Foreign Keys
	ID_User string `json:"id_user"`
	Users   Users  `json:"users" gorm:"foreign_key:ID_User"`
}
