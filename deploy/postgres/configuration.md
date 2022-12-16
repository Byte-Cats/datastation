### Using the `docker-compose.yml`  file

To use this docker-compose.yml file, you will need to create a 

`.env` file that specifies values for the following variables:

  - DB_USERNAME
  - DB_PASSWORD
  - DB_PORT

### For example,

your `.env` file might look like this:
```shell
DB_USERNAME=myusername
DB_PASSWORD=mypassword
DB_PORT=5432
```

### With this setup,

the `PostgreSQL` service will use the values specified in the `.env` file for the following settings 

in the docker-compose.yml file:

  - `POSTGRES_USER`
  - `POSTGRES_PASSWORD`
  - `POSTGRES_DB`

### Conclusion 

This allows you to easily customize 
the `PostgreSQL`
service without having to modify the 
`docker-compose.yml`file directly. 