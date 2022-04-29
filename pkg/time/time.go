package time

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(jt).Format(Layout))
	return []byte(stamp), nil
}
