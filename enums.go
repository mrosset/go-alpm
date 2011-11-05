package alpm

// Install reason of a package.
type PkgReason uint

const (
	PkgReasonExplicit PkgReason = 0
	PkgReasonDepend   PkgReason = 1
)

// Source of a package structure.
type PkgFrom uint

const (
	_           PkgFrom = iota
	FromFile    PkgFrom = iota
	FromLocalDB PkgFrom = iota
	FromSyncDB  PkgFrom = iota
)

// Dependency constraint types.
type DepMod uint

const (
	_         DepMod = iota
	DepModAny DepMod = iota // Any version.
	DepModEq  DepMod = iota // Specific version.
	DepModGE  DepMod = iota // Test for >= some version.
	DepModLE  DepMod = iota // Test for <= some version.
	DepModGT  DepMod = iota // Test for > some version.
	DepModLT  DepMod = iota // Test for < some version.
)

func (mod DepMod) String() string {
	switch mod {
	case DepModEq:
		return "="
	case DepModGE:
		return ">="
	case DepModLE:
		return "<="
	case DepModGT:
		return ">"
	case DepModLT:
		return "<"
	}
	return ""
}

// Signature checking level.
type SigLevel uint

const (
	SigPackage           SigLevel = 1 << 0
	SigPackageOptional   SigLevel = 1 << 1
	SigPackageMarginalOk SigLevel = 1 << 2
	SigPackageUnknownOk  SigLevel = 1 << 3

	SigDatabase           SigLevel = 1 << 10
	SigDatabaseOptional   SigLevel = 1 << 11
	SigDatabaseMarginalOk SigLevel = 1 << 12
	SigDatabaseUnknownOk  SigLevel = 1 << 13

	SigUseDefault = 1 << 31
)
