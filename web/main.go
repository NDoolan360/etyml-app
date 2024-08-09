package main

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"

	"github.com/NDoolan360/etyml-app/web/templates"
)

var handlers = map[string]templ.Component{
	"index.html": templates.BaseLayout(templates.Index(), nil),
}

func main() {
	for file, templComponent := range handlers {
		f, fileErr := os.Create(file)
		if fileErr != nil {
			log.Fatalf("failed to create output file: %v", fileErr)
		}

		renderErr := templComponent.Render(context.Background(), f)
		if renderErr != nil {
			log.Fatalf("failed to write output file: %v", renderErr)
		}
	}
}
