# 다른 포트 허용하기

<ubuntu>
/etc/mysql/mariadb.conf.d/50-server.cnf


[mysqld]
bind-address = 0.0.0.0
port = 3307

