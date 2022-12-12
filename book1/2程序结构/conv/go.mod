module conversion

go 1.19

//导入本地包

//require 是声明我们引入的模块和版本。
require conv v0.0.0

//replace 是进行替换对于的模块指向路径。 位置是相对引用模块的go.mod文件的 与 被引用模块的go.mod
replace conv => ../Testpackage
