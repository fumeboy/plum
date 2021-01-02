package plugins_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nji"
	"nji/plugins"
	"testing"
)

var _ nji.ViewI = &a{}

type a struct {
	A plugins.PathParam
}

func (view *a) Handle(c *nji.Context) {
	c.ResponseWriter.WriteHeader(200)
	_, _ = c.ResponseWriter.Write([]byte("Hello !"))
}

func TestContext(t *testing.T) {
	app := nji.Config{
		UnescapePathValues: true,
	}.New()
	app.GET("/context/:A", nji.Inject(&a{}))
	r, err := http.NewRequest("GET", "/context/2", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)
	assert.Equal(t, "Hello World!", string(w.Body.Bytes()))
}
