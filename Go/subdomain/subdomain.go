package subdomain

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "sync"
)

type SubdomainScanner struct {
    OutputDir     string
    ResultChan    chan string
    WaitGroup     sync.WaitGroup
    Tools         []string
    ValidDomains  []string
}

func NewScanner(outputDir string) *SubdomainScanner {
    return &SubdomainScanner{
        OutputDir:  outputDir,
        ResultChan: make(chan string, 1000),
        Tools: []string{
            "amass",
            "subfinder",
            "assetfinder",
            "findomain",
            "knockpy",
            "massdns",
            "shuffledns",
        },
    }
}

func (s *SubdomainScanner) ScanDomain(domain string) error {
    outputFile := filepath.Join(s.OutputDir, "subdomains", "raw_subdomains.txt")
    
    // Start goroutines for each tool
    for _, tool := range s.Tools {
        s.WaitGroup.Add(1)
        go s.runTool(tool, domain)
    }

    // Start result collector
    go s.collectResults(outputFile)

    // Wait for all scans to complete
    s.WaitGroup.Wait()
    close(s.ResultChan)

    // Validate subdomains
    return s.validateSubdomains(outputFile)
}

func (s *SubdomainScanner) runTool(tool, domain string) {
    defer s.WaitGroup.Done()

    var cmd *exec.Cmd
    switch tool {
    case "amass":
        cmd = exec.Command("amass", "enum", "-d", domain)
    case "subfinder":
        cmd = exec.Command("subfinder", "-d", domain)
    case "assetfinder":
        cmd = exec.Command("assetfinder", domain)
    // Add more tools as needed
    }

    output, err := cmd.Output()
    if err != nil {
        log.Printf("Error running %s: %v", tool, err)
        return
    }

    // Process and send results
    for _, line := range strings.Split(string(output), "\n") {
        if line = strings.TrimSpace(line); line != "" {
            s.ResultChan <- line
        }
    }
}

func (s *SubdomainScanner) collectResults(outputFile string) {
    f, err := os.Create(outputFile)
    if err != nil {
        log.Fatalf("Failed to create output file: %v", err)
    }
    defer f.Close()

    writer := bufio.NewWriter(f)
    seen := make(map[string]bool)

    for result := range s.ResultChan {
        if !seen[result] {
            seen[result] = true
            writer.WriteString(result + "\n")
        }
    }
    writer.Flush()
}

func (s *SubdomainScanner) validateSubdomains(inputFile string) error {
    outputFile := filepath.Join(s.OutputDir, "subdomains", "valid_subdomains.txt")
    
    // Use httpx to validate subdomains
    cmd := exec.Command("httpx", "-l", inputFile, "-o", outputFile)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("error validating subdomains: %v", err)
    }

    return nil
}

