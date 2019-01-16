/**
 * Created by angelina on 2017/4/17.
 */

package ast

type Node interface {
	GetText() string
	Copy() Node
}

type PreparedNode interface {
	GetPrepareParameter() (string, []string)
	Copy() PreparedNode
}

type WhereCondition interface {
	IsEmpty() bool
	AddAndCondition(node PreparedNode) *AndWhereCondition
	PreparedNode
}

func Text(text string) StringNodeImpl {
	return StringNodeImpl(text)
}

type StringNodeImpl string

func (s StringNodeImpl) GetText() string {
	return string(s)
}

func (s StringNodeImpl) Copy() Node {
	return s
}

func Prepare(text string, parameterList ...string) StringPrepareNodeImpl {
	return StringPrepareNodeImpl{Text: text, ParameterList: parameterList}
}

type StringPrepareNodeImpl struct {
	Text          string
	ParameterList []string
}

func (n StringPrepareNodeImpl) GetPrepareParameter() (string, []string) {
	return n.Text, n.ParameterList
}

func (n StringPrepareNodeImpl) Copy() PreparedNode {
	s := StringPrepareNodeImpl{}
	s.ParameterList = make([]string, len(n.ParameterList))
	copy(s.ParameterList, n.ParameterList)
	s.Text = n.Text
	return s
}

func (n StringPrepareNodeImpl) IsEmpty() bool {
	return n.Text == ""
}

func (n StringPrepareNodeImpl) AddAndCondition(node PreparedNode) *AndWhereCondition {
	return &AndWhereCondition{List: []PreparedNode{n, node}}
}

func joinPrepareNode(nodeList []PreparedNode, split string, start string, end string) (text string, parameterList []string) {
	text = start
	for i := range nodeList {
		thisText, parameter := nodeList[i].GetPrepareParameter()
		text += thisText
		parameterList = append(parameterList, parameter...)
		if i < len(nodeList)-1 {
			text += split
		}
	}
	text += end
	return
}
