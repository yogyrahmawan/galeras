#!/usr/bin/env bash

[ "$DEBUG" == 'true' ] && set -x

set -eo pipefail

if [ "$1" == 'mysqld' ]; then 
	echo 'initiating galera config'
	export MYSQL_INITDB_SKIP_TZINFO="yes"
	export MYSQL_ALLOW_EMPTY_PASSWORD="yes"

	# secure install 
	echo "exec mysql_secure_installation"
    gosu root /etc/init.d/mysql start
	> /.temp_pass
    echo "" >> /.temp_pass
    echo "y" >> /.temp_pass
    echo "${MYSQL_ROOT_PASSWORD}" >> /.temp_pass
    echo "${MYSQL_ROOT_PASSWORD}" >> /.temp_pass
    echo "y" >> /.temp_pass
    echo "y" >> /.temp_pass
    echo "y" >> /.temp_pass
    echo "y" >> /.temp_pass
    mysql_secure_installation < /.temp_pass
	gosu root /etc/init.d/mysql stop
	# end secure install

	echo "[mysqld]" >> /etc/my.cnf.d/galera.cnf
	echo "innodb_lock_schedule_algorithm=fcfs" >> /etc/my.cnf.d/galera.cnf
	echo "binlog_format=${BINLOG_FORMAT}" >> /etc/my.cnf.d/galera.cnf
    echo "default_storage_engine=${DEFAULT_STORAGE_ENGINE}" >> /etc/my.cnf.d/galera.cnf
    echo "innodb_autoinc_lock_mode=${INNODB_AUTOINC_LOCK_MODE}" >> /etc/my.cnf.d/galera.cnf
    echo "bind-address=${BIND_ADDRESS}" >> /etc/my.cnf.d/galera.cnf
	echo "" >> /etc/my.cnf.d/galera.cnf
    echo "[galera]" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_on=${WSREP_ON}" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_provider=${WSREP_PROVIDER}" >> /etc/my.cnf.d/galera.cnf
    if [[ ! -z "${WSREP_PROVIDER_OPTIONS// }" ]]; then
        echo "wsrep_provider_options"=${WSREP_PROVIDER_OPTIONS} >> /etc/my.cnf.d/galera.cnf
    fi
    echo "wsrep_cluster_address=${WSREP_CLUSTER_ADDRESS}" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_cluster_name=${WSREP_CLUSTER_NAME}" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_node_address=${WSREP_NODE_ADDRESS}" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_node_name=${WSREP_NODE_NAME}" >> /etc/my.cnf.d/galera.cnf
    echo "wsrep_sst_method=${WSREP_SST_METHOD}" >> /etc/my.cnf.d/galera.cnf

	if [ "$2" == "init" ]; then
		gosu root /etc/init.d/mysql start --wsrep-new-cluster
    elif [ "$2" == "join" ]; then 
        gosu root /etc/init.d/mysql start
    fi

	gosu root tail -f /dev/null
else
	exec "$@";
fi
