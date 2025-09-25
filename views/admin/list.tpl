{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<h1 class="text-4xl">{{.Title}}</h1>
<table class="table w-full">
  {{if eq (len .List) 0}} 
  <tr>
    <td colspan="3">No records</td> 
  </tr>
  {{else}}
  <thead>
    <tr>
    {{range index .List 0}}
      <th class="bg-gray-200">{{.Key}}</th>
    {{end}}
    </tr>
  </thead>
  <tbody>
    {{range .List}}
    <tr>
    {{range .}}
      {{if eq .Key "Id"}}
        <td class="border px-4 py-2">
        <a href="{{$.BaseHref}}/{{.Value}}">
        {{.Value}}
        </a>
        </td>
      </a>
      {{else}}
        <td class="border px-4 py-2">{{.Value}}</td>
      {{end}}
    {{end}}
    </tr>
    </a>
    {{end}}
  </tbody>
  {{end}}
</table>
{{ end }}
