package hunter

type SpecializedAIDetector struct {
    ZeroDayDetector    *ZeroDayAI
    VulnPredictor     *PredictiveAI
    ExploitGenerator  *ExploitAI
    DorkAnalyzer     *DorkAI
}

type ZeroDayAI struct {
    Patterns []string
    Signatures map[string]float64
    RiskScores map[string]int
}

func NewZeroDayAI() *ZeroDayAI {
    return &ZeroDayAI{
        Patterns: []string{
            `(?i)(prototype\.pollution|prototype\.\_\_proto\_\_)`,
            `(?i)(deserialization\.unsafe|object\.deserialize)`,
            `(?i)(template\.injection|ssti\.bypass)`,
            `(?i)(graphql\.introspection|graphql\.mutation\.unrestricted)`,
            `(?i)(ssrf\.cloud|ssrf\.metadata\.aws)`,
        },
        RiskScores: map[string]int{
            "prototype_pollution": 9,
            "unsafe_deserialization": 10,
            "template_injection": 8,
            "graphql_introspection": 7,
            "ssrf_cloud": 9,
        },
    }
}

