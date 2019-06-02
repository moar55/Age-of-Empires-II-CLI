package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func handleInput() string {
	if len(os.Args) == 1 { // check that input is passed
		panic("Unit name or id must be provided!")
	}

	var unit string

	if len(os.Args) == 3 {
		if os.Args[1] == "-v" { //check for verbose flag as first input
			verbose = true
			unit = os.Args[2] //unit name or id passed through command line argument
		} else {
			panic("malformed input")
		}
	} else {
		unit = os.Args[1] //input is of size 1, expecting an id or name as input
		if unit == "-v" {
			panic("malformed input")
		}
	}
	return unit
}

func checkInDB(units *mongo.Collection, unit string) (item Unit, err error) {
	var myUnit Unit
	unitInt, err := strconv.Atoi(unit) //try to convert usre input
	var filter bson.D

	if err != nil {
		filter = bson.D{{"name", unit}} //create a bson document query to search for unit by name
	} else {
		filter = bson.D{{"id", unitInt}} //similar to the above but searching for unity by id
	}

	err = units.FindOne(context.TODO(), filter).Decode(&myUnit) //query database for unit using filter

	if err != nil {
		return myUnit, err // return emtpy Unit struct and err
	}

	return myUnit, nil // return unit fetched from db
}

//Verbose flag, if true less important logs will be printed
var verbose = false

func main() {
	unit := handleInput() // handle user input

	err := godotenv.Load() //attempt to load values in .env file, if present

	if verbose && err != nil {
		fmt.Println("Error loading .env file")
	}

	collection := setupDB() //setup database configuration

	var myUnit Unit

	/*check if item already exists in db*/
	myUnit, err = checkInDB(collection, unit)

	if verbose && err != nil {
		fmt.Println("failed to retreive unit from database")
	} else if err != nil {
		marshalled, _ := json.MarshalIndent(myUnit, "", " ") //marshal myUnit struct with indentation for readability
		fmt.Println(string(marshalled))                      //convert marshalled byte sequence to string
		return
	}

	/*failed to retrieve unit from db*/

	//fetch the desired unit from API
	resp, err := http.Get("https://age-of-empires-2-api.herokuapp.com/api/v1/unit/" + unit)

	if err != nil { // check for errrors in fetching the data
		panic("Error retrieving unit data")
	}

	defer resp.Body.Close() //close reader when main returns

	body, _ := ioutil.ReadAll(resp.Body) //read from reader
	json.Unmarshal(body, &myUnit)        //unmarshall body byte sequence to be stored as a Unit struct

	_, err = collection.InsertOne(context.TODO(), myUnit) //save result in database

	if verbose && err != nil {
		fmt.Println("failed to cache results", err)
	}

	fmt.Println(string(body))

}
