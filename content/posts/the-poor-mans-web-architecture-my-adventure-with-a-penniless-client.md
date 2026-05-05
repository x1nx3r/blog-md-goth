---
title: "The Poor Man's Web Architecture: My Adventure with a Penniless Client"
date: "2026-03-21"
author: "Mega Nugraha"
summary: "So, last year when I was doing my KKN, I had to engineer something so outrageous, never in my mind did I imagine it could actually work. Now that a pretty long time has passed—and the servers somehow haven't caught fire—let me tell you the story of how I built a secure, production-ready web portal when my budget was literally zero."
tags: ["Architecture", "Next.js", "Cloudflare", "Firebase", "Zero-Budget", "Web Development"]
published: true
---

## The Story
So, last year when I was doing my KKN, I had to engineer something so outrageous, never in my mind did I imagine it could actually work. Now that a pretty long time has passed—and the servers somehow haven't caught fire—let me tell you the story of how I built a secure, production-ready web portal when my budget was literally zero.

If you've ever done university KKN (community service), you know the drill: you show up in your jacket, you smile, and suddenly you are the designated "IT Guy" tasked with digitizing an entire village. The grand vision from the _Kelurahan_ (Village) office was massive. They wanted a fully-fledged digital directory for local MSMEs (UMKM), community news feeds, and a custom admin dashboard.

The budget? Absolutely nothing.

The "investors" for this startup were just me and my teammates chipping in to buy the domain: **[kelurahankemayoran.com](https://kelurahankemayoran.com)**. I had exactly four weeks to build a secure, scalable, and performant web app for exactly Rp 0 a month.

Little did I know, this wouldn't be as easy as spinning up a standard VPS and calling it a weekend.

### The Premise: Blast Radius and Broke Architecture

Everyone says you need a robust AWS setup or a beefy DigitalOcean droplet for a production site. But when your client is the local _Kelurahan_ and your budget is nonexistent, you have to reconsider your choices. In this context, AWS essentially stands for _Awas, Warga Susah_ (Careful, the locals are broke). I couldn't leave them with a server bill they wouldn't pay after I went back to campus.

The architecture had to be foolproof and free. More importantly, I needed to limit the blast radius. If I baked the admin panel into the main site, someone would inevitably click the wrong button and take the whole domain offline.

So, I split the architecture into two distinct Next.js apps hosted on Vercel:
1. **The Public Face**: A lightning-fast, static-leaning frontend that could handle the entire sub-district clicking on it at once without costing a cent.
2. **The Control Room**: An isolated admin dashboard, built strictly for the village staff.
    
With the frontends decoupled, I had to build the infrastructure tying them together. And by "infrastructure," I mean throwing more free-tier technology at the problem than absolutely necessary.

### The First Wall: No Time for Tables

In a corporate environment, you might spend a week just planning your database schema. On a KKN project, you have about three days between community health drives to get the backend running.

I didn't want to waste my life building relational SQL tables, writing migrations, and setting up an ORM just to store some article text and image links.

Enter Firebase. It’s the patron saint of the impatient developer. We didn't need complex joins; we just needed a place to dump JSON objects containing village news and UMKM details. Firebase’s free tier meant I could spin up a NoSQL database in minutes, giving the admin dashboard instant read/write capabilities while saving my sanity.

### The Second Wall: The Storage Heist

The biggest hurdle was file storage. Village websites run on photos—every community cleanup, every _Bapak/Ibu PKK_ event, and every new batch of local snacks needs to be documented. Free-tier storage limits are notoriously stingy, and I wasn't about to put my own credit card on the line for S3.

That’s when I remembered a personal side project I had built: a simple image-sharing site living at `uploads.x1nx3r.uk`. You throw an image at it, and it tosses back a pristine `cdn.x1nx3r.uk` link. It was built on Cloudflare R2, meaning generous storage and, crucially, zero egress fees.

So, I put two and two together.

I took my personal image catapult (affectionately named the `r2-object-thrower`), slapped a public API on it, generated a base64 auth key, and quietly tucked that key into the environment variables of the village admin site.

Yes, the official website for Kelurahan Kemayoran is secretly routing its media through a university student’s custom CDN microservice. It's a massive flex, it bypasses the storage wall completely, and the village officials just see their photos load instantly.

### The Final Boss: Defending Against the Dark Arts

Here is where things get serious. If you’ve spent any time looking at Indonesian government or village websites, you know there is a shadow war going on. Local sites are prime real estate for _judol_ (online gambling) syndicates looking to hijack domains for SEO spam.

Security wasn't just a "nice to have." It was a strict requirement.

Building a custom authentication flow with JWTs or dealing with session cookies felt like overkill, and honestly, it left too much surface area. The admin site is a Next.js app. I didn't want to risk some automated scraper nibbling on my JavaScript artifacts, trying to reverse-engineer my API routes or find a poorly scrubbed environment variable.

My solution? I completely removed the admin app from the public internet.

I slapped Cloudflare Access in front of the admin subdomain (**[admin.kelurahankemayoran.com](https://www.google.com/search?q=https://admin.kelurahankemayoran.com)**). I didn't build a login page; I let Cloudflare's edge network act as an enterprise-grade bouncer. I simply took the _Pak Lurah’s_ (Village Head's) email, dropped it into a manual whitelist, and called it a day.

Now, if you try to hit the admin site, Cloudflare intercepts you before a single byte of my Next.js code is ever served. The _Pak Lurah_ just clicks "Login," signs in with his existing Google account, and he's in. Meanwhile, the _judol_ hackers are left knocking on a brick wall at the edge network. Simple. Clean. Overkill.

### The Aftermath

Did I spend an unreasonable amount of time stitching together personal microservices, serverless databases, and edge-network bouncers just to build a digital bulletin board for local snack vendors? Yes. Yes, I did.

But this wasn't really just about building a website. It was about problem-solving under ridiculous constraints. About leaning way too hard into a problem to prove—mostly to myself—that I could orchestrate a highly secure, decoupled architecture for a penniless client.

Now they’ve got a lightning-fast public site, an isolated admin panel, and Zero Trust security, all running on a mildly exaggerated sense of duct-tape engineering. And it won't cost them a single Rupiah to maintain long after my KKN group has graduated.

Catch you next time for whatever weekend "fun" project I end up accidentally turning into a systems architecture exercise.

**Project Links & Tech Stack:**

- **Public Site:** [kelurahankemayoran.com](https://kelurahankemayoran.com)
- **Admin Dashboard:** [admin.kelurahankemayoran.com](https://www.google.com/search?q=https://admin.kelurahankemayoran.com) (Good luck getting past Cloudflare)
- **Frontend:** Next.js (React) hosted on Vercel
- **Database:** Firebase (NoSQL)
- **Storage:** Cloudflare R2 (via `r2-object-thrower` custom CDN)
- **Security:** Cloudflare Access (Zero Trust Auth)