package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

func setupDB() gorm.DB {

	//var username=YOURUSERNAME
	//var pasword=YOURPASSWORD

	db, err := gorm.Open("mysql", strings.Join([]string{username, ":", password, "@/items_database?charset=utf8&parseTime=True"}, ""))

	if err != nil {
		log.Fatal(err)
	}

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//This will give a warning if the table already exists.
	db.CreateTable(&ProductListing{})
	db.CreateTable(&User{})

	return db

}

/**************DB OPERATIONS***************/

func addUser(user *User) {

	//To ensure the ID is set by the DB and not the user.
	user.ID = 0
	db.NewRecord(&user)
	db.Create(&user)
}

func addProductListing(listing *ProductListing) {

	//To ensure the ID is set by the DB and not the user.
	listing.ID = 0
	db.NewRecord(&listing)
	db.Create(&listing)
}

func getProductListingById(id int) *ProductListing {
	//Find all items matching the name
	var listing ProductListing
	db.Where("id = ?", id).Find(&listing)
	return &listing
}

func getProductListingByName(name string) []ProductListing {
	//Find all items matching the name
	var listings []ProductListing
	db.Where("name = ?", name).Find(&listings)
	return listings
}

func deleteProductListingById(id int) {
	db.Delete(ProductListing{ID: id})
}
