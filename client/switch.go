package client

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type listFlag []string

func (list *listFlag) String() string {
	return strings.Join(*list ,",")
}

func (list *listFlag) Set(v string) error {
	*list = append(*list, v)
	return nil
}

type BackandHTTPClient interface {
	Create(title, message string, duration time.Duration) ([]byte,error)
	Edit(id string, title, message string, duration time.Duration) ([]byte,error)
	Fetch(ids []string) ([]byte,error)
	Delete(ids []string) error
	Health(host string) bool
}

type Switch struct {
	client BackandHTTPClient
	backendurl string
	command map[string] func() func(string) error
}

func NewSwitch(uri string) Switch {
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
		fmt.Errorf("Invalid command %s",cmdName)
	}
	return cmd()(cmdName)
}

func (s Switch) Help(){
	var help string
	for name:= range s.command {
		help += name + "\t --help\n"
	}
	fmt.Println("Usage of #{os.Args[0]} <command> [<arg>]\n#{help}");
}

func (s Switch) create() func(string) error{
	return func(cmd string) error{
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		t, m, d := s.reminderFlags(createCmd)
		if err := s.checkArg(3);err!=nil {
			return err
		}
		if err := s.parseCmd(createCmd);err!=nil {
			return err
		}
		res, err := s.client.Create(*t, *m, *d)
		if err !=nil {
			return wrapeError("Could not create reminder",err)
		}
		fmt.Printf("Create reminder %s",string(res))
		return nil
	}
}




func (s Switch) edit() func(string) error{
	return func(cmd string) error{
		ids := listFlag{}
		eidtCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		eidtCmd.Var(&ids,"id","The id reminder of edit")
		t, m, d := s.reminderFlags(eidtCmd)
		if err := s.checkArg(2);err!=nil {
			return err
		}
		if err := s.parseCmd(eidtCmd);err!=nil {
			return err
		}
		lastId := ids[len(ids)-1]
		res, err := s.client.Edit(lastId ,*t, *m, *d)
		if err !=nil {
			return wrapeError("Could not Edit reminder",err)
		}
		fmt.Printf("Edit reminder SuccessFully %s",string(res))
		return nil
	}
}

func (s Switch) delete() func(string) error{
	return func(cmd string) error{
		ids := listFlag{}
		deleteCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		deleteCmd.Var(&ids,"id","The id reminder of edit")
	
		if err := s.checkArg(1);err!=nil {
			return err
		}
		if err := s.parseCmd(deleteCmd);err!=nil {
			return err
		}
	
		err := s.client.Delete(ids)
		if err !=nil {
			return wrapeError("Could not Delete reminder",err)
		}
		fmt.Printf("Delete reminder SuccessFully %v",ids)
		return nil
	}
}

func (s Switch) fetch() func(string) error{
	return func(cmd string) error{
		ids := listFlag{}
		fetchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		fetchCmd.Var(&ids,"id","The id reminder of Fetch")
	
		if err := s.checkArg(1);err!=nil {
			return err
		}
		if err := s.parseCmd(fetchCmd);err!=nil {
			return err
		}
	
		res ,err := s.client.Fetch(ids)
		if err !=nil {
			return wrapeError("Could not Fetch reminder",err)
		}
		fmt.Printf("Fetch reminder SuccessFully %s",string(res))
		return nil
	}
}

func (s Switch) health() func(string) error{
	return func(cmd string) error{
		var host string
		heakthCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		heakthCmd.StringVar(&host,"host",s.backendurl,"The id reminder of Health")
		if err := s.parseCmd(heakthCmd);err!=nil {
			return err
		}

		if !s.client.Health(host){
			fmt.Printf("Host %s is down",host)
		}else{
			fmt.Printf("Host %s is up",host)
		}
		return nil
	}
}

func (s Switch) reminderFlags(f *flag.FlagSet)(*string,*string,*time.Duration){
	t,m,d := "", "", time.Duration(0)
	f.StringVar(&t,"title","","Reminder flag")
	f.StringVar(&t,"t","","Reminder flag")
	f.StringVar(&m,"message","","Reminder flag")
	f.StringVar(&m,"m","","Reminder flag")
	f.DurationVar(&d,"duration",0,"Reminder flag")
	f.DurationVar(&d,"d",0,"Reminder flag")
	return &t, &m, &d
}

func (s Switch) checkArg(minFlag int) error {
	if len(os.Args)==3 && os.Args[2]=="--help"{
		return nil
	}

	if len(os.Args)-2<minFlag {
		fmt.Printf("incorrect use of %s\n%s %s --help\n",os.Args[1], os.Args[0],os.Args[1])
		return fmt.Errorf("%s except at least %d arg(s), %d provided", os.Args[1], minFlag,len(os.Args)-2)
	}
	return nil
}

func (s Switch) parseCmd(f *flag.FlagSet) error {
	err := f.Parse(os.Args[2:])

	if err !=nil {
		return wrapeError("could not parse error '"+f.Name()+"' comand flag",err)
	}
	return nil
}