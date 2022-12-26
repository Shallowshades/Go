package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
字符串
	一个字符串是一个不可改变的字节序列。
	字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。
	文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列

	内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

	试图访问超出字符串索引范围的字节将会导致panic异常

	第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节

	+操作符将两个字符串连接构造一个新字符串

	字符串可以用==和<进行比较；比较通过`逐个字节比较完成`的，因此比较的结果是字符串自然编码的顺序。

	字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然也可以给一个字符串变量分配一个新字符串值。

	因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的

	不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。
	同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的。
	在这两种情况下都没有必要分配新的内存。

	因为Go语言源文件总是用UTF8编码，并且Go语言的文本字符串也以UTF8编码的方式处理，因此可以将Unicode码点也写到字符串面值中

	\a      响铃
	\b      退格
	\f      换页
	\n      换行
	\r      回车	回到本行的开始处
	\t      制表符
	\v      垂直制表符
	\'      单引号（只用在 '\'' 形式的rune符号面值中）
	\"      双引号（只用在 "..." 形式的字符串面值中）
	\\      反斜杠

	可以通过十六进制或八进制转义在字符串面值中包含任意的字节。一个十六进制的转义形式是\xhh，其中两个h表示十六进制数字（大写或小写都可以）。一个八进制转义形式是\ooo，包含三个八进制的o数字（0到7），但是不能超过\377（对应一个字节的范围，十进制为255）。每一个单一的字节表达一个特定的值。

字符串面值
	原生的字符串面值形式是`...`，使用反引号代替双引号
	全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行
	在原生字符串面值内部是无法直接写`字符的，可以用八进制或十六进制转义或+"`"连接字符串常量完成
	唯一的特殊处理是会删除回车以保证在所有平台上的值都是一样的，包括那些把回车也放入文本文件的系统（Windows系统会把回车和换行一起放入文本文件中）。

	原生字符串面值用于编写正则表达式会很方便，因为正则表达式往往会包含很多反斜杠。原生字符串面值同时被广泛应用于HTML模板、JSON面值、命令行提示信息以及那些需要扩展到多行的场景

Unicode
	每个符号都分配一个唯一的Unicode码点，Unicode码点对应Go语言中的rune整数类型
	通用的表示一个Unicode码点的数据类型是int32，也就是Go语言中rune对应的类型；它的同义词rune符文正是这个意思
	可以将一个符文序列表示为一个int32序列。这种编码方式叫UTF-32或UCS-4，每个Unicode码点都使用同样大小的32bit来表示
	这种方式比较简单统一，但是它会浪费很多存储空间，因为大多数计算机可读的文本是ASCII字符，本来每个ASCII字符只需要8bit或1字节就能表示

UTF-8
	UTF8是一个将Unicode码点编码为字节序列的变长编码。
	UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。
	每个符号编码后第一个字节的高端bit位用于表示编码总共有多少个字节。
	如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容
	如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头。更大的Unicode码点也是采用类似的策略处理。

	0xxxxxxx                             runes 0-127    (ASCII)
	110xxxxx 10xxxxxx                    128-2047       (values <128 unused)
	1110xxxx 10xxxxxx 10xxxxxx           2048-65535     (values <2048 unused)
	11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff (other values unused)

	变长的编码无法直接通过索引来访问第n个字符，但是UTF8编码获得了很多额外的优点。
	首先UTF8编码比较紧凑，完全兼容ASCII码，并且可以自动同步：它可以通过向前回朔最多3个字节就能确定当前字符编码的开始字节的位置。它也是一个前缀编码，所以当从左向右解码时不会有任何歧义也并不需要向前查看（像GBK之类的编码，如果不知道起点位置则可能会出现歧义）。
	没有任何字符的编码是其它字符编码的子串，或是其它编码序列的字串，因此搜索一个字符时只要搜索它的字节编码序列即可，不用担心前后的上下文会对搜索结果产生干扰。同时UTF8编码的顺序和Unicode码点的顺序一致，因此可以直接排序UTF8编码序列。同时因为没有嵌入的NUL(0)字节，可以很好地兼容那些使用NUL作为字符串结尾的编程语言。

	Go语言的源文件采用UTF8编码，并且Go语言处理UTF8编码的文本也很出色。
	unicode包提供了诸多处理rune字符相关功能的函数（比如区分字母和数字，或者是字母的大写和小写转换等），unicode/utf8包则提供了用于rune字符序列的UTF8编码和解码的功能。

	有很多Unicode字符很难直接从键盘输入，并且还有很多字符有着相似的结构；有一些甚至是不可见的字符（中文和日文就有很多相似但不同的字）。
	Go语言字符串面值中的Unicode转义字符可以通过Unicode码点输入特殊的字符。有两种形式：\uhhhh对应16bit的码点值，\Uhhhhhhhh对应32bit的码点值，其中h是一个十六进制数字；一般很少需要使用32bit的形式。每一个对应码点的UTF8编码。例如：下面的字母串面值都表示相同的值
	"世界"
	"\xe4\xb8\x96\xe7\x95\x8c"  \x对应8bit码点值
	"\u4e16\u754c"				\u对应16bit码点值
	"\U00004e16\U0000754c"		\U对应32bit码点值

	Unicode转义也可以使用在rune字符中

	对于小于256的码点值可以写在一个十六进制转义字节中，例如\x41对应字符'A'，但是对于更大的码点则必须使用\u或\U转义形式。因此，\xe4\xb8\x96并不是一个合法的rune字符，虽然这三个字节对应一个有效的UTF8编码的码点。

	得益于UTF8编码优良的设计，诸多字符串操作都不需要解码操作。
	可以不用解码直接测试一个字符串是否是另一个字符串的前缀测试、后缀测试、测试包含子串

	每一个UTF8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符\uFFFD，在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号"?"。当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF8字符串

	UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。

	string -> []rune 相当于对它们进行UTF8编码 字符 -> 数字
	[]rune -> string 相当于对它们进行UTF8解码 数字 -> 字符

字符串和Byte切片
	标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
	strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能
	bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效
	strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
	unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。所有的这些函数都是遵循Unicode标准定义的字母、数字等分类规范。strings包也有类似的函数，它们是ToUpper和ToLower，将原始字符串的每个字符都做相应的转换，然后返回新的字符串。

	path和path/filepath包提供了关于文件路径名更一般的函数操作
	使用斜杠分隔路径可以在任何操作系统上工作
	斜杠本身不应该用于文件名，但是在其他一些领域可能会用于文件名，例如URL路径组件
	相比之下，path/filepath包则使用操作系统本身的路径规则，例如POSIX系统使用/foo/bar，而Microsoft Windows使用c:\foo\bar

	字符串和字节slice之间可以相互转换
	从概念上讲，一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组
	编译器的优化可以避免在一些场景下分配和复制字符串数据，但总的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。将一个字节slice转换到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读

	为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数
	strings包
		func Contains(s, substr string) bool
		func Count(s, sep string) int
		func Fields(s string) []string
		func HasPrefix(s, prefix string) bool
		func Index(s, sep string) int
		func Join(a []string, sep string) string

	bytes包
		func Contains(b, subslice []byte) bool
		func Count(s, sep []byte) int
		func Fields(s []byte) [][]byte
		func HasPrefix(s, prefix []byte) bool
		func Index(s, sep []byte) int
		func Join(s [][]byte, sep []byte) []byte

	它们之间唯一的区别是字符串类型参数被替换成了字节slice类型的参数

	bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要初始化，因为零值也是有效的

	当向bytes.Buffer添加任意字符的UTF8编码时，最好使用bytes.Buffer的WriteRune方法，但是WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效

字符串和数字的转换
	除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见。由strconv包提供这类转换功能。
	将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)
	FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要包含有附加额外信息的时
	如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUint函数
	ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。在任何情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型
	使用fmt.Scanf来解析输入的字符串和数字，特别是当字符串和数字混合在一行的时候，它可以灵活处理不完整或不规则的输入
*/

func main() {

	{
		s := "hello, world"
		fmt.Println(len(s))
		fmt.Println(s[0], s[7])

		str := "abc你我他"
		fmt.Println(str)
		fmt.Println(len(str)) //一个中文三个字节,len = 12
		for i := 0; i < len(str); i++ {
			fmt.Printf("%c ", str[i])
		}
		fmt.Println()
		for i := 0; i < len(str); i++ {
			fmt.Printf("%q ", str[i])
		}
		fmt.Println()
		for _, c := range str {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
		for _, c := range str {
			fmt.Printf("%q ", c)
		}
		fmt.Println()

		fmt.Println(s[0:5])
		fmt.Println(s[:5])
		fmt.Println(s[7:])
		fmt.Println(s[:])
		fmt.Println("goodbye" + s[5:])
	}

	// + - < > == !=
	{
		s := "left foot"
		t := s
		s += ", right foot"
		fmt.Println("t = ", t)
		fmt.Println("s = ", s)

		//s[0] = 'L' // compile error: cannot assign to s[0]

		if s >= t {
			fmt.Println("s >= t")
		} else {
			fmt.Println("s < t")
		}

		s1 := "abc"
		s2 := "abcd"
		fmt.Println("s1 = ", s1)
		fmt.Println("s2 = ", s2)
		fmt.Println("s1 < s2 ? ", (s1 < s2))

		s3 := "abd"
		s4 := "abcd"
		fmt.Println("s3 = ", s3)
		fmt.Println("s4 = ", s4)
		fmt.Println("s3 < s4 ? ", (s3 < s4))
	}

	//字面值 ``
	{
		const GoUsage = `Go is a tool for managing Go source code.

		Usage:
			go command [arguments]
			...`
		fmt.Println(GoUsage)
	}

	//utf-8
	{
		s1 := "世界"
		s2 := "\xe4\xb8\x96\xe7\x95\x8c"
		s3 := "\u4e16\u754c"
		s4 := "\U00004e16\U0000754c"

		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)

		c1 := '世'
		c2 := '\u4e16'
		c3 := '\U00004e16'

		fmt.Println("c1 == c2 ? ", (c1 == c2))
		fmt.Println("c1 == c3 ? ", (c1 == c3))

		// c := '\xe4\xb8\x96' error
	}

	{
		//字符串包含13个字节，以UTF8形式编码，但是只对应9个Unicode字符
		s := "Hello, 世界"
		fmt.Println("s = ", s)
		fmt.Println("len(s) = ", len(s))
		fmt.Println(utf8.RuneCountInString(s))

		//为了处理这些真实的字符，需要一个UTF8解码器。unicode/utf8包提供了该功能
		for i := 0; i < len(s); {
			//每一次调用DecodeRuneInString函数都返回一个r和长度
			//r对应字符本身
			//size长度对应r采用UTF8编码后的编码字节数目
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\n", i, r)
			i += size
		}

		//幸运的是，Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串
		//需要注意的是对于非ASCII，索引更新的步长将超过1个字节。
		for i, r := range s {
			fmt.Printf("%d\t%c\n", i, r)
		}

		//使用一个简单的循环来统计字符串中字符的数目
		n := 0
		for range s {
			n++
		}
		fmt.Println("n = ", n)
		fmt.Println("utf8.RuneCountInString(s) = ", utf8.RuneCountInString(s))
	}

	{
		//string -> []rune 相当于对它们进行UTF8解码
		s := "プログラム" //"program" in Japanese katakana
		fmt.Println("s = ", s)
		fmt.Printf("% x\n", s) //% x参数用于在每个十六进制数字前插入一个空格。
		r := []rune(s)
		fmt.Printf("%x\n", r)

		//[]rune -> string 相当于对它们进行UTF8编码
		fmt.Println(string(r))

		//将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串
		fmt.Println(string(rune(65)))
		//fmt.Println(string(65))  warning
		fmt.Println(string(rune(0x4eac)))

		fmt.Println(string(rune(1234567))) // "?"

	}

	{
		fmt.Println(commaInteger("1234567890"))
		fmt.Println(commaFloat("123456.7890"))
	}

	{
		//字符串和字节slice之间可以相互转换
		s := "abc"
		b := []byte(s)
		s2 := string(b)
		fmt.Println(s2)
	}

	{
		fmt.Println(intsToString([]int{1, 2, 3, 4}))

		//将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)
		x := 123
		y := fmt.Sprintf("%d", x)
		fmt.Println(y, strconv.Itoa(x))

		//FormatInt和FormatUint函数可以用不同的进制来格式化数字
		fmt.Println(strconv.FormatInt(int64(x), 2))

		//fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要包含有附加额外信息的时候
		s := fmt.Sprintf("x=%b", x) // "x=1111011"
		fmt.Println(s)

		//如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUint函数
		i1, err := strconv.Atoi("123")
		if err != nil {
			os.Exit(1)
		}
		i2, err := strconv.ParseInt("123", 10, 64) //ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(i1)
		fmt.Println(i2)
	}
	{
		//字符串连接
		s1 := "hello, world"
		s2 := "hello, world"
		s3 := "hello, world"

		fmt.Printf("s1: %v\n", s1)
		fmt.Printf("s2: %v\n", s2)
		fmt.Printf("s3: %v\n", s3)

		fmt.Printf("s1: %p\n", &s1)
		fmt.Printf("s2: %p\n", &s2)
		fmt.Printf("s3: %p\n", &s3)

		fmt.Printf("s1[0] : %q\n", s1[0])

		s := strings.Join([]string{s1, s2, s3}, " | ")
		fmt.Printf("s: %v\n", s)

		//写在缓冲区，效率较高
		var buf bytes.Buffer
		buf.WriteString(s1)
		buf.WriteByte('-')
		buf.WriteString(s2)
		buf.WriteByte('-')
		buf.WriteString(s3)
		fmt.Printf("buf: %v\n", buf)
		fmt.Printf("buf.Bytes(): %v\n", buf.Bytes())
		s4 := buf.String()
		fmt.Printf("s4: %v\n", s4)

		//strings.Builder
		var b strings.Builder
		for i := 3; i >= 1; i-- {
			fmt.Fprintf(&b, "%d...", i)
		}
		b.WriteString("ignition")
		fmt.Println(b.String())

		//strings.Replacer
		r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
		fmt.Println(r.Replace("This is <b>HTML</b>!"))

		fmt.Println(strings.Count("five", ""))

		//Cut 裁剪去一部分
		show := func(s, sep string) {
			before, after, found := strings.Cut(s, sep)
			fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
		}
		show("Gopher", "Go")
		show("Gopher", "ph")
		show("Gopher", "er")
		show("Gopher", "Badger")

		fmt.Printf("%d\n", '\n')

		//Index
		fmt.Printf("strings.Index(\"hello,world\", \"lo\"): %v\x0a", strings.Index("hello,world", "lo"))

		var sd string = "gogogo"
		fmt.Printf("sd: %q\n", sd)
	}

	{
		//test 3.10
		fmt.Println(comma("12345678"))

		//test 3.11
		fmt.Println(comma2("12345.67890"))
		fmt.Println(comma2("+12345.67890"))
		fmt.Println(comma2("-12345.67890"))

		//test 3.12
		fmt.Println(equal("112233", "123123"))
	}
}

// 判断s是否存在prefix前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 判断s是否存在prefix后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 判断s是否包含子串
// 真实的代码包含了一个用哈希技术优化的Contains 实现
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// basename函数灵感源于Unix shell的同名工具。
// 该实现的版本中，basename(s)将看起来像是系统路径的前缀删除，同时将看似文件类型的后缀名部分删除
func basename1(s string) string {
	// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
	// fmt.Println(basename("a/b/c.go")) // "c"
	// fmt.Println(basename("c.d.go"))   // "c.d"
	// fmt.Println(basename("abc"))      // "abc"

	//删除路径
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//删除后缀
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 使用strings.LastIndex库函数
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]                    //若slash=-1, s = s
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//函数的功能是将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。

// 从右到左 3个字符划分一组
func commaRtol(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaRtol(s[:len(s)-3]) + "," + s[len(s)-3:]
}

// 从左到右
func commaLtor(s string) string {
	if len(s) <= 3 {
		return s
	}
	return s[:3] + "," + commaLtor(s[3:])
}

// 整数类型
func commaInteger(s string) string {
	return commaRtol(s)
}

// 浮点数类型
func commaFloat(s string) string {
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		return commaRtol(s)
	} else {
		return commaRtol(s[:dot]) + "." + commaLtor(s[dot+1:])
	}
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作
func reverse(s string) string {
	str := []byte(s)
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}

func comma(s string) string {
	s = reverse(s)
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return reverse(buf.String())
}

// 练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
func comma2(s string) string {
	flag := ""
	if s[0] == '+' || s[0] == '-' {
		flag = string(s[0])
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		return flag + comma(s)
	} else {
		var buf bytes.Buffer
		for i := 0; i < len(s)-dot-1; i++ {
			if i != 0 && i%3 == 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte(s[i+dot+1])
		}
		return flag + comma(s[:dot]) + "." + buf.String()
	}
}

// 练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func equal(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var book []bool
	book = make([]bool, len(s2))
	for _, c := range s1 {
		tail := len(s2)
		for {
			idx := strings.LastIndex(s2[:tail], string(c))
			if idx == -1 {
				return false
			}
			if book[idx] == true {
				tail = idx
				continue
			}
			book[idx] = true
			break
		}
	}
	return true
}
