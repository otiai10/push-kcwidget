package service

import "fmt"

type ErrorClient struct {
	message string
}

func (c *ErrorClient) Send() error {
	return fmt.Errorf(c.message)
}
