package config

type Jwt struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}
