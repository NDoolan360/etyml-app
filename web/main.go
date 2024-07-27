package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/a-h/templ"

	"github.com/NDoolan360/etyml-app/web/templates"
)

var handlers = map[string]templ.Component{
	"index.html": templates.BaseLayout(
		templates.Index(),
	),
}

func main() {
	for file, templComponent := range handlers {
		buf := bytes.NewBufferString("")
		renderErr := templComponent.Render(context.Background(), buf)
		if renderErr != nil {
			log.Fatal(renderErr)
			return
		}

		writeErr := os.WriteFile(file, buf.Bytes(), 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
		} else {
			log.Printf("Generated %s", file)
		}
	}
}
