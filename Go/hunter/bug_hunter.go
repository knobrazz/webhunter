// Copyright 2025 nabar
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hunter

type BugHunter struct {
    Patterns     *attack.PatternMatcher
    Filters      *filter.FilterEngine
    Analyzers    []Analyzer
}

type Analyzer struct {
    Name         string
    Patterns     []string
    Context      map[string]interface{}
    Risk         int
}

func (b *BugHunter) AddCommonVulnPatterns() {
    b.Analyzers = append(b.Analyzers, []Analyzer{
        {
            Name: "Broken Access Control",
            Patterns: []string{
                `(?i)(admin.*panel|dashboard\?|/account/.*id=)`,
                `(?i)(/api/.*token=|/auth/.*key=)`,
            },
            Risk: 8,
        },
        {
            Name: "Sensitive Data Exposure",
            Patterns: []string{
                `(?i)(password|secret|token|key).*=`,
                `(?i)(/backup|/dump|/logs|/temp)`,
            },
            Risk: 9,
        },
        {
            Name: "Security Misconfigurations",
            Patterns: []string{
                `(?i)(test|dev|stage|uat)\.`,
                `(?i)(/phpinfo|/server-status|/admin)`,
            },
            Risk: 7,
        },
        {
            Name: "API Vulnerabilities",
            Patterns: []string{
                `(?i)(/api/v[0-9]+/.*id=|/graphql)`,
                `(?i)(/swagger|/api-docs|/openapi)`,
            },
            Risk: 8,
        },
    }...)
}

