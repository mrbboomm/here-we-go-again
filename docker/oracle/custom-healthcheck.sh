set -Eeuo pipefail

ORACLE_CDB=${1:-}

      # ORACLE_PASSWORD: mysecretpassword
      # ORACLE_DATABASE: godev
      # TARGET_PDB: godev
      # APP_USER: godev_user
      # APP_USER_PASSWORD: godev_pass
db_status=$(sqlplus -S sys/mysecretpassword@localhost:1521/godev as sysdba @custom-healthcheck.sql)

echo db_status ::: $db_status
if [ "${db_status}" == "OPEN" ]; then
   exit 0;
else
   exit 1;
fi;