package graphic

import "github.com/l-vitaly/wb-practice/visitor/pkg/visitor"

type Graphic interface {
	Accept(v visitor.Visitor)
}
