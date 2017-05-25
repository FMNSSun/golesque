package golesque

// 0x10 - add
func Op_Add(context *GLSQContext) error {
	a, b, err := context.PopTwo()

	if err != nil {
		return err
	}

	switch a.GlsqType {
	case GLSQ_TYPE_INT:
		switch b.GlsqType {
		case GLSQ_TYPE_INT:
			context.PushInt(a.GlsqInt + b.GlsqInt)

		default:
			return ErrInvInputTypes
		}

	case GLSQ_TYPE_FLOAT:
		switch b.GlsqType {
		case GLSQ_TYPE_FLOAT:
			context.PushFloat(a.GlsqFloat + b.GlsqFloat)
		}

	case GLSQ_TYPE_BOOL:
		switch b.GlsqType {
		case GLSQ_TYPE_BOOL:
			a_i := boolToInt(a.GlsqBool)
			b_i := boolToInt(b.GlsqBool)
			c_i := (a_i + b_i) % 2

			context.PushBool(intToBool(c_i))
		}

	default:
		return ErrInvInputTypes
	}

	return nil
}
