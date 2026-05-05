package posts

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

//go:embed content/*.md
var embeddedFS embed.FS

// GetEmbeddedFS returns the embedded filesystem for posts, rooted at the content directory.
func GetEmbeddedFS() fs.FS {
	f, err := fs.Sub(embeddedFS, "content")
	if err != nil {
		// This should never happen if the go:embed is correct
		return embeddedFS
	}
	return f
}

// Post represents a blog post with metadata and rendered HTML content.
type Post struct {
	Slug      string
	Title     string
	Date      time.Time
	DateStr   string
	Author    string
	Summary   string
	Tags      []string
	Published bool
	Content   string // raw markdown
	HTML      string // rendered HTML
}

// PostMeta is a lightweight version for listings (no content/HTML).
type PostMeta struct {
	Slug    string
	Title   string
	Date    time.Time
	DateStr string
	Author  string
	Summary string
}

// frontMatter maps to the YAML frontmatter in each markdown file.
type frontMatter struct {
	Title     string   `yaml:"title"`
	Date      string   `yaml:"date"`
	Author    string   `yaml:"author"`
	Summary   string   `yaml:"summary"`
	Tags      []string `yaml:"tags"`
	Published *bool    `yaml:"published"`
}

var slugRe = regexp.MustCompile(`[^a-zA-Z0-9\-_.]`)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
)

// normalizeSlug turns a filename into a URL-safe slug.
func normalizeSlug(filename string) string {
	name := strings.TrimSuffix(filename, ".md")
	name = strings.ReplaceAll(name, "%", "-percent-")
	name = strings.ReplaceAll(name, " ", "-")
	name = slugRe.ReplaceAllString(name, "-")
	return strings.ToLower(name)
}

// renderMarkdown converts markdown bytes to HTML.
func renderMarkdown(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return "", fmt.Errorf("markdown render: %w", err)
	}
	return buf.String(), nil
}

// FormatDate formats a time.Time to "January 2, 2006".
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("January 2, 2006")
}

// LoadPost reads a single markdown file from an fs.FS and returns a Post.
func LoadPost(fds fs.FS, filePath string) (Post, error) {
	f, err := fds.Open(filePath)
	if err != nil {
		return Post{}, fmt.Errorf("open post %s: %w", filePath, err)
	}
	defer f.Close()

	var fm frontMatter
	content, err := frontmatter.Parse(f, &fm)
	if err != nil {
		return Post{}, fmt.Errorf("parse frontmatter %s: %w", filePath, err)
	}

	slug := normalizeSlug(filepath.Base(filePath))

	date, _ := time.Parse("2006-01-02", fm.Date)
	if fm.Date == "" {
		date = time.Now()
	}

	author := fm.Author
	if author == "" {
		author = "Mega Nugraha"
	}

	published := true
	if fm.Published != nil {
		published = *fm.Published
	}

	summary := fm.Summary
	if summary == "" {
		words := strings.Fields(string(content))
		if len(words) > 30 {
			words = words[:30]
		}
		summary = strings.Join(words, " ") + "..."
	}

	html, err := renderMarkdown(content)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Slug:      slug,
		Title:     fm.Title,
		Date:      date,
		DateStr:   FormatDate(date),
		Author:    author,
		Summary:   summary,
		Tags:      fm.Tags,
		Published: published,
		Content:   string(content),
		HTML:      html,
	}, nil
}

// LoadAllPosts reads all .md files from the given fs.FS and directory, sorted newest first.
func LoadAllPosts(fds fs.FS, dir string) ([]Post, error) {
	entries, err := fs.ReadDir(fds, dir)
	if err != nil {
		return nil, fmt.Errorf("read posts dir: %w", err)
	}

	var posts []Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		post, err := LoadPost(fds, filepath.Join(dir, entry.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping %s: %v\n", entry.Name(), err)
			continue
		}
		if post.Published {
			posts = append(posts, post)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

// GetPostBySlug finds a post by its slug from a pre-loaded slice.
func GetPostBySlug(posts []Post, slug string) (Post, bool) {
	for _, p := range posts {
		if p.Slug == slug {
			return p, true
		}
	}
	return Post{}, false
}

// ToMeta converts a Post to PostMeta (strips content).
func (p Post) ToMeta() PostMeta {
	return PostMeta{
		Slug:    p.Slug,
		Title:   p.Title,
		Date:    p.Date,
		DateStr: p.DateStr,
		Author:  p.Author,
		Summary: p.Summary,
	}
}

// AllMeta converts a slice of Posts to PostMeta.
func AllMeta(posts []Post) []PostMeta {
	metas := make([]PostMeta, len(posts))
	for i, p := range posts {
		metas[i] = p.ToMeta()
	}
	return metas
}
