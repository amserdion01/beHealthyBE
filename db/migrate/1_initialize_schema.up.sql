BEGIN;
create extension if not exists "uuid-ossp";
CREATE TABLE recipes (
	id uuid DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
	author VARCHAR(512) NOT NULL,
	author_id TEXT NOT NULL,
	ingredients text NOT NULL,
	image_url text NOT NULL,
	details TEXT NOT NULL,
	portions INT4 NOT NULL,
	preparation INT4 NOT NULL,
	cooking INT4 NOT NULL,
	tools text NOT NULL
);
commit;