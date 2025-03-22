package config

type Config struct {
	App        App
	Server     Server
	DB         DB
	ExtService ExtService
}

type App struct {
	Name    string
	Debug   bool
	Env     string
	Version string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Path     string
}

type ExtService struct {
	BaseUrl string
}
