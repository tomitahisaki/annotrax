package main

import "testing"

func TestContainsAnnotation(t *testing.T) {
	testCases := []struct {
		name            string
		lineText        string
		expectedFound   bool
		expectedKeyword string
	}{
		{
			name:            "line with TODO",
			lineText:        "// TODO: fix this",
			expectedFound:   true,
			expectedKeyword: "TODO",
		},
		{
			name:            "line with FIXME",
			lineText:        "FIXME: handle error properly",
			expectedFound:   true,
			expectedKeyword: "FIXME",
		},
		{
			name:            "line with NOTE",
			lineText:        "# NOTE: temporary workaround",
			expectedFound:   true,
			expectedKeyword: "NOTE",
		},
		{
			name:            "line with HACK",
			lineText:        "// HACK: quick and dirty fix",
			expectedFound:   true,
			expectedKeyword: "HACK",
		},
		{
			name:            "line with XXX",
			lineText:        "// XXX: strange behavior here",
			expectedFound:   true,
			expectedKeyword: "XXX",
		},
		{
			name:            "line without annotation",
			lineText:        "this is just a normal line",
			expectedFound:   false,
			expectedKeyword: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			found, keyword := containsAnnotation(testCase.lineText)

			if found != testCase.expectedFound || keyword != testCase.expectedKeyword {
				t.Fatalf(
					"line=%q: expected (found=%v, keyword=%q), got (found=%v, keyword=%q)",
					testCase.lineText,
					testCase.expectedFound,
					testCase.expectedKeyword,
					found,
					keyword,
				)
			}
		})
	}
}
