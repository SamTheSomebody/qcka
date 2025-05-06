package files

import (
	"fmt"
	"strings"
)

// TODO add grouping
func GetAliases() (map[string]string, error) {
	output := make(map[string]string)
	data, err := Read(".bash_aliases")
	if err != nil {
		return nil, fmt.Errorf("error getting aliases: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			split := strings.Split(line, "=")
			if len(split) > 2 {
				return nil, fmt.Errorf("error parsing aliases: more than 2 '=' in line")
			}
			key := strings.Split(split[0], " ")[1]
			key = strings.TrimSpace(key)
			value := strings.TrimSpace(split[1])
			output[key] = value
		}
	}

	return output, nil
}

func WriteAliases(aliases map[string]string) error {
	data := make([]byte, 0)
	for key, value := range aliases {
		line := fmt.Sprintf("alias %v = \"%v\"\n", key, value)
		data = append(data, []byte(line)...)
	}
	err := Write(data, ".bash_aliases")
	if err != nil {
		return fmt.Errorf("error writing aliases: %v", err)
	}
	return nil
}
