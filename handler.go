package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// secret key
var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
}

//we can thepass the data as username and password from the apis and use it over here
type Credentials struct {
	Username string `json: "username"`
	Password string `json:"password" `
}

// claims struct; claims will be use to create a payload for JWT.
//inside the payload we will pass our username and exp

type Claims struct {
	Username string `json:"username"`
	//jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

}
func Home(w http.ResponseWriter, r *http.Request) {

}
func Refresh(w http.ResponseWriter, r *http.Request) {

}

