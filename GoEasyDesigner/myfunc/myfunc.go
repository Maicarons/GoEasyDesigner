package myfunc

import (
	"github.com/duolabmeng6/goefun/ecore"
	. "github.com/duolabmeng6/goefun/ehttp"
	"regexp"
	"strings"
)

func InsertCode(jscode string, insertionCode string) string {
	// 获取函数名称 窗口.按钮3被点击 = function () { 截取 按钮3被点击
	函数名称 := insertionCode[strings.Index(insertionCode, ".")+1 : strings.Index(insertionCode, "=")]
	// 检查是否已经存在
	if strings.Contains(jscode, 函数名称) {
		return jscode
	}
	// Define the regular expression pattern
	pattern := `(//事件函数位置不要删除[^\n]*)`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Replace the matching pattern with the insertion code
	result := re.ReplaceAllString(jscode, insertionCode+"$1")

	return result
}

func E发送跳转代码到ide(插件URL地址 string, 文件路径 string, 跳转位置 int) {
	文件名 := ecore.E文件取文件名(文件路径, true)
	文件名 = ecore.URL编码(文件名)
	跳转URL := 插件URL地址 + "/myserver?type=target&file=" + 文件名 + "&def=" + ecore.E到文本(跳转位置)
	ecore.E调试输出("调用pycharm代码跳转", 跳转URL)
	eh := NewHttp()
	返回文本, _ := eh.Get(跳转URL)
	println("调用pycharm代码跳转", 返回文本)

}
