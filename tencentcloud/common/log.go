package common

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

func (c *Client) WithLogger(logger Logger) *Client {
	c.logger = logger
	return c
}
