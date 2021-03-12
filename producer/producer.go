package main

import "fmt"

// Contract is the type that the consumer expect to receive
type Contract struct {
	Name    *string  `json:"name"`
	Email   *string  `json:"email"`
	Address *Address `json:"address"`
}

type Address struct {
	Street *string `json:"street"`
	Number *int    `json:"number"`
}

// String method implementation for a Contract
func (c Contract) String() string {
	return fmt.Sprintf(
		"Contract { Name: '%s', Email: '%s', Address: '%s at %d' }",
		*c.Name,
		*c.Email,
		*c.Address.Street,
		*c.Address.Number,
	)
}

// stringToPtr does just that
func stringToPtr(s string) *string {
	return &s
}

// intToPtr does just that
func intToPtr(s int) *int {
	return &s
}

// Handler is the function that we want to test
func Handler() (*Contract, error) {
	return &Contract{
		Name:  stringToPtr("2pac Shakur"),
		Email: stringToPtr("tupac@shakur.com"),
		Address: &Address{
			Street: stringToPtr("Trafalgar"),
			Number: intToPtr(1337),
		},
	}, nil
}

func main() {
	c, _ := Handler()
	fmt.Println(c)
}
