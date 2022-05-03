package utils

func DbgPrint(msg string, dbg bool) {
	if dbg {
		println(msg)
	}
}
