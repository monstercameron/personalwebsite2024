// metadata/metadata.go

package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type RequestMetadata struct {
	Timestamp   string `json:"timestamp"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	RemoteAddr  string `json:"remote_addr"`
	UserAgent   string `json:"user_agent"`
	ContentType string `json:"content_type"`
}

type PageMetrics struct {
	TotalViews       int                       `json:"total_views"`
	UniqueViews      int                       `json:"unique_views"`
	ViewsPerPage     map[string]int            `json:"views_per_page"`
	UserAgentPerPage map[string]map[string]int `json:"user_agent_per_page"`
	MethodCounts     map[string]int            `json:"method_counts"`
	ContentTypes     map[string]int            `json:"content_types"`
}

func CalculateMetrics() (*PageMetrics, error) {
	metadataFile := filepath.Join("cache", "metadata.json")

	data, err := os.ReadFile(metadataFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %v", err)
	}

	var metadataList []RequestMetadata
	if err := json.Unmarshal(data, &metadataList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %v", err)
	}

	metrics := &PageMetrics{
		ViewsPerPage:     make(map[string]int),
		UserAgentPerPage: make(map[string]map[string]int),
		MethodCounts:     make(map[string]int),
		ContentTypes:     make(map[string]int),
	}

	uniqueViews := make(map[string]map[string]bool)

	for _, m := range metadataList {
		// Total views
		metrics.TotalViews++

		// Views per page
		metrics.ViewsPerPage[m.Path]++

		// User agent per page
		if _, ok := metrics.UserAgentPerPage[m.Path]; !ok {
			metrics.UserAgentPerPage[m.Path] = make(map[string]int)
		}
		metrics.UserAgentPerPage[m.Path][m.UserAgent]++

		// Method counts
		metrics.MethodCounts[m.Method]++

		// Content types
		metrics.ContentTypes[m.ContentType]++

		// Unique views
		if _, ok := uniqueViews[m.Path]; !ok {
			uniqueViews[m.Path] = make(map[string]bool)
		}
		uniqueViews[m.Path][m.RemoteAddr] = true
	}

	// Calculate unique views
	for _, ips := range uniqueViews {
		metrics.UniqueViews += len(ips)
	}

	return metrics, nil
}

func GetMetricsJSON() ([]byte, error) {
	metrics, err := CalculateMetrics()
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(metrics, "", "  ")
}