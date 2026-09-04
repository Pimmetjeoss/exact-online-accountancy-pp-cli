// Copyright 2026 pimmetjeoss. Licensed under Apache-2.0. See LICENSE.
// PATCH: Accountancy-specific helpers and workflow commands generated from Exact's official Accountancy docs.
package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	registerNovelCommand(func(root *cobra.Command, flags *rootFlags) {
		var parent *cobra.Command
		for _, sub := range root.Commands() {
			if sub.Name() == "accountancy" {
				parent = sub
				break
			}
		}
		if parent == nil {
			return
		}
		addNovelCommandIfAbsent(parent, newAccountancyResourcesCmd(flags))
		addNovelCommandIfAbsent(parent, newAccountancyDocsCmd(flags))
		addNovelCommandIfAbsent(parent, newAccountancyPracticeSetupCmd(flags))
		addNovelCommandIfAbsent(parent, newAccountancyClientSnapshotCmd(flags))
		addNovelCommandIfAbsent(parent, newAccountancyOwnershipMapCmd(flags))
	})
}

type accountancyODataFlags struct {
	filter       string
	selectClause string
	orderBy      string
	top          int
	skipToken    string
	expand       string
	query        []string
}

func addAccountancyODataFlags(cmd *cobra.Command, f *accountancyODataFlags) {
	cmd.Flags().StringVar(&f.filter, "filter", "", "OData $filter expression, e.g. Account eq guid'...' (sent as $filter)")
	cmd.Flags().StringVar(&f.selectClause, "odata-select", "", "OData $select clause; use root --select for output shaping")
	cmd.Flags().StringVar(&f.orderBy, "orderby", "", "OData $orderby expression")
	cmd.Flags().IntVar(&f.top, "top", 0, "OData $top row limit")
	cmd.Flags().StringVar(&f.skipToken, "skiptoken", "", "OData $skiptoken continuation token")
	cmd.Flags().StringVar(&f.expand, "expand", "", "OData $expand clause")
	cmd.Flags().StringArrayVar(&f.query, "query", nil, "Extra query parameter key=value; may be repeated")
}

func (f accountancyODataFlags) toParams() map[string]string {
	params := map[string]string{}
	if f.filter != "" {
		params["$filter"] = f.filter
	}
	if f.selectClause != "" {
		params["$select"] = f.selectClause
	}
	if f.orderBy != "" {
		params["$orderby"] = f.orderBy
	}
	if f.top > 0 {
		params["$top"] = fmt.Sprint(f.top)
	}
	if f.skipToken != "" {
		params["$skiptoken"] = f.skipToken
	}
	if f.expand != "" {
		params["$expand"] = f.expand
	}
	for _, q := range f.query {
		k, v, ok := strings.Cut(q, "=")
		if ok && strings.TrimSpace(k) != "" {
			params[strings.TrimSpace(k)] = v
		}
	}
	return params
}

type accountancyResourceInfo struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Methods     []string `json:"methods"`
	ReadOnly    bool     `json:"read_only"`
	Docs        string   `json:"docs"`
	Description string   `json:"description"`
}

var accountancyResources = []accountancyResourceInfo{
	{"AccountInvolvedAccounts", "/api/v1/{division}/accountancy/AccountInvolvedAccounts", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountInvolvedAccounts", "Accounts involved for a client account and relation type between accounts."},
	{"AccountOwners", "/api/v1/{division}/accountancy/AccountOwners", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyAccountOwners", "Owners/shareholders for a client account, including percentage shares."},
	{"ClientGroups", "/api/v1/{division}/accountancy/ClientGroups", []string{"GET"}, true, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyClientGroups", "Reference list for classifying clients."},
	{"ClientMainGroups", "/api/v1/{division}/accountancy/ClientMainGroups", []string{"GET"}, true, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyClientMainGroups", "Reference list for classifying client groups."},
	{"InvolvedUserRoles", "/api/v1/{division}/accountancy/InvolvedUserRoles", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUserRoles", "Roles users can have for accountancy clients."},
	{"InvolvedUsers", "/api/v1/{division}/accountancy/InvolvedUsers", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyInvolvedUsers", "Users involved for a client and their roles."},
	{"SolutionLinks", "/api/v1/{division}/accountancy/SolutionLinks", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancySolutionLinks", "Bookkeeping solution linked to a client account."},
	{"TaskTypes", "/api/v1/{division}/accountancy/TaskTypes", []string{"GET", "POST", "PUT", "DELETE"}, false, "https://start.exactonline.nl/docs/HlpRestAPIResourcesDetails.aspx?name=AccountancyTaskTypes", "Custom task types for the accountancy firm."},
}

func accountancyResourceByName(name string) (accountancyResourceInfo, bool) {
	norm := strings.ToLower(strings.ReplaceAll(name, "-", ""))
	for _, r := range accountancyResources {
		if strings.ToLower(r.Name) == strings.ToLower(name) || strings.ToLower(strings.ReplaceAll(r.Name, "-", "")) == norm {
			return r, true
		}
	}
	return accountancyResourceInfo{}, false
}

// pp:data-source local
func newAccountancyResourcesCmd(flags *rootFlags) *cobra.Command {
	var method string
	cmd := &cobra.Command{
		Use:         "resources",
		Short:       "List the covered Exact Online Accountancy resources and methods",
		Example:     "  exact-online-accountancy-pp-cli accountancy resources --json\n  exact-online-accountancy-pp-cli accountancy resources --method GET --json",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			rows := make([]accountancyResourceInfo, 0, len(accountancyResources))
			for _, r := range accountancyResources {
				if method == "" {
					rows = append(rows, r)
					continue
				}
				for _, m := range r.Methods {
					if strings.EqualFold(m, method) {
						rows = append(rows, r)
						break
					}
				}
			}
			data, _ := json.Marshal(rows)
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	cmd.Flags().StringVar(&method, "method", "", "Filter by HTTP method (GET, POST, PUT, DELETE)")
	return cmd
}

// pp:data-source local
func newAccountancyDocsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "docs <resource>",
		Short:       "Show the official Exact docs URL and local endpoint metadata for an Accountancy resource",
		Example:     "  exact-online-accountancy-pp-cli accountancy docs AccountOwners --json",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			r, ok := accountancyResourceByName(args[0])
			if !ok {
				return usageErr(fmt.Errorf("unknown accountancy resource %q", args[0]))
			}
			data, _ := json.Marshal(r)
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	return cmd
}

// pp:data-source live
func newAccountancyPracticeSetupCmd(flags *rootFlags) *cobra.Command {
	var q accountancyODataFlags
	cmd := &cobra.Command{
		Use:         "practice-setup <division>",
		Short:       "Fetch accountancy setup/reference lists: client groups, main groups, user roles, task types",
		Example:     "  exact-online-accountancy-pp-cli accountancy practice-setup 123456 --top 25 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return accountancyMultiGet(cmd, flags, args[0], []string{"ClientGroups", "ClientMainGroups", "InvolvedUserRoles", "TaskTypes"}, q.toParams())
		},
	}
	addAccountancyODataFlags(cmd, &q)
	return cmd
}

// pp:data-source live
func newAccountancyClientSnapshotCmd(flags *rootFlags) *cobra.Command {
	var accountID string
	var q accountancyODataFlags
	cmd := &cobra.Command{
		Use:         "client-snapshot <division>",
		Short:       "Fetch client-related Accountancy records: involved users, solution links, owners, and involved accounts",
		Example:     "  exact-online-accountancy-pp-cli accountancy client-snapshot 123456 --account-id 00000000-0000-0000-0000-000000000000 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			params := q.toParams()
			if accountID != "" && params["$filter"] == "" {
				params["$filter"] = "Account eq guid'" + accountID + "'"
			}
			return accountancyMultiGet(cmd, flags, args[0], []string{"InvolvedUsers", "SolutionLinks", "AccountOwners", "AccountInvolvedAccounts"}, params)
		},
	}
	cmd.Flags().StringVar(&accountID, "account-id", "", "Client account GUID; builds an Account eq guid'...' filter when --filter is absent")
	addAccountancyODataFlags(cmd, &q)
	return cmd
}

// pp:data-source live
func newAccountancyOwnershipMapCmd(flags *rootFlags) *cobra.Command {
	var accountID string
	var q accountancyODataFlags
	cmd := &cobra.Command{
		Use:         "ownership-map <division>",
		Short:       "Read AccountOwners for ownership/shareholder analysis, optionally filtered by client account",
		Example:     "  exact-online-accountancy-pp-cli accountancy ownership-map 123456 --account-id 00000000-0000-0000-0000-000000000000 --agent",
		Annotations: map[string]string{"mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			params := q.toParams()
			if accountID != "" && params["$filter"] == "" {
				params["$filter"] = "Account eq guid'" + accountID + "'"
			}
			return accountancyMultiGet(cmd, flags, args[0], []string{"AccountOwners"}, params)
		},
	}
	cmd.Flags().StringVar(&accountID, "account-id", "", "Client account GUID; builds an Account eq guid'...' filter when --filter is absent")
	addAccountancyODataFlags(cmd, &q)
	return cmd
}

func accountancyMultiGet(cmd *cobra.Command, flags *rootFlags, division string, resourceNames []string, params map[string]string) error {
	c, err := flags.newClient()
	if err != nil {
		return err
	}
	out := map[string]any{"division": division, "resources": map[string]any{}}
	for _, name := range resourceNames {
		r, _ := accountancyResourceByName(name)
		path := replacePathParam(r.Path, "division", division)
		// PATCH(pp4.31.1): migrated to context-aware client + writer-scoped errors.
		data, err := c.Get(cmd.Context(), path, params)
		if err != nil {
			return classifyAPIError(cmd.OutOrStdout(), err, flags)
		}
		var parsed any
		if json.Unmarshal(data, &parsed) != nil {
			parsed = string(data)
		}
		out["resources"].(map[string]any)[name] = map[string]any{"path": path, "docs": r.Docs, "data": parsed}
	}
	data, _ := json.Marshal(out)
	return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
}
