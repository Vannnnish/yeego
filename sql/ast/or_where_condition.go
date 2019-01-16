/**
 * Created by angelina on 2017/4/17.
 */

package ast

func NewOrWhereCondition() *OrWhereCondition {
	return &OrWhereCondition{}
}

type OrWhereCondition struct {
	List []PreparedNode
}

func (n *OrWhereCondition) GetPrepareParameter() (string, []string) {
	return joinPrepareNode([]PreparedNode(n.List), ") OR (", "(", ")")
}

func (n *OrWhereCondition) Copy() PreparedNode {
	s := make([]PreparedNode, len(n.List))
	for i := range n.List {
		s[i] = n.List[i].Copy()
	}
	return &OrWhereCondition{List: s}
}

func (n *OrWhereCondition) IsEmpty() bool {
	return len(n.List) == 0
}

func (n *OrWhereCondition) AddAndCondition(node PreparedNode) *AndWhereCondition {
	return NewAndWhereCondition().AddNode(n).AddNode(node)
}

func (n *OrWhereCondition) AddPrepare(text string, parameterList ...string) *OrWhereCondition {
	n.List = append(n.List, Prepare(text, parameterList...))
	return n
}

func (n *OrWhereCondition) AddCondition(node PreparedNode) *OrWhereCondition {
	n.List = append(n.List, node)
	return n
}
