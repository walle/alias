package alias

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInMartini(t *testing.T) {
	m := martini.New()
	Add("^/foo/bar", "/bar", SERVE)
	m.Use(Handler())
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bar" {
			t.Fatalf("Alias Failed")
		}
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar", nil)
	m.ServeHTTP(w, r)
}

func TestRemoveInMartini(t *testing.T) {
	m := martini.New()
	Add("^/foo/bar", "/bar", SERVE)
	Remove("^/foo/bar")
	m.Use(Handler())
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/foo/bar" {
			t.Fatalf("Alias Failed")
		}
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar", nil)
	m.ServeHTTP(w, r)
}
