package config

type Configuration struct {
	HTTP  HTTP  `toml:"http"`
	Log   Log   `toml:"log"`
	Mongo Mongo `toml:"mongo"`
	Redis Redis `toml:"redis"`
}

type HTTP struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Log struct {
	Type     string `toml:"type"`
	Level    string `toml:"level"`
	Path     string `toml:"path"`
	Filename string `toml:"filename"`
	Maxage   int    `toml:"maxage"`
	Rotation int    `toml:"rotation"`
}

type Mongo struct {
	Host    string `toml:"host"`
	Port    string `toml:"port"`
	Pass    string `toml:"pass"`
	Maxpool uint64 `toml:"maxpool"`
}

type Redis struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	Auth string `toml:"auth"`
}
