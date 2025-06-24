const todo_app_outer_container = document.getElementById("todo_app_outer_container")
const todo_app_title = document.getElementById("todo_app_title")
const todo_app_ordered_list = document.getElementById("todo_app_ordered_list")
const todo_app_add_task_heading = document.getElementById("todo_app_add_task_heading")
const todo_app_input_controls_container = document.getElementById("todo_app_input_controls_container")
const todo_app_add_task_input = document.getElementById("todo_app_add_task_input")
const todo_app_add_task_button = document.getElementById("todo_app_add_task_button")
let generated_tasks_record = []

todo_app_add_task_button.addEventListener("click", () => {
  if (todo_app_add_task_input.value === "") {
    alert("Add a value to the task")
    return
  }
  
  let new_task_text = todo_app_add_task_input.value.trim()
  const now = new Date()
  const task_time_formatted = now.toLocaleString()  
  generated_tasks_record = [...generated_tasks_record, [new_task_text, task_time_formatted]]
  regenerate_todo_list()
})

generate_ui_heights();

regenerate_todo_list();

//////////////////////////
//////////////////////////
// METHODS
//////////////////////////
//////////////////////////

function generate_ui_heights() {
  const screen_height = window.visualViewport.height
  todo_app_outer_container.style.height = `${screen_height}px`
  todo_app_ordered_list.style.height = `${screen_height*.5}px`
  todo_app_add_task_input.style.height = `${screen_height*.15}px`
  hide_loading_screen()
}

function regenerate_todo_list() {
  todo_app_ordered_list.innerHTML = ""

  for (let x = 0; x < generated_tasks_record.length; x++) {
    const current_task = generated_tasks_record[x]
    todo_app_ordered_list.appendChild(task_markup(x, current_task))
  }

  todo_app_add_task_input.value = ""
}

function task_markup(new_task_index, current_task_data) {
  const new_task_container = document.createElement("li")
  new_task_container.className = "new_task_container"

  const new_task_item_number = document.createElement("span")
  new_task_item_number.textContent = new_task_index + 1
  new_task_item_number.className = "new_task_item_number"

  const new_task_data_items_container = document.createElement("div")
  new_task_data_items_container.className = "new_task_data_items_container"

  const new_task_text = document.createElement("p")
  new_task_text.textContent = current_task_data[0]
  new_task_text.className = "new_task_text"

  const new_task_date_and_delete_button_container = document.createElement("div")
  new_task_date_and_delete_button_container.className = "new_task_date_and_delete_button_container"

  const new_task_time = document.createElement("span")
  new_task_time.textContent = current_task_data[1]
  new_task_time.className = "new_task_time"

  const new_task_delete_button = document.createElement("button")
  new_task_delete_button.innerHTML = "X"
  new_task_delete_button.addEventListener("click", () => {
    generated_tasks_record.splice(new_task_index, 1)
    regenerate_todo_list()
  })
  new_task_delete_button.className = "new_task_delete_button"

  new_task_data_items_container.appendChild(new_task_text)

  new_task_date_and_delete_button_container.appendChild(new_task_time)
  new_task_date_and_delete_button_container.appendChild(new_task_delete_button)
  new_task_data_items_container.appendChild(new_task_date_and_delete_button_container)

  new_task_container.appendChild(new_task_item_number)
  new_task_container.appendChild(new_task_data_items_container)

  return new_task_container
}