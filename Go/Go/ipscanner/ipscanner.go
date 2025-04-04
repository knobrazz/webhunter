package ipscanner

import (
    "fmt"
    "log"
    "os/exec"
    "path/filepath"
    "sync"
)

type IPScanner struct {
    OutputDir  string
    WaitGroup  sync.WaitGroup
}

func NewIPScanner(outputDir string) *IPScanner {
    return &IPScanner{
        OutputDir: outputDir,
    }
}

func (s *IPScanner) ScanIP(ip string) error {
    outputFile := filepath.Join(s.OutputDir, "ip_scan", "ports.txt")
    
    // Run Nmap for port scanning
    cmd := exec.Command("nmap", "-sS", "-sV", "-p-", "-T4", ip, "-oN", outputFile)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("nmap scan failed: %v", err)
    }

    // Reverse DNS lookup
    return s.reverseDNSLookup(ip)
}

func (s *IPScanner) ScanCIDR(cidr string) error {
    outputFile := filepath.Join(s.OutputDir, "ip_scan", "cidr_hosts.txt")
    
    // Run masscan for fast port scanning
    cmd := exec.Command("masscan", cidr, "-p1-65535", "--rate=10000", "-oL", outputFile)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("masscan failed: %v", err)
    }

    return nil
}

func (s *IPScanner) reverseDNSLookup(ip string) error {
    outputFile := filepath.Join(s.OutputDir, "ip_scan", "reverse_dns.txt")
    
    cmd := exec.Command("dig", "-x", ip, "+short")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("reverse DNS lookup failed: %v", err)
    }

    // Save results
    return s.saveResults(outputFile, string(output))
}

func (s *IPScanner) saveResults(file, data string) error {
    // Implementation for saving results
    return nil
}

