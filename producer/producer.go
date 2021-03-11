package main

import "fmt"

type Contract struct {
	Name string `json:"name"`
}

func (c Contract) String() string {
	return fmt.Sprintf("Contract { Name: \"%s\" }", c.Name)
}

func main() {
	fmt.Println(Contract{Name: "test"})
}
