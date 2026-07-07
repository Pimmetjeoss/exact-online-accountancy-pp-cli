---
name: pp-exact-online-accountancy
description: "Printing Press CLI for Exact Online Accountancy. Accountancy-only Exact Online REST/OData API surface generated from the official Exact Online REST API..."
author: "Pimmetjeoss"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - exact-online-accountancy-pp-cli
---

# Exact Online Accountancy — Printing Press CLI

## Prerequisites: Install the CLI

This skill drives the `exact-online-accountancy-pp-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. Install via the Printing Press installer:
   ```bash
   npx -y @mvanhorn/printing-press install exact-online-accountancy --cli-only
   ```
2. Verify: `exact-online-accountancy-pp-cli --version`
3. Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on `$PATH`.

If the `npx` install fails before this CLI has a public-library category, install Node or use the category-specific Go fallback after publish.

If `--version` reports "command not found" after install, the install step did not put the binary on `$PATH`. Do not proceed with skill commands until verification succeeds.

Accountancy-only Exact Online REST/OData API surface generated from the official Exact Online REST API documentation. Covers the Accountancy service resources and all documented methods. OAuth bearer token required.

## Command Reference

**accountancy** — Manage accountancy

- `exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete` — This resource will show all accounts which are involved for a specific account and will also show the type of the...
- `exact-online-accountancy-pp-cli accountancy account-involved-accounts-get` — This resource will show all accounts which are involved for a specific account and will also show the type of the...
- `exact-online-accountancy-pp-cli accountancy account-involved-accounts-post` — This resource will show all accounts which are involved for a specific account and will also show the type of the...
- `exact-online-accountancy-pp-cli accountancy account-involved-accounts-put` — This resource will show all accounts which are involved for a specific account and will also show the type of the...
- `exact-online-accountancy-pp-cli accountancy account-owners-delete` — The account owners are accounts which represents the owners or shareholders for a specific account. In this resource...
- `exact-online-accountancy-pp-cli accountancy account-owners-get` — The account owners are accounts which represents the owners or shareholders for a specific account. In this resource...
- `exact-online-accountancy-pp-cli accountancy account-owners-post` — The account owners are accounts which represents the owners or shareholders for a specific account. In this resource...
- `exact-online-accountancy-pp-cli accountancy account-owners-put` — The account owners are accounts which represents the owners or shareholders for a specific account. In this resource...
- `exact-online-accountancy-pp-cli accountancy client-groups-get` — The client groups are used to classify the clients. Official docs:...
- `exact-online-accountancy-pp-cli accountancy client-main-groups-get` — The client main groups are used to classify the client groups. Official docs:...
- `exact-online-accountancy-pp-cli accountancy involved-user-roles-delete` — An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be...
- `exact-online-accountancy-pp-cli accountancy involved-user-roles-get` — An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be...
- `exact-online-accountancy-pp-cli accountancy involved-user-roles-post` — An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be...
- `exact-online-accountancy-pp-cli accountancy involved-user-roles-put` — An involved user role represents a role which a user can have in an accountancy firm. This involved user role can be...
- `exact-online-accountancy-pp-cli accountancy involved-users-delete` — The involved users keeps track on all users which are involved for a certain client of the accountant. The involved...
- `exact-online-accountancy-pp-cli accountancy involved-users-get` — The involved users keeps track on all users which are involved for a certain client of the accountant. The involved...
- `exact-online-accountancy-pp-cli accountancy involved-users-post` — The involved users keeps track on all users which are involved for a certain client of the accountant. The involved...
- `exact-online-accountancy-pp-cli accountancy involved-users-put` — The involved users keeps track on all users which are involved for a certain client of the accountant. The involved...
- `exact-online-accountancy-pp-cli accountancy solution-links-delete` — Solution links are used to store which bookkeeping solution is used for a certain account. This can be the...
- `exact-online-accountancy-pp-cli accountancy solution-links-get` — Solution links are used to store which bookkeeping solution is used for a certain account. This can be the...
- `exact-online-accountancy-pp-cli accountancy solution-links-post` — Solution links are used to store which bookkeeping solution is used for a certain account. This can be the...
- `exact-online-accountancy-pp-cli accountancy solution-links-put` — Solution links are used to store which bookkeeping solution is used for a certain account. This can be the...
- `exact-online-accountancy-pp-cli accountancy task-types-delete` — The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types...
- `exact-online-accountancy-pp-cli accountancy task-types-get` — The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types...
- `exact-online-accountancy-pp-cli accountancy task-types-post` — The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types...
- `exact-online-accountancy-pp-cli accountancy task-types-put` — The task types are specific types defined by the user of the accountancy firm.This resource shows the custom types...


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
exact-online-accountancy-pp-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Auth Setup

Run `exact-online-accountancy-pp-cli auth setup` for the URL and steps to obtain a token (add `--launch` to open the URL). Then store it:

```bash
exact-online-accountancy-pp-cli auth set-token YOUR_TOKEN_HERE
```

Or set `EXACT_ONLINE_ACCOUNTANCY_OAUTH2` as an environment variable.

Run `exact-online-accountancy-pp-cli doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  exact-online-accountancy-pp-cli accountancy account-involved-accounts-delete mock-value --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success, and `--ignore-missing` only when a missing delete target should count as success

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
exact-online-accountancy-pp-cli feedback "the --since flag is inclusive but docs say exclusive"
exact-online-accountancy-pp-cli feedback --stdin < notes.txt
exact-online-accountancy-pp-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.exact-online-accountancy-pp-cli/feedback.jsonl`. They are never POSTed unless `EXACT_ONLINE_ACCOUNTANCY_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `EXACT_ONLINE_ACCOUNTANCY_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration - HeyGen's "Beacon" pattern.

```
exact-online-accountancy-pp-cli profile save briefing --json
exact-online-accountancy-pp-cli --profile briefing accountancy account-involved-accounts-delete mock-value
exact-online-accountancy-pp-cli profile list --json
exact-online-accountancy-pp-cli profile show briefing
exact-online-accountancy-pp-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `exact-online-accountancy-pp-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary from this CLI's published public-library entry or pre-built release, then register it:

```bash
claude mcp add exact-online-accountancy-pp-mcp -- exact-online-accountancy-pp-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which exact-online-accountancy-pp-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   exact-online-accountancy-pp-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `exact-online-accountancy-pp-cli <command> --help`.

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
