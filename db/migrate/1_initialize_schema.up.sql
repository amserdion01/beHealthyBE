BEGIN;
create extension if not exists "uuid-ossp";
CREATE TABLE recipes (
	id uuid DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
	author VARCHAR(512) NOT NULL,
	author_id TEXT NOT NULL,
	ingredients JSON NOT NULL,
	details TEXT NOT NULL,
	portions INT4 NOT NULL,
	preparation TIME NOT NULL,
	cooking TIME NOT NULL,
	tools text[] NOT NULL
);
commit;