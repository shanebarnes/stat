// References:
//     https://golang.org/pkg/syscall/?GOOS=windows#Win32FileAttributeData

package stat

import (
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/djherbis/times"
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

func (fs *FileStat) Stat(name string) (*StatInfo, error) {
	si := StatInfo{Name: name}
	abs, err := filepath.Abs(name)
	if err == nil {
		si.Name = abs
		var fi os.FileInfo
		if fi, err = os.Stat(name); err == nil {
			if ss, ok := fi.Sys().(*syscall.Win32FileAttributeData); ok {
				if ts, err := times.Stat(name); err == nil {
					//if ts.HasChangeTime() {
					si.Ctime = ts.ChangeTime()
					//}
				}
				si.Device = uint64(0)
				si.Mode = fi.Mode()
				si.Size = fi.Size()
				si.Atime = time.Unix(0, ss.LastAccessTime.Nanoseconds())
				si.Mtime = time.Unix(0, ss.LastWriteTime.Nanoseconds())
				si.Btime = time.Unix(0, ss.CreationTime.Nanoseconds())
				si.Blocks = 0
				si.BlockSize = 0
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
