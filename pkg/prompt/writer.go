package prompt

import (
	"os"
	"time"
)

func (pr *prompt) Write(p []byte) (n int, err error) {
	t := ""
	if pr.timePrefix {
		loc, err := time.LoadLocation("Europe/Paris")
		if nil != err {
			panic(err)
		}
		t = time.Now().In(loc).String() + "  "
	}

	i, e := pr.writer.Print(pr.Sprintf( "%s%s", t, string(p)))
	if "fatal" == pr.Profile {
		os.Exit(1)
	}

	return i, e
}

func (pr *prompt) Print(s string) {
	_, _ = pr.Write([]byte(s))
}

func (pr *prompt) Println(s string) {
	pr.Print(s + "\n")
}

func (pr *prompt) Printf(s string, args ...interface{}) {
	pr.Print(pr.Sprintf(s, args...))
}

func (pr *prompt) Sprintf(f string, args ...interface{}) string {
	return pr.writer.Sprintf(f, args...)
}

func (pr *prompt) Printfln(s string, args ...interface{}) {
	pr.Println(pr.Sprintf(s, args...))
}
