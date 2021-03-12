package main

import "fmt"

// Contract is the type that the consumer expect to receive
type Contract struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

// String method implementation for a Contract
func (c Contract) String() string {
	return fmt.Sprintf(
		"Contract { Name: '%s', Email: '%s' }",
		*c.Name,
		*c.Email,
	)
}

// stringToPtr does just that
func stringToPtr(s string) *string {
	return &s
}

// Handler is the function that we want to test
func Handler() (*Contract, error) {
	return &Contract{
		Name:  stringToPtr("2pac Shakur"),
		Email: stringToPtr("tupac@shakur.com"),
	}, nil
}

func main() {
	c, _ := Handler()
	fmt.Println(c)
}
