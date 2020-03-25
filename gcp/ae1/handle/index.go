package handle

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	con := `
<!doctype html>
<html>
<body>
<h1>
    Hello, World!
</h1>
</body>
</html>
`
	w.Header().Set("ETag", fmt.Sprintf("W/\"%x\"", md5.Sum([]byte(con))))
	_, _ = fmt.Fprint(w, con)
}
