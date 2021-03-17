package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shanebarnes/stat/internal/version"
)

const RFC3339NanoZero = "2006-01-02T15:04:05.000000000Z07:00"

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
	Error     string `json:"error,omitempty"`
}

func main() {
	printVersion := flag.Bool("version", false, "Print version information")
	flag.Parse()
	if *printVersion {
		fmt.Fprintf(os.Stdout, "stat version %s\n", version.String())
	} else if len(os.Args) > 1 {
		out := os.Stdout
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		io.WriteString(out, "[\n")
		for i, arg := range os.Args[1:] {
			si, err := getStatInfo(arg)
			if err != nil {
				si = &statInfo{Name: arg, Error: err.Error()}
			}
			enc.Encode(si)

			if (i + 2) < len(os.Args) {
				io.WriteString(out, ",")
			}
		}
		io.WriteString(out, "]\n")
	} else {
		fmt.Fprintf(os.Stdout, "stat [file ...]\n")
	}
}
