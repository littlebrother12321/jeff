{{ template "layout.tpl" . }}
{{ define "content" }}
        <h2 class="text-3xl font-bold">{{ .Title }}</h2>
        <h3>This is luke's amazing website made in Advanced Web Dev</h3>
	<br>
	<a href="/static/WeThePeople.ttf" class="link link-primary">Please download and install this font for website</a>
{{ end }}