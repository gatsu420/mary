-- migrate:up
alter sequence food_locations_id_seq minvalue 0 restart with 0;

-- migrate:down
alter sequence food_locations_id_seq minvalue 1 restart with 1;
