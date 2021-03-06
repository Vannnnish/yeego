/**
 * Created by angelina on 2017/5/20.
 */

package yeexss

import (
	"github.com/vannnnish/yeego/yeestrings"
	"regexp"
)

// 黑名单标签
var BlackLabel = []string{"<iframe>", "</iframe>", "<script>", "</script>", "javascript", "xssm", "script"}

// XssFilter
// 过滤Xss
func XssBlackLabelFilter(s string) string {
	reg, err := regexp.Compile("(?i)" + yeestrings.StringArrayToString(BlackLabel, "|"))
	if err != nil {
		return ""
	}
	if reg.MatchString(s) {
		s = reg.ReplaceAllString(s, "")
	}
	return s
}
