package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

var version = "dev"

func usage() {
    fmt.Fprintf(os.Stderr, `polyenv %s -- run any programming language via tio.run

Usage:
  polyenv run <language> <file>
  polyenv search <query>
  polyenv languages
  polyenv version
`, version)
}

func searchLanguages(query string) error {
    langs, err := FetchLanguages()
    if err != nil {
        return err
    }

    query = strings.ToLower(query)
    var matches []string
    for cmd, name := range langs {
        if strings.Contains(strings.ToLower(cmd), query) || strings.Contains(strings.ToLower(name), query) {
            matches = append(matches, fmt.Sprintf("  %s (%s)", name, cmd))
        }
    }

    if len(matches) == 0 {
        fmt.Printf("No languages matching '%s'.\n", query)
        return nil
    }

    sort.Strings(matches)
    fmt.Printf("Languages matching '%s':\n\n", query)
    for _, m := range matches {
        fmt.Println(m)
    }
    return nil
}

func runFile(lang, path string) error {
    code, err := os.ReadFile(path)
    if err != nil {
        return err
    }

    result, err := Run(lang, string(code))
    if err != nil {
        return err
    }

    if result.Error != "" {
        fmt.Fprintf(os.Stderr, "error: languages '%s' not found\n", lang)
        searchLanguages(lang)
        os.Exit(1)
    }

    if result.Stdout != "" {
        fmt.Print(result.Stdout)
    }

    if !strings.Contains(result.Debug, "Exit code: 0") && result.Debug != "" {
        fmt.Fprint(os.Stderr, result.Debug)
        os.Exit(1)
    }

    return nil
}

func listLanguages() error {
    langs, err := FetchLanguages()
    if err != nil {
        return err
    }

    entries := make([]string, 0, len(langs))
    for cmd, name := range langs {
        entries = append(entries, fmt.Sprintf("  %s (%s)", name, cmd))
    }
    sort.Strings(entries)
    for _, e := range entries {
        fmt.Println(e)
    }
    return nil
}

func run() error {
    if len(os.Args) < 2 {
        usage()
        return nil
    }

    switch os.Args[1] {
    case "run":
        if len(os.Args) != 4 {
            return fmt.Errorf("usage: polyenv run <language> <file>")
        }
        return runFile(os.Args[2], os.Args[3])
    case "search":
        if len(os.Args) != 3 {
            return fmt.Errorf("usage: polyenv search <query>")
        }
        return searchLanguages(os.Args[2])
    case "languages":
        return listLanguages()
    case "version":
        fmt.Printf("polyenv %s\n", version)
        return nil
    default:
        usage()
        return fmt.Errorf("unknown command: %s", os.Args[1])
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
