package dorks

type DorkEngine struct {
    Patterns    map[string][]string
    Analyzers   []DorkAnalyzer
    Results     chan DorkResult
}

func NewDorkEngine() *DorkEngine {
    return &DorkEngine{
        Patterns: map[string][]string{
            "critical": {
                `site:{domain} ext:sql "INSERT INTO" -"keep-alive"`,
                `site:{domain} ext:log username password`,
                `site:{domain} inurl:admin intitle:"index of /"`,
                `site:{domain} ext:env OR ext:yml OR ext:yaml OR ext:config`,
                `site:{domain} intext:"api_key" OR intext:"client_secret"`,
            },
            "high": {
                `site:{domain} ext:php intitle:"phpinfo()"`,
                `site:{domain} inurl:wp-config.php`,
                `site:{domain} inurl:/.git`,
                `site:{domain} filetype:bak OR filetype:backup`,
                `site:{domain} inurl:"/phpMyAdmin/"`,
            },
            "medium": {
                `site:{domain} ext:action OR ext:struts`,
                `site:{domain} inurl:error OR inurl:debug`,
                `site:{domain} "Index of /" +.svn`,
                `site:{domain} intext:"warning" AND "error"`,
                `site:{domain} inurl:signup OR inurl:register`,
            },
        },
    }
}

