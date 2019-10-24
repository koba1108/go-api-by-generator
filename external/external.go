package external

func Init() {
	initDB()
}

func Finalize() {
	closeDB()
}
