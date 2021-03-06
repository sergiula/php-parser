package stmt_test

import (
	"github.com/z7zmey/php-parser/node/name"
	"bytes"
	"testing"
	
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTraitUse(t *testing.T) {
	src := `<? class Foo { use Bar; }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName: &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestTraitsUse(t *testing.T) {
	src := `<? class Foo { use Bar, Baz; }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName: &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestTraitsUseAdaptions(t *testing.T) {
	src := `<? class Foo { use Bar, Baz { Bar::one insteadof Baz, Quux; Baz::one as two; } }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				PhpDocComment: "",
				ClassName: &node.Identifier{Value: "Foo"},
				Stmts: []node.Node{
					&stmt.TraitUse{
						Traits: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "Baz"},
								},
							},
						},
						Adaptations: []node.Node{
							&stmt.TraitUsePrecedence{
								Ref: &stmt.TraitMethodRef{
									Trait: &name.Name{
										Parts: []node.Node{
											&name.NamePart{Value: "Bar"},
										},
									},
									Method: &node.Identifier{Value: "one"},
								},
								Insteadof: []node.Node{
									&name.Name{
										Parts: []node.Node{
											&name.NamePart{Value: "Baz"},
										},
									},
									&name.Name{
										Parts: []node.Node{
											&name.NamePart{Value: "Quux"},
										},
									},
								},
							},
							&stmt.TraitUseAlias{
								Ref: &stmt.TraitMethodRef{
									Trait: &name.Name{
										Parts: []node.Node{
											&name.NamePart{Value: "Baz"},
										},
									},
									Method: &node.Identifier{Value: "one"},
								},
								Alias: &node.Identifier{Value: "two"},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
