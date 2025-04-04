package filter

type SmartFilter struct {
    Rules       []FilterRule
    Processors  []Processor
    Cache       *Cache
}

type FilterRule struct {
    Name        string
    Pattern     string
    Context     map[string]string
    Action      string
    Priority    int
}

func NewSmartFilter() *SmartFilter {
    filter := &SmartFilter{}
    filter.LoadRules()
    return filter
}

func (f *SmartFilter) LoadRules() {
    f.Rules = []FilterRule{
        {
            Name: "API Key Exposure",
            Pattern: `(?i)(api[_-]?key|access[_-]?token|secret[_-]?key)[\s]*[=:]\s*['"][0-9a-zA-Z-_\.]+['"]`,
            Priority: 10,
        },
        {
            Name: "Sensitive File Access",
            Pattern: `(?i)((\/etc\/passwd|\/etc\/shadow|web\.config|\.env|\.git\/config))`,
            Priority: 9,
        },
        {
            Name: "Server Information Disclosure",
            Pattern: `(?i)(phpinfo\(\)|server-status|server-info|[\/]\.git)`,
            Priority: 8,
        },
    }
}

