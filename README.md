# go-hik
> 海康威视OpenAPI安全认证库 - Golang版本实现
# 官网

接口调用认证：[文档说明](https://open.hikvision.com/resourceCenter)

其他语言版本：[下载链接](https://open.hikvision.com/download/5c67f1e2f05948198c909700?type=10)
# 快速使用
````
> go get github.com/Jetereting/go-hik
````
#### 从运行管理中心获取/创建appKey和secret
* 默认地址账号：
http://192.168.x.x:8001/center sysadmin Hik12345+/hik12345!@#
* 点击到合作方管理
http://192.168.x.x:9017/artemis-web/consumer/index

#### 从综合安防平台验证接口情况
* 默认地址账号：
https://192.168.x.x admin hik12345+/hik12345!@#

#### 海康服务器ssh
* 默认用户： hik/root
* 默认密码： 123456/Hik12345+/Hik12345=!@/hik12345!@#
* 默认端口： 55555

# 示例代码
````
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
		"pageSize":  1000,
		"orderBy":   "name",
		"orderType": "asc",
	})
	fmt.Println(result, err)
	// 方式二：封装调用
	result, err = hik.DoorList()
	fmt.Println(result, err)
}

```` 
# 输出结果
````
{
  "total": 152,
  "pageNo": 1,
  "pageSize": 2,
  "list": [
    {
      "indexCode": "88c3f05cee794e5eb91348af1cb85069",
      "name": "11.12 28F两门控制器",
      "resourceType": "acsDevice",
      "devTypeCode": "201926400",
      "devTypeDesc": "DS-K2602",
      "ip": "192.168.11.12",
      "port": "8000",
      "regionIndexCode": "76150874-fac2-4b82-8800-90f6c648a0b4",
      "treatyType": "hiksdk_net",
      "capability": "@support_485@support_antisneak@support_antisneakhost@support_biosignature@support_card@support_cardtype@support_doorlock@support_doorstatus@support_finger@support_leadercard@support_m1cardEncrypt@support_multicard@support_readerverify@",
      "cardCapacity": 100000,
      "fingerCapacity": 3000,
      "faceCapacity": 3000,
      "doorCapacity": 2,
      "netZoneId": "0",
      "isCascade": 0,
      "dataVersion": "0",
      "createTime": "2023-06-19T09:10:04.211+08:00",
      "updateTime": "2023-06-28T12:59:05.165+08:00",
      "manufacturer": "hikvision",
      "acsReaderVerifyModeAbility": "1023",
      "devSerialNum": "DS-K260220191231V020006CHJ97538410",
      "sort": 1745,
      "disOrder": 1745,
      "regionName": "28F",
      "regionPath": "@root000000@f5d94e48-623e-45e4-8354-7e8712a0ada5@70d8f6e0-c646-45a6-b220-f9580096f5df@ccdfd7f1-1b22-423c-bbb6-884b1c8dec25@76150874-fac2-4b82-8800-90f6c648a0b4@",
      "regionPathName": "根节点/T1塔楼-办公/T1办公-门禁/T1办公-防火门/28F"
    },
    {
      "indexCode": "8965e4a2edac471f8e54ce445c2b2d3c",
      "name": "11.13 27F双门控制器",
      "resourceType": "acsDevice",
      "devTypeCode": "201926400",
      "ip": "192.168.11.13",
      "port": "8000",
      "regionIndexCode": "b7d32c79-8e0f-40ac-a851-7ecc562d5e94",
      "treatyType": "hiksdk_net",
      "cardCapacity": 50000,
      "fingerCapacity": 3000,
      "faceCapacity": 3000,
      "doorCapacity": 2,
      "netZoneId": "0",
      "isCascade": 0,
      "dataVersion": "0",
      "createTime": "2023-06-25T09:32:57.398+08:00",
      "updateTime": "2023-06-25T09:32:57.398+08:00",
      "manufacturer": "hikvision",
      "acsReaderVerifyModeAbility": "0",
      "sort": 2200,
      "disOrder": 2200,
      "regionName": "27F",
      "regionPath": "@root000000@f5d94e48-623e-45e4-8354-7e8712a0ada5@70d8f6e0-c646-45a6-b220-f9580096f5df@ccdfd7f1-1b22-423c-bbb6-884b1c8dec25@b7d32c79-8e0f-40ac-a851-7ecc562d5e94@",
      "regionPathName": "根节点/T1塔楼-办公/T1办公-门禁/T1办公-防火门/27F"
    }
  ]
}
````
