package gin

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralColonWithRun(t *testing.T) {
	engine := New()
	engine.GET(`/r\:r`, func(c *Context) { c.String(http.StatusOK, "it worked") })

	var engineBaseURL string
	{
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		assert.NoError(t, err)
		defer listener.Close()

		go func() {
			// err is meaningless here
			_ = engine.RunListener(listener)
		}()
		engineBaseURL = "http://" + listener.Addr().String()
	}

	testRequest(t, engineBaseURL+`/r:r`, "", "it worked")
}

func TestLiteralColonWithDirectServeHTTP(t *testing.T) {
	engine := New()
	engine.GET(`/r\:r`, func(c *Context) { c.String(http.StatusOK, "it worked") })

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, `/r:r`, nil)
	assert.NoError(t, err)

	http.Handler(engine).ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "it worked")
}

func TestLiteralColonWithHandler(t *testing.T) {
	engine := New()
	engine.GET(`/r\:r`, func(c *Context) { c.String(http.StatusOK, "it worked") })

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, `/r:r`, nil)
	assert.NoError(t, err)

	http.Handler(engine.Handler()).ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "it worked")
}

func TestLiteralColonWithHTTPServer(t *testing.T) {
	// After remove engine.updateRouteTrees in TestEscapedColon(),
	// the httptest.NewServer(engine) already test &http.Server{Handler: engine} with correct behavior.
	// So we don't need to test it again.
}
