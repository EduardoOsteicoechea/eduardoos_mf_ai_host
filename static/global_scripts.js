const loading_screen_outer_container = document.getElementById("loading_screen_outer_container")
loading_screen_outer_container.style.opacity = 1

function show_loading_screen() {
  toggle_class_names(
    "loading_screen_outer_container_visible", 
    "loading_screen_outer_container_hidden"
  )
}

function hide_loading_screen() {
  toggle_class_names(
    "loading_screen_outer_container_hidden", 
    "loading_screen_outer_container_visible"
  )
}

function toggle_class_names(class_name_to_add, class_name_to_remove) {
  if (!loading_screen_outer_container.classList.contains(class_name_to_add)) {
    loading_screen_outer_container.classList.add(class_name_to_add)
  }
  if (loading_screen_outer_container.classList.contains(class_name_to_remove)) {
    loading_screen_outer_container.classList.remove(class_name_to_remove)
  }
}

function setViewportHeight() {
  if (window.visualViewport) {
    let visualViewportHeight = window.visualViewport.height;
    document.documentElement.style.setProperty('--vh-visual', `${visualViewportHeight}px`);
  } else {
    // Fallback for older browsers
    let vh = window.innerHeight * 0.01;
    document.documentElement.style.setProperty('--vh', `${vh}px`);
  }
}

/////////////////////
/////////////////////
// DEFAULT ACTIONS
/////////////////////
/////////////////////

setViewportHeight();

// Update on resize (especially useful for keyboard appearance/disappearance)
if (window.visualViewport) {
  window.visualViewport.addEventListener('resize', setViewportHeight);
} else {
  window.addEventListener('resize', setViewportHeight);
}