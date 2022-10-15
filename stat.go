package stat

import (
	"os"
	"time"
)

type StatInfo struct {
	Name      string            `csv:"name"               json:"name"               yaml:"name"`
	Device    uint64            `csv:"device"             json:"device"             yaml:"device"`
	Mode      os.FileMode       `csv:"mode"               json:"mode"               yaml:"mode"`
	User      uint32            `csv:"user"               json:"user"               yaml:"user"`
	Group     uint32            `csv:"group"              json:"group"              yaml:"group"`
	Size      int64             `csv:"size"               json:"size"               yaml:"size"`
	Atime     time.Time         `csv:"aTime"              json:"aTime"              yaml:"aTime"`
	Mtime     time.Time         `csv:"mTime"              json:"mTime"              yaml:"mTime"`
	Ctime     time.Time         `csv:"cTime"              json:"cTime"              yaml:"cTime"`
	Btime     time.Time         `csv:"bTime"              json:"bTime"              yaml:"bTime"`
	Blocks    int64             `csv:"blocks"             json:"blocks"             yaml:"blocks"`
	BlockSize int32             `csv:"blockSize"          json:"blockSize"          yaml:"blockSize"`
	Flags     uint32            `csv:"flags"              json:"flags"              yaml:"flags"`
	Metadata  map[string]string `csv:"metadata,omitempty" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Error     error             `csv:"error,omitempty"    json:"error,omitempty"    yaml:"error,omitempty"`
}

type StatInfoPretty struct {
	Name      string            `csv:"name"               json:"name"               yaml:"name"`
	Device    uint64            `csv:"device"             json:"device"             yaml:"device"`
	Mode      string            `csv:"mode"               json:"mode"               yaml:"mode"`
	User      uint32            `csv:"user"               json:"user"               yaml:"user"`
	Group     uint32            `csv:"group"              json:"group"              yaml:"group"`
	Size      int64             `csv:"size"               json:"size"               yaml:"size"`
	Atime     string            `csv:"aTime"              json:"aTime"              yaml:"aTime"`
	Mtime     string            `csv:"mTime"              json:"mTime"              yaml:"mTime"`
	Ctime     string            `csv:"cTime"              json:"cTime"              yaml:"cTime"`
	Btime     string            `csv:"bTime"              json:"bTime"              yaml:"bTime"`
	Blocks    int64             `csv:"blocks"             json:"blocks"             yaml:"blocks"`
	BlockSize int32             `csv:"blockSize"          json:"blockSize"          yaml:"blockSize"`
	Flags     uint32            `csv:"flags"              json:"flags"              yaml:"flags"`
	Metadata  map[string]string `csv:"metadata,omitempty" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Error     string            `csv:"error"              json:"error,omitempty"    yaml:"error,omitempty"`
}

func (si *StatInfo) Pretty(dateLayout string) *StatInfoPretty {
	err := ""
	if si.Error != nil {
		err = si.Error.Error()
	}

	return &StatInfoPretty{
		Name:      si.Name,
		Device:    si.Device,
		Mode:      si.Mode.String(),
		User:      si.User,
		Group:     si.Group,
		Size:      si.Size,
		Atime:     si.Atime.Format(dateLayout),
		Mtime:     si.Mtime.Format(dateLayout),
		Ctime:     si.Ctime.Format(dateLayout),
		Btime:     si.Btime.Format(dateLayout),
		Blocks:    si.Blocks,
		BlockSize: si.BlockSize,
		Flags:     si.Flags,
		Metadata:  si.Metadata,
		Error:     err,
	}
}
