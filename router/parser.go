package router

import (
	"fmt"
	"github.com/dev2choiz/f7k/controllers"
	"github.com/dev2choiz/f7k/interfaces"
	"regexp"
	"strings"
)

const paramPattern = "(?P<name>[a-z][a-zA-Z0-9-_]+)"
const valueParamPattern = "(?P<name>[a-zA-Z0-9-_]+)"

func (r *router) PopulateWithControllersAnnotations() {
	cr := controllers.RegistryInstance()
	for fullPackName, anns := range cr.Annotations {
		for _, ann := range anns {
			if "route" != ann.Name() {
				continue
			}

			aPack := strings.Split(fullPackName, ".")
			pack := aPack[0]
			ctrl := aPack[1]
			action := aPack[2]



			route := &Route{routeMetadata: &routeMetadata{}}
			route.SetName(ann.DataAssoc()["name"].(string))
			route.SetPath(ann.DataAssoc()["path"].(string))
			route.SetPackage(pack)
			route.SetController(ctrl)
			route.SetAction(action)
			route.SetMethods(ann.DataAssoc()["methods"].([]string))
			r.routerMetadata.Routes[route.name] = route
		}
	}
}


func (r *router) ParseParameters(route interfaces.Route) {
	regEx, url := fmt.Sprintf("(?s){%s}", paramPattern), route.Path()
	checkPattern:= fmt.Sprintf("(?s)^%s$", url)

	compRegEx := regexp.MustCompile(regEx)
	matches := compRegEx.FindAllStringSubmatch(url, -1)

	if nil == matches {
		route.SetCheckPattern(checkPattern)
		return
	}

	var k int
	for i, regexMarker := range compRegEx.SubexpNames() {
		if "name" == regexMarker {
			k = i
			break
		}
	}

	for i, match := range matches {
		param := &RouteParam{}
		param.SetRequirements(make(map[string]interfaces.ParamRequirement))
		param.SetName(match[k])
		param.SetOrder(uint8(i + 1))
		pp := strings.Replace(valueParamPattern, "name", param.Name(), 1)
		checkPattern = strings.Replace(checkPattern, "{" + param.Name() + "}", pp, 1)

		route.AddParam(param)
	}
	route.SetCheckPattern(checkPattern)

	return
}

func (r *router) ParseRequirements(route interfaces.Route) {
	return
}
