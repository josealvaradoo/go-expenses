package utils

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var (
	environment      = map[string]string{}
	enviromnentMutex = sync.RWMutex{}
)

func AccessEnv(key string) string {
	enviromnentMutex.RLock()
	val := environment[key]
	enviromnentMutex.RUnlock()

	if environment[key] != "" {
		return val
	}

	val = os.Getenv(key)

	if val == "" || len(val) <= 0 {
		return ""
	}

	enviromnentMutex.Lock()
	environment[key] = val
	enviromnentMutex.Unlock()

	return val
}
