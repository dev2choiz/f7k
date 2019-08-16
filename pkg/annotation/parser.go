package annotation

import (
	"encoding/json"
	"regexp"
)

const annotationPattern = `(?Us)\@(?P<name>[a-zA-Z][a-zA-Z-0-9]*)\s*\=\s*(?P<value>(\[.*\]|\{.*\}))\n`

func parseMethodAnnotations(com string) []IAnnotation {
	lines := extractAnnotationsStrings(com)
	annotations := make([]IAnnotation, 0)
	for _, line := range lines {
		annotation := &Annotation{}
		annotation.plainText = line

		var re = regexp.MustCompile(annotationPattern)
		matches := re.FindStringSubmatch(line)
		names := re.SubexpNames()
		params := make(map[string]string)
		for i, name := range names {
			if i > 0 && i <= len(matches) {
				params[name] = matches[i]
			}
		}

		err := json.Unmarshal([]byte(params["value"]), &annotation.dataAssoc)
		if nil != err {
			err = json.Unmarshal([]byte(params["value"]), &annotation.data)
			if nil != err {
				continue
			}
		}
		annotation.name = params["name"]
		annotations = append(annotations, annotation)
	}

	return annotations
}

func extractAnnotationsStrings(comStr string) []string {
	return regexp.MustCompile(annotationPattern).FindAllString(comStr, -1)
}
