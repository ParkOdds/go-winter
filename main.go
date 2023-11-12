package main

import (
	"fmt"
)

func GetMember() {
	fmt.Println("please wait...")
}
func IsInSystem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {
	return 201, "manager"
}

func getDeparture(username string, departure *string) {
	if username != "" {
		*departure = "home"
	}
}

func CheckLogin(username string, password string) {
	if IsInSystem(username) {
		fmt.Println("found user in system")
		GetUserDetail(username)
		departure := ""
		getDeparture(username, &departure)
	}
}

func main() {
	GetMember()

}
