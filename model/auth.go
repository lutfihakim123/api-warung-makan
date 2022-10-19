package model

import "github.com/golang-jwt/jwt"

var AplicationName = "ENIGMACAMP"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("root")

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type Credential struct {
	Id       string `json:"id"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Gaji     int    `json:"gaji"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
