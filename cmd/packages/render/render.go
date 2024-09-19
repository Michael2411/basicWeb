package render

import (
	models "GO-WEB/cmd/packages/Models"
	"GO-WEB/cmd/packages/config"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the templates packages. it gets the input from main()
func NewTemplates(a *config.AppConfig) {
	app = a
}

// function to add default data to pages while renedering

func AddDefaultData(data *models.TemplateData) *models.TemplateData {
	return data
}
func RenderTemp(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template
	var err error

	// if I don't want to use the cached html pages

	if !app.UseCache {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("failed to create template cache")
		}
	} else {
		// fetch template cache created in main from AppConfig which was passed using NewTemplates()
		templateCache = app.TemplateCache
	}

	//get requested template from cache
	template, exists := templateCache[tmpl]
	if !exists {
		log.Fatal("Template Doesn't exist")
	}
	// buffer is used to be able to execute and get the error without writing directly to response writer
	// it gives better handling of this template
	buffer := new(bytes.Buffer)

	data = AddDefaultData(data)

	err = template.Execute(buffer, data)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// called in main to create cached renedred templates only once
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache:=make(map[string]*template.Template)
	//--- you can also make map like this
	myCache := map[string]*template.Template{}
	//get all the files ending in page.tmpl from the Templates folder
	pages, err := filepath.Glob("./Templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//check if any layouts exists in the directory
	layouts, err := filepath.Glob("./Templates/*.layout.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through these templates
	for _, page := range pages {
		// base return the last element of the page just to get the template name
		templateName := filepath.Base(page)

		//parse the template path in page and store it in a template named according to base path using New
		parsedTemplate, err := template.New(templateName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			/* ParseGlob matches the layouts with the templates according to teh definitions in the templates themsevles what why we
			we parse the templates first */
			parsedTemplate, err = parsedTemplate.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[templateName] = parsedTemplate
	}
	return myCache, nil
}

////// This is simple template caching Code.
// var templateCache = make(map[string]*template.Template)

// func RenderTempTest(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
// 	_, inMap := templateCache[t]
// 	if !inMap {
// 		log.Println("Creating Template")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("Using Cached Template")
// 	}
// 	tmpl = templateCache[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./Templates/%s", t),
// 		"./Templates/base.layout.tmpl",
// 	}

// 	//parse template
// 	parsedTemplate, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	templateCache[t] = parsedTemplate
// 	return nil
// }
