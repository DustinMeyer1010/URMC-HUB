package logger

import "github.com/LostProgrammer1010/URMC-HUB/internal/models"

func LogSearchResults(matches models.AllResults) {

	ServerLogger.Infof("# of Users Found: %d", len(matches.Users))
	ServerLogger.Infof("# of Computers Found: %d", len(matches.Computers))
	ServerLogger.Infof("# of Printers Found: %d", len(matches.Printers))
	ServerLogger.Infof("# of Drives Found: %d", len(matches.Shares))
	ServerLogger.Infof("# of Groups Found: %d", len(matches.Groups))

}
