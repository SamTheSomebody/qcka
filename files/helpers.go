// Copyright © 2025 Sam Muller gamedevsam@pm.me
package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samthesomebody/qcka/settings"
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

	if settings.Settings.Verbose {
		fmt.Printf("Reading file at: %v\n", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading '.bashrc': %v", err)
	}
	return data, nil
}

func Write(data []byte, fileName string, isAdding bool) error {
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

	if settings.Settings.Verbose {
		fmt.Printf("Wrote aliases to: %v\n", path)
	}
	if !settings.Settings.Quiet {
		if isAdding {
			fmt.Printf("To use this alias in the current session, run: source ~/.bash_aliases\n")
		} else {
			fmt.Printf("This alias is still active in the current session, to re-execute bash run: exec bash\n")
		}
	}

	return nil
}

func awaitYesNo() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read input: %v", err)
		}
		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "y":
			return true, nil
		case "n":
			return false, nil
		default:
			fmt.Printf("Invalid input. Continue? y/n: ")
		}
	}
}
