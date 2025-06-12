package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/kalifun/pmindt/internal"
	"github.com/spf13/cobra"
)

var toMarkDown = &cobra.Command{
	Use:     "toMd",
	Short:   "",
	Long:    "PingCode Mind to  markdown",
	Example: "pmindt toMD",
	Run: func(cmd *cobra.Command, args []string) {
		mindToMd(args)
	},
}

func mindToMd(args []string) {
	var result []string
	for i, v := range args {
		data, err := internal.PmindToMarkdown([]byte(v))
		if err != nil {
			fmt.Printf("The current line [%d] parsing failed. err msg %s\n", i, err.Error())
			continue
		}
		fmt.Println(data)
		result = append(result, data)
	}
	clipboard.WriteAll(strings.Join(result, "\n"))
	fmt.Printf(yellow, "拷贝到其他平台思维导图即可 (已复制到剪贴板）")
}
