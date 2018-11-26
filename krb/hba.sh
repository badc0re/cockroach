#!/bin/bash
set -e

kinit -k -t /etc/krb5.keytab postgres/localhost@MY.EX
klist

cat << EOF > /var/lib/postgresql/data/pg_hba.conf
local all all trust
host all all 0.0.0.0/0 gss include_realm=0 krb_realm=MY.EX
EOF

"${psql[@]}" <<- 'EOSQL'
    SELECT pg_reload_conf();
EOSQL
