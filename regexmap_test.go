package regexmap

import (
	"testing"

	. "github.com/onsi/gomega"
)

type testcase struct {
	name         string
	inputPattern string
	inputString  string

	expectedMap     map[string]string
	expectedToMatch bool
}

func TestMatch(t *testing.T) {
	testcases := []testcase{
		{
			name:         "basic",
			inputPattern: `I'm (?P<Name>\w+)`,
			inputString:  "Hello, I'm Niko",

			expectedToMatch: true,
			expectedMap: map[string]string{
				"Name": "Niko",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			actualMap, actualHasMatched := MatchPattern(tc.inputPattern, tc.inputString)
			g.Expect(actualHasMatched).Should(Equal(tc.expectedToMatch))
			g.Expect(actualMap).To(Equal(tc.expectedMap))
		})
	}
}
