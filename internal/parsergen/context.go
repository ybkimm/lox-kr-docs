package parsergen

type context struct {
	errs Errors
}

func newContext() *context {
	return new(context)
}

func (c *context) Fail(err error) {
	c.errs = append(c.errs, err)
}

func (c *context) Err() error {
	if len(c.errs) == 0 {
		return nil
	}
	return c.errs
}
