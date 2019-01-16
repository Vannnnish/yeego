/**
 * Created by angelina on 2017/4/17.
 */

package ast

type SelectCommand struct {
	selectExprList     Node
	tableReferenceList Node
	whereCondition     WhereCondition
	groupByList        Node
	have               WhereCondition
	orderByList        Node
	limit              Node
}

func NewSelectCommand() *SelectCommand {
	return &SelectCommand{}
}

func (c *SelectCommand) GetPrepareParameter() (output string, parameterList []string) {
	output = "SELECT "
	if c.selectExprList == nil {
		output += "*"
	} else {
		output += c.selectExprList.GetText()
	}
	if c.tableReferenceList != nil {
		output += " FROM " + c.tableReferenceList.GetText()
	}
	if c.whereCondition != nil && !c.whereCondition.IsEmpty() {
		text, parameter := c.whereCondition.GetPrepareParameter()
		output += " WHERE " + text
		parameterList = append(parameterList, parameter...)
	}
	if c.groupByList != nil {
		output += " GROUP BY " + c.groupByList.GetText()
	}
	if c.have != nil && c.have.IsEmpty() {
		text, parameter := c.have.GetPrepareParameter()
		output += " HAVING " + text
		parameterList = append(parameterList, parameter...)
	}
	if c.orderByList != nil {
		output += " ORDER BY " + c.orderByList.GetText()
	}
	if c.limit != nil {
		output += " LIMIT " + c.limit.GetText()
	}
	return
}

func (c *SelectCommand) Copy() *SelectCommand {
	s := &SelectCommand{}
	if c.selectExprList != nil {
		s.selectExprList = c.selectExprList.Copy()
	}
	if c.tableReferenceList != nil {
		s.tableReferenceList = c.tableReferenceList.Copy()
	}
	if c.whereCondition != nil {
		s.whereCondition = c.whereCondition.Copy().(WhereCondition)
	}
	if c.groupByList != nil {
		s.groupByList = c.groupByList.Copy()
	}
	if c.have != nil {
		s.have = c.have.Copy().(WhereCondition)
	}
	if c.orderByList != nil {
		s.orderByList = c.orderByList.Copy()
	}
	if c.limit != nil {
		s.limit = c.limit.Copy()
	}
	return s
}

func (c *SelectCommand) Select(text string) *SelectCommand {
	c.selectExprList = Text(text)
	return c
}

func (c *SelectCommand) From(text string) *SelectCommand {
	c.tableReferenceList = Text(text)
	return c
}

func (c *SelectCommand) Where(text string, parameterList ...string) *SelectCommand {
	c.whereCondition = Prepare(text, parameterList...)
	return c
}

func (c *SelectCommand) WhereObj(obj WhereCondition) *SelectCommand {
	c.whereCondition = obj
	return c
}

func (c *SelectCommand) GroupBy(text string) *SelectCommand {
	c.groupByList = Text(text)
	return c
}

func (c *SelectCommand) Have(text string, parameterList ...string) *SelectCommand {
	c.have = Prepare(text, parameterList...)
	return c
}

func (c *SelectCommand) OrderBy(text string) *SelectCommand {
	c.orderByList = Text(text)
	return c
}

func (c *SelectCommand) Limit(text string) *SelectCommand {
	c.limit = Text(text)
	return c
}

func (c *SelectCommand) GetWhereCondition() WhereCondition {
	return c.whereCondition
}

func (c *SelectCommand) AddAndWhereCondition(node PreparedNode) *SelectCommand {
	c.whereCondition = c.whereCondition.AddAndCondition(node)
	return c
}
