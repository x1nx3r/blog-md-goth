---
title: "The Human Backup Bot: A Story of Tiredness and Trust Issues"
date: "2026-03-25"
author: "Mega Nugraha"
summary: "I was tired of being the designated backup-fetching bot for my team, so I built a self-service S3 browser with Monaco Editor, streaming downloads, and Cloudflare Access — all because I couldn't trust anyone with the Coolify dashboard."
tags: ["DevOps", "Next.js", "Cloudflare", "DigitalOcean", "Automation"]
published: true
---

I love my team, but I’m tired. I’m tired of the message pings, tired of the context switching, and mostly tired of being the only person who knows how to fetch a database backup without accidentally setting the production server on fire.

For months, I was basically a glorified file transfer protocol. Every other day, someone would realize they broke something in dev and ask, _"Hey, can you grab me yesterday’s SQL dump?"_ Technically, they could do it themselves. But the backups live in Coolify, and giving everyone access to Coolify is like giving a group of toddlers the keys to a nuclear silo. It’s a jungle of toggles and "Destroy" buttons. I don't trust anyone—least of all myself after three hours of sleep—not to click the wrong thing.

So, when I think about it I have three solution I can hatch from my tired brain at 3 AM. What are those you ask? Well prepare to be entertained as those are:
- **Give them access to the Coolify dashboard or the DigitalOcean Spaces dashboard.** As I mentioned a few pixels up, this is a no-go. I don’t trust them enough to not nuke the several servers that live through the mercy of Coolify acting sane. And as for the DigitalOcean dashboard? I’m not going to be bothered to learn IAM there just to let them access the Spaces.
- **Cook up a simple API.** Trigger the API, get a backup. I could even add a secret token that I only share through WhatsApp. It's simple, but it’s an unhinged solution and I don’t want to risk it. Coolify's API probably does this already, but even though my team are all pretty good fullstack engineers, I don't trust them to invoke any API call to the main server. Knowing that almost all of them use GitHub Desktop, they’d surely use Postman—which buries throwing a token under so many layers of clicks just to invoke a single API.
- **Build a "Look but don't touch" cage.** needed a way for them to browse, preview, and download `.dmp` files without ever touching the infrastructure that actually keeps the lights on. It had to be fast, it had to be blue, and it had to be impossible for them to accidentally click a "Delete" button.
### The Problem: Hex Curses and 1999 UI
The backups live in DigitalOcean Spaces, which is just S3 with a different logo. Now, if I were a responsible engineer, I would have given the database instances sane, human-readable names. But I’m not, and I didn’t.

Instead, I have to deal with what I call the **Hex Curse**. You aren't looking for "KMP Backups"; you're looking for `mysql-database-uc4w4000k4sgo40k0s4ccws4`. To make matters worse, Coolify has this quirk where it buries the actual `.dmp` files under about five layers of nested directories. Browsing it manually feels like an archaeological dig where every level is just another UUID-named folder.

I didn't have the patience to click through a dozen folders just to find one backup, so I built the browser to be smarter (and lazier) than me.

I wrote the logic to essentially "skip the boring parts." The app crawls the bucket, ignores the useless nesting, and flattens the paths into something a human can actually parse. It’s got a folder tree, breadcrumbs (because I didn't want to fix the browser's back button), and a file counter at the bottom so I can see exactly how much data is haunting our storage limits. It basically turns a digital maze into a clean list, purely because I was too lazy to name the database properly in the first place.

### Streaming: Offloading the Suffer
Database dumps are massive. If you try to load a 300MB SQL file into a standard web response, the server just gives up and dies. Since I’m cheap, this whole thing is hosted on **Vercel**. Next.js is the fastest way to vibecode my way out of a problem at 2 AM, but Vercel’s serverless functions have RAM limits that don't like massive text blobs.

Instead of paying for a better tier, I set up the app to stream the file content. It pulls chunks from the bucket and shoves them directly into the Monaco Editor. Why the Monaco Editor, you ask? Because I needed syntax highlighting and a reliable diff view, and shoving Microsoft's entire IDE engine into a Next.js app was the fastest way to get it done. Are there more lightweight, elegant solutions out there? Probably. Do I care? No. This works, so that's that.

This is a conscious architectural choice: I am offloading the memory burden from my short-lived Vercel instance to my teammates' miserable laptops. My philosophy is simple: if you’re coming here, it’s because you broke something and you need to fix it. You can suffer through a bit of browser lag while the SQL loads. Consider it a penance for the production downtime.

Is it a perfect solution? Probably not. If our backups ever get truly gargantuan, the browser will probably leak memory like a sieve. But our current dumps are all under 100MB, so that is officially a problem for **Future Me** to solve. Current Me has a terminal to look at.

### The "Security"
Here’s the part that would make a real security auditor cry: the app has **zero authentication code**. There is no login page. No "Sign in with Google" button. No password hashing.

Building auth is a chore. I’m an overworked student; I’m not writing a session handler at 2 AM.

Instead, I used the. I threw the whole thing behind a Cloudflare Access tunnel. Cloudflare is the bouncer. It checks the email whitelist, handles the OTP, and keeps the _judol_ (online gambling) bots away from my database dumps. If you aren't on the list, the app doesn't exist to you. I didn't have to write a single line of security logic, and it's probably safer than anything I could have built myself.

### The Verdict: Less Annoyed
The UI is blue. Deep, high-contrast, terminal blue. It’s not a "design choice"—I just picked the first hex codes that didn't hurt my eyes in a dark room and I'm never changing them. Here's a screenshoot if you wanna look at it.
![S3 Browser Screenshot](/static/s3-browser-screenshot.webp)
It renames `.dmp` files to `.sql` because file extensions are just suggestions. It streams files so my laptop doesn't explode. And most importantly, I am no longer a backup bot. The team can find their own SQL dumps, compare their own disasters, and I can finally go back to whatever I was doing before someone asked me for a file.

If you’re the person in your group who’s tired of being the "Server Guy," feel free to steal the repo. It's not a masterpiece, but it’ll save you a few pings. Or just write something better and tell me about it.

**The Stack (Keyword Bingo):**

- **Framework:** Next.js + Bun (because I'm tired of npm)
- **Editor:** Monaco (The VS Code thing)
- **Storage:** DigitalOcean Spaces (S3 but cheaper, i think)
- **Auth:** Cloudflare Access (The "I'm too lazy for this" solution)