package main

import (
	// . "fmt"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func doScanf() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	fmt.Println(a + b)
}

// 带有 IO 缓冲区的输入输出：适用于绝大多数题目
func bufferIO(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var a, b int
	fmt.Fscan(in, &a, &b)
	fmt.Fprintln(out, a, b)
}

// 快读：适应于输入量巨大的题目
func fastIO(r io.Reader, w io.Writer) {
	in := bufio.NewScanner(r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(w)
	defer out.Flush()

	// 读一个整数
	op := func() int {
		in.Scan()
		x, _ := strconv.Atoi(string(in.Bytes()))
		return x
	}

	// 读一个非负整数
	op = func() int {
		in.Scan()
		x := 0
		for _, b := range in.Bytes() {
			x = x*10 + int(b&15)
		}
		return x
	}

	// 支持负数
	op = func() int {
		in.Scan()
		data := in.Bytes()
		x, sig := 0, 1
		if data[0] == '-' {
			sig = -1
			data = data[1:]
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return x * sig
	}

	// 浮点数
	opf := func() float64 {
		in.Scan()
		f, _ := strconv.ParseFloat(string(in.Bytes()), 64)
		return f
	}

	in.Buffer(nil, 1e9)
	op()
	opf()
}

func main() {
	doScanf()
	// bufferIO(os.Stdin, os.Stdout)
	// 32 << 0 或 32 << 1
	//
}
