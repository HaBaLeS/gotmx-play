# Playground for HTMX and Go with GO-Templates
Want to play with Go, Go-Templates and HTMX without the boilerplate of setting up a project. Use this very simple Setup. Default just has the Quickstart Example form https://htmx.org/ functional
With a matching endpoint to trigger the swap
```
  <script src="/static/js/htmx.min.js"></script>
  <!-- have a button POST a click via AJAX -->
  <button hx-post="/clicked" hx-swap="outerHTML">
      Click Me
  </button>
```

- Place static content in `data/static`
- Place Templates in `data/apptpl`
- Place Routes in srv/srv.go
  - [Gin-Gonic](https://github.com/gin-gonic/gin) router is used

## Development
For Development files will be read from the data directory if it's present. Files are parsed for every requiest, this may have an impact of the performance of responses. Change `App#GetTemplate()` if this is a hazzle to you. 

## Production
Run `go build` to create `gohtmx-play` this binary will run standalone and include all resources. Only works if there is no `data/` folder present in the work dir.

