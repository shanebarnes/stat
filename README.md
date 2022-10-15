# stat
Portable command-line utility that displays detailed information about given
files including nanosecond resolution for file timestamp fields.

![build workflow](https://github.com/shanebarnes/stat/workflows/stat/badge.svg)

## Examples

### File Storage

```
./bin/stat-darwin-arm64 go.*
name: /Users/sb/stat/go.mod
device: 16777233
mode: -rw-r--r--
user: sb
group: staff
size: 1721
aTime: "2022-10-15T10:51:29.330700669-03:00"
mTime: "2022-10-15T10:51:22.216926407-03:00"
cTime: "2022-10-15T10:51:22.216926407-03:00"
bTime: "2022-10-15T07:24:15.466480847-03:00"
blocks: 8
blockSize: 4096
flags: 0
---
name: /Users/sb/stat/go.sum
device: 16777233
mode: -rw-r--r--
user: sb
group: staff
size: 6666
aTime: "2022-10-15T10:51:29.331089282-03:00"
mTime: "2022-10-15T10:51:22.216331301-03:00"
cTime: "2022-10-15T10:51:22.216331301-03:00"
bTime: "2022-03-16T20:17:00.892278002-03:00"
blocks: 16
blockSize: 4096
flags: 0
```

### S3 Storage

```
./bin/stat-darwin-arm64 -o json -s s3 https://s3.amazonaws.com/mybucket/myobject
{
  "name": "https://s3.amazonaws.com/mybucket/myobject",
  "device": 0,
  "mode": "----------",
  "user": "",
  "group": "",
  "size": 804,
  "aTime": "0001-01-01T00:00:00.000000000Z",
  "mTime": "2022-10-15T11:05:25.000000000Z",
  "cTime": "0001-01-01T00:00:00.000000000Z",
  "bTime": "0001-01-01T00:00:00.000000000Z",
  "blocks": 0,
  "blockSize": 0,
  "flags": 0,
  "metadata": {
    "hello": "world"
  }
}

```