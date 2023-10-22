package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}


func (a authenticationInfo) authInfo() string {
	return "Autoriation: Basic" + a.username + ":" + a.password

}

func test(a authenticationInfo) {
	fmt.Println(a.())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}