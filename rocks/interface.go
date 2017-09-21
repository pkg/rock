package rocks

import "reflect"

// A Rocker is like an EJB for Go
type Rocker interface {
	RockName() string
	RockImportPath() string
	RockInterface() interface{}
	RockInterfaceImportPath() string
	RockAuthor() string
	RockLicense() string
	RockVersion() string
}

func New(t reflect.Type, iface reflect.Type, name string, author string, license string, version string) Rock {
	var r Rock
	// fill in all the rock details here
	r.Name = name
	r.Author = author
	r.License = license
	r.Version = version
	r.Interface = iface
	r.InterfaceImportPath = iface.PkgPath()
	r.ImportPath = t.PkgPath()
	return r
}

// Rock is a struct that implements Rocker
type Rock struct {
	Name                string
	ImportPath          string
	Interface           interface{}
	InterfaceImportPath string
	Author              string
	License             string
	Version             string
}

//RockName returns the Rock's name
func (r Rock) RockName() string {
	return r.Name
}

// RockImportPath returns the Rock's Import Path
func (r Rock) RockImportPath() string {
	return r.ImportPath
}

// RockInterfaces returns the interfaces the rock implements
func (r Rock) RockInterface() interface{} {
	return r.Interface
}

// RockAuthor returns the Rock's Author
func (r Rock) RockAuthor() string {
	return r.Author
}

// RockLicense returns the Rock's license
func (r Rock) RockLicense() string {
	return r.License
}

// RockVersion returns the Rock's version
func (r Rock) RockVersion() string {
	return r.Version
}
