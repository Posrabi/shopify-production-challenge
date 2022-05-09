FROM postgres

# Setup Schema
COPY /migrations/*up.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
