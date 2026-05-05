---
title: "The 99 Performance Score: Why I Traded Next.js for a GOTTH Press"
date: "2026-05-05"
author: "Mega Nugraha"
summary: "Next.js is great, but my personal blog felt like a tank where I needed a bicycle. I nuked the repo and rebuilt the entire thing in Go just to save a few hundred milliseconds of my life. Was it worth it? Probably not, but look at the score."
tags: ["Go", "HTMX", "Performance", "WebDev", "Vercel"]
published: true
---

I used to think my blog was fast. It was a standard Next.js deployment on Vercel, the metrics were mostly green, and it looked fine. But every time I opened the network tab, I felt like a fraud. Under the hood, a massive React hydration cycle was spinning up just so someone could read a few paragraphs of serif text.

Next.js is a marvel for complex apps. But for a blog styled like a 1954 newspaper? It’s like firing up a V8 engine to drive to the end of the driveway. I was shipping a megabyte of JavaScript just to render some borders and static text.

I got tired of paying the "modern web" tax. So, I nuked the repo. I rebuilt the entire thing using the **GOTTH stack**: Go, Templ, Tailwind, and a practically microscopic sprinkle of HTMX.

The result is an LCP of 0.6s, a 99 Performance score, and a site that loads so aggressively fast it almost feels jarring.

### The Hydration Tax

React has this exhausting habit of looking at the perfectly functional HTML the server already sent, waking up, and deciding, "No, I need to control this." It then burns 300ms of CPU time marking its territory.

For a newspaper theme, this is an insult. A newspaper doesn't "hydrate." It’s ink on paper. It sits there and works. I wanted the digital equivalent of a morning edition—static, uncompromising, and completely stripped of VDOM overhead. I just needed pure HTML and CSS that didn't weigh as much as an Electron app. 

In my old Next.js setup, I had `react-markdown`, `rehype-raw`, and `remark-gfm` all fighting for CPU cycles just to turn a string into a list item. It was a 280MB `node_modules` grave just to host a few thoughts. Now? I have a 14MB binary that doesn't even know what a "Node" is.

### Why Not a "Lightweight" Framework?

People usually ask: *"Why didn't you just use Hugo? Or Astro? Or Jekyll?"*

Sane people use those. Hugo is incredibly fast. Astro is brilliant at stripping out JS. But I didn't want "sane." I wanted "outrageous." 

If I used Astro, I'd still be in the JavaScript ecosystem. I'd still be dealing with `npm install`, `package-lock.json` conflicts, and a build pipeline that feels like a Rube Goldberg machine. If I used Hugo, I’d be restricted by its templating engine and its specific way of doing things. 

I wanted to own the engine. I wanted to write Go because Go is what I use when I want things to stay up and run fast. Building a custom blog engine is the ultimate act of "aesthetic procrastination"—it’s doing something incredibly complex just to achieve something simple, but doing it exactly your way. I didn't want a "blog framework"; I wanted a printing press I built myself with scrap metal and spite.

### Avoiding node_modules

I wanted a stack that felt as analog as the theme looks, which led me to Go and Templ.

Go is the engine. It compiles down to a single binary. There’s no dependency hell, no node_modules black hole eating my drive, and it handles routing and markdown parsing with absolute, cold efficiency.

Templ is essentially JSX for gophers, but it compiles into pure Go code. No diffing, no runtime overhead—it just spits out strings and dies. Paired with Tailwind v4 (minified to within an inch of its life), it’s exactly the kind of strict, compiler-enforced environment I wanted.

And HTMX? Honestly, it’s barely doing anything. I have it polling for a "News Wire" in the sidebar mostly for the aesthetic. If I stripped it out tomorrow, the site would still be entirely functional. It’s a fancy sticker on a typewriter. It provides just enough "live" feeling to keep the newspaper vibe alive without requiring a single line of client-side logic from me.

### The Doomsday Approach to Assets

One of the biggest bottlenecks with serverless deployments like Vercel is the cold start and the subsequent network chaining. Normally, you’d load fonts from Google and pull scripts from a CDN. But what if the DNS lookup takes 200ms?

I went full prepper mode. I used Go’s embed package to pack everything directly into the executable binary.

Every markdown post? A string in the binary. I'm not querying a database, and I'm not even reading from the filesystem. I’m hitting memory. I downloaded all 36 font files for EB Garamond, Playfair Display, and Crimson Text so I wouldn't have to make a single external DNS handshake. Even the minified HTMX script is embedded.

By the time the binary runs on Vercel’s edge, it doesn't need to look at the outside world. It’s a self-contained, 14MB digital press. It refuses to talk to anyone, which is exactly why it’s fast.

### "CMS" is a Strong Word

People keep asking what CMS I'm using. I don't have one. I’m an overworked dev; I’m not writing a custom auth handler or a rich text editor at 2 AM.

My publishing pipeline is primitive:
1. Open a new .md file in `internal/posts/content/`.
2. Type until I'm done.
3. `git push`.

Vercel catches the push, Go compiles the markdown into the binary, and it's live. There’s no database to migrate and no dashboard. If I need to fix a typo, I’m committing straight to main. High-stakes journalism.

### A Beautiful Disaster

I'll be honest, the codebase itself is a mess. If you surveyed the repo, you'd see the remnants of the old Next.js app sitting right next to the new Go engine like a abandoned construction site. It’s a Frankenstein’s monster of Go templates, hacked CSS, and pure spite. 

I spent four hours manually tree-shaking 53KB of CSS down to 20KB using Perl scripts just because I couldn't get a standard bundler to do exactly what I wanted. I have functions that shouldn't exist and data structures that would get me failed in a CS 101 class. There's a lot of "I'll refactor this later" that I know I never will.

But the code being a disaster doesn't matter when the page hits the browser in 0.4 seconds. I’d rather have a messy engine that wins the race than a beautifully architected Next.js app still waiting for the main thread to unblock. In this world, functional garbage that hits 99 is better than a "clean" 85.

### The Verdict

Chasing a 100 Performance score is a sickness. I spent half an hour figuring out how to defer a 16KB script just to shave off 240ms of Total Blocking Time.

Is it worth it? Probably not for a normal project. But for this? Absolutely. In an era where reading a simple article requires downloading a 5MB payload of tracking scripts and hydration logic, serving a fully functional 40KB newspaper feels like a revolutionary act.

It was worth every messy line of Go.
