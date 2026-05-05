package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"blog-gotth/internal/posts"
	"blog-gotth/templates"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Resolve content directory
	contentDir := filepath.Join("content", "posts")
	if env := os.Getenv("CONTENT_DIR"); env != "" {
		contentDir = env
	}

	// Load all posts at startup
	allPosts, err := posts.LoadAllPosts(contentDir)
	if err != nil {
		log.Fatalf("Failed to load posts: %v", err)
	}
	log.Printf("Loaded %d posts from %s", len(allPosts), contentDir)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	// Static files
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Home page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		folioDate := now.Format("Monday, January 2, 2006")
		editionNumber := fmt.Sprintf("%02d", len(allPosts))

		component := templates.HomePage(allPosts, folioDate, editionNumber)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Individual post pages
	r.Get("/post/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		post, found := posts.GetPostBySlug(allPosts, slug)
		if !found {
			http.NotFound(w, r)
			return
		}
		component := templates.PostPage(post)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// HTMX fragments
	r.Get("/api/wire-posts", func(w http.ResponseWriter, r *http.Request) {
		wirePosts := make([]posts.PostMeta, 0)
		if len(allPosts) > 11 {
			for _, p := range allPosts[11:] {
				wirePosts = append(wirePosts, p.ToMeta())
			}
		}
		component := templates.WireFragment(wirePosts)
		templ.Handler(component).ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
