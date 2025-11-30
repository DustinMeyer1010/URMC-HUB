package logger

import "github.com/LostProgrammer1010/URMC-HUB/internal/models"

func LogSearchResults(matches models.AllResults) {

	serverLogger.Infof("# of Users Found: %d", len(matches.Users))
	serverLogger.Infof("# of Computers Found: %d", len(matches.Computers))
	serverLogger.Infof("# of Printers Found: %d", len(matches.Printers))
	serverLogger.Infof("# of Drives Found: %d", len(matches.Shares))
	serverLogger.Infof("# of Groups Found: %d", len(matches.Groups))

}
