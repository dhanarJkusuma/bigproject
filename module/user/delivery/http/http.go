package http

import (
	"bigproject/entity"
	"bigproject/module/user"
	"bigproject/util/nonpanic"
	"bigproject/util/render"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type userHttp struct {
	um *user.UserModule
}

func (uh *userHttp) FetchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var page, size int
	var err error
	queryValues := r.URL.Query()
	pg, sz := queryValues.Get("page"), queryValues.Get("size")

	if pg == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pg)
		nonpanic.HandleRequestError(err)
	}


	if sz == "" {
		size = 10
	} else {
		size, err = strconv.Atoi(sz)
		nonpanic.HandleRequestError(err)
	}


	data := uh.um.FetchData(page, size)
	// data := []entity.User{}
	count := uh.um.FetchCount()
	// count := 5
	res := struct {
		Page int `json:"page"`
		Size int `json:"size"`
		Data []entity.User `json:"data"`
		RecordsTotal int `json:"recordsTotal"`
		RecordsFiltered int `json:"recordsFiltered"`

	}{
		page,
		size,
		data,
		count,
		count,
	}

	render.JSONRender(w, res)
}

func (uh *userHttp) FetchSearchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var page, size int
	var err error
	queryValues := r.URL.Query()
	pg, sz := queryValues.Get("page"), queryValues.Get("size")

	if pg == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pg)
		nonpanic.HandleRequestError(err)
	}


	if sz == "" {
		size = 10
	} else {
		size, err = strconv.Atoi(sz)
		nonpanic.HandleRequestError(err)
	}


	data := uh.um.FetchSearch(page, size, queryValues.Get("q"))
	count := uh.um.FetchCountSearch(queryValues.Get("q"))
	res := struct {
		Page int `json:"page"`
		Size int `json:"size"`
		Data []entity.User `json:"data"`
		RecordsTotal int `json:"recordsTotal"`
		RecordsFiltered int `json:"recordsFiltered"`
	}{
		page,
		size,
		data,
		count,
		count,
	}

	render.JSONRender(w, res)
}

//func (uh *userHttp) FetchTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	var page, size int
//	var err error
//	queryValues := r.URL.Query()
//	pg, sz := queryValues.Get("page"), queryValues.Get("size")
//
//	if pg == "" {
//		page = 0
//	} else {
//		page, err = strconv.Atoi(pg)
//		nonpanic.HandleRequestError(err)
//	}
//
//
//	if sz == "" {
//		size = 10
//	} else {
//		size, err = strconv.Atoi(sz)
//		nonpanic.HandleRequestError(err)
//	}
//
//	data := []entity.User{
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//		entity.User{
//			FullName:"Dhanar J Kusuma",
//			UserEmail:"dhanar.j.kusuma@gmail.com",
//			BirthDate: "",
//			Msisdn: "12341234",
//			UserID:1,
//		},
//	}
//	res := struct {
//		Page 			int `json:"page"`
//		Size 			int `json:"size"`
//		Draw 			int `json:"draw"`
//		Data 			[]entity.User `json:"data"`
//		RecordsTotal 	int `json:"recordsTotal"`
//		RecordsFiltered int `json:"recordsFiltered"`
//
//	}{
//		page,
//		size,
//		1,
//		data[0:5],
//		15,
//		15,
//	}
//	render.JSONRender(w, res)
//}

func RegisterUserHttpDelivery(um *user.UserModule, router *httprouter.Router) {
	handler := &userHttp{
		um,
	}

	router.GET("/users", handler.FetchHandler)
	router.GET("/users/search", handler.FetchSearchHandler)
	// router.GET("/users/test", handler.FetchTest)
}