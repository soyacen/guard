package util

import "github.com/yacen/guard/util/log"

func CheckWriteReturn(n int, err error) {
	if err != nil {
		log.Error(err)
	}
}

func CheckErrorReturn(err error) {
	if err != nil {
		log.Error(err)
	}
}
