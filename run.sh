#!/bin/bash
service postgresql start
psql -U postgres -c "ALTER USER postgres PASSWORD 'postgres'"
service postgresql restart
/app/main