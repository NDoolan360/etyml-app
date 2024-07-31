MODULES = ./web ./web/templates ./api/health ./api/puzzle
MAKE_HTML_MODULE = ./web
MAKE_HTML_BIN = make_html
STATIC_DIR = web/static
DEV_PORTS = 3999 8888

.PHONY: generate install test dev clean

go.work:
	go work init
	go work use $(MODULES)

generate: go.work
ifneq (, $(shell which templ))
	templ generate -path ./web/templates
else
	go run github.com/a-h/templ/cmd/templ@latest generate -path ./web/templates
endif

install: go.work generate
	go install $(MODULES)

dist: install
	rm -rf $@; cp -R $(STATIC_DIR)/. $@
	mkdir $@/scripts
ifneq (, $(shell which minify))
	minify -o $@/styles.css -b web/styles/*.css
	minify -o $@/scripts/ -a web/scripts/*.js
else
	go run github.com/tdewolff/minify/v2/cmd/minify@latest -o $@/styles.css -b web/styles/*.css
	go run github.com/tdewolff/minify/v2/cmd/minify@latest -o $@/scripts/ -a web/scripts/*.js
endif
	go build -o $(MAKE_HTML_BIN) $(MAKE_HTML_MODULE)
	cd $@; ../$(MAKE_HTML_BIN);


test: install
	go test $(MODULES)

dev: clean dist
	parallel -tmux ::: 'watchexec -e go,templ,css,js "make dist"' 'netlify dev'

clean:
	$(foreach port, $(DEV_PORTS), lsof -i:$(port) -t | xargs kill;)
	git clean -fdX
