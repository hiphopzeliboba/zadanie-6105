-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table tenders
(
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at       TIMESTAMPTZ      DEFAULT NOW(),
    updated_at       TIMESTAMPTZ      DEFAULT NOW(),
    name             text,
    description      text,
    service_type     varchar(36),
    status           varchar(16),
    organization_id  varchar(36),
    creator_username text,
    version          integer

);


-- +goose Down
DROP TABLE tenders;
