package filter

type FilterEngine struct {
    Patterns     *PatternMatcher
    Rules        []FilterRule
    Processors   []Processor
}

type FilterRule struct {
    Name         string
    Priority     int
    Conditions   []Condition
    Action       string
}

type Processor struct {
    Type         string
    Config       map[string]interface{}
    Handler      func([]byte) ([]byte, error)
}

func (f *FilterEngine) AddSecurityRules() {
    f.Rules = append(f.Rules, []FilterRule{
        {
            Name: "Authentication Bypass",
            Priority: 10,
            Conditions: []Condition{
                {Pattern: `(?i)(auth.*bypass|admin.*login|\.\.\/admin)`},
            },
        },
        {
            Name: "Information Disclosure",
            Priority: 8,
            Conditions: []Condition{
                {Pattern: `(?i)(phpinfo|debug=|show.*errors|\.git)`},
            },
        },
        {
            Name: "Security Misconfiguration",
            Priority: 7,
            Conditions: []Condition{
                {Pattern: `(?i)(config\.|\.env|\.conf|\.ini)`},
            },
        },
    }...)
}

