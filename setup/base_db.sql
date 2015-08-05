-- load the extension that allows us to generate UUIDs
CREATE EXTENSION "uuid-ossp";

-- create the tables
CREATE TABLE invitees (
    invitee_id uuid DEFAULT uuid_generate_v1mc() PRIMARY KEY
);
