package main

import (
	"bigproject/module/user"
	userHttp "bigproject/module/user/delivery/http"
	visitorHttp "bigproject/module/visitor/delivery/http"
	"bigproject/module/visitor"
	"bigproject/util/database"
	"bigproject/util/memory"
	"bigproject/util/render"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	log.Println("[BigProject] : Init Config. ")
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main(){
	router := httprouter.New()
	router.GET("/", handleIndex)

	db := database.ConnectDB()
	pool := memory.GetRedisConnectionPool()

	userModule := user.RegisterUserModule(db)
	userHttp.RegisterUserHttpDelivery(userModule, router)

	visitorModule := visitor.RegisterVisitorModule(pool)
	visitorHttp.RegisterUserHttpDelivery(visitorModule, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}


func handleIndex(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	render.HTMLRender(w, r, "index.gohtml", nil)
}