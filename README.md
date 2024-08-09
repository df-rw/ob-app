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
mkdir ./src/js
wget -O ./src/js/htmx.min.js https://unpkg.com/htmx.org/dist/htmx.min.js
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

## Handling CORS requests

Requests from the frontend to the backend will be halted due to CORS, as the
frontend will be making calls to a backend with different port number. Update
the server:

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

```diff
index 0efb1a6..1eccf40 100644
--- a/src/index.md
+++ b/src/index.md
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

# Setting the `hx-get` path

Our development environment uses two servers: one each for front and back.

An example production deploy could involve hosting the Observable Framework
application inside the backend application, reducing the number of servers to
one. Since Observable Framework builds to a static site, the frontend would be
served by the backend application, off a filesystem path as static files.

This makes the `hx-get` path incorrect, since it currently references a local
development path and not a production one - the hostname of the backend
application, or a relative path only without a hostname.

One way to set the path at build time is to pass the root URL for the backend
service to Observable Framework. This can be done through environment variables.

## Passing environment variables to Observable Framework

We can pass an environment variable to Observable Framework when we start
the Framework server:

```shell
APPSERVER=http://localhost:8080 npm run dev
```

We can make this variable availble to pages in the `<head>` section of a page:

```diff
index 156a78b..c5d9c4a 100644
--- a/observablehq.config.js
+++ b/observablehq.config.js
@@ -17,10 +17,16 @@ export default {
   // ],
 
   // Content to add to the head of the page, e.g. for a favicon:
-  head: '<link rel="icon" href="observable.png" type="image/png" sizes="32x32"><script src="/js/htmx.min.js"></script>',
+  head: `
+    <link rel="icon" href="observable.png" type="image/png" sizes="32x32">
+    <script src="/js/htmx.min.js"></script>
+    <script>
+    var APPSERVER = '${process.env.APPSERVER ?? ""}';
+    </script>
+  `,
 
   // The path to the source root.
   root: "src",
```

We pick this up in a page:

```diff
index 1eccf40..9ac80f3 100644
--- a/src/index.md
+++ b/src/index.md
@@ -4,6 +4,8 @@ This is the home page of your new Observable Framework project.
 
 For more, see <https://observablehq.com/framework/getting-started>.
 
+The appserver path is: ${APPSERVER}
+
 <button
     hx-get="http://localhost:8080/now"
     hx-target="#now"
```

## Replacing the `hx-get` path

It appears that javascript variables aren't expanded inside DOM elements, which
makes sense, since we are building pages statically. We can get around this
creating a function that makes our control. And since we are changing the DOM
on the fly, we need to get htmx to process after the changes are done:

```diff
index 1eccf40..a37b5a0 100644
--- a/src/index.md
+++ b/src/index.md
@@ -4,9 +4,22 @@ This is the home page of your new Observable Framework project.
 
 For more, see <https://observablehq.com/framework/getting-started>.
 
-<button
-    hx-get="http://localhost:8080/now"
-    hx-target="#now"
-    hx-swap="innerHTML">get now</button>
+The appserver path is: ${APPSERVER}
+
+```js
+const makeGetNow = (label) => {
+    return html`<button
+        hx-get="${APPSERVER}/now"
+        hx-target="#now"
+        hx-swap="innerHTML"
+    >${label}</button>`;
+};
+```
+
+```js
+display(makeGetNow("get now"));
+await visibility()
+htmx.process(document.body);
+```
 
 <div id="now"></div>
```

# TODO

- Write up how production deploy will work:
  - npm run build
  - adjust `cmd/web/main.go`
    - remove cors middleware
    - mount the static site
- Try and figure out a better way to handle dynamic htmx element creation:
  - move the makers to a library
  - figure out how to group everything so `htmx.process` only has to run once
