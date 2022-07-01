# golang 如何写注解文档
> [作为 Gopher，你知道 Go 的注释即文档应该怎么写吗？](https://cloud.tencent.com/developer/article/1959696)

## 命令
```shell
$ go install golang.org/x/tools/cmd/godoc
$ godoc
```
默认端口：6060 
```
cd XXXX; godoc -http=:6060
```
其中 XXXX 是包含 go.mod 的一个仓库目录。  
假设 XXX 是我的 jsonvalue 库的本地目录，根据 go.mod，这个库的地址是 `github.com/Andrew-M-C/go.jsonvalue`。  
那么我就可以在浏览器中打开 `http://${IP}:${PORT}/pkg/github.com/Andrew-M-C/go.jsonvalue/`，就可以访问我的 jsonvalue 库的 GoDoc 页面了。  
并且可以访问 `go.mod` 中所有库的 godoc。 
