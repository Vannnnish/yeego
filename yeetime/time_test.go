/**
 * Created by angelina-zf on 17/2/27.
 */
package yeetime

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	println(DateFormat(time.Now(), "YYYY-MM-DD HH:mm:ss"))
}
