package logrusltsv

import (
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type UserValue struct{}

func (*UserValue) String() string {
	return ".String() was\tused\n"
}

func TestFormatter(t *testing.T) {
	assert := assert.New(t)

	loc, err := time.LoadLocation("Asia/Tokyo")
	assert.NoError(err)

	entry := &logrus.Entry{
		Data: logrus.Fields{
			"app":    "test",
			"time":   "field time",
			"msg":    "field msg",
			"level":  "field level",
			"custom": &UserValue{},

			"field\nwith\tspecial": "value\nwith\tspecial",
		},
		Time:    time.Date(2015, 1, 30, 16, 01, 47, 89, loc),
		Message: "test\a message\n",
		Level:   logrus.ErrorLevel,
	}

	// Fields are sorted by alphabetical
	// "time", "msg", "level" fields are prefiexed with "field."
	// Special charactors are quoted
	// Field values are converted to string with /String() if it has.
	// Line ends with "\n"
	f := &Formatter{}
	output, err := f.Format(entry)
	assert.NoError(err)
	assert.Equal(
		strings.Join([]string{
			`time:2015-01-30T16:01:47+09:00`,
			`level:error`,
			`msg:test\a message\n`,
			`app:test`,
			`custom:.String() was\tused\n`,
			`field.level:field level`,
			`field.msg:field msg`,
			`field.time:field time`,
			`field\nwith\tspecial:value\nwith\tspecial`,
		}, "\t")+"\n",
		string(output),
	)
}
