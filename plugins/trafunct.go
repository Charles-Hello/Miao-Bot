package plugins

import (
	"fmt"
	"sync"
	"wechat_http/ihttp"
)

func ifcon() {
	fmt.Println(ihttp.Con)
	if ihttp.Con == true {
		<-ihttp.Ch
	}

}

func TraverseFunction(object ihttp.AddOrder) {
	defer ifcon()
	Wg.Add(len(ihttp.PluginQueue))
	for _, i := range ihttp.PluginQueue {
		i := i
		go func() {
			i(object) //这里的d可以把前端的值传进去
			Wg.Done()
		}()
	}
	Wg.Wait()

}

var Wg sync.WaitGroup
