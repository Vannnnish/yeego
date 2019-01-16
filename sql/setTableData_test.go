/**
 * Created by angelina on 2017/4/16.
 */

package sql_test

import (
	"github.com/vannnnish/yeego/sql"
	"testing"
)

func TestMustSetTableDataToml(t *testing.T) {
	initTestDbTable()
	sql.MustSetTableDataToml(tomlData)
}
