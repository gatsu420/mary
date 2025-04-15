-- migrate:up
insert into food_feeders (name) values
    ('Unspecified'),
    ('Kak Upi'),
    ('Aqis'),
    ('Other');

-- migrate:down
truncate table food_feeders;
