package core_http_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestRequestIDGenerated(t *testing.T) {
	var requestId string
	var responseId string

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseId = w.Header().Get(requestIDHeader)
		requestId = r.Header.Get(requestIDHeader)
	})

	middleware := RequestID()

	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal("cannot get new request", err)
	}

	middleware(next).ServeHTTP(recorder, request)

	if len(responseId) == 0 || len(requestId) == 0 {
		t.Fatalf("responseId is emty %s or responseId is emty %s", responseId, requestId)
	}

	if responseId != requestId {
		t.Fatalf("requestId is %s and responseId is %s ", requestId, responseId)
	}
}

func TestRequestIDExisted(t *testing.T) {
	var requestId = uuid.NewString()
	var responseId = requestId

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(requestIDHeader, responseId)
		r.Header.Set(requestIDHeader, requestId)
	})

	middleware := RequestID()

	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatal("cannot get new request", err)
	}

	middleware(next).ServeHTTP(recorder, request)

	if request.Header.Get(requestIDHeader) != requestId {
		t.Fatalf("requestId: %s is not the same as request Header id: %s", requestId, request.Header.Get(requestIDHeader))
	}
}
