-- +goose Up
-- +goose StatementBegin
CREATE TABLE bench (
    id serial primary key,
    title text,
    body text default '',
    tt timestamp,
    count int,
    jj jsonb NOT NULL default '{"name":"firstname"}'::jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bench;
-- +goose StatementEnd
