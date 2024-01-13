package model

/*
create table authority
(
    id          int          not null,
    master_id   varchar(100) null,
    overseer_id varchar(100) null,
    worker_id   varchar(100) null,
    project_id  int          null,
    constraint authority_id_uindex
        unique (id)
);

alter table authority
    add primary key (id);

create table collection
(
    id          int auto_increment,
    item_id     bigint      not null,
    description varchar(10) null,
    icon        int         null,
    constraint collection_id_uindex
        unique (id)
);

alter table collection
    add primary key (id);

create table items
(
    id              bigint auto_increment,
    item_name       varchar(20)  not null,
    description     varchar(100) null,
    from_project_id int          null,
    deadline        timestamp    null,
    if_important    tinyint(1)   not null,
    if_done         tinyint(1)   null,
    if_myDay        tinyint(1)   null,
    created_time    timestamp    null,
    if_node         tinyint(1)   null,
    if_checkPoint   tinyint(1)   null,
    collection_id   int          null,
    constraint item_itemId_uindex
        unique (id)
);

alter table items
    add primary key (id);

create table my_days
(
    id      int null,
    user_id int not null,
    item_id int not null
);

create table project
(
    id               int auto_increment,
    project_name     varchar(20) not null,
    project_items_id int         null,
    description      varchar(50) null,
    authority_id     int         not null,
    created_time     timestamp   not null,
    end_time         timestamp   null,
    password         varchar(15) not null,
    constraint project_password_uindex
        unique (password),
    constraint projection_projectionID_uindex
        unique (id)
);

alter table project
    add primary key (id);

create table project_item_conv
(
    project_id int null,
    item_id    int null
);

create table project_user_convs
(
    projects_id int not null,
    users_id    int not null
);

create table users
(
    id           int auto_increment,
    username     varchar(20)  not null,
    password     varchar(20)  not null,
    apartment_id int          not null,
    projects_id  varchar(50)  null,
    link         varchar(20)  null,
    bio          varchar(100) null,
    avatar       varchar(200) null,
    created_at   datetime     null,
    updated_at   datetime     null,
    deleted_at   datetime     null,
    constraint user_apartment_uindex
        unique (apartment_id),
    constraint user_userID_uindex
        unique (id)
);

create index idx_users_deleted_at
    on users (deleted_at);

alter table users
    add primary key (id);


*/

type Users struct {
	Id          int    `json:"id" sql:"id"`
	Username    string `json:"username" sql:"username"`
	Password    string `json:"password" sql:"password"`
	ApartmentId int    `json:"apartment_id" sql:"apartment_id"`
	ProjectsId  string `json:"projects_id" sql:"projects_id"`
	Link        string `json:"link" sql:"link"`
	Bio         string `json:"bio" sql:"bio"`
	Avatar      string `json:"avatar" sql:"avatar"`
}

func (user *Users) CheckPassword(password string) bool {
	return password == user.Password
}
