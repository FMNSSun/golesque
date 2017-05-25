package golesque

/*
 * Contains comparison functions.
 */

// Returns true if the objects are equal.
// Obviously if a and b point to the same object
// they must be equal. Does not do type conversions!
func equ(a *GLSQObj, b *GLSQObj) bool {
	if a == b {
		// If they point to the same location they obviously are equal.
		return true
	}

	if a.GlsqType != b.GlsqType {
		return false // can't be equal if types are not the same
	}

	switch a.GlsqType {
	case GLSQ_TYPE_INT:
		return a.GlsqInt == b.GlsqInt

	case GLSQ_TYPE_FLOAT:
		return a.GlsqFloat == b.GlsqFloat

	case GLSQ_TYPE_BOOL:
		return a.GlsqBool == b.GlsqBool
	}

	return false
}
