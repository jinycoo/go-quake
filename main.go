package main

import (
	"context"
	"fmt"

	"github.com/jinycoo/go-quake/quake"
)

func main() {
	var ctx = context.Background()
	client := quake.NewQuakeClient(nil)
	//res, err := client.HostSearch(ctx, `service:"http/ssl"`, 1, 10, false, false)
	//res, err := client.ServiceIPListDepthSearch(ctx, `2.233.127.6,2.229.167.121`, "611f5974c35468ee50bac1c3" , 10, false)
	//res, err := client.ServiceIPRuleDepthSearch(ctx, `scan`, "" , 10, false)
	//res, err := client.ServiceDepthSearch(ctx, `service:"http/ssl"`, "" , 10, false)
	//res, err := client.ServiceIPRuleSearchAgg(ctx, `scan`, 10, []string{"service", "asn"}, false)
	//res, err := client.ServiceIPListSearchAgg(ctx, `2.233.127.6,2.229.167.121`, 10, []string{"service", "asn"}, false)
	//res, err := client.ServiceSearchAgg(ctx, `service:"http/ssl"`, 10, []string{"service", "asn"}, false)
	//res, err := client.GetServiceAggFields(ctx)
	//res, err := client.ServiceIPListSearch(ctx, `2.233.127.6,2.229.167.121`, 1, 10, false, false)
	//res, err := client.ServiceIPRuleSearch(ctx, `scan`, 1, 10, false, false)
	//res, err := client.ServiceSearch(ctx, `service:"http/ssl"`, 1, 10, false, false)
	res, err := client.GetVulnInfo(ctx, `CNNVD-199612-001`)
	//res, err := client.GetUserInfo(ctx)
	//res, err := client.GetVendors(ctx, "Huawei华为技术有限公司")
	if res != nil {
		for _, asset := range res.VulCpe {
			fmt.Println(asset)
		}
	}
	fmt.Println(err)
	//fmt.Println(res, err)
}
