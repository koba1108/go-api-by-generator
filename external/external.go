package external

func Init() {
	// initDB() :todo あとで直す
	initFirebase()
}

func Finalize() {
	closeDB()
}
