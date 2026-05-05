---
title: "MTL No Longer Sucks, So Let's Do Something Fun With It"
date: "2026-04-03"
author: "Mega Nugraha"
summary: "I spent days building a Tauri-based desktop app just so I could read the next untranslated volume of a Japanese light novel. It's a collection of extremely pragmatic duct tape holding together a volatile AI pipeline, but it works."
tags: ["AI", "Tauri", "Rust", "LLM", "Automation"]
published: true
---

First of all, the purists at NovelUpdates would tie me to a stake and light me on fire if they ever caught me doing this, so I don't really plan to market this thing anywhere except for my own personal use. With that out of the way, let's start this story.

Look, I like reading Japanese light novels. You know what I don't like? Spending five years of my life memorizing Kanji just to figure out what happens in chapter 12.

My descent into this specific brand of madness started with *The Secrets of the Silent Witch*. I wanted to read ahead, and my strategy to bypass the wait was peak manual labor. I would crack open an EPUB editor like Sigil, rip the raw text out on a grueling, per-chapter basis, throw it into Gemini, and then painstakingly copy-paste it back. Then, I’d sit there sweating, praying I could somehow edit the resulting English translation back into the original markup without breaking the entire file.

It was an absolutely miserable workflow, but it proved one dangerous concept: modern AI machine translation (MTL) is actually good enough to read. 

But doing that manual copy-paste loop for an entire book? You are stepping into a special kind of hell.

### The Problem: A Web-Dev Nightmare in a Zip File
If you’ve never unzipped an EPUB file, consider yourself blessed. It is not a magical book format. It is a digital trench coat hiding a bunch of 2005-era XHTML files, messy OPF metadata, and CSS that actively hates you.

And Japanese EPUBs? They are a special breed of awful. Beyond just forcing vertical right-to-left (vertical-rl) layouts, they are absolutely infested with legacy typography markup. I am talking about the curse of `<ruby>` tags (furigana) that shatter a single word into four different XML nodes, completely lobotomizing any standard machine translation attempt because the context is physically split apart. Then there are the em-sesame emphasis dots, which inject dozens of useless `<span>` wrappers into the prose just to make a word look mildly important.

Trying to manually wrangle translated English text back into that deeply fragmented, heavily nested structure while keeping the formatting intact is basically a digital hazmat situation.

### The Solution: Lazypub
The initial idea was simple: write a script, feed the EPUB to an LLM, translate the text, pack it back up. Done, right? Absolutely not. If you just dump a chapter into an AI and say "translate this," it will hallucinate, destroy your XML tags, and hand you back a block of plain text that breaks your e-reader.

I needed a local app—something that could actually act as a persistent workspace. So, I grabbed Tauri (because I'm tired of Electron eating my RAM), slapped a React frontend on it, and decided to build a proper pipeline. Now, when you import an EPUB into Lazypub, it unzips the whole thing into a dedicated folder on your hard drive. Just like a real developer workspace.

### The First Wall: The CSS Exorcism
Before we could even translate the words, I had to fix the structure. Reading vertical right-to-left text translated into English is a great way to induce a migraine. 

Now, if you open a web development dictionary, the "correct" way to fix this is to parse the XHTML chapters, strip out the legacy Japanese layout classes, write a clean, modern stylesheet, and apply new, semantic class names. But remember my engineering ethos: minimum personal effort, maximum structural integrity. 

I didn't want to write a complex XML parser, and I definitely didn't want to use Regex. So, I committed a cardinal web-dev sin instead.

You click a button called **Normalize Layout**, and Lazypub feeds only the CSS and OPF files to Gemini with a very specific, highly cursed instruction: *Convert the layout to a standard horizontal-tb format, but do not change a single class name.*

Gemini goes in, surgically guts the properties of every layout class, and replaces them with standard horizontal rules. The XHTML chapter files are completely ignored. Yes, this means my final, translated EPUB contains CSS classes literally named `.vertical-text-box` that are actively rendering standard horizontal text. It is a semantic nightmare. It would make a senior frontend developer cry. 

But does it render perfectly on my e-reader app without me having to write a single line of layout logic? Yes. I replaced the devil of vertical CSS with the slightly less intimidating devil of semantic mismatch. Functional garbage is still functional.

### The Second Wall: The "Dave" Protocol 
If you’ve ever suffered through machine translation, you know the biggest architectural flaw of LLMs: they have absolutely zero object permanence. A protagonist named "Ryotaro" on page one will become "Dragon Boy" on page three, and for some inexplicable reason, "Dave" by the end of the chapter.

I needed a glossary to enforce state. But building a glossary manually means reading the raw Japanese text, which completely defeats the purpose of being lazy.

In the Enterprise world, solving LLM hallucination is done using RAG (Retrieval-Augmented Generation). I didn't want to build any of that. So, I built a **Poor Man's RAG** instead.

Enter the Entity Extraction loop. Before translating a chapter, Lazypub scans the raw text and extracts proper nouns. But instead of relying on the LLM to just guess the translation, I added a feature where you can paste a URL to a fandom wiki. Lazypub actively crawls that wiki, cross-references the extracted Japanese names, figures out who "That One Guy" actually is, and proposes the official, localized English spelling.

These terms get dumped into a raw `glossary.json` file sitting directly in your project root. Once you click "Approve," that JSON file becomes an absolute, unyielding guardrail. Suddenly, "Dave" is consistently "Ryotaro" again, and you didn't even have to pay for a Pinecone vector database subscription to make it happen.

### The Jank & A Pack of Duct Tapes
Is this a flawless, Enterprise-grade piece of software? Absolutely not. It is a collection of extremely pragmatic duct tape. 

- **The State Manager is Just Your Hard Drive:** I didn't want to deal with complex database schemas. When you import an EPUB, the Rust backend violently unzips the archive directly into a folder. That folder is your database. If the app crashes, nothing is lost because the file system is the single source of truth.
- **Temperature Tuning & The Monaco Safety Net:** To stop Gemini from getting "creative," the API calls are heavily temperature-tuned. I essentially lobotomized the LLM's creativity to turn it into a strict, deterministic parsing engine. 
- **The Collateral Damage:** In order to get this monstrosity to compile and run on Windows, I had to make sacrifices. The biggest casualty was the epubjs live preview. If you use Lazypub right now, you are essentially editing blind until you hit export.

### The Verdict
I spent days building a persistent, Tauri-based desktop application, writing Rust backend functions, and engineering complex AI prompts... just so I can read the next untranslated volume of a light novel.

But it works. I have a dashboard of recent projects, my glossaries sit there neatly along with the unzipped epub artifacts, and I never have to look at vertical Japanese CSS ever again.

If you also want to pretend you're "working" on a translation while an AI does 99% of the job, feel free to clone the repo. Good luck, you'll need it.

**The Stack (Keyword Bingo):**
- **Framework:** Tauri (Vite + React + Rust)
- **Package Manager:** Bun (because npm install takes too long)
- **AI Brain:** Gemini (Flash/Pro via Google AI Studio)
- **Architecture:** Persistent Local Workspaces (No cloud DBs here)
