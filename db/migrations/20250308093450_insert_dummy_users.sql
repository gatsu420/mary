-- migrate:up
insert into users (username, type_id) values
    ('coco_dummy', 1),
    ('cayo_dummy', 1),
    ('ceyi_dummy', 2);

-- migrate:down
delete from users
where username like '%_dummy'
