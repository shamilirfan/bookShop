package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub      int    `json:"sub"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{Alg: "HS256", Typ: "JWT"} // instance/object of Header class/struct

	headerByteArr, err := json.Marshal(header) // header
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	headerBase64 := base_64(headerByteArr) // call base_64 function to to convert base64

	dataByteArr, err := json.Marshal(data) // payload
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	payloadBase64 := base_64(dataByteArr) // call base_64 function to to convert base64
	message := headerBase64 + "." + payloadBase64

	secretByteArr := []byte(secret)   // secret key convert to byte
	messageByteArr := []byte(message) // message(header, payload) key convert to byte

	h := hmac.New(sha256.New, secretByteArr)
	h.Write(messageByteArr)
	signature := h.Sum(nil)               // signature means hash
	signatureBase64 := base_64(signature) // call base_64 function to to convert base64

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64

	return jwt, nil
}

func base_64(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
