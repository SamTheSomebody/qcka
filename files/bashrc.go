package files

import (
	"fmt"
	"strings"
)

const aliasSource = "\n\n# Alias definitions.\n\n if [ -f ~/.bash_aliases ]; then\n    . ~/.bash_aliases\nfi"

func IsAliasSourced() (bool, error) {
	data, err := Read(".bashrc")
	if err != nil {
		return false, fmt.Errorf("error checking if '.bash_aliases' is sourced from '.bashrc': %v", err)
	}
	if strings.Contains(string(data), "~/.bash_aliases") {
		return true, nil
	}
	return false, nil
}

func SourceAlias() error {
	fmt.Printf("'.bash_aliases' not sourced in '.bashrc'! Adding the following lines to '.bashrc': %v\n", aliasSource)
	data, err := Read(".bashrc")
	if err != nil {
		return fmt.Errorf("error sourcing '.bash_alias' in '.bashrc': %v", err)
	}

	data = append(data, []byte(aliasSource)...)
	err = Write(data, ".bashrc")
	if err != nil {
		return fmt.Errorf("error sourcing '.bash_alias' in '.bashrc': %v", err)
	}
	return nil
}
