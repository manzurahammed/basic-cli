package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type idsFlag []string

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("The name of the person %s and born %s",p.name, p.born.String())
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func (ids idsFlag) String() string {
	return strings.Join(ids,",")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids,id)
	return nil
}

func main() {
	var ids idsFlag
	var p person
	flag.Var(&ids,"id","The list of id")
	flag.Var(&p,"name","The list of id")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
}