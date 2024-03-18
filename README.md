# Playground for HTMX and Go with GO-Templates

- Place static content int `data/static`
- Place Templates into `data/apptpl`


## Development
For Development files will be read from the data directory if it's present.

## Production
Run `go build` to create `gohtmx-play` this binary will run standalone and include all resources if the Data folder is not preset at startup
