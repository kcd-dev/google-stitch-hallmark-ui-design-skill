# Google Stitch + Hallmark UI Design Skill

一个面向 Codex / AI Agent 的 UI 设计技能：**用 Google Stitch 的方式组织 UI 生成与评审输入，再用 Hallmark 反 AI-slop 质量闸检查输出**，让 AI 生成的界面更接近可落地的产品设计，而不是模板化的“漂亮废图”。

它解决的问题：把模糊 UI 需求、现有页面截图或产品改版目标，整理成可复制给 Stitch 的 screen prompt / critique prompt / redesign prompt，并补上结构多样性、真实文案、响应式、可访问性、设计 token 和前端验收边界。

> Disclaimer: This is an independent community skill. It is not affiliated with, endorsed by, or sponsored by Google, Google Stitch, or any Hallmark-branded company/product. “Hallmark” here refers to a local anti-AI-slop quality-gate method used by this skill.

## 核心价值

```text
User need / existing UI
  → organize input in a Stitch-first screen-design format
  → add anti-AI-slop quality gates
  → generate / critique / redesign in Stitch
  → review output with checklist
  → turn into design.md / frontend implementation / validation
```

- **Stitch first**：先表达页面目标、用户任务、信息层级、组件、状态和约束。
- **Quality gate second**：再约束结构多样性、真实文案、移动端、token、无障碍和反模板化。
- **Implementation ready**：输出能继续转成 `design.md`、组件清单、前端实现计划和验收 checklist。

## 适合谁

- 用 Codex / Claude / Cursor / Grok 做 UI 原型或前端实现的工程师。
- 想让 AI 先生成高保真方向、再落地到前端代码的独立开发者。
- 需要评审现有 UI、避免 AI 默认模板和“看着高级但用不了”的产品团队。

## 快速安装

把本仓库目录复制或克隆到你的 Codex skills 目录：

```bash
git clone git@github.com:kcd-dev/google-stitch-hallmark-ui-design-skill.git \
  ~/.codex/skills/google-stitch-hallmark-ui-design
```

或手动复制：

```bash
mkdir -p ~/.codex/skills
cp -R google-stitch-hallmark-ui-design-skill ~/.codex/skills/google-stitch-hallmark-ui-design
```

然后在 Codex 中提出类似任务即可触发：

```text
用 google-stitch-hallmark-ui-design 帮我把这个 SaaS dashboard 改成更不像 AI 模板的高保真 UI。
```

## 使用前置条件

**不必须先安装或接入 Google Stitch。** 本 skill 可以先以 prompt-only 方式工作：输出可复制到 Stitch 的 screen prompt、critique prompt、design.md 草案和质量检查清单。

| 模式 | 是否需要 Stitch 账号/接入 | 适合场景 |
|---|---:|---|
| Prompt-only | 否 | 先整理需求、生成可复制 prompt、团队手动放进 Stitch |
| Stitch-assisted | 是 | 要真实跑 Stitch、读取 Stitch 输出、或让 agent 通过 MCP/API/浏览器接入 |

如果当前 Codex 会话里 `list_mcp_resources` / `list_mcp_resource_templates` 为空，只能说明**当前会话没有暴露可调用 Stitch 资源**，不能说明 Stitch 产品不可用。此时应降级交付 prompt 和设计规格，并明确“本轮未真实调用 Stitch”。

## 仓库结构

```text
.
├── SKILL.md                         # Skill 主入口
├── agents/openai.yaml               # Skill UI metadata
└── references/prompt-templates.md   # Stitch prompt / critique / review 模板
```

## 典型用法

### 1. 从零设计

```text
用 google-stitch-hallmark-ui-design 设计一个 B2B API 管理后台首页，目标用户是开发者和运营，要求生成可复制到 Stitch 的 prompt，并避免 hero + 3 cards 的 AI 默认布局。
```

### 2. 现有 UI 评审

```text
用 google-stitch-hallmark-ui-design 评审这张页面截图，输出 Stitch critique prompt、P0/P1/P2 改版建议和不要乱改的业务边界。
```

### 3. Stitch 输出复核

```text
用 google-stitch-hallmark-ui-design 检查这份 Stitch 输出是否有 AI slop：假指标、假浏览器框、移动端横滚、token 混乱、状态缺失。
```

### 4. 移动端 App Screen

```text
用 google-stitch-hallmark-ui-design 设计一个移动端任务详情页。用户要立刻看到任务状态、下一步动作、截止时间和失败原因。要求 320/375/414 px 不横向滚动，主 CTA 不换成两行。
```

### 5. 后台设置页 / 表单页

```text
用 google-stitch-hallmark-ui-design 设计一个团队通知设置页，包含通知渠道、接收人、频率、静默时段和测试发送入口。要求覆盖 error、disabled、permission denied、test failed 状态，不改变已有业务流程。
```

### 6. Stitch 产物转前端任务

```text
用 google-stitch-hallmark-ui-design 把这份 Stitch 结果整理成 design.md 和 Vue + TypeScript + Tailwind 前端实现任务，拆组件、状态、token、响应式规则和验收 checklist。
```

### 7. Stitch 接入能力不确定

```text
用 google-stitch-hallmark-ui-design 判断当前能不能直接调用 Stitch。如果 MCP resources 为空，不要说 Stitch 不可用，只说明当前会话没暴露资源，并给 prompt-only 降级方案。
```

## 开源边界

- 本仓库只提供 agent skill、prompt 模板和质量检查流程。
- 不包含 Google Stitch 的账号、API、MCP token 或任何私有凭证。
- 不保证当前环境一定能直接调用 Stitch；如果没有可用 Stitch MCP，仍可输出 Stitch-ready prompt。
- UI 设计结果必须结合项目业务、真实数据、权限边界和前端验收再落地。

## 质量检查重点

- 不默认使用 `hero → 3 feature cards → CTA → footer`。
- 不编造指标、客户 logo、testimonial、可用率或转化率。
- 不画假浏览器栏、假手机壳、假 IDE 外框。
- 320 / 375 / 414 / 768 px 不横向滚动。
- 颜色、字体、间距、focus ring、状态样式可 token 化。
- loading / empty / error / success / disabled / permission denied 状态有设计。
- 输出能转成 `design.md` 和前端组件任务。

## License

MIT
