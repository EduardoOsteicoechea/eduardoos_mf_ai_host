const todo_app_info_buttons_container = document.getElementById("todo_app_info_buttons_container")
const todo_app_info_buttons_buttons_container = document.getElementById("todo_app_info_buttons_buttons_container")
const todo_app_info_show_description_button = document.getElementById("todo_app_info_show_description_button")
const todo_app_scroll_to_app_button = document.getElementById("todo_app_scroll_to_app_button")
const todo_app_technical_description_description_collapsable_container = document.getElementById("todo_app_technical_description_description_collapsable_container")
const todo_app_technical_description_description_container = document.getElementById("todo_app_technical_description_description_container")
const todo_app_technical_description_description_container_heading = document.getElementById("todo_app_technical_description_description_container_heading")

todo_app_technical_description_description_collapsable_container.classList.add("todo_app_technical_description_description_collapsable_container_collapsed")
todo_app_info_show_description_button.textContent = "View App Narrative";

todo_app_info_show_description_button.addEventListener("pointerup", () => {
  if (todo_app_info_show_description_button.textContent == "View App Narrative") {
    todo_app_info_show_description_button.textContent = "Hide App Narrative"
    toggle_class_names(
      todo_app_technical_description_description_collapsable_container,
      "todo_app_technical_description_description_collapsable_container_expanded",
      "todo_app_technical_description_description_collapsable_container_collapsed",
    )
  } else {
    todo_app_info_show_description_button.textContent = "View App Narrative"
    toggle_class_names(
      todo_app_technical_description_description_collapsable_container,
      "todo_app_technical_description_description_collapsable_container_collapsed",
      "todo_app_technical_description_description_collapsable_container_expanded",
    )
  }
})

todo_app_scroll_to_app_button.addEventListener("pointerup", () => {
  const offset = 50;
  const elementRect = todo_app_input_controls_container.getBoundingClientRect();
  const scrollToPosition = elementRect.top + window.scrollY - offset;
  window.scrollTo({
    top: scrollToPosition,
    behavior: "smooth"
  });
})