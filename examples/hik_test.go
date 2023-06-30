package examples

import (
	"fmt"
	"github.com/Jetereting/go-hik"
	"github.com/Jetereting/go-hik/hik"
	"testing"
)

func init() {
	hik.HIK = go_hik.HKConfig{
		Ip:      "192.168.8.254",
		Port:    443,
		AppKey:  "xxxxxx",
		Secret:  "xxxxxxxxxxxxxxxx",
		TagId:   "kingpark",
		IsHttps: true,
		IsDebug: true,
	}
}

func TestDoorList(t *testing.T) {
	// 方式一：直接调用
	result, err := hik.HIK.HttpPost("/api/resource/v2/acsDevice/search", map[string]interface{}{
		"pageNo":    1,
		"pageSize":  2,
		"orderBy":   "name",
		"orderType": "asc",
	})
	fmt.Println(result, err)
	// 方式二：封装调用
	result, err = hik.DoorList()
	fmt.Println(result, err)
}
