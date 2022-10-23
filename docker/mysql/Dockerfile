FROM mysql:5.7

COPY ./ /docker-entrypoint-initdb.d/

COPY setup.sql .

RUN /etc/init.d/setup.sql start
RUN mysql -u root -p${MYSQL_ROOT_PASSWORD} -e "CREATE DATABASE lba"
RUN mysql -u root -p${MYSQL_ROOT_PASSWORD} -D lba < /setuo.sql

EXPOSE 3306