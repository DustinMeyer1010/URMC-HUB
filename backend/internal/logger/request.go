package logger

func LogRequestInfo(method, route string) {
	serverLogger.Infof("%s %s", method, route)
}
