// main.go
package main

func todo_page_view(pageTop string, pageBody string, pageBottom string) string {
	const page_styles string = `
		<link rel="stylesheet" href="static/todo_app_styles.css">
	`
	const page_javascript string = `
		<script src="static/todo_app_script.js" defer></script>
	`
	
	var todo_view_content string = `
	<div
	id="todo_app_outer_container"
	class="todo_app_outer_container"
	>
		<h1
		id="todo_app_title"
		class="todo_app_title"
		>
			Minimal Framework Todo App
		</h1>	
		<div 
		id="todo_app_input_controls_container" 
		class="todo_app_input_controls_container"
		>
			<textarea 
			id="todo_app_add_task_input"  
			class="todo_app_add_task_input"
			>
			</textarea>			
			<button 
			id="todo_app_add_task_button"  
			class="todo_app_add_task_button"
			>
				Add Task
			</button>
		</div>	
		<ol 
		id="todo_app_ordered_list" 
		class="todo_app_ordered_list"
		></ol>
	</div>
	`

	todo_view := pageTop + page_styles + pageBody + todo_view_content + page_javascript + pageBottom

	return todo_view
}
