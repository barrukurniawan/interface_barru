package main

import (
	"fmt"
	service "interfacebarru/interfacebarru"
)

func main() {
	// declare
	var db []*service.User
	userServiceInstance := service.NewUserService(db)
	fmt.Println(userServiceInstance)

	// register
	response := userServiceInstance.Register(&service.User{Nama: "Barru Kurniawan"})
	fmt.Println(response)
	response2 := userServiceInstance.Register(&service.User{Nama: "budi"})
	fmt.Println(response2)

	// get user
	responseGetUser := userServiceInstance.GetUser()
	fmt.Println("---- RESULT ----")
	for _, v := range responseGetUser {
		fmt.Println(v.Nama)
	}
}
