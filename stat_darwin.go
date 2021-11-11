package main

import (
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
)

func getStatInfo(name string) (*statInfo, error) {
	si := statInfo{Name: name}
	abs, err := filepath.Abs(name)
	if err == nil {
		si.Name = abs
		var fi os.FileInfo
		if fi, err = os.Lstat(name); err == nil {
			if ss, ok := fi.Sys().(*syscall.Stat_t); ok {
				si.Device = uint64(ss.Dev)
				si.Mode = fi.Mode().String()
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
	}
	return &si, err
}
