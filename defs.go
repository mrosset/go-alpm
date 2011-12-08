// godefs -g alpm defs.c

// MACHINE GENERATED - DO NOT EDIT.

package alpm

// Constants
const (
	PKG_REASON_EXPLICIT      = 0
	PKG_REASON_DEPEND        = 0x1
	PKG_FROM_FILE            = 0x1
	PKG_FROM_LOCALDB         = 0x2
	PKG_FROM_SYNCDB          = 0x3
	DEP_MOD_ANY              = 0x1
	DEP_MOD_EQ               = 0x2
	DEP_MOD_GE               = 0x3
	DEP_MOD_LE               = 0x4
	DEP_MOD_GT               = 0x5
	DEP_MOD_LT               = 0x6
	FILECONFLICT_TARGET      = 0x1
	FILECONFLICT_FILESYSTEM  = 0x2
	SIG_PACKAGE              = 0x1
	SIG_PACKAGE_OPTIONAL     = 0x2
	SIG_PACKAGE_MARGINAL_OK  = 0x4
	SIG_PACKAGE_UNKNOWN_OK   = 0x8
	SIG_DATABASE             = 0x400
	SIG_DATABASE_OPTIONAL    = 0x800
	SIG_DATABASE_MARGINAL_OK = 0x1000
	SIG_DATABASE_UNKNOWN_OK  = 0x2000
	SIG_USE_DEFAULT          = -0x80000000
	SIGSTATUS_VALID          = 0
	SIGSTATUS_KEY_EXPIRED    = 0x1
	SIGSTATUS_SIG_EXPIRED    = 0x2
	SIGSTATUS_KEY_UNKNOWN    = 0x3
	SIGSTATUS_KEY_DISABLED   = 0x4
	SIGSTATUS_INVALID        = 0x5
	SIGVALIDITY_FULL         = 0
	SIGVALIDITY_MARGINAL     = 0x1
	SIGVALIDITY_NEVER        = 0x2
	SIGVALIDITY_UNKNOWN      = 0x3
)

// Types

type list struct {
	Data *byte
	Prev *list
	Next *list
}
