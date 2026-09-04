# Exact Online Accountancy CLI

Accountancy-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Accountancy service resources and all documented methods. OAuth bearer token required.

## Install

The recommended path installs both the `exact-online-accountancy-pp-cli` binary and the `pp-exact-online-accountancy` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --agent claude-code
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --agent claude-code --agent codex
```

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-accountancy-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

Install the CLI binary first. The installer writes binaries to a per-user managed bin directory by default: `$HOME/.local/bin` on macOS/Linux and `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows.

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --cli-only
```

Then install the focused Hermes skill.

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-accountancy --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-accountancy --force
```

Restart the Hermes session or gateway if the newly installed skill is not visible immediately.

## Install for OpenClaw
Install both the CLI binary and the focused OpenClaw skill. The installer defaults binaries to a per-user bin directory (`$HOME/.local/bin` on macOS/Linux, `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows):

```bash
npx -y @mvanhorn/printing-press-library install exact-online-accountancy --agent openclaw
```

Restart the OpenClaw session or gateway if the newly installed skill is not visible immediately.

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-accountancy-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `EXACT_ONLINE_ACCOUNTANCY_OAUTH2` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "exact-online-accountancy": {
      "command": "exact-online-accountancy-pp-mcp",
      "env": {
        "EXACT_ONLINE_ACCOUNTANCY_OAUTH2": "<your-key>"
      }
    }
  }
}
```

</details>

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your access token from your API provider's developer portal, then store it:

```bash
exact-online-accountancy-pp-cli auth set-token YOUR_TOKEN_HERE
```

Or set it via environment variable:

```bash
export EXACT_ONLINE_ACCOUNTANCY_OAUTH2="your-token-here"
```

### 3. Verify Setup

```bash
exact-online-accountancy-pp-cli doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value
```

## Usage

Run `exact-online-accountancy-pp-cli --help` for the full command reference and flag list.

## Paths & environment variables

This CLI separates local files into four path kinds:

| Kind | Contents |
|------|----------|
| `config` | User-editable settings such as `config.toml` and saved profiles |
| `data` | Durable local data: `credentials.toml`, `data.db`, cookies, browser-session proof files, and other auth sidecars |
| `state` | Runtime state such as persisted queries, jobs, and `teach.log` |
| `cache` | Regenerable HTTP/cache files |

Each kind resolves independently. The ladder is:

1. Per-kind env var: `EXACT_ONLINE_ACCOUNTANCY_CONFIG_DIR`, `EXACT_ONLINE_ACCOUNTANCY_DATA_DIR`, `EXACT_ONLINE_ACCOUNTANCY_STATE_DIR`, or `EXACT_ONLINE_ACCOUNTANCY_CACHE_DIR`
2. `--home <dir>` for this invocation
3. `EXACT_ONLINE_ACCOUNTANCY_HOME` for a flat relocated root
4. XDG env vars: `XDG_CONFIG_HOME`, `XDG_DATA_HOME`, `XDG_STATE_HOME`, `XDG_CACHE_HOME`
5. Platform defaults matching existing installs

For containers and agent sandboxes, prefer a single relocated root:

```bash
export EXACT_ONLINE_ACCOUNTANCY_HOME=/srv/exact-online-accountancy
exact-online-accountancy-pp-cli doctor
```

Under `EXACT_ONLINE_ACCOUNTANCY_HOME=/srv/exact-online-accountancy`, the four dirs resolve to `/srv/exact-online-accountancy/config`, `/srv/exact-online-accountancy/data`, `/srv/exact-online-accountancy/state`, and `/srv/exact-online-accountancy/cache`.

MCP servers do not receive CLI flags from the host. Put relocation in the host `env` block:

```json
{
  "mcpServers": {
    "exact-online-accountancy": {
      "command": "exact-online-accountancy-pp-mcp",
      "env": {
        "EXACT_ONLINE_ACCOUNTANCY_HOME": "/srv/exact-online-accountancy"
      }
    }
  }
}
```

Precedence matters in fleets: an ambient per-kind variable such as `EXACT_ONLINE_ACCOUNTANCY_DATA_DIR` overrides an explicit `--home` for that kind. Use `EXACT_ONLINE_ACCOUNTANCY_HOME` or the per-kind variables for durable fleet relocation; treat `--home` as the weaker per-invocation lever.

Relocation is one-way. Unsetting `EXACT_ONLINE_ACCOUNTANCY_HOME` does not move files back to platform defaults, and `doctor` cannot find credentials left under a former root. Move the files manually before unsetting relocation variables.

Existing installs keep working because the platform-default rung matches the legacy layout. On the first auth write, stored secrets leave `config.toml` and are consolidated into `credentials.toml` under the data directory. Run `exact-online-accountancy-pp-cli doctor --fail-on warn` to check path and credential-location warnings in automation.

## Commands

### accountancy

Manage accountancy

- **`exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete`** - This resource will show all accounts which are involved for a specific account and will also show the type of the relation between the accounts. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountInvolvedAccounts
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-involved-accounts-get`** - This resource will show all accounts which are involved for a specific account and will also show the type of the relation between the accounts. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountInvolvedAccounts
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-involved-accounts-post`** - This resource will show all accounts which are involved for a specific account and will also show the type of the relation between the accounts. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountInvolvedAccounts
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-involved-accounts-put`** - This resource will show all accounts which are involved for a specific account and will also show the type of the relation between the accounts. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountInvolvedAccounts
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-owners-delete`** - The account owners are accounts which represents the owners or shareholders for a specific account. In this resource also the percentage of shares which the account holds are stored. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountOwners
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-owners-get`** - The account owners are accounts which represents the owners or shareholders for a specific account. In this resource also the percentage of shares which the account holds are stored. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountOwners
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-owners-post`** - The account owners are accounts which represents the owners or shareholders for a specific account. In this resource also the percentage of shares which the account holds are stored. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountOwners
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy account-owners-put`** - The account owners are accounts which represents the owners or shareholders for a specific account. In this resource also the percentage of shares which the account holds are stored. This functionality is only available in a practice management company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountOwners
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy client-groups-get`** - The client groups are used to classify the clients.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyClientGroups
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy client-main-groups-get`** - The client main groups are used to classify the client groups.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyClientMainGroups
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-user-roles-delete`** - An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be used to indicate the role that a user have for a certain client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-user-roles-get`** - An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be used to indicate the role that a user have for a certain client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-user-roles-post`** - An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be used to indicate the role that a user have for a certain client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-user-roles-put`** - An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be used to indicate the role that a user have for a certain client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-users-delete`** - The involved users keeps track on all users which are involved for a certain client of the accountant. The involved user also shows the role that the user has for the involved client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-users-get`** - The involved users keeps track on all users which are involved for a certain client of the accountant. The involved user also shows the role that the user has for the involved client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-users-post`** - The involved users keeps track on all users which are involved for a certain client of the accountant. The involved user also shows the role that the user has for the involved client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy involved-users-put`** - The involved users keeps track on all users which are involved for a certain client of the accountant. The involved user also shows the role that the user has for the involved client. This functionality is only available in a practice company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy solution-links-delete`** - Solution links are used to store which bookkeeping solution is used for a certain account. This can be the bookkeeping solution of Exact Online which is mentioned as ‘internal’ or a bookkeeping solution of another vender which is mentioned as ‘external’.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy solution-links-get`** - Solution links are used to store which bookkeeping solution is used for a certain account. This can be the bookkeeping solution of Exact Online which is mentioned as ‘internal’ or a bookkeeping solution of another vender which is mentioned as ‘external’.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy solution-links-post`** - Solution links are used to store which bookkeeping solution is used for a certain account. This can be the bookkeeping solution of Exact Online which is mentioned as ‘internal’ or a bookkeeping solution of another vender which is mentioned as ‘external’.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy solution-links-put`** - Solution links are used to store which bookkeeping solution is used for a certain account. This can be the bookkeeping solution of Exact Online which is mentioned as ‘internal’ or a bookkeeping solution of another vender which is mentioned as ‘external’.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy task-types-delete`** - The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types defined within a company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyTaskTypes
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy task-types-get`** - The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types defined within a company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyTaskTypes
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy task-types-post`** - The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types defined within a company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyTaskTypes
Scope: Accountancy practicemanagement
- **`exact-online-accountancy-pp-cli accountancy task-types-put`** - The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types defined within a company.

Official docs: https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyTaskTypes
Scope: Accountancy practicemanagement


### Self-learning loop

This CLI caches per-question discovery so repeat queries skip the walk and structurally similar queries get answered via entity substitution. The loop also self-captures: every invocation is journaled locally, and failed-flag corrections plus fresh teaches surface as candidates on the next `recall` for confirm/reject judgment. Agents call `recall` before discovery and fire `teach &` after answering. See the `## Automatic learning` section in `SKILL.md` for the full protocol.

- **`exact-online-accountancy-pp-cli recall <query>`** - Look up cached resources for a query before running discovery
- **`exact-online-accountancy-pp-cli teach`** - Record a query -> resource mapping (silent on success, safe to background with `&`)
- **`exact-online-accountancy-pp-cli learnings list`** - Inspect taught rows
- **`exact-online-accountancy-pp-cli learnings forget <query>`** - Undo a teach
- **`exact-online-accountancy-pp-cli learnings candidates`** - List auto-captured candidates awaiting confirm/reject
- **`exact-online-accountancy-pp-cli learnings stats`** - Local loop metrics: recall hit rate, teach-to-reuse, playbook resolution, candidate counts
- **`exact-online-accountancy-pp-cli teach-pattern`** - Install a query/resource template up front
- **`exact-online-accountancy-pp-cli teach-lookup`** - Add an entity mapping (e.g. country code, team alias) for pattern substitution

Pass `--no-learn` or set `EXACT_ONLINE_ACCOUNTANCY_NO_LEARN=true` to disable the loop for deterministic flows.

The local store's schema version stamp is one-way: once this version of `exact-online-accountancy-pp-cli` opens the database, older binaries refuse it with a version error — upgrade the binary rather than downgrading.

## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value

# JSON for scripting and agents
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value --json
# Filter to specific fields by name
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value --json --select <field>[,<field>...]

# Dry run — show the request without sending
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
exact-online-accountancy-pp-cli accountancy account-involved-accounts-get mock-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select <field>[,<field>...]` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries and add `--ignore-missing` to delete retries when a no-op success is acceptable
- **Explicit confirmation** - `--agent` does not imply `--yes`; pass `--yes` separately only after the target, arguments, and side effects are clear
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Health Check

```bash
exact-online-accountancy-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Run `exact-online-accountancy-pp-cli doctor` to see the resolved config, data, state, and cache directories. The platform-default config path is `~/.config/exact-online-accountancy-pp-cli/config.toml`; `--home`, `EXACT_ONLINE_ACCOUNTANCY_HOME`, and per-kind env vars can relocate it.

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `EXACT_ONLINE_ACCOUNTANCY_OAUTH2` | per_call | Yes | Set to your API credential. |

### agentcookie (optional)

If you use agentcookie to sync secrets across machines, this CLI auto-adopts agentcookie-managed credentials with no extra setup. When the daemon writes to this CLI's config, `exact-online-accountancy-pp-cli doctor` reports `agentcookie: detected` and `auth-status` labels the source as `agentcookie`. Skip this section if you don't use agentcookie - the CLI works the same as any other.

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `exact-online-accountancy-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $EXACT_ONLINE_ACCOUNTANCY_OAUTH2`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
