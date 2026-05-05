---
title: "Structured Reasoning on a Diet: My Experiment With Llama-3.2-3B"
date: "2025-09-28"
summary: "Can a tiny 3B model actually think? My experiment with training Llama-3.2-3B-thinking-8k-v1 to perform structured reasoning, inspired by DeepSeek's R1 approach."
tags: ["AI", "Machine Learning", "LLM", "Reasoning", "Llama"]
---


So… I trained (well, duct-taped) a small reasoning model. It's called Llama-3.2-3B-thinking-8k-v1, and yes, the name is longer than its context window deserves.

## Why?

Everyone's chasing bigger and bigger models, but I wondered: can a tiny one actually think a little? For years, the assumption was that reasoning was an emergent property of scale, something you only get once you cross 100B parameters. That view is now being challenged. Recent surveys show that small language models (SLMs) can, in fact, be trained to reason effectively when given structured training or post-training compression, rather than just brute-force scale (Zhang et al., 2025).

## What?

This model is inspired by DeepSeek's R1, a model famous for its verbose, sometimes spiraling, almost neurotic inner monologues. Call it "productive overthinking." DeepSeek-R1's ability to narrate, self-correct, and occasionally derail is basically Tree-of-Thoughts in action, even if it looks like drama at first glance (DeepSeek-AI et al., 2025, AI et al., 2025, Yao et al., 2023).

I wanted my model to capture some of that. And yes, sometimes it mimics R1's overthinking behavior in ways that are hilarious coming from a 3B model. But I definitely didn't want to fall back into the old "Let's think step-by-step" trope from the early days of reasoning models (Wei et al., 2022). That approach was groundbreaking back then, but it's basically linear counting, and it breaks down when you need exploration, backtracking, or multiple solution paths. Tree-of-Thoughts and its derivatives were created for exactly this reason: reasoning that branches, explores, and self-evaluates instead of plodding down a single line (Yao et al., 2023, Princeton NLP, 2023, IBM Research, 2024).

So, this project is my small-scale attempt at stepping away from "kindergarten CoT" and toward structured, exploratory reasoning.

## How?

1. **Training data**: 8k samples of DeepSeek-R1's reasoning traces. Not just answers, but full "thought monologues." This is essentially knowledge distillation, transferring the reasoning patterns of a big "teacher" into a smaller "student" (Xu et al., 2024, Park, 2023, Biernacki, 2022).

2. **Method**: Supervised Fine-Tuning (SFT). In DeepSeek's own pipeline, this "cold start" SFT stage was essential to give the model baseline readability and reasoning patterns before reinforcement learning kicked in (DeepSeek-AI et al., 2025). So while my setup looks duct-taped, it's actually following a principled, two-stage approach.

3. **Hardware**: One glorious, rented H100 from DigitalOcean. Yes, my wallet cried.

4. **Next step**: Group Relative Policy Optimization (GRPO) - the same reinforcement learning method that powered DeepSeek-R1. GRPO is critic-free, sample-efficient, and proven effective at incentivizing long-form reasoning (DeepSeek-AI et al., 2025, Chen et al., 2025, Li et al., 2025). If SFT is cosplay, GRPO is what turns mimicry into actual capability.

## Where it stands

Right now, this is a preliminary experiment. I haven't benchmarked it yet, so don't expect MMLU charts, ARC scores, or bragging rights. I uploaded it to Hugging Face before the GPU fans even cooled down. But early impressions? Small models can be nudged toward structured reasoning, even if they occasionally roleplay as "DeepSeek-R1 on caffeine." It's slower, sure, but also sharper.

## Conclusion

This isn't the model that dethrones GPT-4. It's not even the model that dethrones Wolfram Alpha. But it is a data point in the growing evidence that reasoning is not just about scale - it's about the training pipeline, the data, and the structure you force the model into.

If you enjoy poking at half-baked but fun experiments: [here it is](https://huggingface.co/x1nx3r/Llama-3.2-3B-thinking-8k-v1)

## References

- [Zhang, C. et al. (2025). Towards Reasoning Ability of Small Language Models](https://arxiv.org/abs/2502.11569)
- [Wei, J. et al. (2022). Chain of Thought Prompting Elicits Reasoning in Large Language Models](https://arxiv.org/abs/2201.11903)
- [Yao, S. et al. (2023). Tree of Thoughts: Deliberate Problem Solving with LLMs](https://arxiv.org/abs/2305.10601)
- [Princeton NLP (2023). Tree-of-Thoughts LLM](https://github.com/princeton-nlp/tree-of-thought-llm)
- [IBM Research (2024). What is Tree of Thoughts Prompting?](https://research.ibm.com/blog/tree-of-thoughts-prompting)
- [Xu, X. et al. (2024). A Survey on Knowledge Distillation of Large Language Models](https://arxiv.org/abs/2402.13116)
- [Park, S. (2023). Unraveling Knowledge Distillation](https://medium.com/@sungwonpark/unraveling-knowledge-distillation-3de8f8d7a8c6)
- [Biernacki, M. (2022). Self-Instruct: Aligning LMs with Self-Generated Instructions](https://arxiv.org/abs/2212.10560)
- [DeepSeek-AI et al. (2025). DeepSeek-R1: Incentivizing Reasoning Capability via Reinforcement Learning](https://arxiv.org/abs/2501.12948)
- [Chen, L. et al. (2025). Multi-Layer GRPO: Enhancing Reasoning and Self-Correction in LLMs](https://arxiv.org/abs/2506.04746)
- [Li, Y. et al. (2025). On the Theory and Practice of GRPO: A Trajectory-Corrected Approach with Fast Convergence](https://arxiv.org/abs/2508.02833)