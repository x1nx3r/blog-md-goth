.PHONY: templ tailwind run build clean dev bundle

# Generate Go code from .templ files
templ:
	./bin/templ generate

# Build Tailwind CSS
tailwind:
	npx @tailwindcss/cli -i internal/assets/static/css/styles.css -o internal/assets/static/css/tailwind.css --minify

# Create the final CSS bundle
bundle: tailwind
	@echo "Bundling CSS and Fonts..."
	cat internal/assets/static/css/tailwind.css \
	    internal/assets/static/css/newspaper.css \
	    internal/assets/static/fonts/google-fonts.css > internal/assets/static/css/bundle.css

# Run the development server
run: templ bundle
	go run main.go

# Build production binary
build: templ bundle
	go build -o bin/server main.go

# Watch mode: regenerate templ on changes
dev:
	./bin/templ generate --watch &
	go run main.go

# Clean generated files
clean:
	rm -f bin/server
	find templates -name '*_templ.go' -delete
