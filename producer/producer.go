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

func main() {
	name := "2pac Shakur"
	email := "tupac@gvng.com"

	fmt.Println(Contract{Name: &name, Email: &email})
}
