const loading_screen_outer_container = document.getElementById("loading_screen_outer_container")
loading_screen_outer_container.style.opacity = 1

function show_loading_screen(){
  loading_screen_outer_container.style.opacity = 1
}
function hide_loading_screen(){
  loading_screen_outer_container.style.opacity = 0
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