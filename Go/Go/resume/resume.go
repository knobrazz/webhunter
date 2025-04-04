package resume

import (
    "encoding/json"
    "os"
    "path/filepath"
    "sync"
)

type Progress struct {
    CompletedTasks map[string]bool `json:"completed_tasks"`
    CurrentPhase   string          `json:"current_phase"`
    LastTarget     string          `json:"last_target"`
    mutex          sync.Mutex
}

func (p *Progress) Save(file string) error {
    p.mutex.Lock()
    defer p.mutex.Unlock()

    data, err := json.MarshalIndent(p, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(file, data, 0644)
}

func (p *Progress) Load(file string) error {
    data, err := os.ReadFile(file)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, p)
}

