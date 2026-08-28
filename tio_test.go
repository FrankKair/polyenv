package main

import (
    "os"
    "strings"
    "testing"
)

func TestBuildPayload(t *testing.T) {
    payload, err := buildPayload("python3", `print("hi")`, "")
    if err != nil {
        t.Fatalf("buildPayload: %v", err)
    }
    if len(payload) == 0 {
        t.Fatal("payload is empty")
    }
}

func TestBuildPayloadWithStdin(t *testing.T) {
    payload, err := buildPayload("python3", `x = input()`, "hello")
    if err != nil {
        t.Fatalf("buildPayload: %v", err)
    }
    if len(payload) == 0 {
        t.Fatal("payload is empty")
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
