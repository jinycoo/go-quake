/**------------------------------------------------------------**
 * @filename quake/member.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 16:55
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

import (
	"context"
)

const (
	profilePath = "/v3/user/info"
)

type QUser struct {
	ID string      `json:"id"`

	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	} `json:"user"`

	Credit   uint   `json:"credit"`
	PersistentCredit uint `json:"persistent_credit"`
	Mobile  string `json:"mobile_phone"`
	Role    []*MRole `json:"role"`
}

type MRole struct {
	Fullname string `json:"fullname"`
	Priority uint8 `json:"priority"`
	Credit   uint `json:"credit"`
}

func (c *Client) GetUserInfo(ctx context.Context) (*QUser, error) {
	var user QUser
	req, err := c.NewRequest(GET, profilePath, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = c.Do(ctx, req, &user, nil); err != nil {
		return nil, err
	}
	return &user, err
}
