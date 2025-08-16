package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name" validate:"required,min=2,max=50"`
	Email        string `gorm:"unique" json:"email" validate:"required,email"`
	MobileNumber string `json:"mobile_number" validate:"required,phone"`
	PAN          string `gorm:"unique" json:"pan" validate:"required,pan"`
	Password     string `json:"-" validate:"required,min=6"`
}

func (u *User) SetPassword(pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return false
	}
	return true
}
