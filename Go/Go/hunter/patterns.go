package hunter

type BugPattern struct {
    Name        string
    Category    string
    Severity    int
    AIScore     float64
    Patterns    []string
}

func LoadAdvancedPatterns() []BugPattern {
    return []BugPattern{
        {
            Name: "GraphQL Injection",
            Category: "API",
            Severity: 9,
            Patterns: []string{
                `introspection.*{.*__schema`,
                `mutation.*{.*delete`,
                `query.*{.*all.*}`,
            },
        },
        {
            Name: "Cloud Misconfiguration",
            Category: "Cloud",
            Severity: 8,
            Patterns: []string{
                `(?i)(aws_access|azure_key|gcp_credential)`,
                `(?i)(bucket.*public|container.*access)`,
            },
        },
        {
            Name: "Authentication Bypass",
            Category: "Auth",
            Severity: 9,
            Patterns: []string{
                `(?i)(jwt.*none|cookie.*secure.*false)`,
                `(?i)(auth.*bypass|admin.*override)`,
            },
        },
        {
            Name: "Zero-Day Patterns",
            Category: "ZeroDay",
            Severity: 10,
            Patterns: []string{
                `(?i)(debug.*true|test.*mode|dev.*enabled)`,
                `(?i)(internal.*api|private.*endpoint)`,
            },
        },
    }
}

