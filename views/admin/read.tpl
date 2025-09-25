{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<h1 class="text-4xl">{{.Title}}</h1>

<button class="btn btn-error" onclick="fetch(window.location.href, {method:'DELETE'}).then(()=>{window.location.href = {{.BaseHref}}})">BALEETED!</button>
<table class="table w-full mt-4">
  <thead>
    <tr>
      <th class="bg-gray-200">Key</th>
      <th class="bg-gray-200">Value</th>
    </tr>
  </thead>
  <tbody>
    {{range .Item}}
    <tr>
      <td class="border px-4 py-2">{{.Key}}</td>
      <td class="border px-4 py-2">{{.Value}}</td>
    </tr>
    {{end}}
  </tbody>
</table>
{{ end }}
