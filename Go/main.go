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

package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sync"
    "time"
    "github.com/yourusername/webhunter/subdomain"
)

func banner() {
    fmt.Println(`
    ██╗    ██╗███████╗██████╗       ██╗  ██╗██╗   ██╗███╗   ██╗████████╗███████╗██████╗ 
    ██║    ██║██╔════╝██╔══██╗      ██║  ██║██║   ██║████╗  ██║╚══██╔══╝██╔════╝██╔══██╗
    ██║ █╗ ██║█████╗  ██████╔╝█████╗███████║██║   ██║██╔██╗ ██║   ██║   █████╗  ██████╔╝
    ██║███╗██║██╔══╝  ██╔══██╗╚════╝██╔══██║██║   ██║██║╚██╗██║   ██║   ██╔══╝  ██╔══██╗
    ╚███╔███╔╝███████╗██████╔╝      ██║  ██║╚██████╔╝██║ ╚████║   ██║   ███████╗██║  ██║
     ╚══╝╚══╝ ╚══════╝╚═════╝       ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
                                Created by Nabaraj Lamichhane
    `)
    fmt.Println("\nCAUTION: Use it for ethical purposes only. If your government does not recommend these types of tools,")
    fmt.Println("avoid using them. Try at your own risk.\n")
}

func main() {
    banner()

    config := parseFlags()
    setupOutputDir(config)

    // Initialize logger
    logFile := filepath.Join(config.OutputDir, "web-hunter.log")
    f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    defer f.Close()
    log.SetOutput(f)

    // Start reconnaissance process
    startRecon(config)
}

func parseFlags() *Config {
    config := &Config{}
    
    // Existing flags
    flag.StringVar(&config.Domain, "d", "", "Target domain")
    flag.StringVar(&config.WildcardDomain, "w", "", "Wildcard domain")
    flag.StringVar(&config.IPAddress, "ip", "", "IP address")
    flag.StringVar(&config.CIDR, "cidr", "", "CIDR range")
    flag.StringVar(&config.WildcardList, "wl", "", "List of wildcard domains")
    flag.StringVar(&config.SubdomainList, "list", "", "Subdomain list")
    
    // New flags
    flag.BoolVar(&config.StealthMode, "stealth", false, "Enable stealth mode")
    flag.BoolVar(&config.AIAnalysis, "ai", false, "Enable AI analysis")
    flag.StringVar(&config.ResumeFile, "resume", "", "Resume from checkpoint file")
    flag.IntVar(&config.Threads, "threads", 10, "Number of concurrent threads")
    flag.DurationVar(&config.Timeout, "timeout", 30*time.Second, "Timeout for requests")
    flag.StringVar(&config.ProxyURL, "proxy", "", "Proxy URL")
    flag.IntVar(&config.RateLimit, "rate", 0, "Rate limit requests per second")
    flag.StringVar(&config.OutputFormat, "output", "json", "Output format (json|csv|md)")
    
    // New AI flags
    flag.StringVar(&config.AIModel, "ai-model", "default", "AI model to use for analysis")
    flag.StringVar(&config.AIAPIKey, "ai-key", "", "API key for AI service")
    flag.BoolVar(&config.ComplianceCheck, "compliance", false, "Enable compliance checking")

    // New stealth flags
    flag.BoolVar(&config.RandomizePaths, "random-paths", false, "Randomize scan paths")
    flag.BoolVar(&config.TrafficObfuscation, "obfuscate", false, "Enable traffic obfuscation")
    flag.IntVar(&config.JitterPercent, "jitter", 0, "Add random delays (percentage)")
    flag.StringVar(&config.CustomUA, "user-agent", "", "Custom user agent string")

    // New output flags
    flag.StringVar(&config.OutputFormat, "format", "json", "Output format (json|yaml|xml|html|md|pdf|excel)")
    flag.StringVar(&config.ReportTemplate, "template", "", "Custom report template file")
    
    // Advanced AI flags
    flag.StringVar(&config.ThreatIntelAPI, "threat-intel", "", "Threat Intelligence API key")
    flag.StringVar(&config.MLModel, "ml-model", "default", "Machine Learning model for pattern recognition")
    flag.Float64Var(&config.AIConfidence, "ai-confidence", 0.8, "AI analysis confidence threshold")

    // Advanced stealth flags
    flag.BoolVar(&config.EmulateUser, "emulate-user", false, "Enable user behavior emulation")
    flag.StringVar(&config.BrowserProfile, "browser-profile", "", "Custom browser fingerprint")
    flag.BoolVar(&config.NetworkDeception, "deception", false, "Enable network deception")

    // Advanced visualization flags
    flag.StringVar(&config.ChartType, "charts", "all", "Chart types to generate (comma-separated)")
    flag.StringVar(&config.VisualTheme, "theme", "dark", "Visualization theme")
    flag.StringVar(&config.CustomTemplate, "custom-template", "", "Path to custom report template")
    
    // Advanced AI model flags
    flag.StringVar(&config.DeepLearningModel, "dl-model", "", "Deep learning model path")
    flag.StringVar(&config.NLPModel, "nlp-model", "", "NLP model configuration")
    flag.Float64Var(&config.AnomalyThreshold, "anomaly-threshold", 0.95, "Anomaly detection threshold")

    // Advanced AI model flags
    flag.StringVar(&config.TransformerModel, "transformer", "", "Transformer model configuration")
    flag.StringVar(&config.GNNModel, "gnn", "", "Graph Neural Network model")
    flag.BoolVar(&config.EnableRL, "rl", false, "Enable Reinforcement Learning")
    flag.BoolVar(&config.FederatedLearning, "federated", false, "Enable Federated Learning")

    // Advanced protocol flags
    flag.BoolVar(&config.HTTP2Enabled, "http2", false, "Enable HTTP/2 manipulation")
    flag.BoolVar(&config.QUICEnabled, "quic", false, "Enable QUIC protocol")
    flag.StringVar(&config.DNSResolver, "dns-resolver", "", "Custom DNS resolver")

    // Advanced visualization flags
    flag.BoolVar(&config.Enable3D, "3d", false, "Enable 3D visualization")
    flag.StringVar(&config.D3Template, "d3-template", "", "D3.js template file")
    flag.BoolVar(&config.ReactiveMode, "reactive", false, "Enable reactive visualizations")

    // Specialized AI flags
    flag.BoolVar(&config.ZeroShotEnabled, "zero-shot", false, "Enable zero-shot learning")
    flag.BoolVar(&config.FewShotEnabled, "few-shot", false, "Enable few-shot learning")
    flag.BoolVar(&config.MetaLearning, "meta-learning", false, "Enable meta-learning")
    flag.StringVar(&config.AutoMLConfig, "automl", "", "AutoML configuration file")

    // Advanced protocol flags
    flag.BoolVar(&config.WSEnabled, "websocket", false, "Enable WebSocket manipulation")
    flag.BoolVar(&config.GRPCEnabled, "grpc", false, "Enable gRPC handling")
    flag.BoolVar(&config.MQTTEnabled, "mqtt", false, "Enable MQTT manipulation")
    flag.StringVar(&config.IPSecConfig, "ipsec", "", "IPSec configuration file")

    // Advanced visualization flags
    flag.BoolVar(&config.WebGLEnabled, "webgl", false, "Enable WebGL rendering")
    flag.BoolVar(&config.SVGAnimations, "svg-anim", false, "Enable SVG animations")
    flag.StringVar(&config.FlowLayout, "flow-layout", "force", "Flow visualization layout")

    // Advanced AI technique flags
    flag.BoolVar(&config.TransferLearning, "transfer", false, "Enable transfer learning")
    flag.BoolVar(&config.ActiveLearning, "active", false, "Enable active learning")
    flag.BoolVar(&config.OnlineLearning, "online", false, "Enable online learning")
    flag.Float64Var(&config.AdaptiveRate, "adapt-rate", 0.01, "Adaptive learning rate")

    // Protocol security flags
    flag.StringVar(&config.TLSConfig, "tls-config", "", "TLS hardening configuration")
    flag.StringVar(&config.FilterRules, "filter-rules", "", "Protocol filter rules")
    flag.BoolVar(&config.EnableAnalyzer, "analyzer", false, "Enable traffic analyzer")

    // Advanced visualization flags
    flag.BoolVar(&config.NetworkGraph, "network-graph", false, "Enable network graph")
    flag.BoolVar(&config.Timeline, "timeline", false, "Enable timeline visualization")
    flag.StringVar(&config.HeatmapConfig, "heatmap", "", "Heatmap configuration")

    // Attack feature flags
    flag.StringVar(&config.FuzzConfig, "fuzz", "", "Fuzzing configuration")
    flag.StringVar(&config.InjectionVectors, "vectors", "", "Injection vectors file")
    flag.StringVar(&config.BypassTechniques, "bypass", "", "Bypass techniques file")

    // Advanced attack flags
    flag.StringVar(&config.PatternFile, "patterns", "", "Custom vulnerability patterns file")
    flag.StringVar(&config.FilterRules, "filter-rules", "", "Custom filter rules file")
    flag.BoolVar(&config.DeepScan, "deep-scan", false, "Enable deep scanning")
    flag.IntVar(&config.RiskThreshold, "risk-threshold", 7, "Minimum risk level to report")
    
    // Bug hunting flags
    flag.BoolVar(&config.AutoExploit, "auto-exploit", false, "Enable automatic exploitation")
    flag.StringVar(&config.ExploitDB, "exploit-db", "", "Custom exploit database")
    flag.BoolVar(&config.FuzzAPIs, "fuzz-apis", false, "Enable API fuzzing")
    flag.StringVar(&config.CustomPayloads, "payloads", "", "Custom payload file")

    // Advanced vulnerability scanning flags
    flag.StringVar(&config.VulnDB, "vuln-db", "", "Custom vulnerability database")
    flag.Float64Var(&config.MinCVSS, "min-cvss", 7.0, "Minimum CVSS score")
    flag.BoolVar(&config.ActiveScan, "active-scan", false, "Enable active vulnerability scanning")
    flag.StringVar(&config.ScanProfile, "scan-profile", "normal", "Scan profile (light|normal|aggressive)")

    // Smart filtering flags
    flag.StringVar(&config.FilterProfile, "filter-profile", "standard", "Filter profile (standard|strict|custom)")
    flag.StringVar(&config.CustomRules, "custom-rules", "", "Custom filtering rules file")
    flag.BoolVar(&config.EnableCache, "enable-cache", false, "Enable filter caching")
    flag.IntVar(&config.CacheSize, "cache-size", 1000, "Filter cache size")

    // Advanced bug hunting flags
    flag.BoolVar(&config.APIFuzz, "api-fuzz", false, "Enable API fuzzing")
    flag.BoolVar(&config.CloudScan, "cloud-scan", false, "Enable cloud resource scanning")
    flag.BoolVar(&config.AuthTest, "auth-test", false, "Enable authentication testing")
    flag.StringVar(&config.CustomPatterns, "custom-patterns", "", "Custom pattern file")

    // AI Bug Hunter flags
    flag.BoolVar(&config.ZeroDayAI, "zero-day-ai", false, "Enable Zero-Day AI detection")
    flag.BoolVar(&config.PatternAI, "pattern-ai", false, "Enable Pattern Learning AI")
    flag.BoolVar(&config.BehaviorAI, "behavior-ai", false, "Enable Behavior Analysis AI")
    flag.Float64Var(&config.AIThreshold, "ai-threshold", 0.85, "AI detection threshold")
    
    // Advanced Bug Detection flags
    flag.BoolVar(&config.CloudVulnScan, "cloud-vuln", false, "Enable cloud vulnerability scanning")
    flag.BoolVar(&config.APISecurityScan, "api-security", false, "Enable API security scanning")
    flag.BoolVar(&config.AuthBypassTest, "auth-bypass", false, "Enable authentication bypass testing")
    flag.StringVar(&config.CustomAIModel, "custom-ai", "", "Custom AI model configuration")

    // Advanced AI Detection flags
    flag.StringVar(&config.AIPatterns, "ai-patterns", "", "Custom AI detection patterns file")
    flag.Float64Var(&config.PatternConfidence, "pattern-conf", 0.85, "Pattern matching confidence")
    flag.BoolVar(&config.ZeroDayDetection, "zero-day", false, "Enable zero-day pattern detection")
    
    // Enhanced Google Dorks flags
    flag.StringVar(&config.DorkCategories, "dork-cats", "all", "Dork categories (sensitive|panels|api|cloud|security|all)")
    flag.BoolVar(&config.CustomDorks, "custom-dorks", false, "Enable custom dork patterns")
    flag.IntVar(&config.DorkThreads, "dork-threads", 5, "Number of dork scanning threads")
    
    // Advanced Exploitation flags
    flag.StringVar(&config.ExploitMode, "exploit-mode", "safe", "Exploitation mode (safe|aggressive)")
    flag.StringVar(&config.CustomPayloads, "custom-payloads", "", "Custom payload file")
    flag.BoolVar(&config.AutoValidate, "auto-validate", false, "Enable automatic exploit validation")
    flag.Float64Var(&config.ExploitSuccess, "exploit-success", 0.75, "Exploit success threshold")

    flag.Parse()
    return config
}

func startRecon(config *Config) {
    // Initialize progress tracking
    var progress *resume.Progress
    if config.ResumeFile != "" {
        progress = &resume.Progress{
            CompletedTasks: make(map[string]bool),
        }
        if err := progress.Load(config.ResumeFile); err != nil {
            log.Printf("Could not load resume file: %v", err)
        }
    }

    // Initialize stealth mode if enabled
    var stealthScanner *stealth.StealthScanner
    if config.StealthMode {
        stealthScanner = &stealth.StealthScanner{
            Config: stealth.StealthConfig{
                RateLimit: config.RateLimit,
                ProxyRotation: config.ProxyURL != "",
            },
        }
        stealthScanner.SetupClient()
    }

    // Initialize all scanners
    subdomainScanner := subdomain.NewScanner(config.OutputDir)
    ipScanner := ipscanner.NewIPScanner(config.OutputDir)
    endpointScanner := endpoint.NewEndpointScanner(config.OutputDir)
    vulnScanner := vulnscanner.NewVulnScanner(config.OutputDir)
    nucleiScanner := nuclei.NewNucleiScanner(config.OutputDir)
    
    // Initialize webhook client if configured
    var webhookClient *webhook.WebhookClient
    if config.SlackWebhook != "" || config.DiscordWebhook != "" {
        webhookClient = &webhook.WebhookClient{
            SlackURL:   config.SlackWebhook,
            DiscordURL: config.DiscordWebhook,
        }
    }

    // Handle different input types
    switch {
    case config.SubdomainList != "":
        if err := copyFile(config.SubdomainList, 
            filepath.Join(config.OutputDir, "subdomains", "valid_subdomains.txt")); err != nil {
            log.Fatalf("Failed to copy subdomain list: %v", err)
        }
    
    case config.Domain != "":
        if err := subdomainScanner.ScanDomain(config.Domain); err != nil {
            log.Printf("Error scanning domain %s: %v", config.Domain, err)
        }
        
        // Scan endpoints after subdomain discovery
        if err := endpointScanner.ScanEndpoints(config.Domain); err != nil {
            log.Printf("Error scanning endpoints for %s: %v", config.Domain, err)
        }
    
    case config.IPAddress != "":
        if err := ipScanner.ScanIP(config.IPAddress); err != nil {
            log.Printf("Error scanning IP %s: %v", config.IPAddress, err)
        }
    
    case config.CIDR != "":
        if err := ipScanner.ScanCIDR(config.CIDR); err != nil {
            log.Printf("Error scanning CIDR %s: %v", config.CIDR, err)
        }
    }

    // Run vulnerability scans
    if err := vulnScanner.ScanTarget(target); err != nil {
        log.Printf("Error in vulnerability scan: %v", err)
    }

    // Run Nuclei scans
    if err := nucleiScanner.ScanTarget(target); err != nil {
        log.Printf("Error in Nuclei scan: %v", err)
    }

    // Run AI analysis if enabled
    if config.AIAnalysis {
        analyzer := &aianalyzer.AIAnalyzer{
            OutputDir: config.OutputDir,
        }
        analysis, err := analyzer.AnalyzeResults()
        if err != nil {
            log.Printf("Error in AI analysis: %v", err)
        } else {
            // Save AI analysis results
            analysisFile := filepath.Join(config.OutputDir, "ai_analysis.json")
            if data, err := json.MarshalIndent(analysis, "", "  "); err == nil {
                os.WriteFile(analysisFile, data, 0644)
            }
        }
    }

    // Save progress if resume is enabled
    if progress != nil {
        if err := progress.Save(config.ResumeFile); err != nil {
            log.Printf("Error saving progress: %v", err)
        }
    }

    // Send notifications if webhook is configured
    if webhookClient != nil {
        // Send scan completion notification
        message := map[string]interface{}{
            "text": fmt.Sprintf("Scan completed for target: %s", target),
            // Add more details as needed
        }
        if err := webhookClient.SendToSlack(message); err != nil {
            log.Printf("Error sending Slack notification: %v", err)
        }
    }

    // Initialize AI Hunter if enabled
    if config.ZeroDayAI || config.PatternAI || config.BehaviorAI {
        aiHunter := hunter.NewAIHunter()
        if analysis, err := aiHunter.AnalyzeTarget(target); err != nil {
            log.Printf("Error in AI hunting: %v", err)
        } else {
            // Save AI hunting results
            aiResultsFile := filepath.Join(config.OutputDir, "ai_hunting_results.json")
            if data, err := json.MarshalIndent(analysis, "", "  "); err == nil {
                os.WriteFile(aiResultsFile, data, 0644)
            }
        }
    }

    // Initialize Google Dorks scanner if enabled
    if config.AutoDork {
        dorkEngine := dorks.NewDorkEngine()
        if results, err := dorkEngine.ScanTarget(target); err != nil {
            log.Printf("Error in Google dorks scanning: %v", err)
        } else {
            // Save dorks results
            dorksFile := filepath.Join(config.OutputDir, "dorks_results.json")
            if data, err := json.MarshalIndent(results, "", "  "); err == nil {
                os.WriteFile(dorksFile, data, 0644)
            }
        }
    }
}

func copyFile(src, dst string) error {
    input, err := os.ReadFile(src)
    if err != nil {
        return err
    }
    return os.WriteFile(dst, input, 0644)

