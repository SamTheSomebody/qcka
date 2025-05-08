// Copyright © 2025 Sam Muller gamedevsam@pm.me
package files

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/samthesomebody/qcka/settings"
)

const (
	groupPrefix            = "#Group:"
	fileName               = ".bash_aliases"
	invalidAliasChars      = "/$`="
	invalidAliasCharsPrint = "/, $, `, ="
)

type Aliases struct {
	Groups        map[string]AliasGroup
	KeyGroupPairs map[string]string
}

type AliasGroup struct {
	Values map[string]string
}

func newAliasGroup() AliasGroup {
	return AliasGroup{
		Values: make(map[string]string),
	}
}

func GetAliases() (*Aliases, error) {
	if settings.Settings.Verbose {
		fmt.Println("Retrieving aliases!")
	}

	a := Aliases{
		Groups:        make(map[string]AliasGroup),
		KeyGroupPairs: make(map[string]string),
	}
	a.Groups["DEFAULT"] = newAliasGroup()
	data, err := Read(fileName)
	if err != nil {
		return nil, fmt.Errorf("error getting aliases: %v", err)
	}
	currentGroup := "DEFAULT"
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, groupPrefix) {
			currentGroup = a.getGroup(line)
		} else if strings.Contains(line, "=") {
			if err := a.getAlias(currentGroup, line); err != nil {
				return nil, err
			}
		}
	}
	return &a, nil
}

func (a *Aliases) getGroup(line string) string {
	currentGroup := strings.Replace(line, groupPrefix, "", 1)
	if _, ok := a.Groups[currentGroup]; !ok {
		a.Groups[currentGroup] = newAliasGroup()
	}
	return currentGroup
}

func (a *Aliases) getAlias(currentGroup, line string) error {
	split := strings.Split(line, "=")
	if len(split) > 2 {
		return fmt.Errorf("error parsing aliases: more than 2 '=' in line")
	}
	key := strings.Split(split[0], " ")[1]
	key = strings.TrimSpace(key)
	value := strings.TrimSpace(split[1])
	if value[0] == '"' {
		value = string(value[1:])
	}
	if value[len(value)-1] == '"' {
		value = string(value[:len(value)-1])
	}
	a.Groups[currentGroup].Values[key] = value
	a.KeyGroupPairs[key] = currentGroup
	return nil
}

func (a *Aliases) Add(group, key, value string) error {
	group = strings.ToUpper(group)
	group = strings.TrimSpace(group)
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	if strings.Contains(key, "2084") || strings.ContainsAny(key, invalidAliasChars) {
		fmt.Printf("Alias %s is invalid. Cannot contain any of: %s\n", key, invalidAliasChars)
		return nil
	}

	cmd := strings.Split(value, " ")[0]
	_, err := exec.Command("bash", "-c", "type -a"+cmd).Output()
	if err != nil && !settings.Settings.Force {
		fmt.Printf("[Warning] %s has no usages in bash. Continue? y/n: ", cmd)
		isContinuing, err := awaitYesNo()
		if err != nil {
			return fmt.Errorf("error adding alias: %v", err)
		}
		if !isContinuing {
			return nil
		}
	}

	output, err := exec.Command("bash", "-c", "type -a "+key).Output()
	if err == nil && !settings.Settings.Force {
		s := strings.TrimSpace(string(output))
		uses := len(strings.Split(s, "\n"))
		fmt.Printf("[Warning] %s has %v other usages in bash, this will cause an override. Continue? y/n: ", key, uses)
		isContinuing, err := awaitYesNo()
		if err != nil {
			return fmt.Errorf("error adding alias: %v", err)
		}
		if !isContinuing {
			return nil
		}
	}

	if g, ok := a.KeyGroupPairs[key]; ok {
		if g == group && !settings.Settings.Force {
			fmt.Printf("[Warning] %s already has an alias (%s). Override it? y/n: ", key, a.Groups[g].Values[key])
			isContinuing, err := awaitYesNo()
			if err != nil {
				return fmt.Errorf("error adding alias: %v", err)
			}
			if !isContinuing {
				return nil
			}
		} else {
			if !settings.Settings.Force {
				fmt.Printf("[Warning] %s already has an alias in group %s. Delete it and add %s to group %s? y/n: ", key, g, key, group)
				isContinuing, err := awaitYesNo()
				if err != nil {
					return fmt.Errorf("error adding alias: %v", err)
				}
				if !isContinuing {
					return nil
				}
			}
			delete(a.Groups[g].Values, key)
			delete(a.KeyGroupPairs, key)
		}
	}

	if _, ok := a.Groups[group]; !ok {
		a.Groups[group] = newAliasGroup()
	}
	a.Groups[group].Values[key] = value

	if settings.Settings.Verbose {
		fmt.Printf("Added alias to %v: %v=\"%v\"\n", group, key, value)
	}

	err = a.Write(true)
	if err != nil {
		return fmt.Errorf("error writing added alias: %v", err)
	}

	return nil
}

func (a *Aliases) Remove(name string) error {
	group, ok := a.KeyGroupPairs[name]
	if !ok {
		if !settings.Settings.Quiet {
			fmt.Printf("'%v' does not exist as an alias in '%v'\n", name, fileName)
		}
		return nil
	}
	delete(a.Groups[group].Values, name)
	delete(a.KeyGroupPairs, name)
	return a.Write(false)
}

func (a *Aliases) Write(isAdding bool) error {
	s := ""
	for groupKey, groupValue := range a.Groups {
		if len(groupValue.Values) == 0 {
			continue
		}
		s += fmt.Sprintf("%v%v\n", groupPrefix, groupKey)
		for key, value := range groupValue.Values {
			s += fmt.Sprintf("alias %v=\"%v\"\n", key, value)
		}
		s += "\n"
	}
	s = strings.TrimSpace(s)

	err := Write([]byte(s), fileName, isAdding)
	if err != nil {
		return fmt.Errorf("error writing aliases: %v", err)
	}

	return nil
}
