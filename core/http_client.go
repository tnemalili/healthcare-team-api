package core

import (
	"billing.api/core/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

type HTTP struct{ HTTPClient *http.Client }

func (receiver *HTTP) REQUEST(url string, method string, headers map[string]interface{}, body interface{}) (*http.Response, error) {

	log.Infof("[core.HTTPClient.REQUEST] TARGET MICROSERVICE URL: %v", url)
	bodyBytes, err := json.Marshal(body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Error("[core.HTTPClient.REQUEST] Error creating http.NewRequest: ", err)
	}
	// NOW WE CAN SET HEADERS
	request.Header.Add("Content-Type", "application/json")
	SETHeaders(request, headers)
	response, err := receiver.HTTPClient.Do(request)
	if err != nil {
		log.Error("[core.HTTPClient.REQUEST]: Error occur on HTTClient.Do()", err)
		return nil, err
	}

	log.Infof("[core.HTTPClient.REQUEST] TARGET MICROSERVICE Response Status: %v", response.Status)

	return response, nil
}

func (receiver *HTTP) FetchDocumentHandler(collection string, id string) (*http.Response, error) {

	url := fmt.Sprintf("%v/%v/%v", StorageUrl, collection, id)
	headers := make(map[string]interface{})
	headers["x-api-key"] = StorageApiKey
	resp, err := receiver.REQUEST(url, "GET", headers, nil)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not Found")
	}

	return resp, err
}

func SETHeaders(request *http.Request, headers map[string]interface{}) *http.Request {

	for key, value := range headers {
		request.Header.Add(key, fmt.Sprintf("%s", value))
	}
	return request
}

func Form(data map[string]interface{}) *url.Values {

	form := url.Values{}
	for key, value := range data {
		form.Add(key, fmt.Sprintf("%s", value))
	}
	return &form
}

func HTTPResponseHandler(response *http.Response) map[string]interface{} {

	var msErrorStatuses = []int{
		http.StatusNotFound,
		http.StatusServiceUnavailable,
		http.StatusBadGateway,
		http.StatusInternalServerError,
		http.StatusBadRequest,
	}

	var r = make(map[string]interface{})

	if contains(msErrorStatuses, response.StatusCode) {
		r["microservice_status"] = response.StatusCode
		r["microservice_error"] = response.Status
		r["url"] = response.Request.URL.String()
		r["message"] = "NB: This is a Microservice Services Related Error!"
		return r
	}
	// READ RESPONSE BODY.
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Errorf("[HTTPResponseHandler] Body ReadAll Error, %v", err.Error())
	}
	if err = json.Unmarshal(body, &r); err != nil {
		log.Errorf("Body parse error, %v", err)
	}

	return r
}

func HTTPErrorHandler(ctx echo.Context, err error, statusCode int) error {
	log.Errorf("[core.HTTPErrorHandler] %v", err)
	return ctx.JSON(statusCode, &models.GenericMessage{Message: err.Error()})
}

func contains(statuses []int, status int) bool {

	for _, s := range statuses {
		if status == s {
			return true
		}
	}
	return false
}

var HTTPClient = &HTTP{HTTPClient: &http.Client{}}
