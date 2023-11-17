package persist

import (
	"bili/backend/engine"
	"bili/backend/tool"
	"fmt"
	"sync"
)

type GetItemChan func(wg *sync.WaitGroup) (chan *engine.Item, error)

func GetItemProcessFun() GetItemChan {
	var itemProcessFun GetItemChan
	if !tool.CheckFfmegStatus() {
		fmt.Println("Can't locate your ffmeg.The video your download can't be merged")
		itemProcessFun = VideoItemCleaner
	} else {
		itemProcessFun = VideoItemProcessor
	}

	return itemProcessFun
}
