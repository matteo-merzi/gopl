package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c,%s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s%c%s", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var b bytes.Buffer
	b.WriteString(c.fn)
	b.WriteRune('(')
	for i, arg := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteRune(')')
	return b.String()
}

func (v variadic) String() string {
	var b bytes.Buffer
	b.WriteString(v.fn)
	b.WriteRune('(')
	for i, arg := range v.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteRune(')')
	return b.String()
}
