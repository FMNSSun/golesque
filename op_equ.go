package golesque

// 0x09 - equ
func Op_Equ(context *GLSQContext) error {
	a, b, err := context.PopTwo()

	if err != nil {
		return err
	}

	if a.GlsqType != b.GlsqType {
		context.PushBool(false)
		return nil
	}

	switch a.GlsqType {
	case GLSQ_TYPE_INT:
		if a.GlsqInt != b.GlsqInt {
			context.PushBool(false)
			return nil
		}

	case GLSQ_TYPE_FLOAT:
		if a.GlsqFloat != b.GlsqFloat {
			context.PushBool(false)
			return nil
		}

	case GLSQ_TYPE_BOOL:
		if a.GlsqBool != b.GlsqBool {
			context.PushBool(false)
			return nil
		}

	default:
		context.PushBool(false)
		return nil
	}

	context.PushBool(true)
	return nil
}
