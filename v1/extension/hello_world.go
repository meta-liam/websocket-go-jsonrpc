package extension

type HelloWorld struct {
	Version string
}

func (h *HelloWorld) Say(name string) string {
	if name == "" {
		return "hello world"
	}
	return name
}

func (h *HelloWorld) AskAndAnswer(name string, answerCallback string) {
	// return
}
