package config

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := Write(c)
	if err != nil {
		return err
	}
	return nil
}
