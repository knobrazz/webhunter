package hunter

type SmartHunter struct {
    Scanners    []Scanner
    Analyzers   []Analyzer
    Exploiters  []Exploiter
}

type Scanner struct {
    Name        string
    Type        string
    Patterns    []string
    Confidence  float64
}

func NewSmartHunter() *SmartHunter {
    hunter := &SmartHunter{}
    hunter.LoadScanners()
    return hunter
}

func (h *SmartHunter) LoadScanners() {
    h.Scanners = []Scanner{
        {
            Name: "API Security Scanner",
            Type: "api",
            Patterns: []string{
                `(?i)(api\/v[0-9]+\/|graphql|swagger|openapi)`,
                `(?i)(\/rest\/|\/soap\/|\/grpc\/)`,
            },
        },
        {
            Name: "Cloud Resource Scanner",
            Type: "cloud",
            Patterns: []string{
                `(?i)(s3\.|blob\.|storage\.|bucket\.)`,
                `(?i)(aws\.|azure\.|gcp\.|digitalocean\.)`,
            },
        },
        {
            Name: "Authentication Scanner",
            Type: "auth",
            Patterns: []string{
                `(?i)(\/login|\/auth|\/oauth|\/sso)`,
                `(?i)(\/reset|\/forgot|\/recover|\/token)`,
            },
        },
    }
}

