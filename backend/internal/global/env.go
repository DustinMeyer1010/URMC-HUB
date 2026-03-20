package global

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/joho/godotenv"
)

//go:embed .env
var enviroment_vairables []byte

var (
	SERVER1,
	SERVER2,
	SERVER3,
	SERVER4,
	SERVER5,
	SERVER6,
	SERVER7,
	SERVER8,
	SERVER9,
	SERVER10,
	LOGON,
	SHARES,
	BASEDN,
	USERNAME_PREFIX,
	URMC_LDAP,
	ADMIN,
	PRINTQUEUEREPORT string
)

func LoadEnv() {
	var found bool
	// Write the embedded .env to a temporary file
	tmpFile := filepath.Join(os.TempDir(), "embedded.env")
	err := os.WriteFile(tmpFile, enviroment_vairables, 0600)

	if err != nil {
		log.Fatal("Error writing temp env file:", err)
	}

	err = godotenv.Load(tmpFile)

	if err != nil {
		panic("Couldn't find .env file to load variables")
	}

	if SERVER1, found = os.LookupEnv("SERVER1"); !found {
		logger.Error("SERVER1 IS MISSING IN ENV")
		panic("SERVER1 IS MISSING IN ENV")
	}
	if SERVER2, found = os.LookupEnv("SERVER2"); !found {
		logger.Error("SERVER2 IS MISSING IN ENV")
		panic("SERVER2 IS MISSING IN ENV")
	}
	if SERVER3, found = os.LookupEnv("SERVER3"); !found {
		logger.Error("SERVER3 IS MISSING IN ENV")
		panic("SERVER3 IS MISSING IN ENV")
	}
	if SERVER4, found = os.LookupEnv("SERVER4"); !found {
		logger.Error("SERVER4 IS MISSING IN ENV")
		panic("SERVER4 IS MISSING IN ENV")
	}
	if SERVER5, found = os.LookupEnv("SERVER5"); !found {
		logger.Error("SERVER5 IS MISSING IN ENV")
		panic("SERVER5 IS MISSING IN ENV")
	}
	if SERVER6, found = os.LookupEnv("SERVER6"); !found {
		logger.Error("SERVER6 IS MISSING IN ENV")
		panic("SERVER6 IS MISSING IN ENV")
	}
	if SERVER7, found = os.LookupEnv("SERVER7"); !found {
		logger.Error("SERVER7 IS MISSING IN ENV")
		panic("SERVER7 IS MISSING IN ENV")
	}
	if SERVER8, found = os.LookupEnv("SERVER8"); !found {
		logger.Error("SERVER8 IS MISSING IN ENV")
		panic("SERVER8 IS MISSING IN ENV")
	}
	if SERVER9, found = os.LookupEnv("SERVER9"); !found {
		logger.Error("SERVER9 IS MISSING IN ENV")
		panic("SERVER9 IS MISSING IN ENV")
	}
	if SERVER10, found = os.LookupEnv("SERVER10"); !found {
		logger.Error("SERVER10 IS MISSING IN ENV")
		panic("SERVER10 IS MISSING IN ENV")
	}
	if LOGON, found = os.LookupEnv("LOGON"); !found {
		logger.Error("LOGON IS MISSING IN ENV")
		panic("LOGON IS MISSING IN ENV")
	}
	if SHARES, found = os.LookupEnv("SHARES"); !found {
		logger.Error("SHARES IS MISSING IN ENV")
		panic("SHARES IS MISSING IN ENV")
	}
	if BASEDN, found = os.LookupEnv("BASEDN"); !found {
		logger.Error("BASEDN IS MISSING IN ENV")
		panic("BASEDN IS MISSING IN ENV")
	}
	if URMC_LDAP, found = os.LookupEnv("URMCLDAP"); !found {
		logger.Error("URMCLDAP IS MISSING IN ENV")
		panic("URMCLDAP IS MISSING IN ENV")
	}
	if USERNAME_PREFIX, found = os.LookupEnv("USERNAME_PREFIX"); !found {
		logger.Error("USERNAME_PREFIX IS MISSING IN ENV")
		panic("USERNAME_PREFIX IS MISSING IN ENV")
	}
	if ADMIN, found = os.LookupEnv("ADMIN"); !found {
		logger.Error("ADMIN IS MISSING IN ENV")
		panic("ADMIN IS MISSING IN ENV")
	}

	if PRINTQUEUEREPORT, found = os.LookupEnv("PRINTQUEUEREPORT"); !found {
		logger.Error("PRINTQUEUEREPORT IS MISSING IN ENV")
		panic("PRINTQUEUEREPORT IS MISSING IN ENV")
	}

}

func AllServers() [10]string {
	return [10]string{
		SERVER1,
		SERVER2,
		SERVER3,
		SERVER4,
		SERVER5,
		SERVER6,
		SERVER7,
		SERVER8,
		SERVER9,
		SERVER10,
	}
}
