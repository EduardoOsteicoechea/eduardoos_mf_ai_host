// page.go
package main

const PageTop = `
	<!DOCTYPE html>
	<html land="en">
		<head>
		 <meta charset="UTF-8">
		<meta name="description" content="t3.medium">
		<meta name="author" content="Eduardo Osteicoechea">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="static/global_styles.css">
		<link rel="stylesheet" href="static/global_article_styles.css">
		<link rel="stylesheet" href="static/skip_to_content.css">
		<script src="static/global_scripts.js"></script>
`
const PageBody = `
		</head>
		<body>
			<div 
			id="loading_screen"
			class="loading_screen_outer_container"
			>
				<p class="loading_screen_outer_container_loading_text">Loading...</p>
			</div>
			
			<a 
			id="skip_to_content_element" 
			id="skip_to_content_outer_container" href="#main"
			>
				Skip to content
			</a>
`

const PageBottom = `
		</body>
	</html>
`
