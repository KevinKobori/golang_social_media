# golang_social_media

Steps to run this project on macOS:
  - Install Golang: https://go.dev/doc/install
  - Install MySQL Comunity Server: https://dev.mysql.com/downloads/mysql/

  - On first terminal from .../golang_social_media/ folder:
    - $ alias mysql=/usr/local/mysql/bin/mysql
    - $ mysql --user=root -p
    - $ <your_mysql_root_password_here>
    - $ mysql> CREATE USER 'golang' IDENTIFIED BY '1234';
    - $ mysql> GRANT ALL PRIVILEGES ON *.* TO 'golang' WITH GRANT OPTION;
    - $ mysql> exit;
    - $ mysql --user=golang -p
    - $ 1234 <put_golang_user_password_here>
    - $ mysql> <copy_and_paste_the_script_on_this_file:/golang_social_media/api/sql/sql.sql>
    - $ mysql> <copy_and_paste_the_script_on_this_file:/golang_social_media/api/sql/dados.sql>
    - $ mysql> show databases; <optional_just_to_check_if_it_works>
    - $ mysql> use devbook; <optional_just_to_check_if_it_works>
    - $ mysql> select * from usuarios; <optional_just_to_check_if_it_works>
    - $ mysql> exit;

    - $ cd api 
    - $ nano .env
    - $ <on_nano_editor_copy_and_paste_the_script_on_this_file:/golang_social_media/api/env.txt>
    - $ go run main.go

  - On second terminal from .../golang_social_media/ folder:
    - $ cd webapp
    - $ go run main.go