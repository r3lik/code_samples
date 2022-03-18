package main

type Config struct {
	HTTPHost string
	HTTPPort int
}
// DefaultConfig returns a Config with sensible options
func DefaultConfig() Config {
	return Config{
		HTTPHost: "127.0.0.1",
		HTTPPort: 8080,
	}
}

func NewConfig() (*Config, error) {
	c := DefaultConfig()
	if v := os.Getenv("HTTP_PORT"); v != "" 
}
