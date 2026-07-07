# Exact Online Accountancy CLI

Accountancy-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Accountancy service resources and all documented methods. OAuth bearer token required.

Printed by [@Pimmetjeoss](https://github.com/Pimmetjeoss) (Pimmetjeoss).

## Install

The recommended path installs both the `exact-online-accountancy-pp-cli` binary and the `pp-exact-online-accountancy` agent skill in one shot:

```bash
npx -y @mvanhorn/printing-press install exact-online-accountancy
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press install exact-online-accountancy --cli-only
```


### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/exact-online-accountancy-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-accountancy --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-exact-online-accountancy --force
```

## Install for OpenClaw

Tell your OpenClaw agent (copy this):

```
Install the pp-exact-online-accountancy skill from https://github.com/mvanhorn/printing-press-library/tree/main/cli-skills/pp-exact-online-accountancy. The skill defines how its required CLI can be installed.
```

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
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value
```

## Usage

Run `exact-online-accountancy-pp-cli --help` for the full command reference and flag list.

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


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value

# JSON for scripting and agents
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value --json

# Filter to specific fields
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value --json --select id,name,status

# Dry run — show the request without sending
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value --dry-run

# Agent mode — JSON + compact + no prompts in one flag
exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries and `--ignore-missing` to delete retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Use with Claude Code

Install the focused skill — it auto-installs the CLI on first invocation:

```bash
npx skills add mvanhorn/printing-press-library/cli-skills/pp-exact-online-accountancy -g
```

Then invoke `/pp-exact-online-accountancy <query>` in Claude Code. The skill is the most efficient path — Claude Code drives the CLI directly without an MCP server in the middle.

<details>
<summary>Use as an MCP server in Claude Code (advanced)</summary>

If you'd rather register this CLI as an MCP server in Claude Code, install the MCP binary first:


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Then register it:

```bash
claude mcp add exact-online-accountancy exact-online-accountancy-pp-mcp -e EXACT_ONLINE_ACCOUNTANCY_OAUTH2=<your-token>
```

</details>

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

## Health Check

```bash
exact-online-accountancy-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/exact-online-accountancy-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `EXACT_ONLINE_ACCOUNTANCY_OAUTH2` | per_call | Yes | Set to your API credential. |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `exact-online-accountancy-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $EXACT_ONLINE_ACCOUNTANCY_OAUTH2`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)

## Local Accountancy Workflow Additions

This build adds Accountancy-specific helpers on top of the generated endpoint mirror:

- `accountancy resources --json` lists the 8 official Accountancy resources and all documented methods (26 operations total).
- `accountancy docs <resource> --json` returns the Exact docs URL, path and supported methods for a resource.
- Every generated `*-get` command supports Exact/OData query flags: `--filter`, `--odata-select`, `--orderby`, `--top`, `--skiptoken`, `--expand`, and repeatable `--query key=value`. Use root `--select` only for output shaping.
- `accountancy practice-setup <division>` fetches `ClientGroups`, `ClientMainGroups`, `InvolvedUserRoles`, and `TaskTypes`.
- `accountancy client-snapshot <division> --account-id <guid>` fetches `InvolvedUsers`, `SolutionLinks`, `AccountOwners`, and `AccountInvolvedAccounts` with an Account filter.
- `accountancy ownership-map <division> --account-id <guid>` fetches shareholder/owner rows from `AccountOwners`.

Examples:

```bash
exact-online-accountancy-pp-cli accountancy resources --json
exact-online-accountancy-pp-cli accountancy docs AccountOwners --json
exact-online-accountancy-pp-cli accountancy account-owners-get 123456 --top 25 --filter "AccountName ne null" --odata-select ID,AccountName,Shares --agent
exact-online-accountancy-pp-cli accountancy practice-setup 123456 --top 25 --agent
exact-online-accountancy-pp-cli accountancy client-snapshot 123456 --account-id 00000000-0000-0000-0000-000000000000 --agent
```

Mutation commands (`*-post`, `*-put`, `*-delete`) call the Exact API only when invoked without `--dry-run`; agents should preview first and only use `--yes --no-input` after target division/body/key semantics are clear.
