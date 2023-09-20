package io

type stdInFile struct{}

var StdInGenerator FileGeneratorFunc = func() ([]FileObj, error) {
	return []FileObj{}, nil
}
