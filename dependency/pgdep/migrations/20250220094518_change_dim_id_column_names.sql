-- migrate:up
alter table food
rename column type to type_id;

alter table food
rename column intake_status to intake_status_id;

alter table food
rename column feeder to feeder_id;

alter table food
rename column location to location_id;

-- migrate:down
alter table food
rename column type_id to type;

alter table food
rename column intake_status_id to intake_status;

alter table food
rename column feeder_id to feeder;

alter table food
rename column location_id to location;
