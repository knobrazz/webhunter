package ai

type AIPatternDetector struct {
    Patterns map[string][]DetectionPattern
}

type DetectionPattern struct {
    Name        string
    Severity    string
    Pattern     string
    Context     map[string]string
}

func NewAIPatternDetector() *AIPatternDetector {
    return &AIPatternDetector{
        Patterns: map[string][]DetectionPattern{
            "cloud_misconfig": {
                {Name: "AWS Key Exposure", Pattern: `(?i)(AKIA[0-9A-Z]{16})`},
                {Name: "Azure Connection", Pattern: `(?i)(DefaultEndpointsProtocol=https?;AccountName=[^;]+;AccountKey=[^;]+)`},
                {Name: "GCP Credentials", Pattern: `(?i)("type": "service_account".*"private_key")`},
            },
            "api_vulnerabilities": {
                {Name: "GraphQL Introspection", Pattern: `(?i)(__schema|__type|__typename)`},
                {Name: "API Key in URL", Pattern: `(?i)(api[_-]?key|access[_-]?token)=([a-zA-Z0-9-_.]+)`},
                {Name: "JWT Weakness", Pattern: `(?i)(alg":"\s*none|alg":"\s*HS256)`},
            },
            "zero_day_patterns": {
                {Name: "Log4j Injection", Pattern: `(?i)(\$\{jndi:ldap|rmi|dns)`},
                {Name: "Spring4Shell", Pattern: `(?i)(class\.module\.classLoader)`},
                {Name: "Prototype Pollution", Pattern: `(?i)(__proto__|constructor\.prototype)`},
            },
        },
    }
}

