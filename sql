sudo dnf install @mysql
sudo systemctl enable --now mysqld
systemctl status mysqld.service
sudo mysql_secure_installation #设置root密码
mysql -uroot -pxxx
systemctl status mysqld
systemctl restart mysqld
systemctl status mysqld
mysql -uroot -pxxx
mysql -hxx.xx.xx.xx -P3306 -uroot -p

create table client_grid
(
    client_id bigint default 0 not null
        primary key,
    grid_id   bigint           not null,
    constraint client_grid_client_id_uindex
        unique (client_id)
);

create table client_info
(
    client_id    bigint      default 0      not null
        primary key,
    company_name varchar(50)                null,
    client_name  varchar(50) default '默认用户' not null,
    longitude    decimal(20, 15)            null,
    latitude     decimal(20, 15)            null,
    constraint client_info_client_id_uindex
        unique (client_id)
);

create table grid_info
(
    grid_id    bigint default 0 not null
        primary key,
    grid_name  varchar(50)      not null,
    applied    bigint default 0 null,
    capacity   bigint default 0 null,
    population bigint default 0 null,
    constraint grid_info_grid_id_uindex
        unique (grid_id)
);
