### settings.go description
defines a set of structs and functions for 
working with database settings in a Go program.

### The DB_Settings struct 
stores the various settings for a database, 
such as its type, `hostname`, and `port`. 

### The Template struct 
stores default values for these settings. 

### The CheckDBSettings function
takes a pointer to a `DB_Settings` struct
and calls a series of "check" functions on it, 
which look at environment variables 
with names like `DATABASE_TYPE` and `DATABASE_HOST`
and set the corresponding fields in the `DB_Settings` struct 
to their values if they are defined, 
or to default values if they are not. 

### The Set functions
| Function          | Field in `DB_Settings` Struct |
|-------------------|-------------------------------|
| `SetDatabaseType` | `Type`                        |
| `SetProtocol`     | `Protocol`                    |
| `SetHostname`     | `Hostname`                    |
| `SetPort`         | `Port`                        |
| `SetName`         | `Name`                        |
| `SetUser`         | `User`                        |

allow you to set the corresponding fields 
in a `DB_Settings` struct directly. 

### The Connect function 
takes a DB_Settings struct and uses it to create a connection 
to a MySQL database using the `database/sql` package 
and the github.com/go-sql-driver/mysql driver.

