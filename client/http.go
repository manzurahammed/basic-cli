package client

import "net/http"

type HTTPClient struct {
	client *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		client: *http.client{},
		BackendURI:uri,
	}
}