package handlers

import (
	"html/template"
	"log"
	"main/internal/controllers"
)

func ParseAllTemplates() *template.Template {
	absPath, err := controllers.GetAbsPath("../static/templates/*.html")
	if err != nil {
		log.Fatal("Error getting absolute path for templates:", err)
	}

	templates, err := template.ParseGlob(absPath)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
	return templates
}
