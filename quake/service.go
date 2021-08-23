/**------------------------------------------------------------**
 * @filename quake/host.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/20 15:30
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
	queryServicesPath   = "/v3/search/quake_service"
	queryScrollServicesPath   = "/v3/scroll/quake_service"
	queryAggServiceFieldsPath = "/v3/aggregation/quake_service"
	queryAggServicesPath = "/v3/aggregation/quake_service"
)

// QueryMatch represents a matched tag.
type QueryMatch struct {
	Query string `json:"query"`
	Rule  string `json:"rule,omitempty"`
	IpList []string `json:"ip_list,omitempty"`
	PageSize int `json:"size,omitempty"`
	IgnoreCache bool `json:"ignore_cache,omitempty"`
	Deduplication bool `json:"latest,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

type QuerySearchMatch struct {
	QueryMatch
	Start  int  `json:"start,omitempty"`
}

type QueryScrollMatch struct {
	QueryMatch
	PaginationId string `json:"pagination_id,omitempty"`
}

type QueryAggMatch struct {
	QueryMatch
	Fields  []string `json:"aggregation_list"`
}

type SearchResult struct {
	Assets []*Asset `json:"assets"`
	*CommonPagination
}

type AggInfo struct {
	Key   interface{} `json:"key"`
	Count float64 `json:"doc_count"`
}

type DepthSearchResult struct {
	Assets []*Asset `json:"assets"`
	Page *DepthPage `json:"page"`
}

type DepthPage struct {
	Total int `json:"total"`
	PaginationID string `json:"page_id"`
}

func (c *Client) ServiceSearch(ctx context.Context, query string, pn, ps int, cache, deduplication bool) (*SearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

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
	req, err := c.NewRequest(POST, queryServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var page Pagination
	if err = c.Do(ctx, req, &serviceAssets, &page); err != nil {
		return nil, err
	}
	return &SearchResult{
		Assets: serviceAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil

}
func (c *Client) ServiceIPRuleSearch(ctx context.Context, rule string, pn, ps int, cache, deduplication bool) (*SearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

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
	req, err := c.NewRequest(POST, queryServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var page Pagination
	if err = c.Do(ctx, req, &serviceAssets, &page); err != nil {
		return nil, err
	}
	return &SearchResult{
		Assets: serviceAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil
}
func (c *Client) ServiceIPListSearch(ctx context.Context, ips string, pn, ps int, cache, deduplication bool) (*SearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

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
	req, err := c.NewRequest(POST, queryServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var page Pagination
	if err = c.Do(ctx, req, &serviceAssets, &page); err != nil {
		return nil, err
	}
	return &SearchResult{
		Assets: serviceAssets,
		CommonPagination: &CommonPagination{
			Page: Page {
				Total: page.Total,
				Num: page.PageIndex,
				Size: page.PageSize,
			},
		},
	}, nil
}

func (c *Client) GetServiceAggFields(ctx context.Context) ([]string, error) {
	var fields = make([]string, 0)

	req, err := c.NewRequest(GET, queryAggServiceFieldsPath, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &fields, nil); err != nil {
		return nil, err
	}

	return fields, nil
}

func (c *Client) ServiceSearchAgg(ctx context.Context, query string, topN int, aggFields []string, cache bool) (interface{}, error) {
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

	req, err := c.NewRequest(POST, queryAggServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ServiceIPRuleSearchAgg(ctx context.Context, rule string, topN int, aggFields []string, cache bool) (interface{}, error) {
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

	req, err := c.NewRequest(POST, queryAggServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ServiceIPListSearchAgg(ctx context.Context, ips string, topN int, aggFields []string, cache bool) (interface{}, error) {
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

	req, err := c.NewRequest(POST, queryAggServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &res, nil); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ServiceDepthSearch(ctx context.Context, query, pid string, ps int, cache bool) (*DepthSearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

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
	req, err := c.NewRequest(POST, queryScrollServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	var meta Meta
	if err = c.Do(ctx, req, &serviceAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchResult{
		Assets: serviceAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}
func (c *Client) ServiceIPRuleDepthSearch(ctx context.Context, rule, pid string, ps int, cache bool) (*DepthSearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

	payload := &QueryScrollMatch{}
	payload.Rule = rule
	payload.PaginationId = pid
	payload.PageSize = ps
	payload.IgnoreCache = cache

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryScrollServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err = c.Do(ctx, req, &serviceAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchResult{
		Assets: serviceAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}
func (c *Client) ServiceIPListDepthSearch(ctx context.Context, ips, pid string, ps int, cache bool) (*DepthSearchResult, error) {
	var serviceAssets = make([]*Asset, 0)

	payload := &QueryScrollMatch{}
	payload.IpList = strings.Split(ips, ",")
	payload.PaginationId = pid
	payload.PageSize = ps
	payload.IgnoreCache = cache

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest(POST, queryScrollServicesPath, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err = c.Do(ctx, req, &serviceAssets, &meta); err != nil {
		return nil, err
	}

	return &DepthSearchResult{
		Assets: serviceAssets,
		Page:  &DepthPage{
			Total: meta.Total,
			PaginationID: meta.PageID,
		},
	}, nil
}