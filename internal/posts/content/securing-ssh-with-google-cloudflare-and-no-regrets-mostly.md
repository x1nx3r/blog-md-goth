---
title: "Securing SSH With Google, Cloudflare, and No Regrets (Mostly)"
date: "2025-05-18"
author: "Muhammad Mega Nugraha"
summary: "How I overengineered a simple SSH setup using Docker, Cloudflare Tunnels, and Google OAuth because I'm lazy but not reckless."
tags: ["Security", "SSH", "Cloudflare", "Docker", "Linux"]
published: true
---

So, I've had some experience fiddling with SSH for a pretty long time — but whenever I do, it's always to some cloud VPS living on someone else's infrastructure. Then, a revelation hit me some random Saturday morning: "Wait, I'm on Linux. I have sshd enabled. I can SSH into this machine from my local network. Why not make it public?" 

Little did I know, this wouldn't be as easy as my other weekend "fun" projects.

## Why?

Because I wanted SSH access to my laptop. I figured, hey — just forward port 22 through a Cloudflare Tunnel to my domain, and call it a day. Easy. Then I remembered: my root password is about as strong as my math skills — which is to say, not at all. 

Suddenly, it didn't feel like such a great idea. The last thing I need is some guy with a Moscow IP poking around my terminal at 3 a.m., making off with my entire $0.73 Steam Wallet fortune. So, naturally, I reconsidered my choices.

## What?

That's when I stumbled across **Guacamole Server**. It looked promising — decent security, built-in RDP and VNC clients, an Apache backend, and a slick web-based admin dashboard. Fancy stuff. I spun it up with Docker Compose, full of hope. 

And just as I was about to tunnel into it, it dropped an **11 GB memory leak** straight into my RAM. The entire system froze. I had to do a hard reset, then review the logs trying to make sense of what might be the culprit for a solid hour, questioning every decision that led me there.

Eventually, I gave up and installed **Wetty** instead — which took five minutes to get running locally. Five. Minutes. Let that be a lesson: stick to the actual requirements and skip the unnecessary shiny stuff. (Says the guy who will absolutely add more shiny stuff later anyway.)

## How?

By throwing more technology at the problem than absolutely necessary, of course. I ran Wetty inside a Docker container — because if you're not containerizing 90% of your personal side projects, are you even overengineering correctly?

Wetty, for the uninitiated, is a web-based terminal that connects to your local machine over port 22. It's basically SSH — but in the browser, and slightly worse. Which, naturally, made it perfect. 

Then I set up a **Cloudflare Tunnel** directly to it. No port forwarding, no NAT gymnastics, no router UI from 2007. Just clean, simple tunneling. Elegant. Portable. Almost civilized.

But wait — it gets better. I stacked a **Cloudflare Access policy** on top, because nothing says "I care about security" like outsourcing auth to a tech giant. The catch? I had to spin up an entirely new Google Cloud project just to set up an OAuth provider for Cloudflare. 

That meant diving headfirst into the Google Cloud Console: flipping switches I barely understood, agreeing to terms I definitely didn't read, and wondering—briefly—if this was all a bit much for remote shell access. 

Now? Only my personal Google account gets through. No email, no shell. Simple. Clean. Overkill.

![Cloudflare Access Gateway Log](https://via.placeholder.com/600x300/24273a/cad3f5?text=Cloudflare+Access+Gateway+Log)
*Here's How the Cloudflare Access Log in Gateway Looks — A Bit Archaic But It Does Its Job*

## Why This Setup?

Because I'm lazy — but not reckless. I didn't want to manage SSH keypairs or rig up some spare USB stick as a makeshift YubiKey that I'd inevitably lose in a drawer full of other USB sticks and dongles. 

I just wanted to open a browser tab, click "Authorize with Google," and type `fastfetch`. And now? I can do exactly that — from any device, anywhere... as long as it's not connected to sketchy coffee shop Wi-Fi or an actual potato.

![SSH Session Through Domain](https://via.placeholder.com/600x400/24273a/cad3f5?text=SSH+Session+in+Browser)
*Here's the Actual SSH Session Accessed Through My Domain*

## The Architecture

Here's what the final setup looks like:

1. **Wetty** running in Docker container
2. **Cloudflare Tunnel** exposing it to my domain
3. **Cloudflare Access** with Google OAuth protection
4. **Zero Trust** network access from anywhere

```bash
# The Docker command that started it all
docker run -d \
  --name wetty \
  -p 3000:3000 \
  wettyoss/wetty \
  --ssh-host localhost
```

## Lessons Learned

- Sometimes the simple solution (Wetty) is better than the fancy one (Guacamole)
- Cloudflare Tunnels are genuinely magical for homelab setups
- Google OAuth integration is surprisingly straightforward once you get past the console UI
- Overengineering can be fun, but know when to stop

## Conclusion

Did I spend an unreasonable number of hours duct-taping together Docker, Cloudflare, and a browser-based terminal just to run `htop` from my phone? Yes. Yes, I did. 

But this wasn't really about SSH. It was about control. About leaning way too hard into a problem that didn't ask for a solution. About overengineering for fun and proving — mostly to myself — that I could stitch together a secure remote terminal using tools designed by people who have trust issues with the public internet.

Now I've got **Zero Trust SSH access**, a web terminal in my browser, and a mildly exaggerated sense of security.

Huge thanks to Google, Cloudflare, Docker, and the countless developers who keep making these powerful tools free — so folks like me can build unnecessarily complicated setups at 2 A.M. and call it productive.

Catch you next time for whatever weekend "fun" project I end up accidentally turning into a systems architecture exercise.

---

**Tech Stack Used:**
- Wetty (Web-based terminal)
- Docker & Docker Compose
- Cloudflare Tunnels
- Cloudflare Access
- Google Cloud OAuth
- Linux (Fedora)