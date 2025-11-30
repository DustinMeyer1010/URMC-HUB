package logger

func LogRequestInfo(method, route string) {
	ServerLogger.Infof("%s %s", method, route)
}
