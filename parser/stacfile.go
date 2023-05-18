package parser

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Pipeline struct {
	Stages []Stage `yaml:"pipeline"`
}

type Stage struct {
	Commands []string `yaml:"commands"`
}

func ParseYaml(path string) ([]Stage, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	m := Pipeline{}
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, err
	}
	return m.Stages, nil
}
