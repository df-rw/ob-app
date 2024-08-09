# ob-app

Observable web app:
- [Observable Framework](https://observablehq.com/framework) frontend.
- [Go](https://go.dev) web server backend.
- Comms between the front and back through [htmx](https://htmx.org).

## Install Observable Framework.

Setup an observable framework application:
- in the _current_ directory
- as an _empty_ project (no demo files)

```shell
npm init @observablehq
```

## Setup Go webserver

```go
mkdir -p ./cmd/web
cat > ./cmd/web/main.go <<EOF
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func now(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "%s", t)
}

func main() {
	port := flag.Int("p", 8080, "webserver port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/now", now)

	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
EOF
```

## Download and install htmx

Install htmx somewhere Observable Framework can find it:

```shell
mkdir ./docs/js
wget -O ./docs/js/htmx.min.js https://unpkg.com/htmx.org/dist/htmx.min.js
```

Update the page header to include htmx:

```shell
sed -i'' -e "s/^\(  head.*\)',/\1<script src=\"\/js\/htmx.min.js\"><\/script>',/" observablehq.config.js
```

## Check everything

Make sure that everything is running correctly at this point. For clarity, run
these in two different shells:

### Check observable frontend application

```shell
npm run dev
```

You should get a browser window opened on port `3000` (or thereabouts) as the
home page for the observable application.

### Check Go backend application

```shell
go run ./cmd/web/main.go 
```

By default this is set for port `8080`; to check do something like:

```shell
curl -D - http://localhost:8080/now
```

and you should see the current time written to stdout.

# Talking between the frontend and the backend
