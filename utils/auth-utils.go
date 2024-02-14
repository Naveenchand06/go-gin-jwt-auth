package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/Naveenchand06/go-gin-jwt-auth/models"
	"github.com/golang-jwt/jwt/v5"
)


func GetSignedJWTToken(user *models.User) (string, error) {
	 token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	wd, _ := os.Getwd()
	fmt.Println("Working Dir ---> ", wd)
	privateKeyPEM, err := os.ReadFile("ecdsa-private.pem")
	
    if err != nil {
        return "", err
    }
	 // Decode PEM file
	 block, _ := pem.Decode(privateKeyPEM)
	 if block == nil {
		 return "", err
	 }

	 // Parse ECDSA private key
	 privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	 if err != nil {
		 return "", err
	 }
	
	return token.SignedString(privateKey)
}
