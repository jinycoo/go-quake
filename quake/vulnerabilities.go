/**------------------------------------------------------------**
 * @filename quake/vulnerabilities.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 18:47
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

import (
	"context"
	"fmt"
)

const (
	vulnerabilitiesPath = "/v3/vulnerability/db/cve/detail"
)

type Vulnerabilities struct {
	VulCpe []string      `json:"vul_cpe"`

	VulCreator struct {
		VulAuthorId string `json:"vul_author_id"`
		VulAuthorName string `json:"vul_author_name"`
	} `json:"vul_creator"`
	VulCveId string `json:"vul_cve_id"`
	VulDescription struct {
		En string `json:"en"`
		Zh string `json:"zh"`
	} `json:"vul_description"`
	VulDetail string `json:"vul_detail"`
	VulHazardRating string `json:"vul_hazard_rating"`
	VulInfluences []*VulInfluence `json:"vul_influence"`
	VulName string `json:"vul_name"`
	VulOtherId []*VulDBObj `json:"vul_other_id"`
	VulPlug string `json:"vul_plug"`
	VulPoc string `json:"vul_poc"`
	VulQhvId string `json:"vul_qhv_id"`
	VulReferences []string `json:"vul_references"`
	VulSample []string `json:"vul_sample"`
	VulSolutionOfficial string `json:"vul_solution_official"`
	VulSolutionTemp string `json:"vul_solution_temp"`
	VulType string `json:"vul_type"`
}

type VulInfluence struct {
	ProductName string `json:"product_name"`
	ProductVersions []string `json:"product_versions"`
	VendorName string `json:"vendor_name"`
}

type VulDBObj struct {
	VulDB string `json:"vul_db"`
	VulDBId string `json:"vul_db_id"`
}

func (c *Client) GetVulnInfo(ctx context.Context, query string) (*Vulnerabilities, error) {
	var vuln Vulnerabilities
	req, err := c.NewRequest(GET, fmt.Sprintf("%s/%s", vulnerabilitiesPath, query) , nil, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &vuln, nil); err != nil {
		return nil, err
	}
	return &vuln, err
}

