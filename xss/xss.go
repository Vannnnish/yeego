/**
 * Created by angelina on 2017/5/20.
 */

package xss

import (
	"github.com/vannnnish/yeego/strings"
	"regexp"
)

// 黑名单标签
var BlackLabel = []string{"<iframe>", "</iframe>", "<script>", "</script>", "javascript", "xssm", "script"}

// XssFilter
// 过滤Xss
func XssBlackLabelFilter(s string) string {
	reg, err := regexp.Compile("(?i)" + strings.StringArrayToString(BlackLabel, "|"))
	if err != nil {
		return ""
	}
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, "")
	}
	return s
}
