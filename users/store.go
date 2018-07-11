package users

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

var (
	mu    sync.Mutex
	store map[string]User
)

type User struct {
	ID       string
	Nicklist []string
	Running  bool
}

func Open() error {
	store = make(map[string]User)

	buf, err := ioutil.ReadFile("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(buf, &store)
	return err
}

func GetOrCreate(id string) (User, error) {
	mu.Lock()
	defer mu.Unlock()

	if user, ok := store[id]; ok {
		return user, nil
	}
	store[id] = User{
		ID: id,
	}
	return store[id], save()
}

func GetAll() []User {
	mu.Lock()
	defer mu.Unlock()

	users := make([]User, 0, len(store))
	for _, val := range store {
		users = append(users, val)
	}
	return users
}

func Update(user User) error {
	mu.Lock()
	defer mu.Unlock()

	store[user.ID] = user
	return save()
}

func save() error {
	buf, err := json.Marshal(store)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("users.json", buf, 0644)
	if err != nil {
		return err
	}

	return nil
}
