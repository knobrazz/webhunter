package vulnscanner

import (
    "encoding/json"
    "fmt"
    "log"
    "os/exec"
    "path/filepath"
    "sync"
)

type VulnScanner struct {
    OutputDir string
    WaitGroup sync.WaitGroup
}

type Vulnerability struct {
    Tool        string      `json:"tool"`
    Severity    string      `json:"severity"`
    Description string      `json:"description"`
    URL         string      `json:"url"`
    POC         string      `json:"poc,omitempty"`
}

func NewVulnScanner(outputDir string) *VulnScanner {
    return &VulnScanner{
        OutputDir: outputDir,
    }
}

func (s *VulnScanner) ScanTarget(target string) error {
    tools := []string{"nuclei", "dalfox", "ghauri", "nikto"}
    
    for _, tool := range tools {
        s.WaitGroup.Add(1)
        go s.runTool(tool, target)
    }

    s.WaitGroup.Wait()
    return s.generateReport()
}

func (s *VulnScanner) runTool(tool, target string) {
    defer s.WaitGroup.Done()
    
    outputFile := filepath.Join(s.OutputDir, "vulnerabilities", tool+"-output.json")
    
    var cmd *exec.Cmd
    switch tool {
    case "nuclei":
        cmd = exec.Command("nuclei", "-u", target, "-json", "-o", outputFile)
    case "dalfox":
        cmd = exec.Command("dalfox", "url", target, "--json-output", outputFile)
    case "nikto":
        cmd = exec.Command("nikto", "-h", target, "-Format", "json", "-o", outputFile)
    }

    if err := cmd.Run(); err != nil {
        log.Printf("Error running %s: %v", tool, err)
    }
}

func (s *VulnScanner) generateReport() error {
    // Implementation for generating consolidated report
    return nil
}

