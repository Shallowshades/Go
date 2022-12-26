package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	//proto()
	//work48()
	work49()
}

func proto() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		//ReadRune方法执行UTF-8解码并返回三个值：解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		//如果输入的是无效的UTF-8编码的字符，返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// 练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
func work48() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	types := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		//ReadRune方法执行UTF-8解码并返回三个值：解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		//如果输入的是无效的UTF-8编码的字符，返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		switch {
		case unicode.IsControl(r):
			types["IsControl"]++
			fallthrough
		case unicode.IsDigit(r):
			types["IsDigit"]++
			fallthrough
		case unicode.IsGraphic(r):
			types["IsGraphic"]++
			fallthrough
		case unicode.IsLetter(r):
			types["IsLetter"]++
			fallthrough
		case unicode.IsMark(r):
			types["IsMark"]++
			fallthrough
		case unicode.IsNumber(r):
			types["IsNumber"]++
			fallthrough
		case unicode.IsPrint(r):
			types["IsPrint"]++
			fallthrough
		case unicode.IsPunct(r):
			types["IsPunct"]++
			fallthrough
		case unicode.IsLower(r):
			types["IsLower"]++
			fallthrough
		case unicode.IsSpace(r):
			types["IsSpace"]++
			fallthrough
		case unicode.IsSymbol(r):
			types["IsSymbol"]++
			fallthrough
		case unicode.IsTitle(r):
			types["IsTitle"]++
			fallthrough
		case unicode.IsUpper(r):
			types["IsUpper"]++

		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("type\t\tcount\n")
	for k, v := range types {
		fmt.Printf("%s\t\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// 练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
func work49() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		word = strings.Trim(word, ",. \t\n\r\f()[]{}'\"“”-_`–")
		if word != "" {
			words[word]++
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println("word\tnum")
	for word, num := range words {
		fmt.Println(word, "\t", num)
	}
}
