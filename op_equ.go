package golesque

// 0x09 - equ
func Op_Equ(context *GLSQContext) error {
	a, b, err := context.PopTwo()

	if err != nil {
		return err
	}

	context.PushBool(equ(a, b))

	return nil
}
