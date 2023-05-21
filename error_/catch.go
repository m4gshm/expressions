package error_

type Catcher struct {
	Err error
}

func (c *Catcher) RunCatch(routine func() error) {
	if c.Err == nil {
		c.Err = routine()
	}
}

func (c *Catcher) Run(routine func()) {
	if c.Err == nil {
		routine()
	}
}

func (c *Catcher) In(other *Catcher) {
	if c.Err == nil && other != nil {
		c.Err = other.Err
	}
}
