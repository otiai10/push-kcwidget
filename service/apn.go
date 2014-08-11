package service

type ApnClient struct {
	// implements PushClient
	set PushSet
	err error
}

func (c *ApnClient) Send() error {
	return c.err
}
