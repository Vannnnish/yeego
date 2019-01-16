/**
 * Created by angelina on 2017/4/17.
 */

package ast

import (
	"github.com/vannnnish/yeego"
	"testing"
)

func TestSelectCommand(t *testing.T) {
	text, parameterList := NewSelectCommand().From("testTable").Where("Id=1").GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM testTable WHERE Id=1")
	yeego.Equal(len(parameterList), 0)
}

func TestSelectCommand2(t *testing.T) {
	s1 := NewSelectCommand().From("Table1").Where("a=1 AND b=?", "c").Limit("1,2")
	text, parameterList := s1.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE a=1 AND b=? LIMIT 1,2")
	yeego.Equal(parameterList, []string{"c"})

	s2 := s1.Copy()
	s2.Limit("2,2")
	text, _ = s1.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE a=1 AND b=? LIMIT 1,2")
	text, _ = s2.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE a=1 AND b=? LIMIT 2,2")
}

func TestAndWhereConditionAddPrepare(t *testing.T) {
	and := NewAndWhereCondition().AddPrepare("a=1").AddPrepare("d=?", "c")
	s1 := NewSelectCommand().From("Table1").WhereObj(and)
	text, parameterList := s1.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE (a=1) AND (d=?)")
	yeego.Equal(parameterList, []string{"c"})

	s2 := s1.Copy()
	text, parameterList = s2.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE (a=1) AND (d=?)")
	yeego.Equal(parameterList, []string{"c"})
}

func TestOrWhereConditionAddPrepare(t *testing.T) {
	and := NewOrWhereCondition().AddPrepare("a=1").AddPrepare("d=?", "c")
	s1 := NewSelectCommand().From("Table1").WhereObj(and)
	text, parameterList := s1.GetPrepareParameter()
	yeego.Equal(text, "SELECT * FROM Table1 WHERE (a=1) OR (d=?)")
	yeego.Equal(parameterList, []string{"c"})
}
