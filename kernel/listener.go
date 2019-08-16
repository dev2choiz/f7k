package kernel

import (
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"net/http"
)

func (k *Kernel) ListenRequests() {
	http.HandleFunc("/", k.Handle)

	p := f7k.AppConfig.GetPort()
	prompt.New("success").Printfln("Listen port %d\n", p)
	prompt.New("fatal").Println(http.ListenAndServe(fmt.Sprintf(":%d", p), nil).Error())
}
