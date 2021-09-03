// The requirements are to:
// - Create a DB package using golang which satisfies the interface (attached). Feel free to use a file (or an in-memory structure/map) as a DB datastore
// - Create a typical golang project, which uses the above DB package and builds a binary that can be run from the cmd-line
// - Provide go-test cases to verify that the DB works

package testdb

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type BackendDatabase interface {
	New() error                 // creates a new database
	Put([]byte, []byte) error   // creates a new entry or updates existing entry
	Get([]byte) ([]byte, error) // gets value in map
	Delete([]byte) error        // deletes key, value in map
	Close()                     // exits the program
	Stats() string              // print out the keys/values in map?
	Flush() error               // clear the map of all keys and values
}

type Database struct {
	database map[string]string
}

func (db *Database) New() error {
	db.database = make(map[string]string)

	t := fmt.Sprintf("%T", db.database)
	if !strings.HasPrefix(t, "map[string]string") {
		return errors.New("error New(): Could not create new database")
	}

	return nil
}

func (db *Database) Put(key []byte, value []byte) error {
	if string(key) == "" {
		return errors.New("error Put(): invalid key")
	} else if string(value) == "" {
		return errors.New("error Put(): invalid value")
	}

	db.database[string(key)] = string(value)
	return nil
}

func (db *Database) Get(key []byte) ([]byte, error) {
	if string(key) == "" {
		return nil, errors.New("error Get(): invalid key")
	}

	// if key is in database...
	if v, found := db.database[string(key)]; found {
		return []byte(v), nil
	}

	return nil, errors.New("error Get(): key does not exist")
}

func (db *Database) Delete(key []byte) error {
	if string(key) == "" {
		return errors.New("error Delete(): invalid key")
	}

	delete(db.database, string(key))
	return nil
}

func (db *Database) Flush() error {
	for key := range db.database {
		delete(db.database, string(key))
	}

	if len(db.database) > 0 {
		return errors.New("error Flush(): could not delete all entries in database")
	}

	return nil
}

func (db *Database) Close() {
	os.Exit(3)
}

func (db *Database) Stats() string {
	var result string = ""

	for key := range db.database {
		result += key + " " + db.database[key] + "\n"
	}

	return result
}
