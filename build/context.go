package build

var (
	Environment string = "unknown"
	Service     string = "unknown"
	Version     string = "v0.0.0-unknown"
)

// Context type.
type Context struct {
	Environment string
	Service     string
	Version     string
}

// NewContext constructor.
func NewContext() *Context {
	return &Context{
		Environment: Environment,
		Service:     Service,
		Version:     Version,
	}
}
