.PHONY: templ tailwind run build clean dev

# Generate Go code from .templ files
templ:
	./bin/templ generate

# Build Tailwind CSS
tailwind:
	npx @tailwindcss/cli -i static/css/styles.css -o static/css/tailwind.css

# Run the development server
run: templ
	go run main.go

# Build production binary
build: templ
	go build -o bin/server main.go

# Watch mode: regenerate templ on changes
dev:
	./bin/templ generate --watch &
	go run main.go

# Clean generated files
clean:
	rm -f bin/server
	find templates -name '*_templ.go' -delete
