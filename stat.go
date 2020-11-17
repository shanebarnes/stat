package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	RFC3339NanoZero = "2006-01-02T15:04:05.000000000Z07:00"
)

type statInfo struct {
	Name      string `json:"name"`
	Device    uint64 `json:"device"`
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

func printStatInfo(si []*statInfo) error {
	buf, err := json.MarshalIndent(si, "", "    ")
	if err == nil {
		fmt.Println(string(buf))
	}
	return err
}
