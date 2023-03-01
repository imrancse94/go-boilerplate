package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"go-boilerplate/Helper"
	"go-boilerplate/constant"
	"go-boilerplate/localize"
	"go-boilerplate/logger"
	"go-boilerplate/models"
	"go-boilerplate/requests"
	"go-boilerplate/response"
	Services "go-boilerplate/services"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	/*req := mux.Vars(r)
	fmt.Println("Test Mux", req["id"], req["name"])*/
	request := requests.LoginRequest{}
	Helper.Request(r, &request)
	userData, message := Services.Login(request)
	logData := logger.LogData{}
	logData.Action = "LOGIN"
	localize.SetLocale("en")

	statusCode := constant.Status("FAILED")
	if userData != "" {
		statusCode = constant.Status("SUCCESS")
	}
	res := response.Response{
		StatusCode: statusCode,
		Message:    localize.Trans(message, ""),
		Data:       userData,
	}

	logData.Data = res
	logger.CreateLog(logData)
	response.SuccessRespond(res, w)

	return
}

func AuthData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("auth_id"))
	user := models.GetUserById(id)
	//fmt.Println("Token", id)
	var data = map[string]interface{}{
		"access_token": Helper.ExtractToken(r.Header.Get("Authorization")),
		"user":         user,
		"permissions":  Services.GetPermissionByUserId(id),
	}
	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("Auth data fetched successfully", ""),
		Data:       data,
	}
	response.SuccessRespond(res, w)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.Header.Get("auth_id"))
	user := models.GetUserById(id)

	data, _ := Services.CreateTokenDataByUser(user)

	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("Token Refreshed successfully", ""),
		Data:       data,
	}
	response.SuccessRespond(res, w)
}

func rsaGeneration() (string, string) {
	// Generate a new RSA key pair with a bit size of 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Extract the public key from the private key
	publicKey := privateKey.PublicKey

	// Encode the private key into PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// Encode the public key into PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	/*fmt.Println("Public Key", string(publicKeyPEM))
	fmt.Println("Private Key", string(privateKeyPEM))
	fmt.Println("RSA key pair generated successfully")*/

	return string(publicKeyPEM), string(privateKeyPEM)
}

func rsaEncryptionAndDecryption() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Convert private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Encode PEM block to base64 string
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEM.Bytes)

	// Print private key in base64 format
	fmt.Println("Private key:")
	fmt.Println(privateKeyBase64)

	// Encrypt plaintext with RSA public key
	plaintext := "Hello, world!"
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, []byte(plaintext))
	if err != nil {
		panic(err)
	}

	// Encode ciphertext to base64 string
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	// Print ciphertext in base64 format
	fmt.Println("Ciphertext:")
	fmt.Println(ciphertextBase64)

	// Decrypt ciphertext with RSA private key
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}

	// Print decrypted plaintext
	fmt.Println("Decrypted plaintext:")
	fmt.Println(string(decrypted))
}

func Test(w http.ResponseWriter, r *http.Request) {

	// Create the keys
	//priv, pub := Helper.GenerateRsaKeyPair()

	// Export the keys to pem string
	//priv_pem := Helper.ExportRsaPrivateKeyAsPemStr(priv)
	//pub_pem, _ := Helper.ExportRsaPublicKeyAsPemStr(pub)

	// Import the keys from pem string
	//priv_parsed, _ := Helper.ParseRsaPrivateKeyFromPemStr(priv_pem)
	//pub_parsed, _ := Helper.ParseRsaPublicKeyFromPemStr(pub_pem)

	// Export the newly imported keys
	//priv_parsed_pem := Helper.ExportRsaPrivateKeyAsPemStr(priv_parsed)
	//pub_parsed_pem, _ := Helper.ExportRsaPublicKeyAsPemStr(pub_parsed)
	/*pub_parsed_pem, priv_parsed_pem, err := Helper.GenerateKey()
	fmt.Println(priv_parsed_pem)
	fmt.Println(pub_parsed_pem)

	// Check that the exported/imported keys match the original keys
	if err != nil {
		fmt.Println(err)
	}*/

	publicKey, privateKey := rsaGeneration()
	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("Key generated successfully", ""),
		Data:       map[string]string{"privateKey": privateKey, "publicKey": publicKey},
	}

	response.SuccessRespond(res, w)
}

func TestEncryptDecrypt(w http.ResponseWriter, r *http.Request) {
	request := requests.TestRequest{}
	Helper.Request(r, &request)
	data := ""
	if request.Type == "encrypt" {
		data, _ = Helper.Encrypt(request.Data, request.Key)
	} else {
		data, _ = Helper.Decrypt(request.Data, request.Key)
	}

	res := response.Response{
		StatusCode: constant.Status("SUCCESS"),
		Message:    localize.Trans("success", ""),
		Data:       data,
	}

	response.SuccessRespond(res, w)
}
