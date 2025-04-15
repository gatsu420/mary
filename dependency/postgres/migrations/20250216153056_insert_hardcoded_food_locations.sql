-- migrate:up
insert into food_locations (name) values
    ('Unspecified'),
    ('Home'),
    ('Restaurant'),
    ('School'),
    ('OTW'),
    ('Relative''s Home'),
    ('Other');

-- migrate:down
truncate table food_locations;
