# customJSONDocDB
A custom NoSQL database using JSON documents written in Go


## What is it
It is a custom-built NoSQL database which supports collections and JSON documents. It also provides an API for performing the CRUD operations on the data.

## Get started
To run the database locally:
1. Clone the repository ```git clone https://github.com/KanzaSheikh/customJSONDocDB.git```.
2. Navigate to the cmd/main directory.
3. Run the project ```go run main.go```.

## Data operations
Upon running the application, an HTTP server is started which can be accessed through the local API at port 8000 to perform simple CRUD operations.

## How to?
Following are a few example operations:

#### READ
The database supports two kinds of read operations i.e. reading all documents and reading documents by ID.

For reading all documents within a collection, run:

```curl --request GET http://localhost:8000/users```

For reading a document by ID:

```curl --request GET http://localhost:8000/users/{id}```

#### CREATE
New documents can be created within a collection by using the following route. If a collection already exists, the document is created inside it, otherwise a new collection is created.

```curl --request POST http://localhost:8000/users --header "Content-Type: application/json" --data '{body}'```

#### DELETE
For deleting existing documents, run:

```curl --request DELETE http://localhost:8000/users/{id}```