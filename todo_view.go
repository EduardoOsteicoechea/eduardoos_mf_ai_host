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
	
	var todo_view_content string = fmt.Sprintf(`
	<div class="todo_app_outer_container">
		<h1 class="todo_app_title">Todo App</h1>
		
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
		
		<h2 class="Pending Tasks_heading">Pending Tasks</h2>
		
		<ol 
		id="todo_app_ordered_list" 
		class="todo_app_ordered_list"
		>

		</ol>
	</div>
	
	<script id="todo_app_script">
		const todo_app_ordered_list = document.getElementById("todo_app_ordered_list")
		const todo_app_add_task_input = document.getElementById("todo_app_add_task_input")
		const todo_app_add_task_button = document.getElementById("todo_app_add_task_button")

		todo_app_add_task_button.addEventListener("click",()=>
		{
			const todo_app_ordered_list_item_count = todo_app_ordered_list.children.length + 1

			const new_task_container = document.createElement("li")
			new_task_container.className = "new_task_container"			

			const new_task_item_number = document.createElement("span")
			new_task_item_number.textContent = todo_app_ordered_list_item_count
			new_task_item_number.className = "new_task_item_number"
			
			const new_task_data_items_container = document.createElement("div")
			new_task_data_items_container.className = "new_task_data_items_container"

			const new_task_text = document.createElement("p")
			new_task_text.innerHTML = todo_app_add_task_input.value
			new_task_text.className = "new_task_text"
			
			const new_task_date_and_delete_button_container = document.createElement("div")
			new_task_date_and_delete_button_container.className = "new_task_date_and_delete_button_container"
			
			const new_task_time = document.createElement("span")
			const now = new Date()
    	const task_time_formatted = now.toLocaleString()
    	new_task_time.textContent = task_time_formatted
			new_task_time.className = "new_task_time"

			const new_task_delete_button = document.createElement("button")
			new_task_delete_button.innerHTML = "X"
			new_task_delete_button.addEventListener("click",()=>{
				new_task_container.remove()
			})
			new_task_delete_button.className = "new_task_delete_button"

			new_task_data_items_container.appendChild(new_task_text)
			new_task_date_and_delete_button_container.appendChild(new_task_time)
			new_task_date_and_delete_button_container.appendChild(new_task_delete_button)
			new_task_data_items_container.appendChild(new_task_date_and_delete_button_container)

			new_task_container.appendChild(new_task_item_number)
			new_task_container.appendChild(new_task_data_items_container)
			
			todo_app_ordered_list.appendChild(new_task_container)

			todo_app_add_task_input.value = ""
		})
	</script>
	`, formatted_date)

	todo_view := pageTop + page_styles + pageBody + todo_view_content + pageBottom

	return todo_view
}
