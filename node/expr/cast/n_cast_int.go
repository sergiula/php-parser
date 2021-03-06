package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// CastInt node
type CastInt struct {
	Expr node.Node
}

// NewCastInt node constuctor
func NewCastInt(Expr node.Node) *CastInt {
	return &CastInt{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *CastInt) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastInt) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
