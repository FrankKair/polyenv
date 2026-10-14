package main

import (
    "bytes"
    "compress/flate"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func TestBuildPayload(t *testing.T) {
    for _, tt := range []struct {
        name, code, stdin, want string
    }{
        {"code", `puts :HELLO`, "", "Vlang\x001\x00ruby\x00F.code.tio\x0011\x00puts :HELLOR"},
        {"stdin", "puts gets", "hello", "Vlang\x001\x00ruby\x00F.code.tio\x009\x00puts getsF.input.tio\x005\x00helloR"},
        {"unicode", "puts 'é'", "", "Vlang\x001\x00ruby\x00F.code.tio\x009\x00puts 'é'R"},
        {"empty", "", "", "Vlang\x001\x00ruby\x00F.code.tio\x000\x00R"},
    } {
        t.Run(tt.name, func(t *testing.T) {
            payload, err := buildPayload("ruby", tt.code, tt.stdin)
            if err != nil {
                t.Fatal(err)
            }
            reader := flate.NewReader(bytes.NewReader(payload))
            defer reader.Close()
            got, err := io.ReadAll(reader)
            if err != nil {
                t.Fatal(err)
            }
            if string(got) != tt.want {
                t.Fatalf("payload = %q, want %q", got, tt.want)
            }
        })
    }
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req)
}

func TestRunFileReportsResponseError(t *testing.T) {
    original := http.DefaultTransport
    t.Cleanup(func() { http.DefaultTransport = original })
    http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
        if req.URL.Path != "/cgi-bin/run/api/" {
            t.Errorf("unexpected API path: %s", req.URL.Path)
        }
        return &http.Response{
            StatusCode: http.StatusOK,
            Body:       io.NopCloser(strings.NewReader("")),
            Header:     make(http.Header),
        }, nil
    })
    path := filepath.Join(t.TempDir(), "test.rb")
    if err := os.WriteFile(path, []byte("puts :HELLO"), 0600); err != nil {
        t.Fatal(err)
    }
    err := runFile("ruby", path)
    if err == nil || err.Error() != "unexpected empty response from tio.run" {
        t.Fatalf("runFile error = %v, want response error", err)
    }
}

func skipIfNoNetwork(t *testing.T) {
    t.Helper()
    if os.Getenv("POLYENV_NETWORK_TESTS") == "" {
        t.Skip("set POLYENV_NETWORK_TESTS=1 to run network tests")
    }
}

func TestRunPython(t *testing.T) {
    skipIfNoNetwork(t)
    result, err := Run("python3", `print("hello from tio")`)
    if err != nil {
        t.Fatalf("Run: %v:", err)
    }
    if !strings.Contains(result.Stdout, "hello from tio") {
        t.Errorf("expected stdout to contain 'hello from tio', got: %q", result.Stdout)
    }
    if result.Error != "" {
        t.Errorf("unexpected error: %s", result.Error)
    }
}

func TestRunRuby(t *testing.T) {
    skipIfNoNetwork(t)
    result, err := Run("ruby", "puts :HELLO")
    if err != nil {
        t.Fatal(err)
    }
    if result.Stdout != "HELLO\n" || result.Error != "" || !strings.Contains(result.Debug, "Exit code: 0") {
        t.Fatalf("unexpected Ruby result: %+v", result)
    }
}

func TestRunBadLanguage(t *testing.T) {
    skipIfNoNetwork(t)
    result, err := Run("nonexistent_lang_xyz", "code")
    if err != nil {
        t.Fatalf("Run: %v:", err)
    }
    if result.Error == "" {
        t.Error("expected error for bad language, got none")
    }
    if !strings.Contains(result.Error, "could not be found") {
        t.Errorf("unexpected error message: %s", result.Error)
    }
}

func TestFetchLanguages(t *testing.T) {
    skipIfNoNetwork(t)
    langs, err := FetchLanguages()
    if err != nil {
        t.Fatalf("FetchLanguages: %v", err)
    }
    if len(langs) < 100 {
        t.Errorf("expected 100+ languages, got %d", len(langs))
    }
    if _, ok := langs["python3"]; !ok {
        t.Error("python3 not found in languages")
    }
}
