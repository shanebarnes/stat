// References:
//     https://golang.org/pkg/syscall/?GOOS=windows#Win32FileAttributeData

package main

import (
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
	"golang.org/x/sys/windows"
)

func getGroupName(sd *windows.SECURITY_DESCRIPTOR) string {
	groupName := ""
	if gsid, _, err := sd.Group(); err == nil {
		groupName = gsid.String()
	}
	return groupName
}

func getSecurityInfo(name string) (*windows.SECURITY_DESCRIPTOR, error) {
	return windows.GetNamedSecurityInfo(name, windows.SE_FILE_OBJECT, windows.OWNER_SECURITY_INFORMATION)
}

func getStatInfo(name string) (*statInfo, error) {
	si := statInfo{Name: name}
	abs, err := filepath.Abs(name)
	if err == nil {
		si.Name = abs
		var fi os.FileInfo
		if fi, err = os.Stat(name); err == nil {
			if ss, ok := fi.Sys().(*syscall.Win32FileAttributeData); ok {
				if ts, err := times.Stat(name); err == nil {
					//if ts.HasChangeTime() {
					si.Ctime = ts.ChangeTime().Format(RFC3339NanoZero)
					//}
				}
				si.Device = uint64(0)
				si.Mode = fi.Mode().String()
				si.Size = humanize.Comma(fi.Size())
				si.Atime = time.Unix(0, ss.LastAccessTime.Nanoseconds()).Format(RFC3339NanoZero)
				si.Mtime = time.Unix(0, ss.LastWriteTime.Nanoseconds()).Format(RFC3339NanoZero)
				si.Btime = time.Unix(0, ss.CreationTime.Nanoseconds()).Format(RFC3339NanoZero)
				si.Blocks = ""
				si.BlockSize = ""
				si.Flags = ss.FileAttributes
			}
			if sd, err := getSecurityInfo(si.Name); err == nil {
				si.User = getUserName(sd)
				si.Group = getGroupName(sd)
			}
		}
	}
	return &si, err
}

func getUserName(sd *windows.SECURITY_DESCRIPTOR) string {
	userName := ""
	if osid, _, err := sd.Owner(); err == nil {
		userName = osid.String()
	}
	return userName
}
