FROM postgres:10

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
		krb5-admin-server \
		krb5-kdc \
		less \
		vim \
	&& rm -rf /var/lib/apt/lists/*

COPY krb5.conf /etc/krb5.conf

RUN kdb5_util create -s -P kpass \
	&& kadmin.local -q "addprinc -pw pgpass1 postgres/admin" \
	&& kadmin.local -q "addprinc -pw pgpass2 postgres/localhost@MY.EX" \
	&& kadmin.local -q "ktadd -k /etc/krb5.keytab postgres/localhost@MY.EX" \
	&& chmod +r /etc/krb5.keytab

COPY hba.sh /docker-entrypoint-initdb.d/hba.sh

COPY start.sh /usr/local/bin/
ENTRYPOINT ["start.sh"]
CMD ["postgres", "-c", "krb_server_keyfile='/etc/krb5.keytab'"]
