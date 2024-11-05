package primitives

import (
	"strings"

	"github.com/bearz-io/bzdev/lib/omap"
	"gopkg.in/yaml.v3"
)

type Expr struct {
	Value     any
	Template  string
	Evaluated bool
	Line      int
	Col       int
}

type ExprsMap struct {
	omap.OrderedMap[string, Expr]
}

func (in *Expr) UnmarshalYAML(node *yaml.Node) error {

	if node.Kind == yaml.ScalarNode {
		in.Col = node.Column
		in.Line = node.Line
		if strings.Contains("${{", node.Value) {
			in.Template = node.Value
		} else {
			in.Value = node.Value
			in.Evaluated = true
		}
		return nil
	}

	return nil
}
