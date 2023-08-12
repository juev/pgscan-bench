-- +goose Up
-- +goose StatementBegin
INSERT INTO bench(
    title, 
    body, 
    tt, 
    count, 
    jj)
SELECT
    md5(random()::text), 
    md5(random()::text), 
    timestamp '1900-01-01 00:00:00' + random() * (timestamp '2023-01-01 00:00:00' - timestamp '1900-01-01 00:00:00'),
    (random() * 70 + 10)::integer,
    '{"name":{"firstname":"first","lastname":"first2"}}'::jsonb
FROM generate_series(1, 100000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
