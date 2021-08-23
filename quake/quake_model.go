/**------------------------------------------------------------**
 * @filename quake/quake_model.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/19 21:05
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

import "time"

type QRes struct {
	Code interface{} `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

type Meta struct {
	Total     int `json:"total,omitempty"`
	PageID    string `json:"pagination_id,omitempty"`
	Page      *Pagination `json:"pagination,omitempty"`

}

type DepthMeta struct {
	Total     int `json:"total,omitempty"`
	PageID    string `json:"pagination_id"`
}

type Pagination struct {
	Count int `json:"count,omitempty"`
	PageIndex int `json:"page_index,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	Total     int `json:"total,omitempty"`
}

type Service struct {
	HTTP *HTTP `json:"http"`
	Version string `json:"version"`
	Response string `json:"response"`
	Name string `json:"name"`
	Banner string `json:"banner"`
	Cert string `json:"cert"`
}

type Favicon struct {
	Hash string `json:"hash"`
	Location string `json:"location"`
	Data string `json:"data"`
}

type HTTP struct {
	Favicon *Favicon `json:"favicon"`
	ResponseHeaders string `json:"response_headers"`
	Path string `json:"path"`
	Sitemap string `json:"sitemap"`
	HTMLHash string `json:"html_hash"`
	Body string `json:"body"`
	Robots string `json:"robots"`
	Server string `json:"server"`
	XPoweredBy string `json:"x_powered_by"`
	MetaKeywords string `json:"meta_keywords"`
	SecurityText string `json:"security_text"`
	SitemapHash string `json:"sitemap_hash"`
	RobotsHash string `json:"robots_hash"`
	Host string `json:"host"`
	StatusCode int `json:"status_code"`
	Title string `json:"title"`
}

type Location struct {
	Isp string `json:"isp"`
	Radius float64 `json:"radius"`
	Gps []float64 `json:"gps"`
	DistrictCn string `json:"district_cn"`
	StreetEn string `json:"street_en"`
	Owner string `json:"owner"`
	ProvinceEn string `json:"province_en"`
	CountryEn string `json:"country_en"`
	CountryCn string `json:"country_cn"`
	CountryCode string `json:"country_code"`
	CityCn string `json:"city_cn"`
	StreetCn string `json:"street_cn"`
	DistrictEn string `json:"district_en"`
	ProvinceCn string `json:"province_cn"`
	CityEn string `json:"city_en"`
}

type ServiceItem struct {
	ServiceID  string `json:"service_id"`
	Name string `json:"name"`
	Components []*Product `json:"components"`
	Port float64 `json:"port"`
	Product string `json:"product"`
	Transport string `json:"transport"`
	Tags []interface{} `json:"tags"`
	Version string `json:"version"`
	Time time.Time `json:"time"`
}

type Host struct {
	IP string `json:"ip"`
	IsIpv6 bool `json:"is_ipv6"`
	Services []*ServiceItem `json:"services"`
	Asn int `json:"asn"`
	Org string `json:"org"`
	Domains []string `json:"domains"`
	Tags []interface{} `json:"tags"`
	OsName string `json:"os_name"`
	OsVersion string `json:"os_version"`
	Hostname string `json:"hostname"`
	Location *Location `json:"location"`
	CTime time.Time `json:"time"`
}

type Asset struct {
	Service *Service `json:"service"`
	Location *Location `json:"location"`
	Images []interface{} `json:"images"`
	Asn int `json:"asn"`
	Tags []interface{} `json:"tags"`
	OsName string `json:"os_name"`
	Org string `json:"org"`
	OsVersion string `json:"os_version"`
	Hostname string `json:"hostname"`
	Transport string `json:"transport"`
	IP string `json:"ip"`
	Port float64 `json:"port"`
	IsIpv6 bool `json:"is_ipv6"`
	Components []*Product `json:"components"`
	CTime time.Time `json:"time"`
}

type Product struct {
	ID string `json:"id"`
	ProductType []string `json:"product_type"`
	ProductCatalog []string `json:"product_catalog"`
	ProductNameCn string `json:"product_name_cn"`
	ProductNameEn string `json:"product_name_en,omitempty"`
	ProductDescription string `json:"product_description,omitempty"`
	ProductIndustry []string `json:"product_industry,omitempty"`
	ProductLevel  string `json:"product_level"`
	ProductVendor string `json:"product_vendor"`
	ProductDork *ProductDork `json:"product_dork,omitempty"`
	Version string `json:"version"`
}

type ProductDork struct {
	DorkSource string `json:"dork_source"`
	DorkQuery  string `json:"dork_str"`
}
