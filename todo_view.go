// main.go
package main

import (
	"fmt"
	"time"
)

func todo_page_view(pageTop string, pageBody string, pageBottom string) string {
	current_time := time.Now()
	formatted_date := current_time.Format("2006-01-02")

	const page_styles string = `
		<link rel="stylesheet" href="static/todo_app_styles.css">
	`
	const page_javascript string = `
		<script src="static/todo_app_script.js" defer></script>
	`
	
	var todo_view_content string = fmt.Sprintf(`
	<div class="todo_app_outer_container">
		<h1 class="todo_app_title">Todo App</h1>
		
		<h2 class="Pending Tasks_heading">Pending Tasks</h2>
		
		<ol 
		id="todo_app_ordered_list" 
		class="todo_app_ordered_list"
		>

		</ol>

		
		
		<h2 class="todo_app_date">%s</h2>

		<div 
		id="todo_app_input_controls_container" class="todo_app_input_controls_container"
		>
			<textarea 
			id="todo_app_add_task_input"  
			class="todo_app_add_task_input"
			></textarea>
			
			<button 
			id="todo_app_add_task_button"  
			class="todo_app_add_task_button"
			>
				Add Task
			</button>
		</div>
	</div>
	`, formatted_date)

	todo_view := pageTop + page_styles + pageBody + todo_view_content + page_javascript + pageBottom

	return todo_view
}
