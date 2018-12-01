package controller

import (
	"fmt"
	"github.com/skiy/comicFetch/library"
	"log"
)

type shenmanhua struct {
}

func (t *shenmanhua) Init() {
	log.Println("shenmh init")
	var site = "http://www.shenmanhua.com/juepinxiaoshenyi/1.html"
	var sel = ".mh_comicpic img"

	res, err := library.ChromedpText(site, sel, site)
	fmt.Println(res, err)
}
