package logrusltsv

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
)

type (
	Formatter struct{}

	byKey [][2]string
)

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	commonItems := [][2]string{
		{"time", entry.Time.Format(time.RFC3339)},
		{"level", entry.Level.String()},
		{"msg", escape(entry.Message)},
	}

	items := [][2]string{}
	for k, v := range entry.Data {
		if k == "time" || k == "msg" || k == "level" {
			k = "field." + k
		}
		items = append(items, [2]string{escape(k), escape(fmt.Sprint(v))})
	}
	sort.Sort(byKey(items))

	return encodeLTSV(append(commonItems, items...)), nil
}

func (items byKey) Len() int           { return len(items) }
func (items byKey) Less(i, j int) bool { return items[i][0] < items[j][0] }
func (items byKey) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }

func encodeLTSV(items [][2]string) []byte {
	b := bytes.Buffer{}
	for i, item := range items {
		b.WriteString(item[0])
		b.WriteString(":")
		b.WriteString(item[1])

		if i == len(items)-1 {
			b.WriteString("\n")
		} else {
			b.WriteString("\t")
		}
	}
	return b.Bytes()
}

func escape(s string) string {
	v := strconv.Quote(s)
	return v[1 : len(v)-1]
}
