package swayipc

import "os"

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if err != nil {
		return false
	}

	return !info.IsDir()
}
