CREATE SCHEMA IF NOT EXISTS identity;

CREATE TABLE IF NOT EXISTS identity.users (
	id UUID PRIMARY KEY,
	email VARCHAR(255) NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	full_name VARCHAR(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	weight FLOAT DEFAULT NULL,
	height FLOAT DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_identity_users_created_at ON identity.users(id, created_at);
