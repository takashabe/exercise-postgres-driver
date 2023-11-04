CREATE ROLE app WITH LOGIN PASSWORD 'password';
GRANT pg_read_all_data,pg_write_all_data TO app;
