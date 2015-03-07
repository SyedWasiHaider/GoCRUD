package main

import (
"github.com/jinzhu/gorm"
"log"
)

func setupDB()  (gorm.DB) {

	db, err := gorm.Open("mysql", "root:aw3s0m3@/items_database?charset=utf8&parseTime=True")

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

	return db;

}


/**************DB OPERATIONS***************/

	func addProductListing(listing * ProductListing){

	//To ensure the ID is set by the DB and not the user.
		listing.ID = 0; 
		db.NewRecord(&listing)
		db.Create(&listing)
	}

	func getProductListingById(id int) (*ProductListing){
	//Find all items matching the name
		var listing ProductListing;
		db.Where("id = ?", id).Find(&listing)
		return &listing;
	}

	func getProductListingByName(name string) ([] ProductListing){
	//Find all items matching the name
		var listings [] ProductListing;
		db.Where("name = ?", name).Find(&listings)
		return listings;
	}