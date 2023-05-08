package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	MyCache map[string]*template.Template
}
