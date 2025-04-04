package output

import (
    "encoding/json"
    "encoding/xml"
    "gopkg.in/yaml.v2"
    "html/template"
)

type OutputManager struct {
    Format     string
    OutputDir  string
    Templates  map[string]*template.Template
}

func (o *OutputManager) GenerateReport(data interface{}) error {
    switch o.Format {
    case "json":
        return o.generateJSON(data)
    case "yaml":
        return o.generateYAML(data)
    case "xml":
        return o.generateXML(data)
    case "html":
        return o.generateHTML(data)
    case "markdown":
        return o.generateMarkdown(data)
    case "pdf":
        return o.generatePDF(data)
    case "excel":
        return o.generateExcel(data)
    }
    return fmt.Errorf("unsupported format: %s", o.Format)
}

