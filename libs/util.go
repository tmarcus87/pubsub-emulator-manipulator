package libs

func MustNotEmpty(pV *string, errorMessage string, onError func()) {
	err := false
	defer func() {
		if onError != nil && err {
			onError()
		}
	}()
	if pV == nil {
		err = true
		defer panic("[nil] " + errorMessage)
	}
	if *pV == "" {
		err = true
		defer panic("[empty] " + errorMessage)
	}
}
