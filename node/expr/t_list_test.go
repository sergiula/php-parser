package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign_op"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestList(t *testing.T) {
	src := `<? list($a) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign_op.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestListArrayIndex(t *testing.T) {
	src := `<? list($a[]) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign_op.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.ArrayDimFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestListList(t *testing.T) {
	src := `<? list(list($a)) = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign_op.Assign{
					Variable: &expr.List{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											ByRef: false,
											Val:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
