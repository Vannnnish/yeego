/**
 * Created by angelina on 2017/4/16.
 */

package yeeSql_test

import (
	"github.com/yeeyuntech/yeego/yeeSql"
	"testing"
)

func TestMustSetTableDataToml(t *testing.T) {
	initTestDbTable()
	yeeSql.MustSetTableDataToml(tomlData)
}
