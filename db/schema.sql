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
