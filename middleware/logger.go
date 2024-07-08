// middleware/logger.go

package middleware

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "time"
)

type RequestMetadata struct {
    Timestamp   string `json:"timestamp"`
    Method      string `json:"method"`
    Path        string `json:"path"`
    RemoteAddr  string `json:"remote_addr"`
    UserAgent   string `json:"user_agent"`
    ContentType string `json:"content_type"`
}

var (
    metadataFile = filepath.Join("cache", "metadata.json")
    mutex        sync.Mutex
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        metadata := RequestMetadata{
            Timestamp:   time.Now().Format(time.RFC3339),
            Method:      r.Method,
            Path:        r.URL.Path,
            RemoteAddr:  r.RemoteAddr,
            UserAgent:   r.UserAgent(),
            ContentType: r.Header.Get("Content-Type"),
        }

        // Log metadata asynchronously
        go func() {
            if err := logMetadata(metadata); err != nil {
                fmt.Printf("Error logging metadata: %v\n", err)
            }
        }()

        // Call the next handler
        next.ServeHTTP(w, r)
    }
}

func logMetadata(metadata RequestMetadata) error {
    mutex.Lock()
    defer mutex.Unlock()

    // Ensure the cache directory exists
    if err := os.MkdirAll(filepath.Dir(metadataFile), 0755); err != nil {
        return fmt.Errorf("failed to create cache directory: %v", err)
    }

    // Read existing metadata
    var metadataList []RequestMetadata
    data, err := os.ReadFile(metadataFile)
    if err == nil {
        if err := json.Unmarshal(data, &metadataList); err != nil {
            return fmt.Errorf("failed to unmarshal existing metadata: %v", err)
        }
    }

    // Append new metadata
    metadataList = append(metadataList, metadata)

    // Write updated metadata back to file
    updatedData, err := json.MarshalIndent(metadataList, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal updated metadata: %v", err)
    }

    if err := os.WriteFile(metadataFile, updatedData, 0644); err != nil {
        return fmt.Errorf("failed to write metadata to file: %v", err)
    }

    return nil
}