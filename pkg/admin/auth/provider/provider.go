package provider

type Provider interface {
	Name() string
	SetCredentials(credentials Config) error
	Authenticate(name, pass string) (Principal, error)
}

type Config struct {
	Type       string `yaml:"type"`
	JwtSecret  string `yaml:"jwtSecret"`
	RawDetails any    `yaml:"details"`
}

type Principal struct {
	Name string
	Role string
}
