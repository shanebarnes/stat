package stat

import (
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func (fs *FileStat) Stat(name string) (*StatInfo, error) {
	si := StatInfo{Name: name}
	abs, err := filepath.Abs(name)
	if err == nil {
		si.Name = abs
		var fi os.FileInfo
		if fi, err = os.Lstat(name); err == nil {
			if ss, ok := fi.Sys().(*syscall.Stat_t); ok {
				si.Device = ss.Dev
				si.Mode = fi.Mode()
				si.User = getUserName(ss.Uid)
				si.Group = getGroupName(ss.Gid)
				si.Size = ss.Size
				si.Atime = time.Unix(ss.Atim.Sec, ss.Atim.Nsec)
				si.Mtime = time.Unix(ss.Mtim.Sec, ss.Mtim.Nsec)
				si.Ctime = time.Unix(ss.Ctim.Sec, ss.Ctim.Nsec)
				si.Btime = time.Time{}
				si.Blocks = ss.Blocks
				si.BlockSize = ss.Blksize
				si.Flags = uint32(0)
			}
		}
	}
	return &si, err
}
