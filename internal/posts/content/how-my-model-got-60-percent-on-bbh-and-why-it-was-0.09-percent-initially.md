---
title: How My Model Got 60% on BBH and Why It Was 0.09% Initially
date: "2025-12-30"
summary: "Standard benchmarks often fail to capture a Small Language Model’s true reasoning ability, punishing models for verbose or unconventional output. By using a powerful LLM as a judge, we can accurately evaluate logic rather than formatting. This approach reveals that models once deemed “failures” may actually reason correctly, highlighting the need to prioritize thinking over compliance in AI evaluation."
tags: ["AI", "Machine Learning", "LLM", "Reasoning", "Llama", "Benchmarks", "BBH"]
---

## How My Model Got 60% on BBH and Why It Was 0.09% Initially
Ad plug here: *If you want to test this "Thinking" capability yourself, the model is fully open-source. You can download **Llama-3.2-3B-Thinking** on Hugging Face, i advise you use LM-Studio and just search for my repo (Safetensors & GGUF available) here: [https://huggingface.co/x1nx3r/Llama-3.2-3B-thinking-v3-13.9K, https://huggingface.co/x1nx3r/Llama-3.2-3B-thinking-v3-13.9K-GGUF]*
#### LLM's and Benchmarks, What's The Deal With It
So, before everything, we need to talk about why we only have benchmarks as a reliable way to measure an LLM performance in a given task.  
#### The Nature of Generative Models
At its core, the nature of an LLM is drastically different from the traditional 'AI Models' we are used to. Traditional models focus on classification; you feed the model an input, and it spits out a definite output based on what it learned during training. For example, you would expect a CNN to classify an image into fixed categories it has already learned.
In contrast, a generative model aims to produce an output that simply 'makes sense' based on its training data. With an LLM, you feed it an unhealthy amount of text, and the model learns to predict a word that is 'good enough' to follow the previous one, creating something new, rather than exactly replicating the data it saw during training.
#### Evaluation Concerns of an LLM
In a traditional model, you can measure performance using figures that are simple to understand, easy to implement, and theoretically sound metrics like Precision, F1-Score, and Recall. This works because the categories are narrow enough to count and easily verifiable by humans. A generative model breaks this framework because we are no longer aiming for a single, 'exactly true' prediction from a fixed set and the scope of possibilities expands from a limited set of categories to the entire vocabulary of possible tokens. Now humor me if you could build a confusion matrix for 128.000+ tokens.
#### Then How Can We Reliably Measure an LLM Performance
The truth is, we essentially lack any reliable method to objectively quantify an LLM's general capabilities. Since we can't calculate 'intelligence,' we have been forced to abandon traditional metrics and settle for a simpler, albeit cruder, approach: acting like a high school teacher who gives a student a test and hopes the grade reflects their actual knowledge and those tests are what we call benchmarks.

Benchmarks are diverse; some focus on measuring performance in one specific subject, while others measure a specific ability like reasoning or coding. The list is essentially inexhaustible. This, in turn, creates a daunting task in the post-training stage. You might see that the model speaks and acts well enough, but you can't scientifically claim it works until you quantify its performance. 

This current approach creates a major issue: comparability. You simply cannot say a specific model is better than another unless you benchmark both of them on the same ground. This is especially painful when you are building a Small Language Model (SLM). An SLM naturally won't fare well on 'god-tier' benchmarks like Humanity's Last Exam, GPQA, or AIME yet those are the only numbers they are publishing. On the flip side, for them to go back to older, static benchmarks is meaningless, because modern models ace them without breaking a sweat. But even if you find that perfect 'middle-ground' benchmark for your SLM, there is a hidden trap waiting for you in the implementation details: The Grading Machine.
#### The Trap of Rigid Evaluation: Why I Got 0.09%
Most standard evaluation suites (like lm-eval-harness) rely on Regex Parsing to grade answers. Think of this like a Scantron machine from high school: if you don't bubble in the circle exactly right, the machine marks it zero, even if you wrote the correct answer in the margin. 

This is a nightmare for Small Language Models (SLMs), but the issue is exacerbated significantly when dealing with a Thinking Model like the one I built. By design, my model produces a long 'Chain of Thought', an internal monologue detailing its reasoning, before outputting a final answer. A rigid regex parser treats this reasoning as noise. It sees a wall of text instead of a simple 'A' or 'B' and marks it wrong. This structural mismatch is exactly why my model, which was essentially showing its work and reasoning correctly, initially scored a practically non-existent 0.09% on BBH
#### The Solution: LLM-as-a-Judge
To fix this, we have to fire the Scantron machine and hire a Teaching Assistant. This leads us to the LLM-as-a-Judge methodology. Instead of relying on a brittle regex script that chokes on verbosity, we ask a much capable frontier model (like GPT-4o or Claude 3.5 Sonnet) to evaluate the SLM's response.

While there is often valid skepticism about using AI to grade AI (the 'circular evaluation' problem), this argument falls flat when applied to Small Language Models. The dynamic here is not 'peer-grading'; it is a Professor grading a Freshman. The benchmarks we are using (BBH) are relatively trivial for a massive frontier model. The Judge can easily parse through the chaotic, verbose 'thinking' process of a 3B model, ignore the formatting quirks, and accurately determine if the logic holds up, something a regex parser could never do. This method essentially allows us to grade the _reasoning_, not just the _formatting_.
#### Case Studies: The "Scantron" vs. The Judge
To understand the 60% discrepancy, we have to look at the raw output. Below are three actual samples from the BBH benchmark run. In all three cases, the standard Regex evaluation gave a score of 0.0. These are taken from the first three samples of the Boolean Expression task in BBH, the LLM evaluator is the newly released Google's Gemini 3 Pro Preview accessed through Google AI Studio.
##### Case 1 : The Verbose Correct Answer
**Question:** `not ( True ) and ( True ) is` **Ground Truth:** `False`
Model Response: 
	\<think> To evaluate the Boolean expression "not (True) and (True)", let's break it down step by step. First, we have the innermost parentheses: "not (True)"... this part evaluates to False (0). Next, we have the outer parentheses... The expression becomes "(False) and (True)". Thus, the result of the expression "not (True) and (True)" is False. \</think>
	The result is False.

The Verdict:
- **Regex Score:** `0.0`  (Failed because of the prefix _"The result is..."_)
- **Judge Score:** `Correct`  (Identified the correct logic and final conclusion.)
The model is effectively punished for answering in a complete sentence. It successfully identified the Boolean operators and executed the logic, proving it _knows_ the answer.

##### Case 2: The Genuine Logic Failure
**Question:** `True and not not ( not False ) is` **Ground Truth:** `True`
Model Response: 
	\<think> ..Let's tackle the innermost parentheses first: "not False". ...Therefore, ¬F = T. ...Now, applying the NOT operator again: ¬T = F. ...So, putting it all together, the second part simplifies to "not not ( not False )" = F. \</think>
	A: False

The Verdict:
- **Regex Score:** `0.0` 
- **Judge Score:** `Incorrect`  (Correctly identified that the model failed the triple negation logic.)
- Note: This proves the Judge is not a "yes-man"; it penalizes actual errors.
The model got dizzy. It correctly solved the first negation (`not False` = `True`), but then confused itself during the subsequent negations (`not not`). The Judge spotted this logical breakdown and rightfully failed the sample, proving that the 60% score isn't inflated.

##### Case 3: The "Multiple Choice" Hallucination
**Question:** `not True or False or ( False ) is` **Ground Truth:** `False`
Model Response: 
	\<think>..."not True" is equivalent to the Boolean NOT operator applied to True... this part evaluates to False. ...Therefore, the expression simplifies to: "False or False or False" ...Thus, the entire expression evaluates to False. \</think>
	The final answer is A: False.
	
The Verdict:
- **Regex Score:** `0.0`  (Failed because the model invented a multiple-choice format _"A:"_ not present in the prompt.)
- **Judge Score:** `Correct`  (Ignored the formatting hallucination and graded the correct boolean logic.)
This is a classic "distribution shift" error. The model is used to answering multiple-choice questions (A, B, C), so it hallucinated a letter "A" to accompany its answer. The logic was flawless, but the formatting was fatal for the regex parser.

If we looked only at the Regex column, we would conclude the model is broken (0% accuracy). By looking at the Model Logic and Judge Verdict, we see a model that is reasoning correctly 66% of the time but struggling with output formatting, a classic trait of distilled Small Language Models.

### So What? The Implications I Realized
This discrepancy forces us to confront an uncomfortable truth about the current state of Open Source AI. If I had simply run `lm-eval-harness`, looked at the **0.09%** score, and called it a day, I would have scrapped this project. I would have deleted the checkpoints, assumed the distillation failed, and moved on.
 Yet, the model was actually performing at 60% accuracy, a massive achievement for a 3B parameter model on a hard benchmark like BBH.

This begs the question: How many other "failed" models are gathering dust in the Hugging Face graveyard, hidden behind a bad regex score?

We are currently in an era where we prioritize Compliance over Capabilities. We train models to be good test-takers, to output "A" instead of "The answer is A", often at the cost of their reasoning potential. My Llama-3.2-3B might be a messy, chatty student that fails the Scantron, but it _understands_ the material. And in the quest for AGI, I would argue that a model that thinks correctly but formats poorly is infinitely more valuable than a model that formats perfectly but thinks vacantly.

So, the next time you see a low score on a leaderboard, ask yourself: Is the model actually dumb, or is the teacher just using the wrong answer key?
