# GOTTH Blog Bootstrap

This is a Go, Templ, Tailwind, and HTMX (GOTTH) stack bootstrap for your blog migration.

## Structure
- `main.go`: Local development server (using `chi` router).
- `api/index.go`: Vercel Serverless Function entry point.
- `templates/`: `.templ` files for type-safe HTML.
- `static/`: Static assets (Tailwind CSS, HTMX, images).
- `vercel.json`: Configuration for Vercel deployment.

## Development
1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
2. **Generate Templates**:
   ```bash
   ./bin/templ generate
   ```
3. **Build Tailwind**:
   ```bash
   npx @tailwindcss/cli -i static/css/styles.css -o static/css/output.css
   ```
4. **Run Locally**:
   ```bash
   go run main.go
   ```

## Vercel Deployment
This project is configured to work with the [Vercel Go Runtime](https://vercel.com/docs/functions/serverless-functions/languages/go).
The `vercel.json` file routes all requests to `api/index.go`.
