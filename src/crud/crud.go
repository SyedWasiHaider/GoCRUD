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
	setupRoutesAndServe();
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

