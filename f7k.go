package f7k

import (
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"github.com/dev2choiz/f7k/utils"
)

var AppConfig interfaces.ConfigInterface
var AppLoader interfaces.AppLoader
var AppPath interfaces.AppPath
var Dispatcher interfaces.EventDispatcher
var Kernel interfaces.Kernel
var Router interfaces.Router
var ViewConfig interfaces.ConfViewer

// Without dependencies
var Prompt = prompt.New("info")
var Utils = utils.Instance()

var Name = "f7k"
var PrettyName = "F7k"
var Version string

// Variable set in cobra command only, because his value come from flags
var Verbose = false
