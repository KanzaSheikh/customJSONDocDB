//The models package defines the structure for the
//`user` type which is used for implementing
//documents within the collections. JSON mapping is also
//incorporated to enable encoding and decoding
//the JSON-based HTTP requests. 

package models

import(
	"encoding/json"
)
type (
	Address struct{
		City string `json:"city"`
		State string `json:"state"`
		Country string `json:"country"`
		Pincode json.Number `json:"pincode"`
	}
	
	User struct{
		ID string `json:"id"`
		Name string `json:"name"`
		Age json.Number `json:"age"`
		Contact string `json:"contact"`
		Company string `json:"company"`
		Address Address `json:"address"`
	}
)