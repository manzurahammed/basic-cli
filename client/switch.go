package client

type BackandHTTPClient interface {

}

type Switch struct {
	client BackandHTTPClient
	backendurl string
	comand map[string] func() func(string) error
}

func NewSwitch(uri string){
	httpClient := NewHTTPClient(uri);
	return Switch{}
}