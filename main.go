package main

import (
	"bigproject/module/user"
	userHttp "bigproject/module/user/delivery/http"
	"bigproject/module/visitor"
	visitorHttp "bigproject/module/visitor/delivery/http"
	"bigproject/util/config"
	"bigproject/util/database"
	"bigproject/util/memory"
	"bigproject/util/messaging"
	"bigproject/util/render"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func init() {
	log.Println("[BigProject] : Init Config ")
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		errStr := fmt.Sprintf("[%v][BigProject][Error] : %v", time.Now(), err.Error())
		panic(errStr)
	}
}

func main(){
	router := httprouter.New()
	router.GET("/", handleIndex)

	db := database.ConnectDB()
	pool := memory.GetRedisConnectionPool()

	// register user module
	userModule := user.RegisterUserModule(db)
	userHttp.RegisterUserHttpDelivery(userModule, router)

	// register visitor module
	visitorModule := visitor.RegisterVisitorModule(pool)

	// add nsq
	conf := config.GetConfig()
	log.Println("[BigProject] : Init MQ ")
	messagingOptions := messaging.Options{
		LookupAddress:  []string{conf.MQLookupAddr},
		PublishAddress: conf.MQPublishAddr,
	}
	ConsumeEngine := messaging.NewConsumer(messagingOptions)
	handleIncomeMsg := func(m *nsq.Message)error {
		visitorModule.IncrementCount()
		m.Finish()
		return nil
	}
	consumerOption := &messaging.ConsumerOptions{
		Topic:       "bigproject",
		Channel:     "dhanarjkusuma",
		Handler:     handleIncomeMsg,
		MaxAttempts: uint16(conf.MQDefaultConsumerMaxAttempts),
		MaxInFlight: conf.MQDefaultConsumerMaxInFlight,
	}
	ConsumeEngine.RegisterConsumer(consumerOption)
	ConsumeEngine.RunConsumer()

	publisherEngine := messaging.NewPublisher(messagingOptions)
	// register delivery rest visitor
	visitorHttp.RegisterUserHttpDelivery(visitorModule, &publisherEngine, router)

	log.Printf("[BigProject] : App Running in port :%v ", conf.AppPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.AppPort), router))
}


func handleIndex(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	conf := config.GetConfig()
	data := map[string]interface{}{
		"Host": conf.AppHost,
		"Port": conf.AppPort,
	}
	render.HTMLRender(w, r, "index.gohtml", data)
}