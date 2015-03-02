package main

import (
"fmt"
"bytes"
"strconv"
"log"
"encoding/json"
"github.com/jinzhu/gorm"
"github.com/gorilla/mux"
_ "github.com/go-sql-driver/mysql"
"net/http")


type ProductListing struct {
	ID         int  `sql:"not null;unique;auto_increment;primary key"`
	Name     string
	Description      string
	Price float64
}

var db gorm.DB;

func main(){

	db = setupDB();
	r := mux.NewRouter()
	r.HandleFunc("/Create", create).Methods("POST");
	r.HandleFunc("/Update/{id}", update).Methods("PUT");
	r.HandleFunc("/Find/{name}", getListing).Methods("GET");
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}

func setupDB()  (gorm.DB) {

	db, err := gorm.Open("mysql", "YOURUSERNAME:YOURPASSWORD@/items_database?charset=utf8&parseTime=True")

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

/**************CRUD OPERATIONS*************/
//Some repetitive code here that can be refactored...

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
	var t ProductListing;   
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "Could not decode JSON", 400);
	}else{
    	//The response writer automagically sends 200 as the
    	//response code.
		addProductListing(&t);
		fmt.Fprintf(w, "All Good")

	}
	log.Println(t.Name)
}

func update(w http.ResponseWriter, r *http.Request){
	//Gets the id parameter
	params := mux.Vars(r);
	idStr := params["id"]
	id, errParse := strconv.Atoi(idStr)
    if errParse != nil {
        // handle error
        http.Error(w, "Not a valid ID", 400);
	}


	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
	var t ProductListing;   
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "Could not decode JSON", 400);
	}else{

		listing := getProductListingById(id);
		if (listing!=nil && listing.ID > 0){
			t.ID = listing.ID;
			db.Where("id=?", id).Save(&t);
			fmt.Fprintf(w, "All Good")
		}else{
			http.Error(w, "No listing with that id exists", 400);
		}
	}

}


func getListing(w http.ResponseWriter, r *http.Request){
	
	//Gets the name parameter
	params := mux.Vars(r);
	name := params["name"]

	//Find all items matching the name
	listings := getProductListingByName(name);

	//Some golang magic here:
	/*
		Coverts the array of listings to a json array
		and returns it.
	*/
		var buffer bytes.Buffer;
		jsonBytes, err := json.Marshal(listings)

		if err != nil {
			http.Error(w, "Couldn't encode JSON", 500);
		}

		buffer.WriteString(string(jsonBytes));
		fmt.Fprintf(w, "%s", buffer.String());
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


