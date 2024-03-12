package regexmap

import "regexp"

func MatchPattern(r, s string) (map[string]string, bool) {
	return MatchRegex(regexp.MustCompile(r), s)
}

func MatchRegex(r *regexp.Regexp, s string) (map[string]string, bool) {
	groupNames := r.SubexpNames()
	matches := r.FindStringSubmatch(s)
	if nameCount := len(groupNames); nameCount != len(matches) || nameCount == 0 {
		return nil, false
	}
	groupNames = groupNames[1:]
	matches = matches[1:]
	captureMap := make(map[string]string)
	for index, currentGroupName := range groupNames {
		captureMap[currentGroupName] = matches[index]
	}
	return captureMap, true
}
