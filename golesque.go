package golesque

import "errors"
import "unicode/utf8"
import "fmt"

var ErrInvEOF error = errors.New("Invalid EOF. Expected more bytes!")
var ErrInvStackSize error = errors.New("Invalid stack size!")
var ErrInvInputTypes error = errors.New("Invalid input types!")
var ErrInvOp error = errors.New("Invalid operation!")
var ErrAssertionFailed error = errors.New("Assertion failed!")
var ErrWantBool error = errors.New("Invalid type... wanted bool!")

const GLSQ_TYPE_INT uint8 = 1
const GLSQ_TYPE_BOOL uint8 = 2
const GLSQ_TYPE_CHAR uint8 = 3
const GLSQ_TYPE_STR uint8 = 4
const GLSQ_TYPE_FLOAT uint8 = 5

var DebugDump bool = true

type GLSQObj struct {
	GlsqType  uint8
	GlsqInt   int64
	GlsqBool  bool
	GlsqChar  rune
	GlsqStr   string
	GlsqRunes []rune
	GlsqFloat float64
}

func (gobj *GLSQObj) decodeUTF8() {
	runeCount := utf8.RuneCountInString(gobj.GlsqStr)

	runes := make([]rune, runeCount)

	for index, runeValue := range gobj.GlsqStr {
		runes[index] = runeValue
	}

	gobj.GlsqRunes = runes
}

func (gobj *GLSQObj) Runes() []rune {
	if gobj.GlsqRunes == nil {
		gobj.decodeUTF8()
	}

	return gobj.GlsqRunes
}

func (gobj *GLSQObj) RuneCount() int {
	if gobj.GlsqRunes == nil {
		gobj.decodeUTF8()
	}

	return len(gobj.GlsqRunes)
}

func (gobj *GLSQObj) GetNthRune(index int) rune {
	if gobj.GlsqRunes == nil {
		gobj.decodeUTF8()
	}

	return gobj.GlsqRunes[index]
}

type GLSQContext struct {
	Stack  []*GLSQObj
	Sp     int
	Failed bool
}

func (c *GLSQContext) PushObj(g *GLSQObj) {
	c.Stack[c.Sp] = g
	c.Sp++
}

func (c *GLSQContext) PopObj() (*GLSQObj, error) {
	if c.Sp == 0 {
		return nil, ErrInvStackSize
	}

	c.Sp--

	obj := c.Stack[c.Sp]

	return obj, nil
}

func (c *GLSQContext) PopBool() (*GLSQObj, error) {
	a, err := c.PopObj()

	if err != nil {
		return nil, err
	}

	if a.GlsqType != GLSQ_TYPE_BOOL {
		return nil, ErrWantBool
	}

	return a, nil
}

func (c *GLSQContext) PushBool(value bool) {
	c.PushObj(&GLSQObj{GlsqType: GLSQ_TYPE_BOOL, GlsqBool: value})
}

func (c *GLSQContext) PushChar(value rune) {
	c.PushObj(&GLSQObj{GlsqType: GLSQ_TYPE_CHAR, GlsqChar: value})
}

func (c *GLSQContext) PushCharStr(value rune) {
	str := string(value)
	c.PushObj(&GLSQObj{GlsqType: GLSQ_TYPE_STR, GlsqStr: str})
}

func (c *GLSQContext) PushInt(value int64) {
	c.PushObj(&GLSQObj{GlsqType: GLSQ_TYPE_INT, GlsqInt: value})
}

func (c *GLSQContext) PushFloat(value float64) {
	c.PushObj(&GLSQObj{GlsqType: GLSQ_TYPE_FLOAT, GlsqFloat: value})
}

func (c *GLSQContext) PopTwo() (*GLSQObj, *GLSQObj, error) {
	if c.Sp < 2 {
		return nil, nil, ErrInvStackSize
	}

	b, _ := c.PopObj()
	a, _ := c.PopObj()

	return a, b, nil
}

func Dump(code []byte) error {
	slice := code

	for {
		if len(slice) < 1 {
			return nil
		}

		c1 := slice[0]

		switch c1 {
		case 0x00:
			// Read in another byte
			if len(slice) < 2 {
				return ErrInvEOF
			}
			c2 := slice[1]

			slice = slice[2:]

			name, ok := OpNames[(uint32(uint32(c2) << 8))]

			if !ok {
				fmt.Println("n/a")
			}

			fmt.Println(name)

		case 0x01:
			// Read in two other bytes
			if len(slice) < 3 {
				return ErrInvEOF
			}

			c2 := slice[1]
			c3 := slice[2]

			slice = slice[3:]

			name, ok := OpNames[uint32(uint32(c2)<<24|uint32(c3)<<16)]

			if !ok {
				fmt.Println("n/a")
			}

			fmt.Println(name)

		case 0x02:
			fmt.Println("ldc.b.f")

			slice = slice[1:]

		case 0x03:
			// Push constant true
			fmt.Println("ldc.b.t")

			slice = slice[1:]

		case 0x04:
			// Push 32bit int

			if len(slice) < 5 {
				return ErrInvEOF
			}

			c2 := slice[1]
			c3 := slice[2]
			c4 := slice[3]
			c5 := slice[4]

			i := uint32(uint32(c5)<<24 | uint32(c4)<<16 | uint32(c3)<<8 | uint32(c2))

			var si int32 = 0

			if i&0x80000000 != 0 {
				si = -1 * int32(i&0x7FFFFFFF)
			} else {
				si = int32(i & 0x7FFFFFFF)
			}

			slice = slice[5:]

			fmt.Printf("ldc.i.4\t0x%02x%02x%02x%02x\t\t;%d\n", c5, c4, c3, c2, si)

		default:
			slice = slice[1:]

			name, ok := OpNames[uint32(c1)]

			if !ok {
				fmt.Println("n/a")
			}

			fmt.Println(name)
		}
	}
}

func Run(code []byte, context *GLSQContext) error {
	slice := code

	for {
		if len(slice) < 1 {
			return nil
		}

		c1 := slice[0]

		switch c1 {
		case 0x00:
			// Read in another byte
			if len(slice) < 2 {
				return ErrInvEOF
			}
			c2 := slice[1]

			slice = slice[2:]

			err := op(uint32(uint32(c2)<<8), context)

			if err != nil {
				return err
			}

		case 0x01:
			// Read in two other bytes
			if len(slice) < 3 {
				return ErrInvEOF
			}

			c2 := slice[1]
			c3 := slice[2]

			slice = slice[3:]

			err := op(uint32(uint32(c2)<<24|uint32(c3)<<16), context)

			if err != nil {
				return err
			}

		case 0x02:
			// Push constant false
			context.PushBool(false)

			slice = slice[1:]

		case 0x03:
			// Push constant true
			context.PushBool(true)

			slice = slice[1:]

		case 0x04:
			// Push 32bit int

			if len(slice) < 5 {
				return ErrInvEOF
			}

			c2 := slice[1]
			c3 := slice[2]
			c4 := slice[3]
			c5 := slice[4]

			i := uint32(uint32(c5)<<24 | uint32(c4)<<16 | uint32(c3)<<8 | uint32(c2))

			var si int32 = 0

			if i&0x80000000 != 0 {
				si = -1 * int32(i&0x7FFFFFFF)
			} else {
				si = int32(i & 0x7FFFFFFF)
			}

			context.PushInt(int64(si))

			slice = slice[5:]

		default:
			slice = slice[1:]

			err := op(uint32(c1), context)

			if err != nil {
				return err
			}
		}
	}
}

var ops map[uint32](func(*GLSQContext) error) = map[uint32](func(*GLSQContext) error){
	0x00000009: Op_Equ,
	0x00000010: Op_Add,
	0xFFFF0000: Op_Assert_True,
}

var OpNames map[uint32]string = map[uint32]string{
	0x00000009: "equ",
	0x00000010: "add",
	0xFFFF0000: "assert_true",
}

func op(opc uint32, context *GLSQContext) error {
	opf, ok := ops[opc]

	if !ok {
		return ErrInvOp
	}

	return opf(context)
}
