
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
drop table if exists `artist_private_info_mast`;
create table `medici_master`.`artist_private_info_mast` (
  `artist_private_id` bigint not null,
  `artist_id` bigint not null,
  `phone_number` int,
  `mail_address` varchar(256),
  `postal_code` varchar(50),
  `address` varchar(300),
  `work_place_postal_code` varchar(50),
  `work_place_address` varchar(300),
  `formal_name` varchar(200),
  primary key(`artist_private_id`, `artist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `art_work_mast`;
create table `medici_master`.`art_work_mast` (
  work_id bigint not null primary key,
  artist_id bigint not null,
  start_date date not null,
  end_date date not null,
  work_title_ja varchar(200),
  work_title_en varchar(200),
  work_description_ja blob,
  work_description_en blob,
  work_image_path varchar(200),
  create_year int,
  size varchar(100),
  size_x_position int,
  size_y_position int,
  size_z_position int,
  material varchar(200)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `art_related_item_mast`;
create table `medici_master`.`art_related_item_mast` (
  item_id bigint not null primary key,
  artist_id bigint not null,
  start_date date not null,
  end_date date not null,
  item_name_ja varchar(200),
  item_name_en varchar(200),
  item_description_ja blob,
  item_description_en blob,
  item_image_path varchar(200),
  size varchar(100),
  size_x_position int,
  size_y_position int,
  size_z_position int,
  material varchar(200)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `artist_event_mast`;
create table `medici_master`.`artist_event_mast` (
  `event_id` bigint(20) not null,
  `artist_id` bigint(20) not null,
  `start_date` date not null,
  `end_date` date not null,
  `event_name_ja` varchar(200) default null,
  `event_name_en` varchar(200) default null,
  `event_description_ja` blob,
  `event_description_en` blob,
  `event_image_path` varchar(200) default null,
  `event_start_at` datetime null default null,
  `event_end_at` datetime null default null,
  `event_phone` int(11) default null,
  `event_mail_address` varchar(256) default null,
  `event_url` blob,
  primary key (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `art_media_mast`;
create table `medici_master`.`art_media_mast` (
  media_id bigint not null primary key,
  artist_id bigint not null,
  start_date date not null,
  end_date date not null,
  media1_type int,
  media_name_ja varchar(200),
  media_name_ varchar(200),
  media_description_ja blob,
  media_description_en blob,
  media_url blob
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table if exists `artist_private_info_mast`;
drop table if exists `art_work_mast`;
drop table if exists `art_related_item_mast`;
drop table if exists `artist_event_mast`;
drop table if exists `art_media_mast`;
