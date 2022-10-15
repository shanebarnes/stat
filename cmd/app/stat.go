package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shanebarnes/stat"
	"github.com/shanebarnes/stat/internal/encoding"
	"github.com/shanebarnes/stat/internal/version"
	"gopkg.in/yaml.v2"
)

const (
	formatCsv  = "csv"
	formatJson = "json"
	formatTsv  = "tsv"
	formatXml  = "xml"
	formatYaml = "yaml"

	rfc3339NanoZero = "2006-01-02T15:04:05.000000000Z07:00"

	storageFile = "file"
	storageS3   = "s3"
)

var (
	flagDateLayout   string
	flagOutputFormat string
	flagStorageType  string
	flagPrintVersion bool
)

func newFlagSet() *flag.FlagSet {
	flagset := flag.NewFlagSet("", flag.ExitOnError)

	flagset.StringVar(&flagDateLayout, "d", rfc3339NanoZero, "date layout")
	flagset.StringVar(&flagOutputFormat, "o", formatJson, "output format (json, yaml)")
	flagset.StringVar(&flagStorageType, "s", storageFile, "storage type (file, s3)")
	flagset.BoolVar(&flagPrintVersion, "version", false, "Print version information")

	return flagset
}

func validateOutputFormat(format *string) error {
	var err error
	*format = strings.ToLower(*format)
	switch *format {
	case formatCsv, formatJson, formatTsv, formatXml, formatYaml:
	default:
		err = fmt.Errorf("unrecognized output format '%s'", *format)
	}
	return err
}

func validateStorageType(storage *string) error {
	var err error
	*storage = strings.ToLower(*storage)
	switch *storage {
	case storageFile, storageS3:
	default:
		err = fmt.Errorf("unrecognized storage type '%s'", *storage)
	}
	return err
}

func main() {
	flagSet := newFlagSet()

	if len(os.Args) < 2 {
		flagSet.Usage()
		os.Exit(1)
	}

	flagSet.Parse(os.Args[1:])

	if flagPrintVersion {
		fmt.Fprintf(os.Stdout, "stat version %s\n", version.String())
	} else if err := validateOutputFormat(&flagOutputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "stat: %v\n", err)
	} else if err := validateStorageType(&flagStorageType); err != nil {
		fmt.Fprintf(os.Stderr, "stat: %v\n", err)
	} else if flagSet.NArg() == 0 {
		fmt.Fprintf(os.Stdout, "stat [file ...]\n")
	} else {
		writer := os.Stdout
		var encoder encoding.Encoder
		switch flagOutputFormat {
		case formatCsv: // TODO: fix
			csvEncoder := encoding.NewDelimiter(',', writer)
			csvEncoder.EncodeHeader(stat.StatInfo{})
			encoder = csvEncoder
		case formatJson:
			jsonEncoder := json.NewEncoder(writer)
			jsonEncoder.SetIndent("", "  ")
			encoder = jsonEncoder
		case formatTsv: // TODO: fix
			tsvEncoder := encoding.NewDelimiter('\t', writer)
			tsvEncoder.EncodeHeader(stat.StatInfo{})
			encoder = tsvEncoder
		case formatXml: // TODO: fix
			xmlEncoder := xml.NewEncoder(writer)
			xmlEncoder.Indent("", "  ")
			encoder = xmlEncoder
		case formatYaml:
			encoder = yaml.NewEncoder(writer)
		}

		for _, file := range flagSet.Args() {
			// TODO: range over directories? if recursive option set?
			var err error
			var si *stat.StatInfo
			switch flagStorageType {
			case storageFile:
				si, err = stat.NewFileStat().Stat(file)
			case storageS3:
				si, err = stat.NewS3Stat(nil).Stat(file)
			}

			if err != nil {
				si = &stat.StatInfo{Name: file, Error: err}
			}

			encoder.Encode(si.Pretty(flagDateLayout))
		}
	}
}
