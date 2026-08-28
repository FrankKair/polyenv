package main

import (
    "bytes"
    "compress/flate"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

const (
    apiURL       = "https://tio.run/cgi-bin/run/api"
    languagesURL = "https://tio.run/languages.json"
    tokenLen     = 16
)

// Result holds the ouput from a TIO execution.
type Result struct {
    Stdout string
    Debug  string
    Error  string
}

func buildPayload(lang, code, stdin string) ([]byte, error) {
    var payload strings.Builder
    fmt.Fprintf(&payload, "Vlang\x001\x00%s\x00", lang)
    fmt.Fprintf(&payload, "F.code.tio\x00%d\x00%s\x00", len(code), code)
    if stdin != "" {
        fmt.Fprintf(&payload, "F.input.tio\x00%d\x00%s\x00", len(stdin), stdin)
    }
    payload.WriteString("R")

    var buf bytes.Buffer
    w, err := flate.NewWriter(&buf, flate.BestCompression)
    if err != nil {
        return nil, err
    }
    if _, err := w.Write([]byte(payload.String())); err != nil {
        return nil, err
    }
    if err := w.Close(); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// Run sends code to tio.run and returns the result.
func Run(lang, code string) (*Result, error) {
    payload, err := buildPayload(lang, code, "")
    if err != nil {
        return nil, fmt.Errorf("building payload: %w", err)
    }

    client := &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Post(apiURL, "application/octet-stream", bytes.NewReader(payload))
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("reading response: %w", err)
    }

    text := string(body)
    if len(text) < tokenLen {
        return &Result{Error: "unexpected empty response from tio.run"}, nil
    }

    token := text[:tokenLen]
    parts := strings.Split(text, token)

    r := &Result{}
    if len(parts) > 1 {
        r.Stdout = parts[1]
    }
    if len(parts) > 2 {
        r.Debug = parts[2]
    }

    notFound := fmt.Sprintf("The language '%s' could not be found on the server.\n", lang)
    if r.Stdout == notFound {
        r.Error = notFound
    }

    return r, nil
}

// FetchLanguages returns a mapa of command -> display name for all available languages.
func FetchLanguages() (map[string]string, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(languagesURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close() 

    var data map[string]struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    langs := make(map[string]string, len(data))
    for cmd, info := range data {
        langs[cmd] = info.Name
    }
    return langs, nil
}
