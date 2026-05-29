# Changelog

## v0.1.2 - 2026-05-29

- Add explicit logged-in Google Stitch account requirement for Stitch-assisted mode; otherwise downgrade to ordinary Prompt-only mode.
- Add Stitch native export / Figma / Code handoff workflow.
- Clarify that existing Stitch projects should prefer native Export / Share / Code / Figma entries before third-party tools.
- Add usable-first v0.1 handoff rule: if Figma export is not observed, fall back to Code / Share / design.md / frontend task package with explicit boundary.
- Add privacy rule to avoid persisting browser window names, account names, project IDs, session URLs, cookies, tokens, or private prompts.
- Add README usage example and prompt template for native export handoff.
- Add Grok-reviewed browser / BitBrowser lifecycle rules: open only on real Stitch/export tasks, protect login state, close task-opened resources when requested.
- Add Done / Partial / Not accepted validation language, including browser cleanup and export boundary requirements.
- Add confirmed design direction locking for continuing homepage and page-suite work without visual drift.
- Add Stitch Code export post-process Go script for handoff packages and page-suite acceptance matrix.

## v0.1.1 - 2026-05-26

- Add usage prerequisites: prompt-only vs Stitch-assisted modes.
- Clarify that Google Stitch installation/account/MCP is not required for prompt-only usage.
- Add fallback rules when MCP resources/templates are empty.
- Add more usage cases for dashboard, landing page critique, settings forms, mobile app screens, Stitch output review, frontend handoff, and capability checks.
- Add prompt templates for capability checks, mobile app screens, and admin settings/form pages.

## v0.1.0 - 2026-05-26

- Initial open-source release.
- Add Stitch-first UI workflow.
- Add Hallmark-style anti-AI-slop quality gates.
- Add prompt templates for new UI, critique, redesign, review, and design.md conversion.
