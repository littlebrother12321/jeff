{{ template "admin/layout.tpl" . }}
{{ define "content" }}
  <h2 class="text-2xl font-bold mb-4">Login</h2>
  <form method="POST" action="/admin/login">
    <div class="form-control mb-4">
      <label class="label">
        <span class="label-text">Email</span>
      </label>
      <input type="email" name="email" placeholder="Your Email" class="input input-bordered" required />
    </div>
    <div class="form-control mb-4">
      <label class="label">
        <span class="label-text">Password</span>
      </label>
      <input type="password" name="password" class="input input-bordered" required />
    </div>
    <div class="form-control">
      <button type="submit" class="btn btn-primary">Log In</button>
    </div>
  </form>
  {{ if .Result }}
    <div role="alert" class="alert mt-4 pe-8 w-fit">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info h-6 w-6 shrink-0">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
      </svg>
      <span>{{ .Result }}</span>
    </div>
  {{ end }}
{{ end }}
