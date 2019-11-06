package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
)

type FormFile struct {
	Roots []FormTreeNode
	Nodes []FormTreeNode
	Enums []*FormEnum
}

type FormTreeNode interface {
	Message() *Message
	Extension() *form.Node
	Root() FormTreeNode
	Parent() FormTreeNode
	Children() []FormTreeNode
	IsRoot() bool
}

type formTreeNode struct {
	message   *Message
	extension form.Node
	parent    FormTreeNode
	children  []FormTreeNode
}

func (nd *formTreeNode) Message() *Message        { return nd.message }
func (nd *formTreeNode) Extension() *form.Node    { return &nd.extension }
func (nd *formTreeNode) Root() FormTreeNode       { return getRoot(nd) }
func (nd *formTreeNode) Parent() FormTreeNode     { return nd.parent }
func (nd *formTreeNode) Children() []FormTreeNode { return nd.children }

func getRoot(nd FormTreeNode) FormTreeNode {
	for {
		if nd.IsRoot() {
			return nd
		}
		nd = nd.Parent()
	}
}

func (nd *formTreeNode) IsRoot() bool {
	return nd.parent == nil
}

type FormField struct {
	*Field
	Extension form.Field
}

// TODO remove
type FormOneOf struct {
	*OneOf
	Required bool
}

type FormEnum struct {
	*Enum
	Extension form.Field
}

type FormEnumValue struct {
	*EnumValue
	Extension form.InputOption
}
