package router

import (
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello test!"))
}

func TestRootMatch(t *testing.T) {
	// Create route
	r := New("/")
	r.Add("GET", "/", http.HandlerFunc(handler))

	// Matching routes
	matches := []string{"/", ""}

	// Check
	for _, match := range matches {
		req, _ := http.NewRequest("GET", match, nil)
		h := r.Match(req)
		if h == nil {
			t.Errorf("'%s' should match against '/'", match)
		}
	}
}

func TestRouteMatch(t *testing.T) {
	r := New("/v1")
	r.Add("GET", "/test", http.HandlerFunc(handler))
	r.Add("GET", "/test/1/2/3/4/5/6", http.HandlerFunc(handler))

	matches := []string{
		"http://example.com/v1/test",
		"http://example.com/v1/test/",
		"http://example.com/v1/test/1/2/3/4/5/6",
		"http://example.com/v1/test/1/2/3/4/5/6/",
	}

	for _, match := range matches {
		req, _ := http.NewRequest("GET", match, nil)
		h := r.Match(req)
		if h == nil {
			t.Errorf("%s should have matched our routes", match)
		}
	}

	nomatches := []string{
		"http://example.com/v1",
		"http://example.com/v1/",
		"http://example.com/v1/test/1/2/3/4/5",
		"http://example.com",
		"http://example.com/",
		"http://example.com/something/else",
	}

	for _, nomatch := range nomatches {
		req, _ := http.NewRequest("GET", nomatch, nil)
		h := r.Match(req)
		if h != nil {
			t.Errorf("%s shouldn't have matched our routes", nomatch)
		}
	}
}

func TestRouteParam(t *testing.T) {
	r := New("/")
	r.Add("GET", "/:test", http.HandlerFunc(handler))
	r.Add("GET", "/:test/1", http.HandlerFunc(handler))
	r.Add("GET", "/:test/1/2", http.HandlerFunc(handler))
	r.Add("GET", "/1/2/:param", http.HandlerFunc(handler))
	r.Add("GET", "/1/2/:param1/3/4/:param2", http.HandlerFunc(handler))

	req, _ := http.NewRequest("GET", "http://example.com/value", nil)
	h := r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/value/1", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value/1")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/value/1/2", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value/1/2")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/1/2/value", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/1/2/value")
	} else if Param(req, "param") != "value" {
		t.Errorf("Param :param should be set to 'value'. Got %s", Param(req, "param"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/1/2/value1/3/4/value2", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/1/2/value1/3/4/value2")
	} else if Param(req, "param1") != "value1" {
		t.Errorf("Param :param1 should be set to 'value1'. Got %s", Param(req, "param1"))
	} else if Param(req, "param2") != "value2" {
		t.Errorf("Param :param2 should be set to 'value2'. Got %s", Param(req, "param2"))
	}

	//dumpTree(r.(*router).tree, "")
}

func TestCatchAllRoute(t *testing.T) {
	r := New("/")
	r.Add("GET", "/:test", http.HandlerFunc(handler))
	r.Add("GET", "/:test/1/*", http.HandlerFunc(handler))
	r.Add("GET", "/1/2/*", http.HandlerFunc(handler))
	r.Add("GET", "/wrong/but/*/valid", http.HandlerFunc(handler))

	req, _ := http.NewRequest("GET", "http://example.com/value", nil)
	h := r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/value/1/something", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value/1/something")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/value/1/2/3/4/5/6/7/8/9/0", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value/1/2/3/4/5/6/7/8/9/0")
	} else if Param(req, "test") != "value" {
		t.Errorf("Param :test should be set to 'value'. Got %s", Param(req, "test"))
	}

	req, _ = http.NewRequest("GET", "http://example.com/1/2/value", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/1/2/value")
	}

	req, _ = http.NewRequest("GET", "http://example.com/wrong/but/something/valid/or/else", nil)
	h = r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/wrong/but/something/valid/or/else")
	}
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello test!"))
}

func TestParam(t *testing.T) {
	r := New("/")
	r.Add("GET", "/:param", http.HandlerFunc(paramHandler))

	req, _ := http.NewRequest("GET", "http://example.com/value", nil)
	h := r.Match(req)
	if h == nil {
		t.Errorf("%s should have matched our routes", "http://example.com/value")
	} else if Param(req, "param") != "value" {
		t.Errorf("Param :param should be set to 'value'. Got %s", Param(req, "param"))
	}
}

func TestGetWrongParam(t *testing.T) {
	r := New("/")
	r.Add("GET", "/:param", http.HandlerFunc(paramHandler))

	req, _ := http.NewRequest("GET", "http://example.com/value", nil)
	h := r.Match(req)
	if h == nil {
		t.Fatalf("%s should have matched our routes", "http://example.com/value")
	} else if Param(req, "invalid") != "" {
		t.Errorf("Param :invalid should be set to ''. Got %v", Param(req, "invalid"))
	}

	if Param(req, "invalid") != "" {
		t.Errorf("Param for :invalid should have been ''. Got %s", Param(req, "invalid"))
	}
}
