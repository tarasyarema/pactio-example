package main

import "fmt"

// Contract is the type that the consumer expect to receive
type Contract struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c Contract) String() string {
	return fmt.Sprintf(
		"Contract { Name: ''%s', Email: '%s' }",
		c.Name,
		c.Email,
	)
}

func main() {
	fmt.Println(Contract{
		Name:  "2pac Shakur",
		Email: "tupac@gvng.com",
	})
}
