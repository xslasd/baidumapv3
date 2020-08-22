package baidumapsdk

type DataRes struct {
	Status int         `json:"status"`  // 返回结果状态值， 成功返回0，其他值请查看下方返回码状态表。	int
	Result interface{} `json:"result"`  // data数据
	Msg    string      `json:"message"` // 错误消息
}

type ReverseGeocodingRes struct {
	Location         location         `json:"location"`
	Business         string           `json:"business"`          // 坐标所在商圈信息，如 "人民大学,中关村,苏州街"。最多返回3个。
	FormattedAddress string           `json:"formatted_address"` // 结构化地址信息
	AddressComponent addressComponent `json:"addressComponent"`
	Pois             []*pois          `json:"pois"`
	Roads            []*roads         `json:"roads"`
	PoiRegions       []*poiRegions    `json:"poiRegions"`
	SematicDesc      string           `json:"sematic_description"` //当前位置结合POI的语义化结果描述
}

type GeocoderRes struct {
	Location      location `json:"location"`
	Precise       int64    `json:"precise"`
	Confidence    int64    `json:"confidence"`
	Comprehension int64    `json:"comprehension"`
	Level         string   `json:"level"`
}

type IPLocationRes struct {
	Status  int     `json:"status"`
	Address string  `json:"address"`
	Content content `json:"content"`
	Msg     string  `json:"message"` // 错误消息
}
type content struct {
	Address       string        `json:"address"`
	AddressDetail addressDetail `json:"address_detail"`
}
type addressDetail struct {
	City         string `json:"city"`
	CityCode     int64  `json:"city_code"`
	District     string `json:"district"`
	Province     string `json:"province"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

// 根据经纬度坐标获取地址。
type location struct {
	Lng float64 `json:"lng"` // 经度
	Lat float64 `json:"lat"` // 纬度
}

// 注意，国外行政区划，字段仅代表层级
type addressComponent struct {
	Country         string `json:"country"`           // 国家
	CountryCode     int64  `json:"country_code"`      // 国家编码
	CountryCodeIso  string `json:"country_code_iso"`  // 国家英文缩写（三位）
	CountryCodeIso2 string `json:"country_code_iso2"` // 国家英文缩写（两位）
	Province        string `json:"province"`          // 省名
	City            string `json:"city"`              // 城市名
	CityLevel       int64  `json:"city_level"`        // 城市所在级别
	District        string `json:"district"`          // 区县名
	Town            string `json:"town"`              // 乡镇名
	TownCode        string `json:"town_code"`         // 乡镇id
	Street          string `json:"street"`            // 街道名（行政区划中的街道层级）
	StreetNumber    string `json:"street_number"`     // 街道门牌号
	Adcode          string `json:"adcode"`            // 行政区划代码
}

// pois（周边poi数组）
type pois struct {
	Addr      string `json:"addr"`      // 地址信息
	direction string `json:"direction"` // 和当前坐标点的方向
	Distance  int64  `json:"distance"`  // 离坐标点距离
	Name      string `json:"name"`      // poi名称
	Tag       string `json:"tag"`       // poi类型，如’美食;中餐厅’
	Point     point  `json:"point"`     // poi坐标{x,y}
	Tel       string `json:"tel"`       // 电话
	UID       string `json:"uid"`       // poi唯一标识
	Zip       string `json:"zip"`       // 邮编
}

type point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type roads struct {
	Name     string `json:"name"`     // 周边道路名称
	Distance string `json:"distance"` // 传入的坐标点距离道路的大概距离
}
type poiRegions struct {
	DirectionDesc string `json:"direction_desc"` // 请求中的坐标与所归属区域面的相对位置关系
	Name          string `json:"name"`           // 归属区域面名称
	Tag           string `json:"tag"`            // 归属区域面类型
}
