//go:build ignore
package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "runtime"
)

func main() {
    fmt.Println("Running static file generation for _mf_sccmp_todo_app...")
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        fmt.Printf("Error: Failed to get current file path for module root discovery\n")
        os.Exit(1)
    }
    currentModuleDir := filepath.Dir(filename)
    destDir := filepath.Join(".", "static")

    if err := os.MkdirAll(destDir, 0755); err != nil {
        fmt.Printf("Error: Failed to create static directory %s: %v\n", destDir, err)
        os.Exit(1)
    }

    filesToCopy := []string{
        "todo_app_description.css",
        "todo_app_description.js",
        "todo_app_heading.css",
        "todo_app_inputs.css",
        "todo_app_inputs.js",
        "todo_app_script.js",
        "todo_app_styles.css",
        "todo_app_view.html",
    }

    for _, fileName := range filesToCopy {
        srcPath := filepath.Join(currentModuleDir, fileName)
        destPath := filepath.Join(destDir, fileName)

        if _, err := os.Stat(srcPath); os.IsNotExist(err) {
            fmt.Printf("Warning: Source file not found: %s. Skipping.\n", srcPath)
            continue
        }
        srcFile, err := os.Open(srcPath)
        if err != nil {
            fmt.Printf("Error: Failed to open source file %s: %v\n", srcPath, err)
            os.Exit(1)
        }
        defer srcFile.Close()
        destFile, err := os.Create(destPath)
        if err != nil {
            fmt.Printf("Error: Failed to create destination file %s: %v\n", destDir, err)
            os.Exit(1)
        }
        defer destFile.Close()
        if _, err = io.Copy(destFile, srcFile); err != nil {
            fmt.Printf("Error: Failed to copy file from %s to %s: %v\n", srcPath, destPath, err)
            os.Exit(1)
        }
        fmt.Printf("Copied %s to %s\n", fileName, destPath)
    }
    fmt.Println("Static file copying complete.")
}