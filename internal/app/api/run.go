package api

func Run() {
	d := Dependency{}
	d.Inject()

	err := routing(d)
	if err != nil {
		panic(err)
	}
}
