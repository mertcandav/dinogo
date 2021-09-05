package utils

import "strings"

// GetCommandName returns name of command by specified separator.
func GetCommandName(cmd, sep string) string {
	for index := range cmd {
		if strings.HasPrefix(cmd[index:], sep) {
			return cmd[:index]
		}
	}
	return cmd
}

// SplitCommandName splits from command name by specified separator.
func SplitCommandName(cmd, sep string) (string, string) {
	for index := range cmd {
		if strings.HasPrefix(cmd[index:], sep) {
			return cmd[:index], cmd[index+1:]
		}
	}
	return cmd, ""
}

// Args returns arguments by specified separator and argument prefix.
func Args(cmd, prefix, sep string) []string {
	args := make([]string, 0)
	for _, part := range strings.SplitN(cmd, sep, -1) {
		if part != prefix && strings.HasPrefix(part, prefix) {
			args = append(args, part)
		}
	}
	return args
}
