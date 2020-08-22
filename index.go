package baidumapsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// 全球逆地理编码服务Url
	ReverseGeocoding_Url = "http://api.map.baidu.com/reverse_geocoding/v3"
	// 地理编码服务
	Geocoder_Url = "http://api.map.baidu.com/geocoding/v3"
	// 普通IP定位
	IPLocation_Url = "http://api.map.baidu.com/location/ip"
)

type BaiduMapSDK struct {
	ak string
}

func NewBaiduMapClient(ak string) *BaiduMapSDK {
	return &BaiduMapSDK{ak}
}

func (s *BaiduMapSDK) GetAk() string {
	return s.ak
}

// 全球逆地理编码服务
// 全球逆地理编码服务（又名Geocoder）是一类Web API接口服务；
// 逆地理编码服务提供将坐标点（经纬度）转换为对应位置信息（如所在行政区划，周边地标点分布）功能。
// 服务同时支持全球行政区划位置描述及周边地标POI数据召回（包括中国在内的全球200多个国家地区）；
// 逆地理编码境外POI服务于2020年7月2日上线，在此之前注册的AK默认无使用权限。
// 若在此之前注册的AK需使用该服务，请进入API控制台为AK勾选“逆地理编码境外POI服务”，则可正常使用。
func (s *BaiduMapSDK) ReverseGeocoding(lat, lng float64) (*ReverseGeocodingRes, error) {
	params := url.Values{}
	params.Set("location", fmt.Sprintf("%v,%v", lat, lng))
	params.Set("output", "json")
	return s.ReverseGeocodingParams(&params)
}

func (s *BaiduMapSDK) ReverseGeocodingParams(params *url.Values) (*ReverseGeocodingRes, error) {
	Url, err := url.Parse(ReverseGeocoding_Url)
	params.Set("ak", s.ak)
	Url.RawQuery = params.Encode()
	if err != nil {
		return nil, err
	}
	fmt.Println(Url.String())
	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res DataRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Status == 0 {
		data, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		var obj ReverseGeocodingRes
		err = json.Unmarshal(data, &obj)
		if err != nil {
			return nil, err
		}
		return &obj, nil
	}
	return nil, errors.New(res.Msg)
}

// 地理编码服务
// 地理编码服务（又名Geocoder）是一类Web API接口服务；
// 地理编码服务提供将结构化地址数据（如：北京市海淀区上地十街十号）转换为对应坐标点（经纬度）功能；
// 地理编码服务当前未推出国际化服务，解析地址仅限国内；
func (s *BaiduMapSDK) Geocoder(address string) (*GeocoderRes, error) {
	params := url.Values{}
	params.Set("address", address)
	params.Set("output", "json")
	return s.GeocoderParams(&params)
}

func (s *BaiduMapSDK) GeocoderParams(params *url.Values) (*GeocoderRes, error) {
	Url, err := url.Parse(Geocoder_Url)
	params.Set("ak", s.ak)
	Url.RawQuery = params.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res DataRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Status == 0 {
		data, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		var obj GeocoderRes
		err = json.Unmarshal(data, &obj)
		if err != nil {
			return nil, err
		}
		return &obj, nil
	}
	return nil, errors.New(res.Msg)
}

// 普通IP定位
// 普通IP定位是一套以HTTP/HTTPS形式提供的轻量级定位接口，用户可以通过该服务，根据IP定位来获取大致位置。
// 目前该服务同时支持 IPv4 和 IPv6 来获取位置信息。
func (s *BaiduMapSDK) IPLocation(ip string) (*IPLocationRes, error) {
	params := url.Values{}
	params.Set("ip", ip)
	params.Set("coor", "bd09ll")
	return s.IPLocationParams(&params)
}
func (s *BaiduMapSDK) IPLocationParams(params *url.Values) (*IPLocationRes, error) {
	Url, err := url.Parse(IPLocation_Url)
	params.Set("ak", s.ak)
	Url.RawQuery = params.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res IPLocationRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Status != 0 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
