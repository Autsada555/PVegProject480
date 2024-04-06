package entity

import (
	"time"

	// "github.com/asaskevich/govalidator"
	// unit "github.com/chadawan9913/sa-lab4-66/AppData/Roaming/Code/User/History/5d60e7a8"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Firstname string `valid:"required~FirstName is required"`
	Lastname  string `valid:"required~LastName is required"`
	Phone     string `valid:"required~Phone number is required,stringlength(10|10)~Phone must be at 10 characters"`
	Email     string `gorm:"uniqueIndex" valid:"required~Email is required, email~Email is invalid"`
	Profile   string `gorm:"type:longtext"`
	Password  string `valid:"required~Password is required, stringlength(8|100)~Password ต้องมากกว่า 7 แต่น้อยกว่า 100 ตัว"`

	GenderID uint
	Gender   *Gender

	AddressID uint
	Address   *Address

	Order   []Order
	Payment []Payment
}

type Gender struct {
	gorm.Model
	TypeName string `gorm:"unique"`

	Member []Member
}

// type Role struct {
// 	gorm.Model
// 	Name string `gorm:"unique"`

// 	Member []Member
// }

type Address struct {
	gorm.Model
	Address  string `valid:"required~Address is required"`
	District string `valid:"required~District is required"`
	Province string `valid:"required~Province is required"`
	Postcode string `valid:"required~Postcode is required"`

	MemberID *uint  `valid:"-"`
	Member   Member `gorm:"references:id" valid:"-"`
}

type Order struct {
	gorm.Model
	TotalAmount float32
	Date        time.Time
	Quantity    int `valid:"required~Quantity is required, range(1|10)~Quantity should be between 1 and 10"`

	MemberID uint    `valid:"-"`
	Member   *Member `gorm:"references:id" valid:"-"`

	MenuID uint  `valid:"-"`
	Munu   *Menu `gorm:"references:id" valid:"-"`

	AddressID *uint   `valid:"-"`
	Address   Address `gorm:"references:id" valid:"-"`

	Payment []Payment
}

type Menu struct {
	gorm.Model
	MenuName  string
	MenuCost  float32
	MenuImage string `gorm:"type:longtext"`

	MenuTypeID uint `valid:"-"`
	MunuType   *MenuType
}

type MenuType struct {
	gorm.Model
	TypeName string `gorm:"unique"`

	Menu []Menu `gorm:"foreignKey:MenuTypeID"`
}

type Payment struct {
	gorm.Model
	Totalmount float32
	Time       time.Time
	SlipImage  string

	DeliveryID uint `valid:"-"`
	Delivery   *Delivery

	BankTypeID uint `valid:"-"`
	BankType   *BankType

	OrderID uint `valid:"-"`
	Order   *Order

	MemberID uint `valid:"-"`
	Member   *Member
}

type BankType struct {
	gorm.Model
	BankName string `gorm:"unique"`

	BankType []BankType `gorm:"foreignKey:BankTypeID"`
}

type Delivery struct {
	gorm.Model
	Name string `gorm:"unique"`

	Delivery []Delivery `gorm:"foreignKey:DeliveryID"`
}
