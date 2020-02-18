package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
)

const (
	RFC3339NanoZero = "2006-01-02T15:04:05.000000000Z07:00"
)

type statInfo struct {
	Name      string `json:"name"`
	Device    int32  `json:"device"`
	Mode      string `json:"mode"`
	User      string `json:"user"`
	Group     string `json:"group"`
	Size      string `json:"size"`
	Atime     string `json:"aTime"`
	Mtime     string `json:"mTime"`
	Ctime     string `json:"cTime"`
	Btime     string `json:"bTime"`
	Blocks    string `json:"blocks"`
	BlockSize string `json:"blockSize"`
	Flags     uint32 `json:"flags"`
}

func main() {
	if len(os.Args) > 1 {

		stats := []*statInfo{}
		for _, arg := range os.Args[1:] {
			if si, err := getStatInfo(arg); err == nil {
				stats = append(stats, si)
			}
		}

		printStatInfo(stats)
	} else {
		fmt.Fprintf(os.Stdout, "stat [file ...]\n")
	}
}

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
			si.Atime = (time.Unix(ss.Atimespec.Sec, ss.Atimespec.Nsec)).Format(RFC3339NanoZero)
			si.Mtime = (time.Unix(ss.Mtimespec.Sec, ss.Mtimespec.Nsec)).Format(RFC3339NanoZero)
			si.Ctime = (time.Unix(ss.Ctimespec.Sec, ss.Ctimespec.Nsec)).Format(RFC3339NanoZero)
			si.Btime = (time.Unix(ss.Birthtimespec.Sec, ss.Birthtimespec.Nsec)).Format(RFC3339NanoZero)
			si.Blocks = humanize.Comma(ss.Blocks)
			si.BlockSize = humanize.Comma(int64(ss.Blksize))
			si.Flags = ss.Flags
		}
	}

	return &si, err
}

func printStatInfo(si []*statInfo) error {
	buf, err := json.MarshalIndent(si, "", "    ")
	if err == nil {
		fmt.Println(string(buf))
	}
	return err
}

func getUserName(uid uint32) string {
	userName := strconv.FormatUint(uint64(uid), 10)
	if user, err := user.LookupId(userName); err == nil {
		userName = user.Username
	}
	return userName
}
