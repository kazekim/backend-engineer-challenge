#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    create table gr_git_repositories
    (
        id         varchar(20)              not null
            constraint gr_git_repository_pk
                primary key,
        created_at timestamp with time zone not null,
        updated_at timestamp with time zone not null,
        name       varchar(200)             not null,
        url        varchar(500)             not null,
        is_active  boolean default true     not null
    );

    create table gr_git_repository_scan_results
    (
        id            varchar(20)              not null
            constraint gr_git_repository_scan_results_pk
                primary key,
        created_at    timestamp with time zone not null,
        updated_at    timestamp with time zone not null,
        deleted_at    timestamp with time zone,
        repository_id varchar(20)              not null
            constraint gr_git_repository_scan_results_gr_git_repositories_id_fk
                references gr_git_repositories
                on delete cascade,
        status        varchar(20)              not null,
        findings      jsonb,
        queued_at     timestamp with time zone not null,
        scanning_at   timestamp with time zone,
        finished_at   timestamp with time zone
    );

    create index gr_git_repository_scan_results_repositories_id_index
        on gr_git_repository_scan_results (repository_id);
  COMMIT;
EOSQL