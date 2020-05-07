package main

import (
	// TheGoApiLerarnNote 项目有go.mod import必须与 module 相同
	"github.com/PlagueCat-Miao/TheGoApiLerarnNote/GINtest"
	// 只要TheGoRpcLerarnNote 项目里没go.mod 路径大小写都行 go get 会在 pkg 自动写好路径 ，且你不要管带不带！ 大不大写
	"github.com/plaguecat-miao/TheGoRpcLerarnNote/a"
)

func main() {
	var kk GINtest.Login;
	A.Hello()
	println(kk)
}
