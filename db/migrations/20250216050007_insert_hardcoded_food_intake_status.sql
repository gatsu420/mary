-- migrate:up
insert into food_intake_status (name) values
    ('Unspecified'),
    ('Not Eaten'),
    ('Barely'),
    ('Half'),
    ('Mostly'),
    ('Full');

-- migrate:down
truncate table food_intake_status;
