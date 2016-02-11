package users

import "encoding/json"
import "errors"
import "io/ioutil"
import "os"

type UserStore struct {
	Users    map[string]User
	filename string
}

// NewUserStore creates a new UserStore instance, and loads
// the users from the JSON file.
func NewUserStore(filename string) (s *UserStore, err error) {
	s = new(UserStore)
	s.filename = filename
	err = s.Load()

	return
}

// Load users from the JSON file
func (s *UserStore) Load() error {

	file, err := os.Open(s.filename)

	if err != nil {
		return errors.New("Could not open file: " + err.Error())
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&s.Users)

	if err != nil {
		return errors.New("Could not parse file: " + err.Error())
	} else {
		return nil
	}

}

// Save user database to the JSON file
func (s *UserStore) Save() error {

	json, err := json.MarshalIndent(s.Users, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.filename, json, 0644)
	if err != nil {
		return errors.New("Couldn't write JSON to disk: " + err.Error())
	}

	return nil

}

// Add user to store (and Save)
func (s *UserStore) Add(u User) error {
	if s.Exists(u.Name) {
		return errors.New("User " + u.Name + " already exists")
	} else {
		s.Users[u.Name] = u
	}
	return s.Save()
}

// Exists checks whether a user exists in the database or not
func (s *UserStore) Exists(name string) bool {
	if _, ok := s.Users[name]; ok {
		return true
	} else {
		return false
	}
}

// Get a user from the database
func (s *UserStore) Get(name string) *User {
	if u, ok := s.Users[name]; ok {
		return &u
	} else {
		return nil
	}
}

// Delete a user from the database (and Save)
func (s *UserStore) Delete(name string) error {
	delete(s.Users, name)
	return s.Save()
}
