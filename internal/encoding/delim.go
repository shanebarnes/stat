package encoding

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

type Delimiter struct {
	encoder *csvutil.Encoder
	writer  *csv.Writer
}

func NewDelimiter(separator rune, writer io.Writer) *Delimiter {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = separator
	return &Delimiter{
		encoder: csvutil.NewEncoder(csvWriter),
		writer:  csvWriter,
	}
}

func (d *Delimiter) Encode(v interface{}) error {
	err := d.encoder.Encode(v)
	d.writer.Flush()
	return err
}

func (d *Delimiter) EncodeHeader(v interface{}) error {
	err := d.encoder.EncodeHeader(v)
	d.writer.Flush()
	return err
}
