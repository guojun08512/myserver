package db

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	//gorm.Model
	ID  		string		`gorm:"primary_key;type:varchar(36);not null;"`
	Name     	string 		`gorm:"type:varchar(32)"`
	Password 	string    	`gorm:"type:varchar(32)"`
	Roles    	string
	Disable  	bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time 	`gorm:"index" json:"-"`
	Email       string     	`gorm:"type:varchar(32)"`
	Phone       string      `gorm:"type:varchar(20)"`
}

func CreateUser(db *gorm.DB, name, password, email, role, phone string) *User {
	u := User{
		ID: uuid.NewV4().String(),
		Name: name,
		Password: password,
		Email: email,
		Roles: role,
		Phone: phone,
	}
	db.Create(&u)
	return &u
}

func GetUser(db *gorm.DB, name, password string) (*User, error) {
	var user *User
	t := db.First(&user).Where("name=(?) and password=(?)", name, password)
	if t.Error != nil {
		return nil, t.Error
	}
	return user, nil
}
