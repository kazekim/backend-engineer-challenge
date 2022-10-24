package grcrouters

func (r *defaultRouter) Start() error {
	err := r.Gin().Start(r.options.Environment.ServerPort)
	return err
}
