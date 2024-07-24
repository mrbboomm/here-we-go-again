set -Eeuo pipefail

ORACLE_CDB=${1:-}

db_status=$(sqlplus -S sys/oracle@localhost:1521/xe as sysdba @custom-healthcheck.sql)

echo db_status ::: $db_status
if [ "${db_status}" == "OPEN" ]; then
   exit 0;
else
   exit 1;
fi;