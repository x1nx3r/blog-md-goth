---
title: "The 99 Performance Score: Why I Traded Next.js for a GOTTH Press"
date: "2026-05-05"
author: "Mega Nugraha"
summary: "Next.js is great, but my personal blog felt like a tank where I needed a bicycle. I nuked the repo and rebuilt the entire thing in Go just to save a few hundred milliseconds of my life. Was it worth it? Probably not, but look at the score."
tags: ["Go", "HTMX", "Performance", "WebDev", "Vercel"]
published: true
---

I used to think my blog was fast. It was built with Next.js, hosted on Vercel, and it looked fine. But every time I ran a PageSpeed report, I felt like a fraud. The metrics were "green," but I knew the truth. Under the hood, there was a massive hydration cycle happening every time someone just wanted to read 500 words of my rambling.

Next.js is a marvel of engineering for big apps, but for a blog that looks like a 1954 newspaper? It’s like using a nuclear reactor to power a flashlight. I was tired of sending a megabyte of JavaScript just to render some serif text and a couple of borders. So, I decided to be difficult. I nuked the repo and rebuilt the entire thing using the **GOTTH stack**: Go, Templ, Tailwind, and a tiny, almost apologetic sprinkle of HTMX.

The result? A **99 Performance score**, an LCP of **0.6s**, and a blog that loads so fast it actually startles me.

### Death by Hydration
React has this habit of waking up, looking at the perfectly functional HTML the server already sent, and saying, *"I need to control this."* It then spends 300ms of the user's CPU time essentially marking its territory. For a newspaper theme, this is an insult. A newspaper doesn't "hydrate." It sits there, it's ink on paper, and it works.

I wanted the digital equivalent of a morning edition—static, heavy on the serifs, and absolutely uncompromising on load speed. I don't need a VDOM to tell me how to render a paragraph. I just need a string of HTML and some CSS that doesn't weigh as much as a small car. I was tired of the "modern web" tax. I wanted to pay in cash.

### The Stack (or: How to avoid Node_Modules)
I wanted a stack that felt as "analog" as the theme looks. I landed on the GOTTH stack. It sounds like something a Victorian goth would build in a basement during a fever dream, and that’s exactly the vibe I wanted.

- **Go:** The engine. It compiles to a single binary. No dependency hell. No `node_modules` folder eating my disk space. It handles the routing, the markdown parsing, and the static file serving with the cold efficiency of a corporate accountant.
- **Templ:** This is JSX for gophers. It compiles to pure Go code. There’s no diffing, no runtime overhead. It just spits out strings and dies. If you haven't tried it, it's like writing HTML but with the safety of a compiler that actually hates you.
- **Tailwind v4:** For the styling, but minified to within an inch of its life. I'm using the latest v4 engine because it's even faster and more opinionated, which matches the rest of the stack.
- **HTMX:** Honestly? It's barely doing anything. I have it polling for a "News Wire" in the sidebar because I thought it looked cool. If I removed it tomorrow, the site would still be 99.9% as functional. It’s the digital equivalent of a fancy sticker on a typewriter. It’s there for the "vibe" of real-time news, even if the "news" is just me rambling about CSS.

### The "Prepper" Strategy
One of the biggest bottlenecks for Vercel’s serverless functions is cold starts and network chaining. Usually, you’d load fonts from Google and JS from a CDN. But I’m paranoid. What if Google is slow? What if the DNS lookup takes 200ms?

I went full "prepper mode." I used Go’s `embed` package to pack **everything** into the binary. 
- **The Markdown Content:** Every post is a string in the binary. I'm not hitting a database. I'm not even hitting the disk. I'm hitting memory.
- **The Typography Obsession:** I downloaded all 36 font files for EB Garamond, Playfair Display, and Crimson Text. I didn't want a single external DNS handshake. My CSS points to `/static/fonts/`, and those files are living inside the executable.
- **Dependencies:** I even embedded the HTMX minified JS. I don't want to talk to a CDN. I don't want to talk to anyone. 

By the time the binary runs on Vercel, it doesn't need to look at the filesystem or the outside world. It is a self-contained, 14MB digital press. It doesn't want to talk to anyone, and that's why it's fast.

### The Workflow (The "I'm Too Tired for a CMS" Edition)
People ask me what CMS I use. I don't use a CMS. I’m an overworked dev; I’m not writing a session handler or a login page at 2 AM. 

My workflow for "publishing" an article is as follows:
1. Open a new `.md` file in `internal/posts/content/`.
2. Write words until I feel better. 
3. `git push`.

Vercel triggers a build, Go compiles the Markdown into the binary, and the site updates. It's primitive, manual, and has zero features. There's no "Headless CMS" bill, no database to migrate, and no dashboard to get lost in. If I want to fix a typo, I'm committing to main. It's high-stakes journalism.

### The Codebase: A Beautiful Disaster
I'll be honest: the codebase is a mess. It’s a Frankenstein’s monster of Go templates, hacked-together CSS, and "I'll fix this later" comments. I spent four hours manually tree-shaking 53KB of CSS into 20KB using Perl scripts and pure spite because I couldn't get a "modern" tool to work the way I wanted.

I have functions that probably shouldn't exist and data structures that would make a CS professor weep. But when you load the page, it hits you in 0.4 seconds. The code is a disaster, but the *result* is a 99. That's the secret: I'd rather have a messy engine that wins the race than a shiny one that's still at the starting line. Reliability isn't about clean code; it's about the thing working every single time you hit the URL.

### The Performance Cult
Chasing a 100 Performance score is a mental illness. I spent thirty minutes figuring out how to defer a 16KB HTMX script just to save 240ms of "Total Blocking Time." I bundled my CSS into a single file and inlined the critical parts because I hated the idea of a second network request.

Is it worth it? Probably not for most people. But for me, that 99 score is a badge of honor. It’s a statement that says, *"I care about your time more than I care about my convenience."* In an era where every landing page is a 5MB payload of tracking scripts and "engaging" animations, a 40KB newspaper is a revolutionary act.

### The Verdict: Less Annoyed
Running the new site through PageSpeed Insights was the most satisfying thing I’ve done all year. 
- **FCP:** 0.4s
- **LCP:** 0.6s
- **TBT:** 0ms

By removing the "modern web" tax, I actually made the web usable again. It’s fast, it’s private, and it’s robust. If the major CDNs go down tomorrow, my blog will be the only thing left standing, still polling for news that will never come. 

It’s worth every messy line of Go.

**The Stack (Keyword Bingo):**
- **Language:** Go (The Gopher)
- **Templates:** Templ (The JSX-killer)
- **CSS:** Tailwind v4 + Spite
- **Interactivity:** HTMX (The "Vibe" addition)
- **Deployment:** Vercel (Because it's free)
