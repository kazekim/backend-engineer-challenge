package besqlx

//Get query normal select query and return first value found in 'dest' struct
func (c *defaultClient) Get(dest interface{}, query string, args ...interface{}) error {
	return c.db.Get(dest, query, args...)
}
