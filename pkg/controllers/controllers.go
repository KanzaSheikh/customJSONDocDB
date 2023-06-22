//The controllers package implements the handler functions
//for mapping the API routes to the relevant functions.
//The functions receive the HTTP request and response writer
//as parameters to handle the request content.

package controllers

import(
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"net/http"
	"path/filepath"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/configs"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/models"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/utils"
)

var (
	_dir string = "../.."
)

//ReadDocument is the handler for the basic route with the 
//HTTP GET method. The variable path prefix defines the value
//for collection. The collection name is retrived using 
//the mux.Vars() function. 

//The directory for the specified collection is read by
//combining the _dir literal path with the collection using 
//the utils function for reading directory. Each file returned 
//by utils.ReadDir is parsed and the contents are mapped to 
//the internal User struct as defined by the models package.

//An array of documents is created which is encoded by the JSON
//package and returned to the HTTP ResponseWriter.

func ReadDocuments(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	params := mux.Vars(r)
	collection := params["collection"]
	if collection == ""{
		fmt.Println("Collection is missing")
	}
	dir := filepath.Join(_dir, collection)
	if _, err := os.Stat(dir); err != nil {
		fmt.Println("Database does not exist")
	}
	files, _ := utils.ReadDir(dir)
	for _, file := range files{
		user := &models.User{}
		path := filepath.Join(dir, file.Name())
		utils.ParseFile(path, user)
		users = append(users, *user)
	json.NewEncoder(w).Encode(users)
	users = nil
	}
}

//CreateDocument is the handler for the registered 
//collection route with the HTTP POST method. The 
//variable path prefix defines the value for collection. 
//The collection name is retrived using the mux.Vars()
//functions. 

//The content of the HTTP request is decoded using the JSON
//decoder which is mapped to a variable of models.User type.

//A mutex is aquired for the synchronization primitive 
//while creating the document. The mutex is locked such 
//that a new goroutine is not able to acquire the mutex 
//until it is unlocked which is deferred.

//The utils function for writing file is utilized to write
//the .json file containing the API request content
//to the path specified by combining the _dir literal,
//collection name, and the resource name.

func CreateDocument(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	params := mux.Vars(r)
	collection := params["collection"]
	_ = json.NewDecoder(r.Body).Decode(&user)

	mutex := configs.GetMutex(_dir, collection)
	mutex.Lock()
	defer mutex.Unlock()

	resource := user.Name
	user.ID = strconv.Itoa(rand.Intn(100000000))
	dir := filepath.Join(_dir, collection)
	path := filepath.Join(dir, resource+".json")
	utils.WriteFile(dir, path, user)
	json.NewEncoder(w).Encode(user)
}

//ReadDocumentByID is the handler for the registered route 
//with the HTTP GET method and the ID for the document. 
//The variable path prefix defines the values
//for collection and ID. The variable values are retrived 
//using the mux.Vars() function and stored against concrete
//type variabless. 

//The directory for the specified collection is read by
//combining the _dir literal path with the collection using 
//the utils function for reading directory. Each file returned 
//by utils.ReadDir is parsed and the contents are mapped to 
//the internal User struct as defined by the models package.

//The ID field for each User is compared with the ID variable
//in the route variables list. If the two values matched, the 
//document content is encoded and is written to the ResponseWriter.

func ReadDocumentById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := params["collection"]
	if collection == ""{
		fmt.Println("Collection is missing")
	}
	dir := filepath.Join(_dir, collection)
	if _, err := os.Stat(dir); err != nil {
		fmt.Println("Database does not exist")
	}
	files, _ := utils.ReadDir(dir)
	for _, file := range files{
		user := &models.User{}
		path := filepath.Join(dir, file.Name())
		utils.ParseFile(path, user)
		if user.ID == params["id"]{
			json.NewEncoder(w).Encode(*user)
			return
		}
	}
}

//DeleteDocument is the handler for the registered route 
//with the HTTP DELETE method and the ID for the document. 
//The variable path prefix defines the values
//for collection and ID. The variable values are retrived 
//using the mux.Vars() function and stored against concrete
//type variabless. 

//This method is used to delete the document under the
//collections directory with the specified ID variable.
//The directory for the specified collection is read by
//combining the _dir literal path with the collection using 
//the utils function for reading directory. Each file returned 
//by utils.ReadDir is parsed and the contents are mapped to 
//the internal User struct as defined by the models package.

//The ID field for each User is compared with the ID variable
//in the route variables list. If the two values matched, the 
//path of the document is used to delete the file using the
//utils.DeleteFile function which uses the os.RemoveAll()
//function under the hood.

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := params["collection"]
	if collection == ""{
		fmt.Println("Collection is missing")
	}

	mutex := configs.GetMutex(_dir, collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(_dir, collection)
	if _, err := os.Stat(dir); err != nil {
		fmt.Println("Directory not okay")
	}
	files, _ := utils.ReadDir(dir)
	for _, file := range files{
		user := &models.User{}
		path := filepath.Join(dir, file.Name())
		utils.ParseFile(path, user)
		if user.ID == params["id"]{
			utils.DeleteFile(path)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}