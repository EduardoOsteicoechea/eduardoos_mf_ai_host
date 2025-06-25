package _mf_sccmp_todo_app

import (
	"fmt"
	"os"
)

func todo_page_view(pageTop string, pageBody string, pageBottom string) string {
	const page_styles string = `
	<link rel="stylesheet" href="static/todo_app/todo_app_styles.css">
		<link rel="stylesheet" href="static/todo_app/todo_app_heading.css">
		<link rel="stylesheet" href="static/todo_app/todo_app_description.css">
		<link rel="stylesheet" href="static/todo_app/todo_app_inputs.css">
	`
	const page_javascript string = `
		<script src="static/todo_app/todo_app_script.js" defer></script>
		<script src="static/todo_app/todo_app_description.js" defer></script>
		<script src="static/todo_app/todo_app_inputs.js" defer></script>
	`
	const html_filepath = "static/todo_app/todo_app_view.html"

	contentBytes, err := os.ReadFile(html_filepath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", html_filepath, err)
		os.Exit(1)
	}
	var todo_view_content string = string(contentBytes)

	todo_view := pageTop + page_styles + pageBody + todo_view_content + page_javascript + pageBottom

	return todo_view
}
