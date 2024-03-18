# Playground for HTMX and Go with GO-Templates

- Place static content int `data/static`
- Place Templates into `data/apptpl`
- Place Routes in srv/srv.go
  - [Gin-Gonic](https://github.com/gin-gonic/gin) router is used

## Development
For Development files will be read from the data directory if it's present.

## Production
Run `go build` to create `gohtmx-play` this binary will run standalone and include all resources if the Data folder is not preset at startup
