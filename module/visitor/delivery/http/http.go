package http

import (
	"bigproject/module/visitor"
	"bigproject/util/messaging"
	"bigproject/util/render"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type visitorHttp struct {
	uv *visitor.VisitorModule
	publisher *messaging.Publisher
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

func (vh *visitorHttp) IncrVisitor(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	vh.publisher.PublishMessage("bigproject", []byte("inc"))

	res := struct {
		Message string `json:"message"`
	}{
		"increment process is running. ",
	}
	render.JSONRender(w, res)
}

func RegisterUserHttpDelivery(uv *visitor.VisitorModule, publisher *messaging.Publisher, router *httprouter.Router) {
	handler := &visitorHttp{
		uv,
		publisher,
	}

	router.GET("/visitor", handler.FetchVisitor)
	router.POST("/visitor", handler.IncrVisitor)
}