package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"net/http"
	"time"
	"fmt"
)

// reminderBody represents the HTTP client which communicates with reminders backend API
type reminderBody struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	Duration time.Duration `json:"duration"`
}

// HTTPClient represents the HTTP client which communicates with reminders backend API
type HTTPClient struct {
	client *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		client: &http.Client{},
		BackendURI:uri,
	}
}

func (h HTTPClient) Create(title, message string, duration time.Duration) ([]byte,error){
	requestBody := reminderBody{
		Title: title,
		Message: message,
		Duration: duration,
	}
	return h.apiCall(
		http.MethodPost,
		"/reminders",
		&requestBody,
		http.StatusCreated,
	)
}

func (h HTTPClient) Edit(id, title, message string, duration time.Duration) ([]byte,error){
	requestBody := reminderBody{
		ID:id,
		Title: title,
		Message: message,
		Duration: duration,
	}
	return h.apiCall(
		http.MethodPatch,
		"/reminders"+id,
		&requestBody,
		http.StatusOK,
	)
}

func (h HTTPClient) Fetch(ids []string) ([]byte,error){
	idsSet := strings.Join(ids,",")
	return h.apiCall(
		http.MethodGet,
		"/reminders"+idsSet,
		nil,
		http.StatusOK,
	)
}

func (h HTTPClient) Delete(ids []string) ( error){
	idsSet := strings.Join(ids,",")
	_,err := h.apiCall(
		http.MethodDelete,
		"/reminders"+idsSet,
		nil,
		http.StatusNoContent,
	)
	return err
}

func (h HTTPClient) Health(host string) ( bool){
	res, err := http.Get(host+"/health")
	if err  !=nil || res.StatusCode != http.StatusOK{
		return false
	}
	return true
}

func (h HTTPClient) apiCall(method, path string, requestBody interface{},resCode int) ([]byte,error){
	bs, err := json.Marshal(requestBody)
	if err!=nil {
		e := wrapeError("could not marshal request body", err)
		return nil, e
	}

	req, err := http.NewRequest(method, h.BackendURI+path,bytes.NewReader(bs))

	if err != nil {
		e := wrapeError("could not create request", err)
		return []byte{}, e
	}

	res, err := h.client.Do(req)

	if err != nil {
		e := wrapeError("could not make api request", err)
		return []byte{}, e
	}

	resBody, err := h.readResBody(res.Body)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode !=resCode {
		if len(resBody) > 0 {
			fmt.Printf("got this response body:\n%s\n", resBody)
		}
		return []byte{}, fmt.Errorf(
			"expected response code: %d, got: %d",
			resCode,
			res.StatusCode,
		)
	}

	return []byte(resBody), nil
}

func (h HTTPClient) readResBody(b io.Reader) (string, error){
	res, err := ioutil.ReadAll(b)
	if err != nil {
		return "", wrapeError("could not read response body", err)
	}

	if len(res) == 0 {
		return "", nil
	}

	var buff bytes.Buffer
	
	if  err := json.Indent(&buff,res,"","\t"); err !=nil{
		return "", wrapeError("could not indent json", err)
	}

	return buff.String(),nil
}
