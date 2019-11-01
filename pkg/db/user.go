package db

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	gorm.Model
	ID   		int64		`gorm:"primary_key;size:32"`
	Name     	string 		`json:"name"`
	Password 	string    	`json:"password"`
	Roles    	string 		`json:"roles"`
	Disable  	bool  		`json:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time 	`gorm:"index" json:"-"`
	Email       string     	`gorm:"size=32"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

func CreateUser(db *gorm.DB, name, password, email, role string) *User {
	u := User{
		Name: name,
		Password: password,
		Email: email,
		Roles: role,
	}
	db.Save(&u)
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
