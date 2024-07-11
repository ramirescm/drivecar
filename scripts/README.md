docker run --name imersao-postgres --network imersao-golang -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres

docker run -d --hostname aprenda-golang --name imersao-rabbit -p 5672:5672 rabbitmq:3  --network imersao-golang

-- pwd diretorio atual.. tmp diretorio dentro do container onde a local atual sera mapeado
docker run -it --network imersao-golang --rm -v $(pwd):/tmp postgres bash

psql -h imersao-postgres -U postgres

# \q
# psql -h imersao-postgres -U postgres imersao < users.sql    // run script to create table
# psql -h imersao-postgres -U postgres imersao < folders.sql  // run script to create table
# psql -h imersao-postgres -U postgres imersao < files.sql    // run script to create table
# psql -h imersao-postgres -U postgres // conect cotainer
# \c imersao // conect db
# \dt // list tables