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
				si.Device = uint64(ss.Dev)
				si.Mode = fi.Mode()
				si.User = getUserName(ss.Uid)
				si.Group = getGroupName(ss.Gid)
				si.Size = ss.Size
				si.Atime = time.Unix(ss.Atimespec.Sec, ss.Atimespec.Nsec)
				si.Mtime = time.Unix(ss.Mtimespec.Sec, ss.Mtimespec.Nsec)
				si.Ctime = time.Unix(ss.Ctimespec.Sec, ss.Ctimespec.Nsec)
				si.Btime = time.Unix(ss.Birthtimespec.Sec, ss.Birthtimespec.Nsec)
				si.Blocks = ss.Blocks
				si.BlockSize = ss.Blksize
				si.Flags = ss.Flags
			}
		}
	}
	return &si, err
}
