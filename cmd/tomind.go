package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/kalifun/pmindt/internal"
	"github.com/spf13/cobra"
)

var yellow = "\033[33m%s\033[0m\n"

var toMind = &cobra.Command{
	Use:     "toPCM",
	Short:   "",
	Long:    "markdown to PingCode Mind",
	Example: "pmindt toPCM",
	Run: func(cmd *cobra.Command, args []string) {
		mdtoMind()
	},
}

func mdtoMind() {
	fmt.Println("请输入层次结构（按 Ctrl+D 结束输入）：")

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return
	}
	data, err := internal.MdToPMind(lines)
	if err != nil {
		panic(err)
	}
	clipboard.WriteAll(data)
	fmt.Printf(yellow, "拷贝到PingCode思维导图即可 (已复制到剪贴板）")
}
