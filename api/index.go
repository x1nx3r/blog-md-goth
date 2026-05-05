package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"blog-gotth/internal/posts"
	"blog-gotth/templates"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

var allPosts []posts.Post

func init() {
	contentDir := os.Getenv("CONTENT_DIR")
	if contentDir == "" {
		contentDir = filepath.Join("content", "posts")
	}

	var err error
	allPosts, err = posts.LoadAllPosts(contentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to load posts: %v\n", err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		folioDate := now.Format("Monday, January 2, 2006")
		editionNumber := fmt.Sprintf("%02d", len(allPosts))
		component := templates.HomePage(allPosts, folioDate, editionNumber)
		templ.Handler(component).ServeHTTP(w, r)
	})

	router.Get("/post/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		post, found := posts.GetPostBySlug(allPosts, slug)
		if !found {
			http.NotFound(w, r)
			return
		}
		component := templates.PostPage(post)
		templ.Handler(component).ServeHTTP(w, r)
	})

	router.Get("/api/wire-posts", func(w http.ResponseWriter, r *http.Request) {
		wirePosts := make([]posts.PostMeta, 0)
		if len(allPosts) > 11 {
			for _, p := range allPosts[11:] {
				wirePosts = append(wirePosts, p.ToMeta())
			}
		}
		component := templates.WireFragment(wirePosts)
		templ.Handler(component).ServeHTTP(w, r)
	})

	router.ServeHTTP(w, r)
}
