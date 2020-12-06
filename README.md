# stat
Portable file status utility

![build workflow](https://github.com/shanebarnes/stat/workflows/stat/badge.svg)

## Examples

```
./stat go.*
[
    {
        "name": "/Users/sb/stat/go.mod",
        "device": 16777220,
        "mode": "-rw-r--r--",
        "user": "sb",
        "group": "staff",
        "size": "159",
        "aTime": "2020-12-05T12:58:40.363341210-05:00",
        "mTime": "2020-12-05T12:56:18.200584109-05:00",
        "cTime": "2020-12-05T12:56:18.200584109-05:00",
        "bTime": "2020-12-05T12:49:00.168109171-05:00",
        "blocks": "8",
        "blockSize": "4,096",
        "flags": 0
    },
    {
        "name": "/Users/sb/stat/go.sum",
        "device": 16777220,
        "mode": "-rw-r--r--",
        "user": "sb",
        "group": "staff",
        "size": "553",
        "aTime": "2020-12-05T12:58:40.363473576-05:00",
        "mTime": "2020-12-05T12:56:18.200350841-05:00",
        "cTime": "2020-12-05T12:56:18.200350841-05:00",
        "bTime": "2020-12-05T12:49:00.168609398-05:00",
        "blocks": "8",
        "blockSize": "4,096",
        "flags": 0
    }
]
```