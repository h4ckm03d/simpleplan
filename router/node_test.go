package router

import (
	"net/http"
	"testing"
)

type emptyHandler struct{}

func (e emptyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestAdd(t *testing.T) {
	root := rootNode("/", emptyHandler{})
	if root.path != "/" {
		t.Errorf("root.path should be '/'. Got %s", root.path)
	}

	root.add("/some/route/with/five/parts", emptyHandler{})
	if len(root.children) != 1 {
		for _, ch := range root.children {
			t.Errorf("Error data: %s", ch.path)
		}
		t.Fatalf("root.children should have 1 items. Got %d", len(root.children))
	}

	root.add("/test/action", emptyHandler{})
	if len(root.children) != 2 {
		for _, ch := range root.children {
			t.Errorf("Error data: %s", ch.path)
		}
		t.Fatalf("root.children should have 2 items. Got %d", len(root.children))
	}
}
