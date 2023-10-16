CREATE TABLE IF NOT EXISTS users ( 
	id varchar(255) NOT NULL, 
	username varchar(255) not null UNIQUE,
	password varchar(255) not null,
	user_state varchar(255) not null,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	finish_plan_at TIMESTAMPTZ,
	primary key (id)
);