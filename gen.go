package shortulid

import (
	"fmt"
	"strings"
	"time"

	"github.com/richardlehane/crock32"
)

var EPOCH = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

func New() string {
	elapsedtime := uint64(time.Now().Unix() - EPOCH)
	prefix := fmt.Sprintf("%0s", strings.ToUpper(crock32.Encode(elapsedtime)))
	return prefix + RandStringBytesMaskImprSrcUnsafe(3)
}
