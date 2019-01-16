/**
 * Created by angelina on 2017/4/16.
 */

package yeesql_test

import (
	"github.com/vannnnish/yeego/yeesql"
	"testing"
)

func TestMustSetTableDataToml(t *testing.T) {
	initTestDbTable()
	yeesql.MustSetTableDataToml(tomlData)
}
