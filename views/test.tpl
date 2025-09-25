{{ template "layout.tpl" . }}

{{ define "content" }}
        <h2 class="text-3xl font-bold>{{ .Title}}</h2>
        <p>This is a test page!</p>
{{ end }}