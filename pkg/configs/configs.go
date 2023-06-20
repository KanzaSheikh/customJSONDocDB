//The configs package introduces the initial configurations
//for setting up the database and prerequisite functions
//for performing the CRUD operations.

//It sets up the path for and creates the relevant
//directory for the purpose of storing documents within
//collections. The directory is created as per the shortest 
//path by lexical processing of the directory structure.
//A custom driver is implemented for the databses
//which provides functionality for receiver methods in
//implementing mutexes.

//Mutexes are added to provide basic synchronization primitives 
//such as mutual exclusion locks. This enables successfully
//performing DB operations such as WRITE and DELETE
//by locking the mutex, so that no two operations take
//place concurrently potentially causing dirty writes.


package configs

import (
	"fmt"
	"sync"
	"os"
	"path/filepath"
	"github.com/jcelliott/lumber"
)

func New(dir string, options *Options)(*Driver, error){
	dir = filepath.Clean(dir)
	opts := Options{}
	if options != nil {
		opts = *options 
	}
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}
	driver := Driver{
		dir: dir,
		mutexes: make(map[string]*sync.Mutex),
		log: opts.Logger,
	}
	if _, err := os.Stat(dir); err == nil{
		opts.Logger.Debug("Database already exists")
		return &driver, nil
	}
	opts.Logger.Debug("Creating the database")
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex{
	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

func GetOrCreateMutex(dir string, collection string) *sync.Mutex{
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println(err)
	}
	mutex := db.getOrCreateMutex(collection)
	return mutex
}