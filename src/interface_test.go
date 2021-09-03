package testdb

import (
	"testing"
)

type mockDatabase struct {
	db       Database
	actual   error
	expected error
}

type mockDatabasePut struct {
	db       Database
	actual   error
	expected error
}

type mockDatabaseGet struct {
	db       Database
	actual   []byte
	expected []byte
	err      error
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// New()
func TestNewTable(t *testing.T) {

	t.Run("Test 1: should not init database", func(t *testing.T) {
		testCase := mockDatabasePut{
			db: Database{},
		}

		// Database should be nil before calling New()
		if testCase.db.database != nil {
			t.Fatalf("expected %v, got %v", nil, testCase.db.database)
		}

	})

	t.Run("Test 2: should init database", func(t *testing.T) {
		testCase := mockDatabasePut{
			db: Database{},
		}

		testCase.db.New()

		// Database should NOT be nil after calling New()
		if testCase.db.database == nil {
			t.Fatalf("expected %v, got %v", nil, testCase.db.database)
		}

	})

	t.Run("Test 3: should be able to add elements to database", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.database[string("1")] = "0"

		// if database has been initialized correctly, we should be able to add elements to it
		if _, found := testCase.db.database[string("1")]; !found {
			t.Fatalf("Database not initialized correctly. Expected to find key: 1 but didn't")
		}

	})

}

//////////////////////////////////////////////////////////////////////////////////////////////////
// Put()
func TestPutTable(t *testing.T) {

	t.Run("Test 1: should return nil", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte("1"), []byte("0"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		value := testCase.db.database[string("1")]
		if value != "0" {
			t.Fatalf("expected %v, got %v", "0", value)
		}
	})

	t.Run("Test 2: should return nil", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte("0"), []byte("1"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		value := testCase.db.database[string("0")]
		if value != "1" {
			t.Fatalf("expected %v, got %v", "1", value)
		}
	})

	t.Run("Test 3: should return nil", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte("1"), []byte("-1"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		value := testCase.db.database[string("1")]
		if value != "-1" {
			t.Fatalf("expected %v, got %v", "-1", value)
		}
	})

	t.Run("Test 4: should return invalid key error", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte(""), []byte("0"))

		if string(testCase.actual.Error()) != string("error Put(): invalid key") {
			t.Fatalf("expected error, got something else")
		}

	})

	t.Run("Test 5: should return invalid value error", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte("A"), []byte(""))

		if string(testCase.actual.Error()) != string("error Put(): invalid value") {
			t.Fatalf("expected error, got something else")
		}

	})

	t.Run("Test 6: should return nil", func(t *testing.T) {
		testCase := mockDatabasePut{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.actual = testCase.db.Put([]byte("`~!@#$%^&*()_+"), []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ "))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		value := testCase.db.database[string("`~!@#$%^&*()_+")]
		if value != "ABCDEFGHIJKLMNOPQRSTUVWXYZ " {
			t.Fatalf("expected %v, got %v", "ABCDEFGHIJKLMNOPQRSTUVWXYZ ", value)
		}

	})

}

//////////////////////////////////////////////////////////////////////////////////////////////////
// Get()
func TestGetTable(t *testing.T) {
	t.Run("Test 1: should return 0", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db:       Database{},
			expected: []byte("0"),
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))

		testCase.actual, testCase.err = testCase.db.Get([]byte("1"))

		if string(testCase.actual) != string(testCase.expected) || testCase.err != nil {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

	})

	t.Run("Test 2: should return 1", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db:       Database{},
			expected: []byte("1"),
		}

		testCase.db.New()
		testCase.db.Put([]byte("0"), []byte("1"))

		testCase.actual, testCase.err = testCase.db.Get([]byte("0"))

		if string(testCase.actual) != string(testCase.expected) || testCase.err != nil {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

	})

	t.Run("Test 3: should return -1", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db:       Database{},
			expected: []byte("-1"),
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("-1"))

		testCase.actual, testCase.err = testCase.db.Get([]byte("1"))

		if string(testCase.actual) != string(testCase.expected) || testCase.err != nil {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

	})

	t.Run("Test 4: should return invalid key error", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db: Database{},
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))

		testCase.actual, testCase.err = testCase.db.Get([]byte(""))

		if string(testCase.err.Error()) != string("error Get(): invalid key") {
			t.Fatalf("expected error, got something else")
		}

	})

	t.Run("Test 5: should return key does not exist error", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db: Database{},
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))

		testCase.actual, testCase.err = testCase.db.Get([]byte("0"))

		if string(testCase.err.Error()) != string("error Get(): key does not exist") {
			t.Fatalf("expected error, got something else")
		}

	})

	t.Run("Test 6: should return ABCDEFGHIJKLMNOPQRSTUVWXYZ ", func(t *testing.T) {
		testCase := mockDatabaseGet{
			db:       Database{},
			expected: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ "),
		}

		testCase.db.New()
		testCase.db.Put([]byte("`~!@#$%^&*()_+"), []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ "))

		testCase.actual, testCase.err = testCase.db.Get([]byte("`~!@#$%^&*()_+"))

		if string(testCase.actual) != string(testCase.expected) {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

	})

}

//////////////////////////////////////////////////////////////////////////////////////////////////
// Delete()

func TestDeleteTable(t *testing.T) {
	t.Run("Test 1: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))

		testCase.actual = testCase.db.Delete([]byte("1"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if v, found := testCase.db.database[string("1")]; found {
			t.Fatalf("found %v in database", v)
		}

	})

	t.Run("Test 2: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.Put([]byte("0"), []byte("1"))

		testCase.actual = testCase.db.Delete([]byte("0"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if v, found := testCase.db.database[string("0")]; found {
			t.Fatalf("found %v in database", v)
		}

	})

	t.Run("Test 3: should return invalid key error", func(t *testing.T) {
		testCase := mockDatabase{
			db: Database{},
		}

		testCase.db.New()
		testCase.db.Put([]byte("0"), []byte("1"))

		testCase.actual = testCase.db.Delete([]byte(""))

		if string(testCase.actual.Error()) != string("error Delete(): invalid key") {
			t.Fatalf("expected error, got %v", testCase.actual)
		}

	})

	t.Run("Test 4: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.Put([]byte("0"), []byte("1"))

		// Tries to delete a key that doesn't exist
		testCase.actual = testCase.db.Delete([]byte("1"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

	})

	t.Run("Test 5: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))
		testCase.db.Put([]byte("0"), []byte("1"))
		testCase.db.Put([]byte("-1"), []byte("1"))

		testCase.actual = testCase.db.Delete([]byte("-1"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if v, found := testCase.db.database[string("-1")]; found {
			t.Fatalf("found %v in database", v)
		}

	})

	t.Run("Test 5: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		testCase.db.New()
		testCase.db.Put([]byte("`~!@#$%^&*()_+"), []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ "))

		testCase.actual = testCase.db.Delete([]byte("`~!@#$%^&*()_+"))

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if v, found := testCase.db.database[string("`~!@#$%^&*()_+")]; found {
			t.Fatalf("found %v in database", v)
		}

	})

}

//////////////////////////////////////////////////////////////////////////////////////////////////
// Flush()

func TestFlushTable(t *testing.T) {
	t.Run("Test 1: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		// Flush an uninitialized database
		testCase.actual = testCase.db.Flush()

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if len(testCase.db.database) > 0 {
			t.Fatalf("Database is not empty")
		}

	})

	t.Run("Test 2: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		// Flush an empty database
		testCase.db.New()
		testCase.actual = testCase.db.Flush()

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if len(testCase.db.database) > 0 {
			t.Fatalf("Database is not empty")
		}

	})

	t.Run("Test 3: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		// Flush a database with 1 element
		testCase.db.New()
		testCase.db.Put([]byte("1"), []byte("0"))
		testCase.actual = testCase.db.Flush()

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if len(testCase.db.database) > 0 {
			t.Fatalf("Database is not empty")
		}

	})

	t.Run("Test 4: should return nil", func(t *testing.T) {
		testCase := mockDatabase{
			db:       Database{},
			expected: nil,
		}

		// Flush a database with 1000 elements
		testCase.db.New()

		n := 1000
		i := 0
		for i < n {
			testCase.db.Put([]byte(string(rune(i))), []byte(string(rune(i))))
			i++
		}

		testCase.actual = testCase.db.Flush()

		if testCase.actual != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, testCase.actual)
		}

		if len(testCase.db.database) > 0 {
			t.Fatalf("Database is not empty")
		}

	})

}
