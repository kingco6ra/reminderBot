package times

import (
	"errors"
	"strings"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/olebedev/when/rules/ru"
)

var ErrorTimePattern = errors.New("ErrorTimePattern")

func ParseReminderTime(now time.Time, text string) (*time.Time, error) {
	w := when.New(nil)
	w.Add(ru.All...)
    w.Add(en.All...)

	r, e := w.Parse(strings.TrimSpace(text), now)
    if r == nil {
        return nil, ErrorTimePattern
    }
	
	if e != nil {
		return nil, e
	}

	return &r.Time, nil
}
