# PATHS
SASS_SOURCE := internal/sass/main.scss
SASS_OUTPUT := static/css/main.css


.PHONY: sass
sass: 
	sass $(SASS_SOURCE) $(SASS_OUTPUT)

.PHONY: live/sass
live/sass: 
	sass --watch $(SASS_SOURCE):$(SASS_OUTPUT)

.PHONY: templ-build
templ-build:
	templ generate

.PHONY: live/templ
live/templ: 
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false