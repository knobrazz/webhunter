package output

import (
    "html/template"
    "net/http"
)

type InteractiveVisualizer struct {
    Server      *http.Server
    WebSocket   *WSHandler
    Charts      map[string]*Chart
}

type Chart struct {
    Type       string
    Data       interface{}
    Options    map[string]interface{}
    Events     []string
}

type WSHandler struct {
    Clients    map[string]*Client
    Broadcast  chan interface{}
}

func (v *InteractiveVisualizer) StartServer(addr string) error {
    // Setup routes
    http.HandleFunc("/ws", v.WebSocket.HandleConnection)
    http.HandleFunc("/api/data", v.handleDataRequest)
    http.HandleFunc("/chart/update", v.handleChartUpdate)

    // Start server
    return v.Server.ListenAndServe()
}

func (v *InteractiveVisualizer) CreateInteractiveChart(chartType string, data interface{}) *Chart {
    chart := &Chart{
        Type:    chartType,
        Data:    data,
        Options: make(map[string]interface{}),
        Events:  []string{"click", "hover", "select"},
    }

    // Add interactivity options based on chart type
    switch chartType {
    case "network":
        chart.Options["draggable"] = true
        chart.Options["zoomable"] = true
    case "timeline":
        chart.Options["selectable"] = true
        chart.Options["editable"] = true
    case "heatmap":
        chart.Options["tooltip"] = true
        chart.Options["brush"] = true
    }

    return chart
}

