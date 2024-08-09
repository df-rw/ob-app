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

Requests from the frontend to the backend will be halted due to CORS. Update the server:

```diff
index f776624..93f8bc7 100644
--- a/cmd/web/main.go
+++ b/cmd/web/main.go
@@ -14,6 +14,14 @@ func now(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprintf(w, "%s", t)
 }

+func corsMiddleware(next http.Handler) http.Handler {
+	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+		w.Header().Set("Access-Control-Allow-Origin", "*")
+		w.Header().Set("Access-Control-Allow-Headers", "*")
+		next.ServeHTTP(w, r)
+	})
+}
+
 func main() {
 	port := flag.Int("p", 8080, "webserver port")
 	flag.Parse()
@@ -22,6 +30,8 @@ func main() {

 	mux.HandleFunc("/now", now)

+	corsMux := corsMiddleware(mux)
+
 	fmt.Println("Listening on port", *port)
-	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
+	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), corsMux))
 }
```

## Adding htmx to Observable

We can now include a htmx call in the frontend:

```
index 0efb1a6..1eccf40 100644
--- a/docs/index.md
+++ b/docs/index.md
@@ -3,3 +3,10 @@
 This is the home page of your new Observable Framework project.

 For more, see <https://observablehq.com/framework/getting-started>.
+
+<button
+    hx-get="http://localhost:8080/now"
+    hx-target="#now"
+    hx-swap="innerHTML">get now</button>
+
+<div id="now"></div>
```
