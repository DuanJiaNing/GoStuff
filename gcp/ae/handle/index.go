package handle

import (
	"GoStuff/gcp/ae/app"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", app.Handler(index))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprint(w, `<!doctype html>
<html>
<body>
<h1>
    Hello, World!
</h1>
</body>
</html>
`)
}
