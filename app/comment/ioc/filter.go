package ioc

import (
	"log"

	"github.com/importcjj/sensitive"

	"github.com/crazyfrankie/douyin/app/comment/common/constants"
)

func InitFilter() *sensitive.Filter {
	filter := sensitive.New()
	err := filter.LoadWordDict(constants.WordDictPath)
	if err != nil {
		log.Println("InitFilter Fail,Err=" + err.Error())
	}

	return filter
}
