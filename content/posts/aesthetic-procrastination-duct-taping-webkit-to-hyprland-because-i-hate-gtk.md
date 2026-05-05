---
title: "Aesthetic Procrastination: Duct-Taping WebKit to Hyprland Because I Hate GTK"
date: "2026-03-25"
author: "Mega Nugraha"
summary: "I spent an unreasonable amount of time building floating desktop widgets with Tauri, Hyprland window rules, and a system-level keylogger — all because I saw a macOS setup on a stream and refused to let it win."
tags: ["Linux", "Rust", "Tauri", "Hyprland", "Desktop Customization", "Wayland"]
published: true
---

There is a specific kind of brain rot that infects Linux users. It’s the belief that if your desktop environment looks just a _little_ bit cleaner, your productivity will finally skyrocket. Spoiler alert: it won't. But that didn't stop me from spending a ridiculous amount of time building a to-do list widget instead of actually doing my to-do list.

It started, as most terrible time-sinks do, with a stream. I was watching _peppy_ (the creator of osu!) and saw his macOS setup. He had these impossibly clean, floating widgets on his desktop powered by a tool called Übersichts.

I use Linux (Arch, by the way, because of course I do). Übersichts doesn't exist here. Normally, a sane person would say, "Oh well," and go back to work. But I have a pathological need to replicate things I see on other platforms just to prove I can. Also, the **Ramadhan Code Fest 2026** was happening, and I needed an excuse to submit something.

So, I built **Yet Another Slightly Cooler To-Do List App**.

### The Illusion of a Native App
If you look at a screenshot, it looks like a sleek, unified, system-native widget bar sitting at the bottom of my screen. It has a task dock, a keystroke visualizer, and a music visualizer that pulses to MPRIS data.

It is none of those things.

What it _actually_ is, is four completely separate, transparent web pages. I used **Tauri** to spin up borderless webkit2gtk windows, and then I used `hyprctl` (Hyprland’s IPC) to violently force them into specific coordinates on my monitor. It’s an elaborate optical illusion. I am essentially using a literal web browser to display 20 characters of text just because I refuse to touch the abominations that are GTK or Qt.

(And before you ask why I didn't use Electron: I already have five Chromium instances eating my RAM. I draw the line at six.)

### Building a Keylogger for the Vibes
The widget includes a keystroke visualizer—a little bar that shows what you’re typing in real-time before fading out. It’s great for streams, or for realizing exactly how many times you mistype `sudo`.

But how do you get global keystrokes in Wayland? Wayland’s entire security model is built around stopping apps from doing exactly that. High-level APIs wouldn't let me do it.

So, I bypassed the display server entirely. I wrote a Rust backend using `evdev` to listen directly to `/dev/input`. Yes, I wrote a system-level keylogger that interrogates the raw hardware of my keyboard, entirely so I could render some fading CSS text in a transparent web window.

Overkill? Absolutely. But it looks cool.

### The Rest of the Stack (Keyword Bingo)
To round out this monument to procrastination, I tied into the system DBus using `zbus` to interrogate MPRIS. The app literally watches my media players to see if I’m listening to LoFi beats for the 8th hour in a row, adjusting its polling rate to save CPU cycles when the music stops.

I also used **Bun** as the JavaScript runtime. Why? Because `npm` was taking too long, and honestly, Bun has a funnier logo. I’m an overworked open-source contributor; my technical choices are primarily driven by spite and aesthetics at this point.

### The Verdict: Floating Regret
Does this app actually help me finish my tasks? No. Pressing "Enter" on a new task activates it immediately with zero confirmation, mimicking the speed of my own poor life choices.

It floats, it pulses, and it tracks my keys. It is a testament to the modern developer's ability to spend 10 hours building a tool that saves zero minutes of work. But at least when I’m staring blankly at my screen ignoring my deadlines, the bottom 30 pixels of my monitor look incredibly professional.

If you’re on Hyprland and want to experience this specific brand of UI duct-tape, the source code is up. Just don't try to run it on Windows. I can only save you from so much.

**The Procrastination Stack:**
- **Core:** Tauri (Rust + webkit2gtk)
- **Window Management:** Hyprland window rules (The real MVP)
- **Hardware Interrogation:** `evdev` & `zbus`
- **JS Runtime:** Bun (for the logo)