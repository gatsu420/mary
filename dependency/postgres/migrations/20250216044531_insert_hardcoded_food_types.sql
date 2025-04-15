-- migrate:up
insert into food_types (name) values
    ('Unspecified'),
    ('Breakfast'),
    ('Lunch'),
    ('Dinner'),
    ('Snack'),
    ('Fruit'),
    ('Breastmilk'),
    ('Drink');

-- migrate:down
truncate table food_types;
