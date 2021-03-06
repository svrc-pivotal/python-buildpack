package packager

import "github.com/Masterminds/semver"

type Dependencies []struct {
	URI     string   `yaml:"uri"`
	File    string   `yaml:"file"`
	SHA256  string   `yaml:"sha256"`
	Name    string   `yaml:"name"`
	Version string   `yaml:"version"`
	Stacks  []string `yaml:"cf_stacks"`
	Modules []string `yaml:"modules"`
}
type Manifest struct {
	Language     string       `yaml:"language"`
	IncludeFiles []string     `yaml:"include_files"`
	PrePackage   string       `yaml:"pre_package"`
	Dependencies Dependencies `yaml:"dependencies"`
	Defaults     []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"default_versions"`
}

type File struct {
	Name, Path string
}

func (d Dependencies) Len() int      { return len(d) }
func (d Dependencies) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d Dependencies) Less(i, j int) bool {
	if d[i].Name < d[j].Name {
		return true
	} else if d[i].Name == d[j].Name {
		v1, e1 := semver.NewVersion(d[i].Version)
		v2, e2 := semver.NewVersion(d[j].Version)
		if e1 == nil && e2 == nil {
			return v1.LessThan(v2)
		} else {
			return d[i].Version < d[j].Version
		}
	}
	return false
}
