package c8

func checkError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}
