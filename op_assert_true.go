package golesque

// 0x01FFFF - assert_true
func Op_Assert_True(context *GLSQContext) error {
	a, err := context.PopBool()

	if err != nil {
		return err
	}

	if a.GlsqBool != true {
		return ErrAssertionFailed
	}

	return nil
}
