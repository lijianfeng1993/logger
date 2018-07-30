package logger

import "testing"

func TestFileLogger(t *testing.T) {
	//logger := NewFileLogger(LogLevelDebug, "/Users/lijianfeng/GoProject/src/LearnGo/logger/log", "filetest")
	logger := NewFileLogger(LogLevelDebug, "E:/GoProject/src/LearnGo/logger/log", "filetest")
	logger.Debug("user id is come from china.", )
	logger.Warn("test warn log.",)
	logger.Fatal("test fatal log.",)
	logger.Close()
}


func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("this a debug info.")
	logger.Warn("this is warning info.")
	logger.Fatal("this is a fatal info.")
	logger.Error("this is a error info")
}


