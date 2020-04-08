# login_backend

## 描述：
使用gin框架和mysql的简单登录验证后台，前台使用vue来实现。

* go 版本
> go version go1.14.1 windows/amd64

* go mod
> go mod init login_backend
> 
> 执行go build，go run，go test会自动触发更新go.mod文件
>
> go list -u -m all 查看所有直接及间接依赖的可用版本
>
> go get -u login_backend 升级所有相关依赖到最新版本。

* gin
> import "github.com/gin-gonic/gin"

* mysql
> import "database/sql"
>
> import _ "github.com/go-sql-driver/mysql"


