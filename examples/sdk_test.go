package examples

import (
	"github.com/Jetereting/go-hik"
	"testing"
)

func TestSDK(t *testing.T) {
	hk := go_hik.HKConfig{
		Ip:      "127.0.0.1",
		Port:    443,
		AppKey:  "28057000",
		Secret:  "dZztQSS0000kLpURG000",
		IsHttps: true,
		IsDebug: true,
		TagId:   "kingpark",
	}

	body := map[string]string{
		"pageNo":   "1",
		"pageSize": "100",
	}
	result, err := hk.HttpPost("/api/resource/v1/cameras", body)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("OK", result.Str)

	/*body := map[string]string{
		"cameraIndexCode": "71c1e8bd1b0d406a94e7cdf88a251f9b",
		"protocol":        "rtmp",
	}
	result, err := hk.HttpPost("/api/video/v2/cameras/previewURLs", body)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("OK", result.Str)*/
}
