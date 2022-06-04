CREATE TABLE user (
    id bigint(20) not null auto_increment primary key,
    account varchar(50) not null,
    password varchar(50) not null,
    status varchar(1) not null default '0'
);

CREATE UNIQUE INDEX idx_user_account ON user(account);