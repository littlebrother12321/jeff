{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<h1 class="text-4xl mb-4">Profile</h1>

<form method="POST" action="/profile">
  <div class="form-control mb-4">
    <label class="label">
      <span class="label-text">Name</span>
    </label>
    <input value="{{.User.Name}}" type="text" name="name" class="input input-bordered" required />
  </div>
  <div class="form-control mb-4">
    <label class="label">
      <span class="label-text">Email</span>
    </label>
    <input value="{{.User.Email}}" type="email" name="email" class="input input-bordered" required />
  </div>
  <div class="form-control">
    <button type="submit" class="btn btn-primary">Update Profile</button>
  </div>
</form>  
{{ if .Result }}
  <div role="alert" class="alert mt-4 pe-8 w-fit">
    <svg xmlns="http://www.w4.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info h-6 w-6 shrink-0">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
    </svg>
    <span>{{ .Result }}</span>
  </div>
{{ end }}
{{ end }}
