/**
 * Created by angelina on 2017/4/17.
 */

package ast

func NewAndWhereCondition() *AndWhereCondition {
	return &AndWhereCondition{}
}

type AndWhereCondition struct {
	List []PreparedNode
}

func (n *AndWhereCondition) GetPrepareParameter() (string, []string) {
	return joinPrepareNode([]PreparedNode(n.List), ") AND (", "(", ")")
}

func (n *AndWhereCondition) Copy() PreparedNode {
	s := make([]PreparedNode, len(n.List))
	for i := range n.List {
		s[i] = n.List[i].Copy()
	}
	return &AndWhereCondition{List: s}
}

func (n *AndWhereCondition) IsEmpty() bool {
	return len(n.List) == 0
}
func (n *AndWhereCondition) AddAndCondition(node PreparedNode) *AndWhereCondition {
	n.List = append(n.List, node)
	return n
}

func (n *AndWhereCondition) AddPrepare(text string, parameterList ...string) *AndWhereCondition {
	n.List = append(n.List, Prepare(text, parameterList...))
	return n
}

func (n *AndWhereCondition) AddNode(node PreparedNode) *AndWhereCondition {
	n.List = append(n.List, node)
	return n
}
