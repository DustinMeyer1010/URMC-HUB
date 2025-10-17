package global

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

//go:embed .env
var enviroment_vairables []byte

var (
	SERVER1, SERVER2, SERVER3, SERVER4, SERVER5, SERVER6, SERVER7, SERVER8, SERVER9, SERVER10, LOGON, SHARES string
)

func LoadEnv() {
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

	SERVER1 = os.Getenv("SERVER1")
	SERVER2 = os.Getenv("SERVER2")
	SERVER3 = os.Getenv("SERVER3")
	SERVER4 = os.Getenv("SERVER4")
	SERVER5 = os.Getenv("SERVER5")
	SERVER6 = os.Getenv("SERVER6")
	SERVER7 = os.Getenv("SERVER7")
	SERVER8 = os.Getenv("SERVER8")
	SERVER9 = os.Getenv("SERVER9")
	SERVER10 = os.Getenv("SERVER10")
	LOGON = os.Getenv("LOGON")
	SHARES = os.Getenv("SHARES")

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
