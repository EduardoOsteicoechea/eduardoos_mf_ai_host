package _mf_sccmp_todo_app

import (
	"fmt"
	"os"
)

func PrintCssTags() string {
	const markup string = `
		<link rel="stylesheet" href="static/todo_app_styles.css">
		<link rel="stylesheet" href="static/todo_app_heading.css">
		<link rel="stylesheet" href="static/todo_app_description.css">
		<link rel="stylesheet" href="static/todo_app_inputs.css">
	`
	return markup
}

func PrintJavaScriptTags() string {
	const markup string = `
		<script src="static/todo_app_script.js" defer></script>
		<script src="static/todo_app_description.js" defer></script>
		<script src="static/todo_app_inputs.js" defer></script>
	`
	return markup
}

func PrintComponentMarkup() string {
	const html_filepath = "static/todo_app_view.html"

	contentBytes, err := os.ReadFile(html_filepath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", html_filepath, err)
		os.Exit(1)
	}
	var todo_view_content string = string(contentBytes)

	return todo_view_content
}

// func CopyStaticFiles() error {
// 	fmt.Println("Copying .js and .css files to ./static directory...")

// 	destDir := filepath.Join(".", "static")

// 	if err := os.MkdirAll(destDir, 0755); err != nil {
// 		return fmt.Errorf("failed to create static directory %s: %w", destDir, err)
// 	}

// 	// --- Determine the source module's directory (where _sc_comp_todo_app's files are) ---
// 	// This is a robust way to find the module's root from an imported function.
// 	_, filename, _, ok := runtime.Caller(0)
// 	if !ok {
// 		return fmt.Errorf("failed to get current file path for module root discovery")
// 	}
// 	currentModuleDir := filepath.Dir(filename) // Start from the directory of this file (main.go)
// 	for {
// 		// Check if 'go.mod' exists in the current directory
// 		if _, err := os.Stat(filepath.Join(currentModuleDir, "go.mod")); err == nil {
// 			break // Found go.mod, this is the module root.
// 		}
// 		parentDir := filepath.Dir(currentModuleDir)
// 		if parentDir == currentModuleDir { // Reached filesystem root, didn't find go.mod
// 			return fmt.Errorf("could not find go.mod file for the '_sc_comp_todo_app' module. Please ensure this function is called from within a valid Go module structure, or adjust the source path logic")
// 		}
// 		currentModuleDir = parentDir // Move up to the parent directory
// 	}
// 	// --- End of source module directory determination ---

// 	filesToCopy := []string{
// 		"todo_app_description.css",
// 		"todo_app_description.js",
// 		"todo_app_heading.css",
// 		"todo_app_inputs.css",
// 		"todo_app_inputs.js",
// 		"todo_app_script.js",
// 		"todo_app_styles.css",
// 	}

// 	for _, fileName := range filesToCopy {
// 		srcPath := filepath.Join(currentModuleDir, fileName)
// 		destPath := filepath.Join(destDir, fileName)

// 		// Check if source file exists
// 		if _, err := os.Stat(srcPath); os.IsNotExist(err) {
// 			fmt.Printf("Warning: Source file not found: %s. Skipping.\n", srcPath)
// 			continue
// 		}

// 		srcFile, err := os.Open(srcPath)
// 		if err != nil {
// 			return fmt.Errorf("failed to open source file %s: %w", srcPath, err)
// 		}
// 		defer srcFile.Close() // Ensure the source file is closed

// 		destFile, err := os.Create(destPath)
// 		if err != nil {
// 			return fmt.Errorf("failed to create destination file %s: %w", destPath, err)
// 		}
// 		defer destFile.Close() // Ensure the destination file is closed

// 		_, err = io.Copy(destFile, srcFile)
// 		if err != nil {
// 			return fmt.Errorf("failed to copy file from %s to %s: %w", srcPath, destPath, err)
// 		}
// 		fmt.Printf("Copied %s to %s\n", fileName, destPath)
// 	}
// 	fmt.Println("File copying complete.")
// 	return nil
// }