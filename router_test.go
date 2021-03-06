package nji

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试路由组
func TestGroup(t *testing.T) {
	app := NewServer()
	group := app.Group("/api/v2", func(ctx *Context) {
		t.Log(1)
	})
	group.GET("/view1", func(ctx *Context) {
		ctx.Resp.Writer.WriteHeader(http.StatusNoContent)
		t.Log(2, ctx.Request.URL.Path)
	}, func(ctx *Context) {
		t.Log(3, ctx.Request.URL.Path)
	})
	r, err := http.NewRequest("GET", "/api/v2/view1", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}

// 测试中止
func TestAbort(t *testing.T) {
	app := NewServer()
	group := app.Group("/api")
	group.GET("/typ", func(ctx *Context) {
		t.Log(1)
		ctx.Abort()
	}, func(ctx *Context) {
		t.Log(2)
	})
	r, err := http.NewRequest("GET", "/api/typ", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}

// 测试Append
func TestAppend(t *testing.T) {
	app := NewServer()
	app.Append(func(ctx *Context) {
		t.Log(1)
	}, func(ctx *Context) {
		t.Log(2)
	})
	app.GET("/test", func(ctx *Context) {
		t.Log(3, ctx.Request.URL.Path)
	})
	r, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	app.ServeHTTP(httptest.NewRecorder(), r)
}

func TestPanic(t *testing.T) {
	app := NewServer()
	app.Append(func(ctx *Context) {
		t.Log(1)
	}, func(ctx *Context) {
		t.Log(2)
	})
	app.GET("/test/:abc", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
	app.GET("/test/:name", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
}

func TestPanic2(t *testing.T) {
	app := NewServer()
	app.Append(func(ctx *Context) {
		t.Log(1)
	}, func(ctx *Context) {
		t.Log(2)
	})
	app.GET("/test/:abc/1", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
	app.GET("/test/123/1", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
}

func TestPanic3(t *testing.T) {
	app := NewServer()
	app.Append(func(ctx *Context) {
		t.Log(1)
	}, func(ctx *Context) {
		t.Log(2)
	})
	app.GET("/test/:name", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
	app.GET("/test/:name", func(ctx *Context) {
		t.Log(ctx.Request.URL.Path)
	})
}
