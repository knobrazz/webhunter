package attack

type PatternMatcher struct {
    VulnPatterns    []VulnPattern
    CustomPatterns  map[string]Pattern
    Signatures      []Signature
}

type VulnPattern struct {
    Name        string
    Regex       string
    Risk        int
    CVE         []string
    Payloads    []string
}

type Pattern struct {
    Type        string
    Indicators  []string
    Context     map[string]string
    Confidence  float64
}

func NewPatternMatcher() *PatternMatcher {
    return &PatternMatcher{
        VulnPatterns: []VulnPattern{
            {
                Name: "SQL Injection",
                Regex: `(?i)(union.*select|select.*from|select.*where|drop.*table|--.*$)`,
                Risk: 9,
            },
            {
                Name: "XSS",
                Regex: `(?i)(<script.*>|javascript:|onload=|onerror=)`,
                Risk: 8,
            },
            {
                Name: "File Inclusion",
                Regex: `(?i)(\.\.\/|file:\/\/|php:\/\/|data:\/\/)`,
                Risk: 8,
            },
            {
                Name: "Command Injection",
                Regex: `(?i)([;&|`+"`"+`].*\$?\(|\bping\b|\bcat\b|\bbash\b)`,
                Risk: 9,
            },
            {
                Name: "SSRF",
                Regex: `(?i)(localhost|127\.0\.0\.1|0\.0\.0\.0|internal\.)`,
                Risk: 7,
            },
        },
    }
}

