package files
import (
	"os"
	"fmt"
	"strings"
)


func getPath(fileName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading %v: %v", fileName, err)
	}
	return strings.Join([]string{homeDir, fileName}, string(os.PathSeparator)), nil
}

func getPerm(fileName string) (os.FileMode, error) {
	path, err := getPath(fileName)
	if err != nil {
		return os.FileMode(0), fmt.Errorf("error getting %v permissions: %v", fileName, err)
	}
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return os.FileMode(0), fmt.Errorf("error getting %v permissions: %v", fileName, err)
	}
	return fileInfo.Mode().Perm(), err
}

func Read(fileName string) ([]byte, error) {
	path, err := getPath(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading '.bashrc': %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading '.bashrc': %v", err)
	}
	return data, nil
}

func Write(data []byte, fileName string) error {
	path, err := getPath(fileName)
	if err != nil {
		return fmt.Errorf("error sourcing alias in '%v': %v", fileName, err)
	}

	perm, err := getPerm(fileName)
	if err != nil {
		return fmt.Errorf("error sourcing alias in '%v': %v", fileName, err)
	}

	err = os.WriteFile(path, data, perm)
	if err != nil {
		return fmt.Errorf("error sourcing alias in '%v': %v", fileName, err)
	}
	return nil
}
