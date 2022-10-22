package begin

type Gin interface {

	Start(port string) error

	Group(relativePath string)  RouterGroup
	RegisterAPIVersionGroup() RouterGroup
	APIVersionGroup() RouterGroup
}