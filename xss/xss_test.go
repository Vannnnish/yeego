/**
 * Created by angelina on 2017/5/20.
 */

package xss

import (
	"github.com/vannnnish/yeego"
	"testing"
)

var (
	str1 = `<ScrIpt>76sajkhfdjhah<iframe>`
	str2 = `<script>alert(1)</script>`
)

func TestXssBlackLabelFilter(t *testing.T) {
	yeego.Equal(XssBlackLabelFilter(str1), "76sajkhfdjhah")
	yeego.Equal(XssBlackLabelFilter(str2), "alert(1)")
}
