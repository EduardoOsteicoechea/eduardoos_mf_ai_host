// page.go
package main

import (
	"strings"
)

const PageTop = `
	<!DOCTYPE html>
	<html land="en">
		<head>
		 <meta charset="UTF-8">
		<meta name="description" content="t3.medium">
		<meta name="author" content="Eduardo Osteicoechea">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="icon" href="static/favicon.ico">
		<link rel="stylesheet" href="static/global_styles.css">
		<link rel="stylesheet" href="static/global_article_styles.css">
		<link rel="stylesheet" href="static/skip_to_content.css">
		<link rel="stylesheet" href="static/loading_screen.css">
`
const PageBody = `
		</head>
		<body>
			<div 
			id="loading_screen_outer_container"
			class="loading_screen_outer_container"
			>
				<p 
				id="loading_screen_outer_container_loading_text"
				class="loading_screen_outer_container_loading_text"
				>
					Loading...
				</p>
			</div>
			
			<a 
			id="skip_to_content_element" 
			class="skip_to_content_outer_container" 
			href="#main"
			>
				Skip to content
			</a>

			<script 
			src="static/global_scripts.js"
			>
			</script>
`

const PageBottom = `

		</body>
	</html>
`

func PrintPage(css_tags []string, javascript_tags []string, components_markup []string) string {
	composed_css_tags := strings.Join(css_tags, "\n")
	composed_js_tags := strings.Join(javascript_tags, "\n")
	composed_component_markup := strings.Join(components_markup, "\n")
	markup := PageTop + "\n" + composed_css_tags + "\n" + PageBody + "\n" + composed_component_markup + "\n" + composed_js_tags + "\n" + PageBottom

	return markup
}