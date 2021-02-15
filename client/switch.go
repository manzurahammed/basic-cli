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
	s := Switch{
		client:httpClient,
		backendurl:uri,
	}
	s.command = map[string] func() func(string) error {
		"create":s.create,
		"edit":s.edit,
		"delete":s.delete,
		"fetch":s.fetch,
		"health":s.health,
	}
	return s
}

func (s Switch) Switch() error{
	cmdName := os.Args[1]
	cmd, ok := s.command[cmdName]
	if !ok {
		fmt.ErrorF("Invalid command %s",cmdName)
	}
	return cmd()(cmdName)
}

func (s Switch) create() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) create() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) edit() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) delete() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) fetch() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}

func (s Switch) health() func(string) error{
	return func(cmd strinf) error{
		fmt.Println("Create Reminder")
		return nil
	}
}