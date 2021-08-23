/**------------------------------------------------------------**
 * @filename quake/product.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/18 11:43
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

import (
	"context"
	"fmt"
)

const (
	createProductPath = "/user/product"
	getProductIndustryPath = "/product/industry"
	getProductVendorPath = "/product/vendor"
)



type Vendor struct {
	ID string `json:"id"`
	NameZh string `json:"vendor_name_zh"`
	NameEn string `json:"vendor_name_en"`
	Desc string `json:"vendor_description"`
	HomePage string `json:"vendor_homepage"`
	Industries []*Industry `json:"vendor_industry_tag"`
	MTime string `json:"vendor_update_time"`
	ProductCount int `json:"vendor_product_count"`
}

type Industry struct {
	ID string `json:"id"`
	Name string `json:"industry_name"`
	Desc string `json:"industry_description"`
}

func (c *Client) GetVendors(ctx context.Context, vendorName string) ([]*Vendor, error) {
	var vendors = make([]*Vendor, 0)

	req, err := c.NewRequest(GET, fmt.Sprintf("%s?vendor_name=%s", getProductVendorPath, vendorName), nil, nil)
	if err != nil {
		return vendors, err
	}
fmt.Println(req)
	err = c.Do(ctx, req, vendors, nil)

	return vendors, err
}

func (c *Client) GetIndustries(ctx context.Context, size int) ([]*Industry, error) {
	var industries = make([]*Industry, 0)

	req, err := c.NewRequest(GET, fmt.Sprintf("%s?size=%d", getProductIndustryPath, size), nil, nil)
	if err != nil {
		return industries, err
	}

	err = c.Do(ctx, req, industries, nil)

	return industries, err
}

//func (c *Client) AddProduct(ctx context.Context) (string, error) {
//	var product Product
//	payload := &alertCreateRequest{
//		Name:    name,
//		Expires: expires,
//		Filters: &AlertFilters{IP: ip},
//	}
//
//	b, err := json.Marshal(payload)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := c.NewRequest(POST, createProductPath, nil, bytes.NewReader(b))
//	if err != nil {
//		return nil, err
//	}
//
//	if err := c.Do(ctx, req, &product); err != nil {
//		return nil, err
//	}
//
//	return &product, nil
//}
