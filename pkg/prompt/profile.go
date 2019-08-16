package prompt

import "github.com/fatih/color"

type Profile struct {
	Name   string
	Colors []color.Attribute
}

func defaultProfiles() map[string]*Profile {
	s := []*Profile{
		newProfile("info", color.FgWhite, color.BgBlack),
		newProfile("success", color.FgGreen, color.BgBlack, color.Bold),
		newProfile("danger", color.FgHiRed, color.BgBlack, color.Bold),
		newProfile("fatal", color.FgHiRed, color.BgBlack, color.Bold),
		newProfile("primary", color.FgHiBlue, color.BgBlack),
		newProfile("muted", color.FgHiWhite, color.BgBlack, color.Italic),
	}
	m := make(map[string]*Profile)
	for _, p := range s {
		m[p.Name] = p
	}
	return m
}

func newProfile(name string, attributes ...color.Attribute) *Profile {
	return &Profile{Name: name, Colors: attributes}
}
