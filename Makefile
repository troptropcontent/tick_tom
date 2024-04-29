go_server:
	air
sass:
	/opt/dart-sass/sass --watch ./internal/assets/stylesheets/application.scss ./public/css/application.css
tailwind:
	/tailwindcss -o ./internal/assets/stylesheets/utilities/tailwind.css --watch