FROM postgres:17.5-alpine

COPY db_scripts /docker-entrypoint-initdb.d

RUN chmod a+r /docker-entrypoint-initdb.d/*