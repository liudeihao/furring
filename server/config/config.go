package config

var Instance *Config

type Config struct {
    JWT    jwtConfig    `yaml:"jwt"`
    DB     dbConfig     `yaml:"db"`
    Server serverConfig `yaml:"server"`
}

type serverConfig struct {
    Addr string `yaml:"addr"`
    Port string `yaml:"port"`
}

type jwtConfig struct {
    SecretKey string `yaml:"key"`
    Issuer    string `yaml:"issuer"`
}

type dbConfig struct {
    Driver string `yaml:"driver"`
    DSN    string `yaml:"dsn"`
}
