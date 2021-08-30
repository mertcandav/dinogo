package utils

import "strings"

// GetCommandName returns name of command by specified seperator.
func GetCommandName(cmd, sep string) string {
	for index := range cmd {
		if strings.HasPrefix(cmd[index:], sep) {
			return cmd[:index]
		}
	}
	return cmd
}

// SplitCommandName splits from command name by specified seperator.
func SplitCommandName(cmd, sep string) (string, string) {
	for index := range cmd {
		if strings.HasPrefix(cmd[index:], sep) {
			return cmd[:index], cmd[index+1:]
		}
	}
	return cmd, ""
}
