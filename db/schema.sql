create table teamteam_cursors (
   id varchar(255) not null,
   `cursor` bigint not null,
   updated_at datetime(3) not null,

   primary key (id)
);

create table teamteam_events (
  id bigint not null auto_increment,
  foreign_id bigint not null,
  timestamp datetime(3) not null,
  type int not null,

  primary key (id)
);

create table player_match (
  id bigint not null auto_increment,
  status int not null,
  created_at datetime not null,
  updated_at datetime null,
  player_name varchar(255),
  round_num int,
  rank int,
  my_part int,
  player_part int,

  primary key (id)
)
