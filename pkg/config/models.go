package config

type Config struct {
	Dark bool
}

type Connection struct {
	Name string
	Host string
	Port int
	User string
	Pass string
  Database string
}

type History struct {
	Command string
	Date    string
}
