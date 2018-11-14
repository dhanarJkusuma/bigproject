package http

import (
	"bigproject/module/visitor"
	"bigproject/util/render"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type visitorHttp struct {
	uv *visitor.VisitorModule
}

func (vh *visitorHttp) FetchVisitor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var count int
	count = vh.uv.RetrieveCount()
	res := struct {
		Count int `json:"count"`
	}{
		Count:count,
	}

	render.JSONRender(w, res)
}

func RegisterUserHttpDelivery(uv *visitor.VisitorModule, router *httprouter.Router) {
	handler := &visitorHttp{
		uv,
	}

	router.GET("/visitor", handler.FetchVisitor)
}