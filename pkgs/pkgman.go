package pkgman
/*
Package Manager - FetchHost

-> central import management for all standard library and external packages

import the root package name, then declare an exportable variable name to the identifier
with an "_" separating the root package name as exampled:

import "pkg"
var Pkg_FuncSameName = pkg.FuncName
*/

////////////////////////////////////////////////////////////////////
// Imports - Standard Lib
import "fmt"


////////////////////////////////////////////////////////////////////
// Imports - External


////////////////////////////////////////////////////////////////////
// Declarations - Standard Lib
var (
	Fmt_Println = fmt.Println
)

////////////////////////////////////////////////////////////////////
// Declarations - External
