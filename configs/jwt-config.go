package configs

// JwtConfig represents jwt configurations.
type JwtConfig struct {
	Secret []byte
}

// GetJwtConfig provides jwt configurations.
func GetJwtConfig() JwtConfig {
	return JwtConfig{Secret: []byte("bkB6kgc5lKRcNuU-gZzlFXNx53gWDHvqMHPQWINX-Kw")}
}
