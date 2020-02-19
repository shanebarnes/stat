// +build darwin linux

package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
)

func getGroupName(gid uint32) string {
	groupName := strconv.FormatUint(uint64(gid), 10)
	if group, err := user.LookupGroupId(groupName); err == nil {
		groupName = group.Name
	}
	return groupName
}

func getStatInfo(name string) (*statInfo, error) {
	si := statInfo{Name: name}
	abs, err := filepath.Abs(name)
	if err == nil {
		si.Name = abs
		ss := syscall.Stat_t{}
		if err = syscall.Stat(name, &ss); err == nil {
			si.Device = ss.Dev
			si.Mode = (os.FileMode(ss.Mode)).String()
			si.User = getUserName(ss.Uid)
			si.Group = getGroupName(ss.Gid)
			si.Size = humanize.Comma(ss.Size)
			si.Atime = time.Unix(ss.Atimespec.Sec, ss.Atimespec.Nsec).Format(RFC3339NanoZero)
			si.Mtime = time.Unix(ss.Mtimespec.Sec, ss.Mtimespec.Nsec).Format(RFC3339NanoZero)
			si.Ctime = time.Unix(ss.Ctimespec.Sec, ss.Ctimespec.Nsec).Format(RFC3339NanoZero)
			si.Btime = time.Unix(ss.Birthtimespec.Sec, ss.Birthtimespec.Nsec).Format(RFC3339NanoZero)
			si.Blocks = humanize.Comma(ss.Blocks)
			si.BlockSize = humanize.Comma(int64(ss.Blksize))
			si.Flags = ss.Flags
		}
	}
	return &si, err
}

func getUserName(uid uint32) string {
	userName := strconv.FormatUint(uint64(uid), 10)
	if user, err := user.LookupId(userName); err == nil {
		userName = user.Username
	}
	return userName
}
