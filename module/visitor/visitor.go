package visitor

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type VisitorModule struct {
	pool redis.Pool
}

func (vm *VisitorModule) RetrieveCount() int {
	var res int
	var err error

	redisCon := vm.pool.Get()
	defer redisCon.Close()

	res, err = redis.Int(redisCon.Do("GET", "BIGPROJECT_VISITOR"))
	if err != nil {
		res = 0
	}
	return res
}

func (vm *VisitorModule) IncrementCount() {
	redisCon := vm.pool.Get()
	defer redisCon.Close()

	res, _ := redisCon.Do("SET", "hello", "world")
	log.Println("[BigProject] : Visitor increased, result : ", res)
}

func RegisterVisitorModule(pool redis.Pool) *VisitorModule {
	return &VisitorModule{pool:pool}
}
