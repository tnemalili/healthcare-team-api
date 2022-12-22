package core

import (
	"auth.api/core/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (receiver *HTTP) CREATEDareHandler(username string, otp string) (*http.Response, error) {

	// TODO: CHECK IF THE SELECTED BID HAS AN ACCOUNT
	encrypted, _ := EncryptAES(username)
	newDare := models.NewDare(encrypted, encrypted, otp)
	// INJECT TRANSACTION ID
	url := fmt.Sprintf("%v/%v/%v", StorageUrl, DARE, encrypted)
	headers := make(map[string]interface{})
	headers["x-api-key"] = StorageApiKey
	resp, err := receiver.REQUEST(url, "POST", headers, newDare)
	return resp, err
}

func (receiver *HTTP) CREATECredHandler(username string, password string) (*http.Response, error) {

	secret, err := HashIdentity(password)
	encrypted, _ := EncryptAES(username)
	token := GenerateCode()
	newCred := models.NewCred(encrypted, secret, token)
	// INJECT TRANSACTION ID
	url := fmt.Sprintf("%v/%v/%v", StorageUrl, CREDENTIALS, encrypted)
	headers := make(map[string]interface{})
	headers["x-api-key"] = StorageApiKey
	resp, err := receiver.REQUEST(url, "POST", headers, newCred)
	return resp, err
}

func (receiver *HTTP) DareHandler(username string) (*http.Response, error) {

	encrypted, _ := EncryptAES(username)
	resp, err := receiver.FetchDocumentHandler(DARE, encrypted)

	return resp, err
}

func (receiver *HTTP) LoginHandler(username string, password string) bool {

	encrypted, _ := EncryptAES(username)
	resp, err := receiver.FetchDocumentHandler(CREDENTIALS, encrypted)
	if err != nil {
		return false
	}

	verified := AuthHandler(encrypted, password, resp)
	log.Info("[LoginHandler] Success: ", verified)

	return verified
}

func (receiver *HTTP) UPDATECredentialHandler(id string, updates []interface{}) (*http.Response, error) {

	url := fmt.Sprintf("%v/%v/%v", StorageUrl, CREDENTIALS, id)
	headers := make(map[string]interface{})
	headers["x-api-key"] = StorageApiKey
	resp, err := receiver.REQUEST(url, "PUT", headers, updates)
	return resp, err
}

func AuthHandler(username string, password string, resp *http.Response) bool {

	data := HTTPResponseHandler(resp)

	return CheckIdentityHash(password, data["password"].(string)) && Verified(data["username"].(string), username)
}
