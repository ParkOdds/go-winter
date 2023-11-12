package main

import (
	"fmt"
	"time"
)

func GetMember() {
	fmt.Println("please wait...")
	time.Sleep(3 * time.Second)
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

func logEnd() {
	time.Now()

	fmt.Println("completed program")
	fmt.Println(time.Now())

}

func checkServerResponse() {
	fmt.Println("check server time")
	time.Sleep(3 * time.Second)
	panic("server error")
}

func ma() {

	// defer func() {
	// 	fmt.Println("anonymous funtion")
	// }()

	// defer logEnd()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover ")
			fmt.Println(r)

		}
	}()

	GetMember()

	checkServerResponse()

}
