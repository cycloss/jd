package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type templateArgs struct {
	Title string
	Date  string
	// tags range works because if index is not 0, a comma and space are inserted before the next item
	Tags strSliceFlag
}

func writeTemplateToFile(pf *os.File, ta *templateArgs) error {
	t := loadTemplate()
	err := t.Execute(pf, ta)
	return err
}

func loadTemplate() *template.Template {
	templateFullPath := fmt.Sprintf("%s/%s", JD_CONF_PATH, JD_TEMPLATE_NAME)
	t, err := template.ParseFiles(templateFullPath)

	if err != nil {
		log.Print(err)
		exitFatal("could not find or open template at %s", templateFullPath)
	}

	return t
}
