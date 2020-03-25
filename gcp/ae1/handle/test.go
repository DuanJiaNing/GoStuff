package handle

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/xxx/test", test)
}

func test(w http.ResponseWriter, r *http.Request) {
	con := `
<!doctype html>
<html>
<body>
<h1>
    from /xxx/test
</h1>
</body>
</html>
`

	w.Header().Set("Content-Type", "application/x-zip-compressed")
	w.Header().Set("ETag", fmt.Sprintf("%x", md5.Sum([]byte(con))))
	w.Header().Set("X-E-T-a-g", fmt.Sprintf("%x", md5.Sum([]byte(con))))
	_, _ = fmt.Fprint(w, con, w.Header())
}
