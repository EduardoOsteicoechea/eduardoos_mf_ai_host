const todo_app_add_task_name_container = document.getElementById("todo_app_add_task_name_container")
const todo_app_add_task_name_input = document.getElementById("todo_app_add_task_name_input")
const todo_app_add_task_expand_task_data_button = document.getElementById("todo_app_add_task_expand_task_data_button")
const todo_app_add_task_extended_data = document.getElementById("todo_app_add_task_extended_data")
const todo_app_add_task_extended_data_description = document.getElementById("todo_app_add_task_extended_data_description")
const todo_app_add_task_extended_data_description_label = document.getElementById("todo_app_add_task_extended_data_description_label")
const todo_app_add_task_extended_data_description_label_input = document.getElementById("todo_app_add_task_extended_data_description_label_input")
const todo_app_add_task_extended_data_importance = document.getElementById("todo_app_add_task_extended_data_importance")
const todo_app_add_task_extended_data_importance_label = document.getElementById("todo_app_add_task_extended_data_importance_label")
const todo_app_add_task_extended_data_importance_input = document.getElementById("todo_app_add_task_extended_data_importance_input")
const todo_app_add_task_extended_data_substeps = document.getElementById("todo_app_add_task_extended_data_substeps")
const todo_app_add_task_extended_data_substeps_label = document.getElementById("todo_app_add_task_extended_data_substeps_label")
const todo_app_add_task_extended_data_substeps_buttons_container = document.getElementById("todo_app_add_task_extended_data_substeps_buttons_container")
const todo_app_add_task_extended_data_substeps_buttons_container_add_button = document.getElementById("todo_app_add_task_extended_data_substeps_buttons_container_add_button")
const todo_app_add_task_extended_data_substeps_buttons_container_remove_button = document.getElementById("todo_app_add_task_extended_data_substeps_buttons_container_remove_button")
const todo_app_add_task_extended_data_substeps_items_container = document.getElementById("todo_app_add_task_extended_data_substeps_items_container")
const todo_app_add_task_extended_data_global_time = document.getElementById("todo_app_add_task_extended_data_global_time")
const todo_app_add_task_extended_data_global_time_label = document.getElementById("todo_app_add_task_extended_data_global_time_label")
const todo_app_add_task_extended_data_global_time_input = document.getElementById("todo_app_add_task_extended_data_global_time_input")

let current_task_data = {
  task_name: "",
  task_description: "",
  task_importance: 50,
  task_subtasks: [],
  task_time: 0
}


todo_app_add_task_expand_task_data_button.style.transform = "transform: rotate(0deg)"

todo_app_add_task_expand_task_data_button.addEventListener("pointerup", () => {
  if (todo_app_add_task_expand_task_data_button.classList.contains("todo_app_add_task_expand_task_data_button_downward")) {
    toggle_class_names(
      todo_app_add_task_expand_task_data_button,
      "todo_app_add_task_expand_task_data_button_upward",
      "todo_app_add_task_expand_task_data_button_downward",
    )
  } else {
    toggle_class_names(
      todo_app_add_task_expand_task_data_button,
      "todo_app_add_task_expand_task_data_button_downward",
      "todo_app_add_task_expand_task_data_button_upward",
    )
  }

  if (todo_app_add_task_extended_data.classList.contains("todo_app_add_task_extended_data_collapsed")) {
    toggle_class_names(
      todo_app_add_task_extended_data,
      "todo_app_add_task_extended_data_extended",
      "todo_app_add_task_extended_data_collapsed",
    )
  } else {
    toggle_class_names(
      todo_app_add_task_extended_data,
      "todo_app_add_task_extended_data_collapsed",
      "todo_app_add_task_extended_data_extended",
    )
  }
})


todo_app_add_task_extended_data_substeps_buttons_container_add_button.addEventListener("pointerup", () => {
  let current_item_index = todo_app_add_task_extended_data_substeps.children.length + 1
  generate_task_subtask_markup(current_item_index);
})


function generate_task_subtask_markup(current_item_index) {
  const subtask_container = document.createElement("div")
  subtask_container.className = "subtask_container"
  
  const subtask_number = document.createElement("span")
  subtask_number.className = "subtask_number"
  subtask_number.textContent = current_item_index

  const subtask_name_container = document.createElement("div")
  subtask_name_container.className = "subtask_input_container"
  const subtask_name_label = document.createElement("label")
  subtask_name_label.textContent = "name"
  subtask_name_label.className = "subtask_input_label"
  const subtask_name_input = document.createElement("input")
  subtask_name_input.className = "subtask_input_input"
  subtask_name_container.appendChild(subtask_name_label)
  subtask_name_container.appendChild(subtask_name_input)

  const subtask_time_container = document.createElement("div")
  subtask_time_container.className = "subtask_input_container"
  const subtask_time_label = document.createElement("label")
  subtask_time_label.textContent = "time"
  subtask_time_label.className = "subtask_input_label"
  const subtask_time_input = document.createElement("input")
  subtask_time_input.className = "subtask_input_input"
  subtask_time_container.appendChild(subtask_time_label)
  subtask_time_container.appendChild(subtask_time_input)

  const subtask_remove_button = document.createElement("button")
  subtask_remove_button.textContent = "x"
  subtask_remove_button.addEventListener("pointerup", () => {
    subtask_container.remove()
  })
  subtask_remove_button.className = "subtask_remove_button"

  subtask_container.appendChild(subtask_number)
  subtask_container.appendChild(subtask_name_container)
  subtask_container.appendChild(subtask_time_container)
  subtask_container.appendChild(subtask_remove_button)

  todo_app_add_task_extended_data_substeps.appendChild(subtask_container);
}