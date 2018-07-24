package logger

import (
	"runtime"
	"os"
	"time"
	"fmt"
	"path"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3) //0表示该函数，1表示调用该函数的地方，2表示再上一层，堆栈
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func writeLog(file *os.File, level int, format string, args ...interface{}){
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")  //以go的诞生之日为时间点格式化，固定的，不能变

	levelStr := getLevelText(level)

	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}