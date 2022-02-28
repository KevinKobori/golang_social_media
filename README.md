# golang_social_media

## Steps to run this project on macOS:

### Install Golang: https://go.dev/doc/install
### Install MySQL Comunity Server: https://dev.mysql.com/downloads/mysql/

### On first terminal from path .../golang_social_media/:
Make an alias to access mysql bin path comands on terminal:

        - $ alias mysql=/usr/local/mysql/bin/mysql

Access MySQL using root user:

        - $ mysql --user=root -p

Put your root password access here:

        - $ <your_mysql_root_password_here>

Create another user with name 'golang' and password '1234': 

        - mysql> $ CREATE USER 'golang' IDENTIFIED BY '1234';

Grant root privileges to 'golang' user:

        - mysql> $ GRANT ALL PRIVILEGES ON *.* TO 'golang' WITH GRANT OPTION;

Exit MySQL:

        - mysql> $ exit;

Access MySQL using golang user:

        - $ mysql --user=golang -p

Put golang password access here:

        - $ 1234

Create a database with name 'devbook':

        - mysql> $ CREATE DATABASE devbook;

Set MySQL to use this same database:

        - mysql> $ USE devbook;

Copy and paste the script on MySQL:

        - mysql> $ <copy_and_paste_the_script_on_this_file:/golang_social_media/api/sql/sql.sql>

Exit MySQL:
       
        - mysql> $ exit;

Access api folder:

        - $ cd api 

Create a new file named .env:

        - $ touch .env

Copy and paste the script to .env archive and save:

        DB_USUARIO=golang
        DB_SENHA=1234
        DB_NOME=devbook
        API_PORT=5000
        SECRET_KEY=oIOryDxh2AHNwLzKw064B29lroBcP/WXQ2Icr94B86g3ri2NxJVzda4PTaQRpmDYLvPf2X9gwvCSZ6IXHMgwNA==

Run api:

        - $ go run main.go

### On second terminal from path .../golang_social_media/:

Access webapp folder:

        - $ cd webapp

Run webapp:

        - $ go run main.go


### Open on Chrome http://localhost:3000/

### OBS: DOES NOT WORK ON SAFARI.

