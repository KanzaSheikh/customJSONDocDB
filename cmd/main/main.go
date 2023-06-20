//This is the main package for tha application.
//It defines a request router and dispatcher for 
//matching incoming requests to their respective handler
// using the mux package.

//A set of registered routes is introduced against which
// the URL is matched and handler is called. Requests are
//matched on the basis of variable path and query values of the URL.

//Further, a HTTP server is implemented using ListenAndServe
// with a given address and handler. We use the mux router as the
//handler for starting the server.

package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/routes"
	"github.com/KanzaSheikh/customJSONDocDB/pkg/configs"
)

const Version = "1.0.0" 

func main(){
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	_, err := configs.New(configs.Dir, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}