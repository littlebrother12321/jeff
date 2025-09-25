### MVC

### Docs

https://mehedimk.github.io/Beego/index.html

### Tailwind

https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss

Settings: `{"go-template": "html"}`

```bash
pnpm init
pnpm i -D tailwindcss @tailwindcss/cli
touch tailwind.config.js
```

```js
// tailwind.config.js
/** @type {import("tailwindcss").Config} */
module.exports = {
  content: ["./views/**/*.{html,tpl,tmpl}", "./static/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [],
};
```

```json
{
  "name": "myapp",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "build:css": "tailwindcss -i ./styles/main.css -o ./static/css/output.css --minify",
    "watch:css": "tailwindcss -i ./styles/main.css -o ./static/css/output.css --minify --watch"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "packageManager": "pnpm@10.11.0",
  "devDependencies": {
    "@tailwindcss/cli": "^4.1.12",
    "tailwindcss": "^4.1.12"
  }
}
```

```bash
mkdir -p styles
touch styles/main.css
```

```css
@import "tailwindcss";

a {
  @apply text-blue-500;
}

a:hover {
  @apply text-blue-600 underline;
}

.container {
  @apply max-w-5xl mx-auto block;
}
```
# jeff
