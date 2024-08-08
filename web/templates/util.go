package templates

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
)

func guesses(ctx context.Context) []string {
	return ctx.Value("guesses").([]string)
}

func hints(ctx context.Context) []string {
	return ctx.Value("hints").([]string)
}

func assembleHintLink(id string, guesses []string, hints []string) templ.SafeURL {
	out := "?"

	for _, guess := range guesses {
		out += fmt.Sprintf("guess=%s&", guess)
	}

	for _, hint := range hints {
		out += fmt.Sprintf("hint=%s&", hint)
	}

	out += fmt.Sprintf("hint=%s", id)

	return templ.SafeURL(out)
}
