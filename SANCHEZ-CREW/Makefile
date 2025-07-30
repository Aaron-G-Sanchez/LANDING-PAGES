# PATHS
SASS_SOURCE := internal/sass/main.scss
SASS_OUTPUT := static/css/main.css


.PHONY: sass
sass: 
	sass $(SASS_SOURCE) $(SASS_OUTPUT)

.PHONY: live/templ
live/templ: 
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false

.PHONY: live/server
live/server:
	air \
	--build.cmd "go build -o tmp/bin/main" \
	--build.bin "tmp/bin/main" \
	--build.delay "0" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true 

.PHONY: live/sass
live/sass: 
	sass --watch $(SASS_SOURCE):$(SASS_OUTPUT)

.PHONY: live/sync_assets
live/sync_assets: 
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "static" \
	--build.include_ext "css"

.PHONY: dev
dev:
	make -j4 live/templ live/server live/sass live/sync_assets