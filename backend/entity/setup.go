package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("PVegDB.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(
		&Customer{},
	)
db = database

	password, _ := bcrypt.GenerateFromPassword([]byte("admin12345"), 14)

	db.Model(&Customer{}).Create(&Customer{
		Firstname:      "admin",
		Lastname:       "admin",
		Email:          "admin@admin.admin",
		Password:       string(password),
	})

	// database.Where(HotelType{Name: "Hotel"}).FirstOrCreate(&HotelType{Name: "Hotel"})
	// database.Where(HotelType{Name: "Resort"}).FirstOrCreate(&HotelType{Name: "Resort"})
	// database.Where(HotelType{Name: "Villa"}).FirstOrCreate(&HotelType{Name: "Villa"})
	// database.Where(HotelType{Name: "Apartment"}).FirstOrCreate(&HotelType{Name: "Apartment"})
	// database.Where(HotelType{Name: "Tent"}).FirstOrCreate(&HotelType{Name: "Tent"})

	// database.Where(RoomType{Name: "Standard"}).FirstOrCreate(&RoomType{Name: "Standard"})
	// database.Where(RoomType{Name: "Superior"}).FirstOrCreate(&RoomType{Name: "Superior"})
	// database.Where(RoomType{Name: "Deluxe"}).FirstOrCreate(&RoomType{Name: "Deluxe"})
	// database.Where(RoomType{Name: "Suite"}).FirstOrCreate(&RoomType{Name: "Suite"})
	// database.Where(RoomType{Name: "Tent 2 Adults"}).FirstOrCreate(&RoomType{Name: "Tent 2 Adults"})
	// database.Where(RoomType{Name: "Tent 4 Adults"}).FirstOrCreate(&RoomType{Name: "Tent 4 Adults"})

	// gender := []Gender{{
	// 	GenderType: "male",
	// }, {
	// 	GenderType: "female",
	// }, {
	// 	GenderType: "unknown",
	// }}
	// for _, gender := range gender {
	// 	db.Create(&gender)
	// }
	// prefix := []Prefix{{
	// 	PrefixType: "Mr.",
	// }, {
	// 	PrefixType: "Mrs.",
	// }, {
	// 	PrefixType: "Ms.",
	// }, {
	// 	PrefixType: "Miss.",
	// }, {
	// 	PrefixType: "unknown",
	// }}
	// for _, prefix := range prefix {
	// 	db.Create(&prefix)
	// }
	// country := []Country{{
	// 	CountryName: "Loei",
	// }, {
	// 	CountryName: "Nakonrachsima",
	// }, {
	// 	CountryName: "unknown",
	// }}
	// for _, country := range country {
	// 	db.Create(&country)
	// }

	// //bookignFlight
	// foodFlight := []FoodFlight{{
	// 	Maincourse: "nothing",
	// 	Price:      0,
	// }, {
	// 	Maincourse: "padthai",
	// 	Price:      200,
	// }, {
	// 	Maincourse: "hoytod",
	// 	Price:      300,
	// }}
	// for _, foodFlight := range foodFlight {
	// 	db.Create(&foodFlight)
	// }

	// baggeg := []Baggeg{{
	// 	Weight: "nothing",
	// 	Price:  0,
	// }, {
	// 	Weight: "50kg",
	// 	Price:  50,
	// }, {
	// 	Weight: "100kg",
	// 	Price:  100,
	// }}
	// for _, baggeg := range baggeg {
	// 	db.Create(&baggeg)
	// }

	// //payment
	// statuse := []PaymentStatus{{
	// 	Status: "NO",
	// }, {
	// 	Status: "YES",
	// }, {
	// 	Status: "FAIL",
	// }}
	// for _, statuse := range statuse {
	// 	db.Create(&statuse)
	// }

	// TypeName := []PaymentType{{
	// 	TypeName: "Mobile Banking",
	// }, {
	// 	TypeName: "Promptpay",
	// }}
	// for _, TypeName := range TypeName {
	// 	db.Create(&TypeName)
	// }
	// //company
	// company := []Company{{
	// 	CompanyName: "Airthailand",
	// }, {
	// 	CompanyName: "NokAirConditioning",
	// }, {
	// 	CompanyName: "ElysianAirlines",
	// }}
	// for _, company := range company {
	// 	db.Create(&company)
	// }

	// //to
	// to := []To{{
	// 	ToName: "สุวรรณภูมิ",
	// }, {
	// 	ToName: "ภูเก็ต",
	// }, {
	// 	ToName: "เชียงใหม่",
	// }}
	// for _, to := range to {
	// 	db.Create(&to)
	// }

	// //from
	// from := []From{{
	// 	FromName: "กระบี่",
	// }, {
	// 	FromName: "กรุงเทพ",
	// }, {
	// 	FromName: "หาดใหญ่",
	// }}
	// for _, from := range from {
	// 	db.Create(&from)
	// }

	// newBooking := &BookFlight{}
	// db.Create(newBooking)

	// newBookingHotel := &BookHotel{}
	// db.Create(newBookingHotel)

}
