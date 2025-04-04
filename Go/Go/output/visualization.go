package output

import (
    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
)

type Visualizer struct {
    Charts     []ChartType
    Templates  map[string]string
    ThemeData  map[string]interface{}
}

type ChartType struct {
    Type     string
    Data     interface{}
    Options  map[string]interface{}
}

func (v *Visualizer) GenerateVisuals(data interface{}) error {
    charts := []string{"vulnerability_timeline", "risk_heatmap", "attack_surface", "compliance_radar"}
    
    for _, chartType := range charts {
        if err := v.createChart(chartType, data); err != nil {
            return err
        }
    }
    
    return v.combineCharts()
}

func (v *Visualizer) createChart(chartType string, data interface{}) error {
    switch chartType {
    case "vulnerability_timeline":
        return v.createTimelineChart(data)
    case "risk_heatmap":
        return v.createHeatmap(data)
    case "attack_surface":
        return v.createAttackSurfaceMap(data)
    case "compliance_radar":
        return v.createComplianceRadar(data)
    }
    return nil
}

