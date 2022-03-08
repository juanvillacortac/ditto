package proto

// ParseConfig is a config for parser.
type ParseConfig struct {
	Debug      bool
	Permissive bool
}

// Option is an option for ParseConfig.
type Option func(*ParseConfig)

// WithDebug is an option to enable the debug mode.
func WithDebug(debug bool) Option {
	return func(c *ParseConfig) {
		c.Debug = debug
	}
}

// WithPermissive is an option to allow the permissive parsing rather than the just documented spec.
func WithPermissive(permissive bool) Option {
	return func(c *ParseConfig) {
		c.Permissive = permissive
	}
}
