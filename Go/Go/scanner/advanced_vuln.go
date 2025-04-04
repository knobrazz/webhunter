package scanner

type VulnScanner struct {
    Patterns     []VulnPattern
    Detectors    []Detector
    Validators   []Validator
}

type VulnPattern struct {
    ID          string
    Name        string
    Description string
    CVSS        float64
    Patterns    []string
    Payloads    []string
    Validation  string
}

func NewVulnScanner() *VulnScanner {
    scanner := &VulnScanner{}
    scanner.LoadPatterns()
    return scanner
}

func (s *VulnScanner) LoadPatterns() {
    s.Patterns = []VulnPattern{
        {
            ID: "VULN-001",
            Name: "GraphQL Introspection",
            Patterns: []string{
                `{"query":"\{__schema\{types\{name\}\}\}"}`,
                `{"query":"{__type(name:\"Query\"){fields{name}}}"}`,
            },
            CVSS: 7.5,
        },
        {
            ID: "VULN-002",
            Name: "JWT Vulnerabilities",
            Patterns: []string{
                `eyJ[a-zA-Z0-9_-]*\.eyJ[a-zA-Z0-9_-]*\.[a-zA-Z0-9_-]*`,
                `{"alg":"none"}`,
            },
            CVSS: 8.0,
        },
        {
            ID: "VULN-003",
            Name: "Cloud Misconfigurations",
            Patterns: []string{
                `(?i)(aws_access_key|azure_subscription|gcp_credentials)`,
                `(?i)(s3\.amazonaws\.com|storage\.googleapis\.com)`,
            },
            CVSS: 8.5,
        },
    }
}

