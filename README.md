# go-quake

To start working with Quake you have to get your token first. You can do this at [https://quake.360.cn](https://quake.360.cn).

### Usage

```go
import "github.com/jinycoo/go-quake/quake"	// with go modules enabled (GO111MODULE=on or outside GOPATH)
```

Simple example of resolving hostnames:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jinycoo/go-quake/quake" // go modules required
)

func main() {
	var ctx = context.Background()
	client := quake.NewQuakeClient(nil)
	res, err := client.ServiceSearch(ctx, `service:"http/ssl"`, 1, 10, false, false)

	if err != nil {
		log.Panic(err)
	} else {
		if res != nil {
			for _, asset := range res.Assets {
				fmt.Println(asset)
			}
			fmt.Printf("%+v", res.Page)
		}
	}
}
```
Output for above:
```bash
&{0xc00009d440 0xc0001a81e0 [] 12874 []  Fastweb SpA  2-229-167-121.ip197.fastwebnet.it tcp 2.229.167.121 8888 false [0xc0000bc540] 2021-08-20 08:54:00.918 +0000 UTC}
&{0xc00009d4a0 0xc0001a84b0 [] 12874 []  Fastweb SpA   tcp 2.229.164.99 8888 false [0xc0000bc600] 2021-08-20 08:54:00.522 +0000 UTC}
&{0xc00009d500 0xc0001a8780 [] 14618 []  Amazon.com, Inc.  ec2-3-214-118-21.compute-1.amazonaws.com tcp 3.214.118.21 500 false [] 2021-08-20 08:54:00.28 +0000 UTC}
&{0xc00009d560 0xc0001a8a50 [] 12874 []  Fastweb SpA   tcp 2.233.127.6 8888 false [0xc0000bc6c0] 2021-08-20 08:54:00.237 +0000 UTC}
&{0xc00009d5c0 0xc0001a8d20 [] 12874 []  Fastweb SpA   tcp 2.238.16.81 8888 false [0xc0000bc780] 2021-08-20 08:54:00.047 +0000 UTC}
&{0xc00009d620 0xc0001a8ff0 [] 12874 []  Fastweb SpA   tcp 2.239.185.97 8888 false [0xc0000bc840] 2021-08-20 08:53:59.668 +0000 UTC}
&{0xc00009d680 0xc0001a92c0 [] 12874 []  Fastweb SpA   tcp 2.233.127.22 8888 false [0xc0000bc900] 2021-08-20 08:53:59.196 +0000 UTC}
&{0xc00009d6e0 0xc0001a9590 [] 12874 []  Fastweb SpA   tcp 2.226.182.33 8888 false [0xc0000bc9c0] 2021-08-20 08:53:58.698 +0000 UTC}
&{0xc00009d740 0xc0001a9860 [] 12874 []  Fastweb SpA  2-227-82-202.ip185.fastwebnet.it tcp 2.227.82.202 8888 false [0xc0000bca80] 2021-08-20 08:53:58.623 +0000 UTC}
&{0xc00009d7a0 0xc0001a9b30 [] 12874 []  Fastweb SpA   tcp 2.234.11.248 8888 false [0xc0000bcb40] 2021-08-20 08:53:58.279 +0000 UTC}
{Num:1 Size:10 Total:910147168}<nil>
```

### Tips and tricks

Every method accepts context in the first argument so you can easily cancel any request.

You can also set config: `mode="debug"` to see the actual request data (method, url, body).

### Implemented REST API

#### Service Search Methods
- [x] /v3/search/quake_service
- [x] /v3/scroll/quake_service
- [x] /v3/aggregation/quake_service GET
- [x] /v3/aggregation/quake_service POST

#### Host Search Methods
- [x] /v3/search/quake_host
- [x] /v3/scroll/quake_host
- [x] /v3/aggregation/quake_host
- [x] /v3/aggregation/quake_host

#### Account Methods
- [x] /v3/user/info

#### Vulnerabilities Methods
- [x] /v3/vulnerability/db/cve/detail/{query}

#### Product Methods
- [ ] /product/vendor
- [ ] /product/catalog
- [ ] /product/industry
- [ ] /user/product
- [ ] /user/product/vendor


### Links
* [360-Quake](https://quake.360.cn)
* [API Documentation](https://quake.360.cn/quake/#/help)
