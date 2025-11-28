package main

import "testing"

func TestIsProbablyBinary(t *testing.T) {
	testCases := []struct {
		name             string
		path             string
		expectedIsBinary bool
	}{
		{
			name:             "png file is binary",
			path:             "image.png",
			expectedIsBinary: true,
		},
		{
			name:             "jpeg file is binary",
			path:             "photo.jpeg",
			expectedIsBinary: true,
		},
		{
			name:             "zip file is binary",
			path:             "archive.zip",
			expectedIsBinary: true,
		},
		{
			name:             "svg file is binary",
			path:             "icon.svg",
			expectedIsBinary: true,
		},
		{
			name:             "go source file is not binary",
			path:             "main.go",
			expectedIsBinary: false,
		},
		{
			name:             "text file is not binary",
			path:             "README.txt",
			expectedIsBinary: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			isBinary := isNonTextFile(testCase.path)

			if isBinary != testCase.expectedIsBinary {
				t.Fatalf(
					"path=%q: expected isBinary=%v, got %v",
					testCase.path,
					testCase.expectedIsBinary,
					isBinary,
				)
			}
		})
	}
}
