package handler

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"blog-gotth/internal/assets"
	"blog-gotth/internal/posts"
	"blog-gotth/templates"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

var allPosts []posts.Post

func init() {
	var postsFS fs.FS = posts.GetEmbeddedFS()
	contentDir := "."

	if env := os.Getenv("CONTENT_DIR"); env != "" {
		postsFS = os.DirFS(env)
		contentDir = "."
	}

	var err error
	allPosts, err = posts.LoadAllPosts(postsFS, contentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to load posts: %v\n", err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()

	// Static files from centralized embedded FS (for Vercel)
	staticFS, _ := fs.Sub(assets.Static, "static")
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

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
