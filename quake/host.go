/**------------------------------------------------------------**
 * @filename quake/query.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 15:58
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	queryHostsPath   = "/v3/search/quake_host"
	queryScrollHostsPath   = "/v3/scroll/quake_host"
	queryAggHostFieldsPath = "/v3/aggregation/quake_host"
	queryAggHostsPath = "/v3/aggregation/quake_host"
)

type SearchHResult struct {
	Assets []*Host `json:"assets"`
	*CommonPagination
}

type DepthSearchHResult struct {
	Assets []*Host `json:"assets"`
	Page *DepthPage `json:"page"`
}

func (c *Client) HostSearch(ctx context.Context, query string, pn, ps int, cache, deduplication bool) (*SearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QuerySearchMatch{}
	payload.Query = query
	if pn == 0 {
		pn = 1
	}
	if ps > 10000 {
		return nil, errors.New("page size too large")
	}
	payload.Start = (pn - 1) * ps
	payload.PageSize = ps
	payload.IgnoreCache = cache
	payload.Deduplication = deduplication

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
    var page Pagination
	if err = c.Do(ctx, req, &hostAssets, &page); err != nil {
		return nil, err
	}
	return &SearchHResult{
		Assets: hostAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil

}
func (c *Client) HostIPRuleSearch(ctx context.Context, rule string, pn, ps int, cache, deduplication bool) (*SearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QuerySearchMatch{}
	payload.Rule = rule
	if pn == 0 {
		pn = 1
	}
	if ps > 10000 {
		return nil, errors.New("page size too large")
	}
	payload.Start = (pn - 1) * ps
	payload.PageSize = ps
	payload.IgnoreCache = cache
	payload.Deduplication = deduplication

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var page Pagination
	if err = c.Do(ctx, req, &hostAssets, &page); err != nil {
		return nil, err
	}
	return &SearchHResult{
		Assets: hostAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil
}
func (c *Client) HostIPListSearch(ctx context.Context, ips string, pn, ps int, cache, deduplication bool) (*SearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QuerySearchMatch{}
	payload.IpList = strings.Split(ips, ",")
	if pn == 0 {
		pn = 1
	}
	if ps > 10000 {
		return nil, errors.New("page size too large")
	}
	payload.Start = (pn - 1) * ps
	payload.PageSize = ps
	payload.IgnoreCache = cache
	payload.Deduplication = deduplication

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var page Pagination
	if err = c.Do(ctx, req, &hostAssets, &page); err != nil {
		return nil, err
	}
	return &SearchHResult{
		Assets: hostAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil
}

func (c *Client) GetHostAggFields(ctx context.Context) ([]string, error) {
	var fields = make([]string, 0)

	req, err := c.NewRequest(GET, queryAggHostFieldsPath, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &fields, nil); err != nil {
		return nil, err
	}

	return fields, nil
}

func (c *Client) HostSearchAgg(ctx context.Context, query string, topN int, aggFields []string, cache bool) (interface{}, error) {
	var res = make(map[string][]*AggInfo)

	payload := &QueryAggMatch{}
	payload.Query = query
	payload.Fields = aggFields
	payload.PageSize = topN
	payload.IgnoreCache = cache
	payload.Deduplication = false

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(POST, queryAggHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) HostIPRuleSearchAgg(ctx context.Context, rule string, topN int, aggFields []string, cache bool) (interface{}, error) {
	var res = make(map[string][]*AggInfo)

	payload := &QueryAggMatch{}
	payload.Rule = rule
	payload.Fields = aggFields
	payload.PageSize = topN
	payload.IgnoreCache = cache
	payload.Deduplication = false

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(POST, queryAggHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) HostIPListSearchAgg(ctx context.Context, ips string, topN int, aggFields []string, cache bool) (interface{}, error) {
	var res = make(map[string][]*AggInfo)

	payload := &QueryAggMatch{}
	payload.IpList = strings.Split(ips, ",")
	payload.Fields = aggFields
	payload.PageSize = topN
	payload.IgnoreCache = cache
	payload.Deduplication = false

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(POST, queryAggHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) HostDepthSearch(ctx context.Context, query, pid string, ps int, cache bool) (*DepthSearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QueryScrollMatch{}
	payload.Query = query
	if len(pid) > 0 {
		payload.PaginationId = pid
	}
	payload.PageSize = ps
	payload.IgnoreCache = cache

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryScrollHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
fmt.Println(string(b))
	var meta Meta
	if err = c.Do(ctx, req, &hostAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchHResult{
		Assets: hostAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}
func (c *Client) HostIPRuleDepthSearch(ctx context.Context, rule, pid string, ps int, cache bool) (*DepthSearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QueryScrollMatch{}
	payload.Rule = rule
	payload.PaginationId = pid
	payload.PageSize = ps
	payload.IgnoreCache = cache

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryScrollHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err = c.Do(ctx, req, &hostAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchHResult{
		Assets: hostAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}
func (c *Client) HostIPListDepthSearch(ctx context.Context, ips, pid string, ps int, cache bool) (*DepthSearchHResult, error) {
	var hostAssets = make([]*Host, 0)

	payload := &QueryScrollMatch{}
	payload.IpList = strings.Split(ips, ",")
	payload.PaginationId = pid
	payload.PageSize = ps
	payload.IgnoreCache = cache

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryScrollHostsPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err = c.Do(ctx, req, &hostAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchHResult{
		Assets: hostAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}
