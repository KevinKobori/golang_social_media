## Steps to run this project on macOS:

Install Golang: https://go.dev/doc/install

Install MySQL Comunity Server: https://dev.mysql.com/downloads/mysql/

### On first terminal from path .../golang_social_media/api/:
Make an alias to access mysql bin path comands on terminal:

        alias mysql=/usr/local/mysql/bin/mysql

Access MySQL using root user:

        mysql --user=root -p

Put your root password access here:

        <your_mysql_root_password_here>

mysql> Change the 'root' password to '1234':

        ALTER USER 'root'@'localhost' IDENTIFIED BY '1234';

mysql> Copy and paste the script on MySQL:

        <copy_and_paste_the_script_on_this_file:/golang_social_media/api/sql/sql.sql>

mysql> Exit MySQL:
       
        exit;

Create a new file named .env:

        touch .env

Copy and paste the script to .env archive and save:

        DB_USUARIO=root
        DB_SENHA=1234
        DB_NOME=devbook
        API_PORT=5000
        SECRET_KEY=oIOryDxh2AHNwLzKw064B29lroBcP/WXQ2Icr94B86g3ri2NxJVzda4PTaQRpmDYLvPf2X9gwvCSZ6IXHMgwNA==

Run Backend:

        - $ go run main.go

### On second terminal from path .../golang_social_media/webapp/:

Run Frontend:

        - $ go run main.go


### Open on Chrome http://localhost:3000/

### OBS: DOES NOT WORK ON SAFARI.

