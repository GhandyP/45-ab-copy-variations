# 45 — A/B copy variations

## Goal

Create distinct copy variants with an experiment hypothesis from a supplied campaign brief without launching or measuring an experiment automatically.

## Run the current prototype

From this case directory:

```bash
export GOOGLE_API_KEY="your-key"
go run .
```

`internal/adkhelper.Run` uses ADK Go v2.2.0, Gemini, and `gemini-flash-latest`; set `ADK_MODEL` to override the model.

Sample prompt:

```text
Create five distinct subject-line variants and a hypothesis for each.
```

## What exists now

- `main.go (exists)` sends campaign text to one focused Gemini ADK agent.
- The response is free-form; variant IDs, distinctness, character limits, hypothesis, audience, and control assignment are not validated.
- This is intentionally text-only: it does not connect to an experiment platform, scheduler, audience store, or analytics system.

## Missing elements

- **Contract:** Define `ExperimentBrief{channel, audience, objective, control}` -> `VariantSet{variants[{id,copy,hypothesis}], guardrails, successMetric}`.
- **ADK layer:** Add a variant schema and campaign session; keep an optional `get_brand_rules` function tool read-only and exclude launch tools.
- **Domain boundary:** Add brand-rule, deduplication, experiment-config, audience, and analytics adapters.
- **Fixtures and validation:** Cover duplicate wording, limits, forbidden claims, control preservation, locale, and too-few-variant cases.
- **Safety:** Require marketing approval before exposing variants to customers; do not infer experiment results from model output.

## Target structure

```text
45-ab-copy-variations/
├── main.go (exists)
├── README.md (exists)
├── schema.go (next)
├── variants.go (next)
├── adapters/experiment.go (next)
├── testdata/brief.json (next)
└── main_test.go (next)
```

## Implementation order

1. Define variant, control, hypothesis, metric, and guardrail schemas.
2. Add deterministic deduplication and channel-length validation.
3. Generate variants within a campaign session and preserve control text.
4. Test duplicates, locale, limits, forbidden claims, and count requirements.
5. Add a dry-run experiment configuration only after human approval.

## Done when

- [ ] Every variant is distinct, labeled, within limits, and tied to a hypothesis.
- [ ] Control, audience, and success metric are preserved exactly.
- [ ] Fixtures reject duplicates and prohibited copy.
- [ ] Launching an experiment requires a reviewed configuration and human approval.

## Next improvement

Implement variant identity, deduplication, and hypothesis validation before integrating an experiment platform.
