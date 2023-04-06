package main

import (
	"first-project/internal/app"
)

const (
	configDir = "configs"
	envFile   = ".env"
)

func main() {
	app.Run(envFile, configDir)
}

//
//type Store struct {
//	Users []User
//}
//
//type User struct {
//	name string
//}
//
//func main() {
//	users := Store{
//		Users: []User{
//			{name: "1"},
//			{name: "2"},
//			{name: "3"},
//			{name: "4"},
//		},
//	}
//
//	users.Delete(2)
//	get := users.Get()
//	fmt.Println(get)
//}
//
//func (s *Store) Delete(index int) {
//	users := helpers.SliceRemove[User](s.Users, index)
//	s.Users = users
//
//}
//
//func (s *Store) Get() []User {
//	return s.Users
//}
