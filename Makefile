db:
	psql -d rsvp-dev -a -f setup/drop_public_schema.sql
	psql -d rsvp-dev -a -f setup/base_db.sql
