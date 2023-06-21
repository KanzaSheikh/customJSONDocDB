# customJSONDocDB
A custom NoSQL database using JSON documents written in Go


## What is it
It is a custom-built NoSQL database which supports collections and JSON documents. It also provides an API for performing the CRUD operations on the data (currently only CRD!).

## Get Started
To run the database locally:
1. Clone the repository ```git clone https://github.com/KanzaSheikh/customJSONDocDB.git```
2. Navigate to the cmd/main directory
3. Run the project ```go run main.go```

## Data Manipulation
Upon running the application, an HTTP server is started which can be accessed through the local API at port 8000 to perform simple CRUD operations.

## How to?
Following are a few example operations

#### READ
The database supports two kind of read operations i.e. reading all documents and reading documents by ID.

For reading all documents within a collection, run:

```curl -i -s -X GET http://localhost:8000/{collection}```

For reading a document by ID:

```curl -i -s -X GET http://localhost:8000/{collection}/{id}```

#### CREATE
New documents can be created within a collection by using the following route. If a collection already exists, the document is created inside it, otherwise a new collection is created.

```curl -i -s -X POST http://localhost:8000/users -H "Content-Type: application/json" -d {JSON-body}```

#### DELETE
For deleting existing documents, run:

```curl -i -s -X DELETE http://localhost:8000/users/{id}```