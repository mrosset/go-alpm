/**
 * Install reasons.
 * Why the package was installed.
 */
typedef enum _alpm_pkgreason_t {
	/** Explicitly requested by the user. */
	$ALPM_PKG_REASON_EXPLICIT = 0,
	/** Installed as a dependency for another package. */
	$ALPM_PKG_REASON_DEPEND = 1
} alpm_pkgreason_t;

typedef enum _alpm_pkgfrom_t {
	$PKG_FROM_FILE = 1,
	$PKG_FROM_LOCALDB,
	$PKG_FROM_SYNCDB
} alpm_pkgfrom_t;

/** Types of version constraints in dependency specs. */
typedef enum _alpm_depmod_t {
  /** No version constraint */
	$ALPM_DEP_MOD_ANY = 1,
  /** Test version equality (package=x.y.z) */
	$ALPM_DEP_MOD_EQ,
  /** Test for at least a version (package>=x.y.z) */
	$ALPM_DEP_MOD_GE,
  /** Test for at most a version (package<=x.y.z) */
	$ALPM_DEP_MOD_LE,
  /** Test for greater than some version (package>x.y.z) */
	$ALPM_DEP_MOD_GT,
  /** Test for less than some version (package<x.y.z) */
	$ALPM_DEP_MOD_LT
} alpm_depmod_t;

/**
 * File conflict type.
 * Whether the conflict results from a file existing on the filesystem, or with
 * another target in the transaction.
 */
typedef enum _alpm_fileconflicttype_t {
	$ALPM_FILECONFLICT_TARGET = 1,
	$ALPM_FILECONFLICT_FILESYSTEM
} alpm_fileconflicttype_t;

/**
 * PGP signature verification options
 */
typedef enum _alpm_siglevel_t {
	$ALPM_SIG_PACKAGE = (1 << 0),
	$ALPM_SIG_PACKAGE_OPTIONAL = (1 << 1),
	$ALPM_SIG_PACKAGE_MARGINAL_OK = (1 << 2),
	$ALPM_SIG_PACKAGE_UNKNOWN_OK = (1 << 3),

	$ALPM_SIG_DATABASE = (1 << 10),
	$ALPM_SIG_DATABASE_OPTIONAL = (1 << 11),
	$ALPM_SIG_DATABASE_MARGINAL_OK = (1 << 12),
	$ALPM_SIG_DATABASE_UNKNOWN_OK = (1 << 13),

	$ALPM_SIG_USE_DEFAULT = (1 << 31)
} alpm_siglevel_t;

/**
 * PGP signature verification status return codes
 */
typedef enum _alpm_sigstatus_t {
	$ALPM_SIGSTATUS_VALID,
	$ALPM_SIGSTATUS_KEY_EXPIRED,
	$ALPM_SIGSTATUS_SIG_EXPIRED,
	$ALPM_SIGSTATUS_KEY_UNKNOWN,
	$ALPM_SIGSTATUS_KEY_DISABLED,
	$ALPM_SIGSTATUS_INVALID
} alpm_sigstatus_t;

/**
 * PGP signature verification status return codes
 */
typedef enum _alpm_sigvalidity_t {
	$ALPM_SIGVALIDITY_FULL,
	$ALPM_SIGVALIDITY_MARGINAL,
	$ALPM_SIGVALIDITY_NEVER,
	$ALPM_SIGVALIDITY_UNKNOWN
} alpm_sigvalidity_t;


