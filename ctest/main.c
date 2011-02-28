#include <alpm.h>
#include <alpm_list.h>
#include <stdio.h>

static void alpm_init();

int main() {

	alpm_list_t *i;

	if(alpm_initialize() != 0) {
		printf("could not int alpm\n");
		return 1;
	}

	if (alpm_option_set_root("/") != 0) {
		printf("failed setting root option\n");
		alpm_release();
		return 1;
	}

  if (alpm_option_set_dbpath("/var/lib/pacman") != 0) {
		printf("failed setting db path\n");
		alpm_release();
		return 1;
	}

	printf("root = %s\n", alpm_option_get_root());
	printf("dbpath = %s\n", alpm_option_get_dbpath());

	pmdb_t *db_local = alpm_db_register_local();
	alpm_list_t *searchlist = alpm_db_get_pkgcache(db_local);

	for(i = searchlist; i ; i = alpm_list_next(i)) {
		pmpkg_t *pkg = alpm_list_getdata(i);
		printf("local/%s %s\n", alpm_pkg_get_name(pkg), alpm_pkg_get_version(pkg));
	}

	if(alpm_release() != 0) {
		printf("could not release alpm\n");
	}
	return 0;
}

/* vim: set ts=2 sw=2 noet: */
