package nonpanic

import (
	"log"
	"time"
)

func HandleRequestError(err error){
	if err != nil {
		log.Printf("[%v][BigProject][Request Error] : %v", time.Now(), err.Error())
	}
}

func HandleParsingDBValueError(err error){
	if err != nil {
		log.Printf("[%v][BigProject][DB Parsing Data Error] : %v", time.Now(), err.Error())
	}
}

func HandleDBError(err error){
	if err != nil {
		log.Printf("[%v][BigProject][DB/Memory  Error] : %v", time.Now(), err.Error())
	}
}