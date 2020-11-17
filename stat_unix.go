// +build darwin linux

package main

import (
	"os/user"
	"strconv"
)

func getGroupName(gid uint32) string {
	groupName := strconv.FormatUint(uint64(gid), 10)
	if group, err := user.LookupGroupId(groupName); err == nil {
		groupName = group.Name
	}
	return groupName
}

func getUserName(uid uint32) string {
	userName := strconv.FormatUint(uint64(uid), 10)
	if user, err := user.LookupId(userName); err == nil {
		userName = user.Username
	}
	return userName
}
