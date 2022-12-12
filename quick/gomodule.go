package main

/*
	gopath工作模式
	缺点：
		1.无版本控制概念
		2.无法同步一致第三方版本号
		3.无法指定当前项目引用的第三方版本号const

	go mod命令
		download    download modules to local cache				下载go.mod文件中指明的所有依赖
		edit        edit go.mod from tools or scripts			编辑go.mod文件
		graph       print module requirement graph				查看现有依赖结构
		init        initialize new module in current directory	生成go.mod文件
		tidy        add missing and remove unused modules		整理现有的依赖
		vendor      make vendored copy of dependencies			导出项目所有的依赖到vendor目录
		verify      verify dependencies have expected content	校验一个模块是否被篡改过
		why         explain why packages or modules are needed	查看为什么需要依赖某模块

	go env
		$ go env
			GO111MODULE="auto"
			GOPROXY="https://proxy.golang.org,direct"
			GONOPROXY=""
			GOSUMDB="sum.golang.org"
			GONOSUMDB=""
			GOPRIVATE=""
			...

	GO111MODULE

		Go语言提供了 GO111MODULE这个环境变量来作为 Go modules 的开关，其允许设置以下参数：
		auto：只要项目包含了 go.mod文件的话启用Go modules，目前在 Go1.11 至 Go1.14 中仍然是默认值。
		on：启用 Go modules，推荐设置，将会是未来版本中的默认值。
		off：禁用 Go modules，不推荐设置。
		`go env -w GO111MODULE=on`

	GOPROXY

		这个环境变量主要是用于设置 Go 模块代理（Go module proxy）,
		其作用是用于使 Go 在后续拉取模块版本时直接通过镜像站点来快速拉取。
		GOPROXY 的默认值是：https://proxy.golang.org,direct
		proxy.golang.org国内访问不了,需要设置国内的代理.
		阿里云https://mirrors.aliyun.com/goproxy/
		七牛云https://goproxy.cn,direct
		`$ go env -w GOPROXY=https://goproxy.cn,direct`
	GOPROXY 的值是一个以英文逗号 “,” 分割的 Go 模块代理列表，允许设置多个模块代理，假设你不想使用，也可以将其设置为 “off” ，这将会禁止 Go 在后续操作中使用任何 Go 模块代理。

	“direct” 是一个特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取（比如 GitHub 等），场景如下：当值列表中上一个 Go 模块代理返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，也就是回到源地址去抓取，而遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误。

	GOSUMDB（校验）

	它的值是一个 Go checksum database，用于在拉取模块版本时（无论是从源站拉取还是通过 Go module proxy 拉取）保证拉取到的模块版本数据未经过篡改，若发现不一致，也就是可能存在篡改，将会立即中止。
	GOSUMDB 的默认值为：sum.golang.org，在国内也是无法访问的，但是 GOSUMDB 可以被 Go 模块代理所代理

	GONOPROXY/GONOSUMDB/GOPRIVATE

	这三个环境变量都是用在当前项目依赖了私有模块，
	一般建议直接设置 GOPRIVATE，它的值将作为 GONOPROXY 和 GONOSUMDB 的默认值
	并且它们的值都是一个以英文逗号 “,” 分割的模块路径前缀，也就是可以设置多个
*/

func main() {

}
