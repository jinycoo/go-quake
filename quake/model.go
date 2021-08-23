/**------------------------------------------------------------**
 * @filename quake/model.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/18 10:10
 * @desc     go-quake - ..
 **------------------------------------------------------------**/
package quake

type CommonPagination struct {
	Page Page `json:"page"`
}

type Page struct {
	Num   int `json:"num"`
	Size  int `json:"size"`
	Total int `json:"total"`
}


type ResultPagination struct {
	List interface{} `json:"list"`
	*CommonPagination
}