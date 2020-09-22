package happy

import (
	"github.com/orian/wrapgen/internal/test/happy"
)

type DemoType struct{}

type Demo interface {
	Make(param happy.ExportedStruct, second DemoType) happy.NonInterfaceAlias
}
