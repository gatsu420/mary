-- migrate:up
insert into users (name, type_id) values
    ('coco_DUMMY', 1),
    ('cayo_DUMMY', 1),
    ('ceyi_DUMMY', 2);

-- migrate:down
delete from users
where name like '%_DUMMY'
