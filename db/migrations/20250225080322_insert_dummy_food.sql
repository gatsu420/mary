-- migrate:up
insert into food (name, type_id, intake_status_id, feeder_id, location_id, remarks) values
    ('sate_DUMMY', 1, 1, 1, 1, 'DUMMY_FOOD_TESTING'),
    ('ayam goreng_DUMMY', 1, 2, 2, 2, 'DUMMY_FOOD_TESTING'),
    ('kacang rebus_DUMMY', 2, 2, 2, 2, 'DUMMY_FOOD_TESTING');

-- migrate:down
delete from food
where name like '%_DUMMY'
and remarks = 'DUMMY_FOOD_TESTING';
