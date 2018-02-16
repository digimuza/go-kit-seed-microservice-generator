CREATE EXTENSION dblink;

CREATE OR REPLACE FUNCTION create_db(dbname text)
  RETURNS void AS
$func$
BEGIN

IF EXISTS (SELECT 1 FROM pg_database WHERE datname = dbname) THEN
   RAISE NOTICE 'Database already exists'; 
ELSE
   PERFORM dblink_exec('dbname=' || current_database(),
      'CREATE DATABASE ' || quote_ident(dbname));
   PERFORM dblink_exec('dbname=' || dbname,
      'CREATE EXTENSION "uuid-ossp"');
END IF;

END
$func$ LANGUAGE plpgsql;

SELECT create_db('<%= appName %>-db')
