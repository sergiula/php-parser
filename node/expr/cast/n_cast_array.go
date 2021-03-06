package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// CastArray node
type CastArray struct {
	Expr node.Node
}

// NewCastArray node constuctor
func NewCastArray(Expr node.Node) *CastArray {
	return &CastArray{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *CastArray) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastArray) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
