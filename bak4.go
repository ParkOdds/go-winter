package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Address  string
	PostCode string
}

type UserProfile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int
	Height    float32
	Address   Address
	Bill      struct {
		BillAddress string
	}
}

func (u UserProfile) ToFullDesc() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func main2() {
	fmt.Println("xx")
	user := map[string]string{}
	user["username"] = "nattapol"
	user["password"] = "xxxxx"
	fmt.Println(user)
	fmt.Println(user["username"])

	userProfile := UserProfile{
		Firstname: "nattapol1",
		Lastname:  " xxxxxx",
		Age:       21,
	}
	fmt.Println(userProfile)
	userProfile.Address.PostCode = "10250"

	fmt.Println(userProfile.ToFullDesc())

	byteTxtJson, err := json.MarshalIndent(userProfile, "", " ")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(string(byteTxtJson))

	dataJson := `
	{
		"firstname": "nattapol2",
		"lastname": " xxxxxx",
		"Age": 21,
		"Height": 180,
		"Address": {
		 "Address": "",
		 "PostCode": "10330"
		}
	}`

	var nat2Profile UserProfile
	err = json.Unmarshal([]byte(dataJson), &nat2Profile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nat2Profile)

}
