package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Firstname      string
	Lastname       string
	Email          string `gorm:"uniqueIndex" valid:"required~Email is required, email~Email is invalid"`
	Password       string `valid:"required~Password is required, stringlength(8|100)~Password ต้องมากกว่า 7 แต่น้อยกว่า 100 ตัว"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("isPositive", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		n := i.(int)
		return n > 0
	}))
}
