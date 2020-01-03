package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/logging"

	netcontext "golang.org/x/net/context"
)

const (
	appLogName = "app_log"
)

var (
	contextKey             = "holds a *logContent"
	errNotAppEngineContext = errors.New("not an illegal context")
)

type UserAppLogLine struct {
	Timestamp time.Time
	Severity  logging.Severity
	Message   string
}

type logContent struct {
	sync.Mutex

	request     *http.Request
	pendingLogs []*UserAppLogLine
}

func FlushLog(ctx netcontext.Context) {
	content, ok := ctx.Value(&contextKey).(*logContent)
	if !ok {
		return
	}

	client, err := logging.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Printf("got error when create log client: %v", err)
		// TODO logs lost
		return
	}

	trace := content.request.Header.Get("X-Cloud-Trace-Context")
	if len(trace) == 0 {
		log.Print("failed to get X-Cloud-Trace-Context from request heard")
		// TODO logs lost
		return
	}

	traceID := fmt.Sprintf("projects/%s/traces/%s", os.Getenv("GCP_PROJECT_ID"), trace[:32])
	log.Printf("X-Cloud-Trace-Context: %s, traceID: %s", trace, traceID)
	appLogger := client.Logger(appLogName)
	for _, l := range content.pendingLogs {
		appLogger.Log(logging.Entry{
			Timestamp: l.Timestamp,
			Severity:  l.Severity,
			Payload:   l.Message,
			Trace:     traceID,
		})
	}
	if err := appLogger.Flush(); err != nil {
		log.Printf("got error when flush log: %v", err)
		return
	}
}

func WithLogContentValue(r *http.Request) netcontext.Context {
	return netcontext.WithValue(r.Context(), &contextKey, &logContent{
		request:     r,
		pendingLogs: []*UserAppLogLine{},
	})
}

func Logf(ctx netcontext.Context, severity logging.Severity, format string, args ...interface{}) {
	content, ok := ctx.Value(&contextKey).(*logContent)
	if !ok {
		panic(errNotAppEngineContext)
	}

	s := fmt.Sprintf(format, args...)
	s = strings.TrimRight(s, "\n") // Remove any trailing newline characters.

	content.Lock()
	content.pendingLogs = append(content.pendingLogs, &UserAppLogLine{
		Timestamp: time.Now(),
		Severity:  severity,
		Message:   s,
	})
	content.Unlock()
}
