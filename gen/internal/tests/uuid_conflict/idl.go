// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package uuid_conflict

import (
	"go.uber.org/thriftrw/gen/internal/tests/typedefs"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "uuid_conflict",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/uuid_conflict",
	FilePath: "uuid_conflict.thrift",
	SHA1:     "c7ab8450f4c3a548cde8938fe7e150cf1b8f9493",
	Includes: []*thriftreflect.ThriftModule{
		typedefs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./typedefs.thrift\"\n\ntypedef string UUID\n\nstruct UUIDConflict {\n    1: required UUID localUUID\n    2: required typedefs.UUID importedUUID\n}\n"
