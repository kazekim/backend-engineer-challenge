package begin

import "fmt"

//Start begin gin server service
func (c *defaultGin) Start(port string) error {
	server := fmt.Sprintf("localhost:%v", port)
	return c.g.Run(server)
}
