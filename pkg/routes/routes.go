//This package lists the routes to be registered
//against the incoming requests to the HTTP server.
//The routes are mapped to handler functions which
//are based on the controllers to the application.

//The routes include variable query values for 
//creating and querying collections thereby providing
//the function to create multiple collections within the 
//database. 

//Here, we have registered 4 routes mapping URLs to
//handlers. The corresponding controller handler is 
//called passing (http.ResponseWriter, *http.Request) 
//as parameters.

//HTTP methods, including GET, POST, and DELETE,
//are also added as matchers to match similar routes
//to different handlers based on the HTTP method added.

package routes

import (
	"github.com/gorilla/mux"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/controllers"
)

var RegisterRoutes = func(r *mux.Router){

	r.HandleFunc("/{collection}", controllers.ReadDocuments).Methods("GET")
	r.HandleFunc("/{collection}", controllers.CreateDocument).Methods("POST")
	r.HandleFunc("/{collection}/{id}", controllers.ReadDocumentById).Methods("GET")
	r.HandleFunc("/{collection}/{id}", controllers.DeleteDocument).Methods("DELETE")
}