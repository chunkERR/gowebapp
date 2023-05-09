package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache bool
	MyCache  map[string]*template.Template
	InfoLog  *log.Logger
}
