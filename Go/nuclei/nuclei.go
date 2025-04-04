package nuclei

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "path/filepath"
)

type NucleiScanner struct {
    OutputDir string
    Templates []string
}

func NewNucleiScanner(outputDir string) *NucleiScanner {
    return &NucleiScanner{
        OutputDir: outputDir,
        Templates: []string{
            "cves", "vulnerabilities", "misconfiguration",
            "exposures", "technologies", "takeovers",
        },
    }
}

func (s *NucleiScanner) ScanTarget(target string) error {
    outputFile := filepath.Join(s.OutputDir, "nuclei", "scan-results.json")
    
    args := []string{
        "-u", target,
        "-t", filepath.Join("nuclei-templates", "*"),
        "-severity", "critical,high,medium",
        "-json",
        "-o", outputFile,
    }

    cmd := exec.Command("nuclei", args...)
    return cmd.Run()
}

func (s *NucleiScanner) UpdateTemplates() error {
    cmd := exec.Command("nuclei", "-ut")
    return cmd.Run()
}

