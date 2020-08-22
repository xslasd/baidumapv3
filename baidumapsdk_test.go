package baidumapsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

var sdk *BaiduMapSDK

func TestMain(t *testing.M) {
	sdk = NewBaiduMapClient("bt63e7RZrzZ71WxaW4dakcecm4BQg42G")
	t.Run()
}

func TestBaiduMapSDK_Geocoder(t *testing.T) {
	data, err := sdk.Geocoder("杭州市江干区采荷路41号")
	if err != nil {
		t.Error(err)
	}
	res, _ := json.Marshal(data)
	fmt.Println(string(res))
}

func TestBaiduMapSDK_ReverseGeocoding(t *testing.T) {
	data, err := sdk.ReverseGeocoding(120.19834501474797, 30.257413871308195)
	if err != nil {
		t.Error(err)
	}
	res, _ := json.Marshal(data)
	fmt.Println(string(res))
}

func TestBaiduMapSDK_IPLocation(t *testing.T) {
	data, err := sdk.IPLocation("222.72.23.24")
	if err != nil {
		t.Error(err)
	}
	res, _ := json.Marshal(data)
	fmt.Println(string(res))
}
