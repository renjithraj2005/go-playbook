package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need go-sqlite3 https://github.com/mattn/go-sqlite3
)

// InitORM This method will init the db and do the migrations
func InitORM() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		//Panic is a built-in function that stops the ordinary flow of control and begins panicking.
		//When the function F calls panic, execution of F stops, any deferred functions in F are executed normally,
		//and then F returns to its caller
		panic("failed to connect database")
	}
	defer db.Close() //remember the use of defer? A defer statement defers the execution of close until the InitORM function returns

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Renjith", Email: "renjith@sayonetech.com", Gender: "Male", City: "Cochin"})

	// Read
	var user User
	db.First(&user, 1)                     // find user with id 1
	db.First(&user, "name = ?", "Renjith") // find user with name Renjith

	// Update - update user's city to Kochi
	db.Model(&user).Update("City", "Kochi")

	// Delete - delete product
	db.Delete(&user)
}
