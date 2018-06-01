FROM centos:7
MAINTAINER yogy rahmawan <yogy.frestarahmawan@gmail.com>

ENV LC_ALL="C"

RUN groupadd -r mysql --gid=997 \
    && useradd -r -g mysql -d /var/lib/mysql --uid=997 mysql

# install gosu
ENV GOSU_VERSION 1.10
RUN set -x \
    && yum -y install epel-release \
    && yum -y install wget dpkg \
    && dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
    && wget -O /usr/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch" \
    && wget -O /tmp/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc" \
    && export GNUPGHOME="$(mktemp -d)" \
    && gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
    && gpg --batch --verify /tmp/gosu.asc /usr/bin/gosu \
    && rm -r "$GNUPGHOME" /tmp/gosu.asc \
    && chmod +x /usr/bin/gosu \
    && gosu nobody true

RUN yum install -y \
    rsync \
    nmap \
    lsof \
    perl-DBI \
    nc \
    boost-program-options \
    iproute \
    iptables\
    libaio \
    libmnl \
    libnetfilter_conntrack \
    libnfnetlink \
    make \
    openssl \
    which

RUN \
	{ \
		echo "# MariaDB 10.1 CentOS repository list - created 2016-10-28 06:40 UTC"; \
		echo "# http://downloads.mariadb.org/mariadb/repositories/"; \
		echo "[mariadb]"; \
		echo "name = MariaDB"; \
		echo "baseurl = http://yum.mariadb.org/10.1/centos7-amd64"; \
		echo "gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB"; \
		echo "gpgcheck=1"; \
	} > /etc/yum.repos.d/mariadb-10.1.repo && \
	yum clean all

RUN yum -y -v update && yum -y -v install \
	MariaDB-server \
    MariaDB-client \
    MariaDB-compat \
    galera \
    socat \
    jemalloc

ADD ./docker-entrypoint.sh /docker-entrypoint.sh

EXPOSE 3306 4444 4567 4568

VOLUME /var/lib/mysql
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["mysqld"]

