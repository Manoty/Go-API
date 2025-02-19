package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1":{
		Email: "george_bush@gmail.com",
		Id:    "user1",
		Name:  "George Bush",
		Token: "123",
		},

	"user2":{
		Email: "bill_clinton@gmail.com",
		Id:    "user2",
		Name:  "Bill Clinton",
		Token: "456",
		},	
}