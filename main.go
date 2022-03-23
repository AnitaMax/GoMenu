package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type CMD struct {
	short  byte
	long   string
	action func(string) int
	help   string
}

func (cmd *CMD) match(request string) bool {
	if len(request) < 1 {
		return false
	}
	switch len(request) {
	case 1:
		if byte(request[0]) == cmd.short {
			return true
		}
	default:
		if cmd.long == request {
			return true
		}
	}
	return false
}

func (cmd *CMD) runOn(request string, arg string) int {
	var res = cmd.match(request)
	if res {
		return cmd.action(arg)
	}
	return -2
}

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}

func main() {

	testfunc := func(arg string) int {
		fmt.Println("Hello," + arg)
		return 0
	}
	cmd_test := CMD{'t', "test", testfunc, "这是一个测试命令"}

	cmds := list.New()
	cmds.PushBack(cmd_test)

	for {
		fmt.Println("请输入命令与参数,用空格隔开")
		var input string
		Scanf(&input)
		var ss = strings.Split(input, " ")
		// fmt.Println(ss)
		// fmt.Println(len(ss))
		for i := cmds.Front(); i != nil; i = i.Next() {
			var c, ok = i.Value.(CMD)
			if ok {
				c.runOn(ss[0], ss[1])
			}

		}
	}
}
