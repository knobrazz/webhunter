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

package endpoint

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "sync"
)

type EndpointScanner struct {
    OutputDir     string
    WaitGroup     sync.WaitGroup
    ResultChan    chan string
}

func NewEndpointScanner(outputDir string) *EndpointScanner {
    return &EndpointScanner{
        OutputDir:  outputDir,
        ResultChan: make(chan string, 1000),
    }
}

func (s *EndpointScanner) ScanEndpoints(target string) error {
    tools := []string{"waybackurls", "katana", "gau", "hakrawler"}
    
    for _, tool := range tools {
        s.WaitGroup.Add(1)
        go s.runTool(tool, target)
    }

    // Collect and process results
    go s.processResults()

    s.WaitGroup.Wait()
    close(s.ResultChan)
    
    return s.filterEndpoints()
}

func (s *EndpointScanner) runTool(tool, target string) {
    defer s.WaitGroup.Done()

    var cmd *exec.Cmd
    switch tool {
    case "waybackurls":
        cmd = exec.Command("waybackurls", target)
    case "katana":
        cmd = exec.Command("katana", "-u", target)
    case "gau":
        cmd = exec.Command("gau", target)
    case "hakrawler":
        cmd = exec.Command("hakrawler", "-url", target)
    }

    output, err := cmd.Output()
    if err != nil {
        log.Printf("Error running %s: %v", tool, err)
        return
    }

    s.ResultChan <- string(output)
}

func (s *EndpointScanner) filterEndpoints() error {
    // Filter out unwanted extensions
    unwantedExt := regexp.MustCompile(`\.(jpg|jpeg|png|gif|css|js|woff|woff2|ttf|svg)$`)
    
    inputFile := filepath.Join(s.OutputDir, "endpoints", "all-endpoints.txt")
    outputFile := filepath.Join(s.OutputDir, "endpoints", "filtered-endpoints.txt")
    
    input, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer input.Close()

    output, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer output.Close()

    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)

    for scanner.Scan() {
        line := scanner.Text()
        if !unwantedExt.MatchString(line) {
            writer.WriteString(line + "\n")
        }
    }

    return writer.Flush()
}

func (s *EndpointScanner) FindCriticalEndpoints() error {
    // Look for sensitive endpoints
    patterns := []string{
        "api", "admin", "config", "backup", "dev",
        ".env", ".git", "wp-admin", "jenkins",
        "database", "auth", "login", "test",
    }

    inputFile := filepath.Join(s.OutputDir, "endpoints", "filtered-endpoints.txt")
    outputFile := filepath.Join(s.OutputDir, "endpoints", "critical-endpoints.txt")

    // Implementation for finding critical endpoints
    return nil
}

