#!/usr/bin/env bash
docker run -d --rm --name grdb -e POSTGRES_PASSWORD=asdfghjkl -p 5433:5432 postgres