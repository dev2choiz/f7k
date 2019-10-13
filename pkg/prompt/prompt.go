package prompt

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type prompt struct {
	Profiles      map[string]*Profile
	Profile       string
	writer        *color.Color
	timePrefix    bool
	verbosity     uint8
	onlyIfVerbose bool
}

func New(profile string) *prompt {
	color.New().Add(color.BgGreen).Add(color.FgBlack).Add(color.Italic)

	pr := &prompt{}
	pr.timePrefix = true
	pr.verbosity = verbosity()
	pr.Profiles = defaultProfiles()
	if "" == profile {
		profile = "info"
	}
	pr.Profile = profile
	err := pr.loadProfile(profile)
	if nil != err {
		panic(err)
	}

	return pr
}

func verbosity() uint8 {
	flag.Parse()
	for _, arg := range flag.Args() {
		if "-v" == arg {
			return 1
		}
		if "-vv" == arg {
			return 2
		}
		if "--verbose" == arg || "-vvv" == arg {
			return 3
		}
	}

	return 0
}

func (pr *prompt) New(name string) *prompt {
	err := pr.loadProfile(name)
	if nil != err {
		pr.loadProfile("info")
	}

	return pr
}

func (pr *prompt) loadProfile(name string) *error {
	p, ok := pr.Profiles[name]
	if !ok {
		list := strings.Join(pr.ProfileNames(), ", ")
		e := fmt.Errorf("%s profile was not found, Did you meant one of thoses ? %s", name, list)
		return &e
	}

	pr.writer = &color.Color{}
	for _, col := range p.Colors {
		pr.writer.Add(col)
	}

	return nil
}

func (pr *prompt) ProfileNames() []string {
	i, keys := 0, make([]string, len(pr.Profiles))
	for key, _ := range pr.Profiles {
		keys[i] = key
		i++
	}
	return keys
}

func (pr *prompt) SetTime(time bool) *prompt {
	pr.timePrefix = time

	return pr
}

func (pr *prompt) OnlyIfVerbose(v bool) *prompt {
	pr.onlyIfVerbose = v

	return pr
}
