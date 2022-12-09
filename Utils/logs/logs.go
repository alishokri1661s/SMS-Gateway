package logs

import "log"

func LogOnError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PanicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
