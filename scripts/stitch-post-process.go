// stitch-post-process creates a small frontend handoff package from a Stitch Code export.
// It is intentionally dependency-free so it can run with: go run scripts/stitch-post-process.go --input exported.html --output handover --ref references/design-direction-template.md
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	input := flag.String("input", "", "Stitch exported HTML file")
	output := flag.String("output", "handover", "Output directory for handoff files")
	ref := flag.String("ref", "", "Optional design direction reference markdown")
	flag.Parse()

	if *input == "" {
		fatal("missing --input")
	}
	b, err := os.ReadFile(*input)
	if err != nil {
		fatal("read input: %v", err)
	}
	content := string(b)
	if err := os.MkdirAll(*output, 0o755); err != nil {
		fatal("create output: %v", err)
	}

	colors := uniqueMatches(content, regexp.MustCompile(`(?i)#(?:[0-9a-f]{3}|[0-9a-f]{6}|[0-9a-f]{8})\b|rgba?\([^)]{3,80}\)`), 40)
	title := firstSubmatch(content, regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`))
	texts := visibleTextSamples(content, 18)
	warnings := privacyWarnings(content)
	hash := sha256.Sum256(b)

	artifactName := filepath.Base(*input)
	cleanPath := filepath.Join(*output, artifactName)
	if err := os.WriteFile(cleanPath, b, 0o644); err != nil {
		fatal("write copied artifact: %v", err)
	}

	var refNote string
	if *ref != "" {
		if rb, err := os.ReadFile(*ref); err == nil {
			refNote = strings.TrimSpace(string(rb))
		}
	}

	handover := buildHandover(artifactName, title, hex.EncodeToString(hash[:])[:16], colors, texts, warnings, refNote)
	if err := os.WriteFile(filepath.Join(*output, "handover.md"), []byte(handover), 0o644); err != nil {
		fatal("write handover: %v", err)
	}

	fmt.Printf("handover written: %s\n", filepath.Join(*output, "handover.md"))
	if len(warnings) > 0 {
		fmt.Printf("privacy warnings: %d; review handover before sharing\n", len(warnings))
	}
}

func buildHandover(artifact, title, shortHash string, colors, texts, warnings []string, ref string) string {
	var sb strings.Builder
	sb.WriteString("# Stitch Export Handoff\n\n")
	sb.WriteString("测试日期：" + time.Now().Format("2006-01-02") + "\n")
	sb.WriteString("文档版本：v0.1\n\n")
	sb.WriteString("## Artifact\n\n")
	sb.WriteString("- Source file: `" + artifact + "`\n")
	if title != "" {
		sb.WriteString("- HTML title: " + title + "\n")
	}
	sb.WriteString("- SHA256 short: `" + shortHash + "`\n")
	sb.WriteString("- Source: Stitch Code export\n\n")

	sb.WriteString("## Detected visual tokens\n\n")
	if len(colors) == 0 {
		sb.WriteString("- No explicit color tokens detected. Review CSS manually.\n\n")
	} else {
		for _, c := range colors {
			sb.WriteString("- `" + c + "`\n")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## Visible copy samples\n\n")
	for _, t := range texts {
		sb.WriteString("- " + t + "\n")
	}
	if len(texts) == 0 {
		sb.WriteString("- No text samples extracted.\n")
	}
	sb.WriteString("\n")

	sb.WriteString("## Page suite acceptance matrix\n\n")
	sb.WriteString("| Page | Status | Visual continuity | Real task entry | Notes |\n")
	sb.WriteString("|---|---|---|---|---|\n")
	sb.WriteString("| Homepage | Pending | TBD | TBD | Generate next using confirmed design direction |\n")
	sb.WriteString("| Registration | Done / Review | Yes if generated from confirmed Stitch output | Yes if CTA/form present | Replace captcha placeholder with real component during implementation |\n")
	sb.WriteString("| Login | Pending | TBD | TBD | Reuse auth card pattern |\n")
	sb.WriteString("| Forgot password | Pending | TBD | TBD | Reuse form state patterns |\n")
	sb.WriteString("| Dashboard empty state | Pending | TBD | TBD | Reuse illustration/token style |\n\n")

	sb.WriteString("## Privacy scan\n\n")
	if len(warnings) == 0 {
		sb.WriteString("- No obvious token/cookie/session patterns detected by the lightweight scanner. Manual review is still required before publishing.\n\n")
	} else {
		for _, w := range warnings {
			sb.WriteString("- WARNING: " + w + "\n")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## Implementation notes\n\n")
	sb.WriteString("- Treat this export as design input, not production-ready app code.\n")
	sb.WriteString("- Map colors, spacing, radii, shadows, and focus states into the project design tokens.\n")
	sb.WriteString("- Replace placeholders such as captcha/security widgets with real project components.\n")
	sb.WriteString("- Validate responsive widths at 320, 375, 414, and 768 px.\n")
	sb.WriteString("- Do not publish private Stitch URLs, account names, browser IDs, cookies, tokens, or private prompts.\n\n")

	if ref != "" {
		sb.WriteString("## Design direction reference\n\n")
		sb.WriteString(ref + "\n")
	}
	return sb.String()
}

func uniqueMatches(s string, re *regexp.Regexp, limit int) []string {
	seen := map[string]bool{}
	var out []string
	for _, m := range re.FindAllString(s, -1) {
		m = strings.TrimSpace(m)
		if m == "" || seen[m] {
			continue
		}
		seen[m] = true
		out = append(out, m)
		if len(out) >= limit {
			break
		}
	}
	sort.Strings(out)
	return out
}

func firstSubmatch(s string, re *regexp.Regexp) string {
	m := re.FindStringSubmatch(s)
	if len(m) < 2 {
		return ""
	}
	return strings.TrimSpace(html.UnescapeString(stripTags(m[1])))
}

func visibleTextSamples(s string, limit int) []string {
	s = regexp.MustCompile(`(?is)<script.*?</script>|<style.*?</style>`).ReplaceAllString(s, " ")
	parts := regexp.MustCompile(`(?is)>\s*([^<>]{2,80})\s*<`).FindAllStringSubmatch(s, -1)
	seen := map[string]bool{}
	var out []string
	for _, p := range parts {
		t := strings.TrimSpace(html.UnescapeString(p[1]))
		t = regexp.MustCompile(`\s+`).ReplaceAllString(t, " ")
		if len([]rune(t)) < 2 || seen[t] {
			continue
		}
		seen[t] = true
		out = append(out, t)
		if len(out) >= limit {
			break
		}
	}
	return out
}

func privacyWarnings(s string) []string {
	checks := map[string]*regexp.Regexp{
		"possible API/token/secret keyword": regexp.MustCompile(`(?i)(api[_-]?key|access[_-]?token|refresh[_-]?token|secret|cookie|authorization)\s*[:=]`),
		"possible private session URL":      regexp.MustCompile(`(?i)(session|sid|authuser|oauth|token)=`),
		"possible long bearer-like value":   regexp.MustCompile(`(?i)bearer\s+[a-z0-9._~+/=-]{20,}`),
	}
	var out []string
	for name, re := range checks {
		if re.MatchString(s) {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

func stripTags(s string) string {
	return regexp.MustCompile(`(?is)<[^>]+>`).ReplaceAllString(s, "")
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "错误："+format+"\n", args...)
	os.Exit(1)
}
