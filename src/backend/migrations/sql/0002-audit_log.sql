-- +migrate Up

CREATE TABLE public.audit_log (
	id bigserial PRIMARY KEY,
	log JSONB,
	created_at timestamp WITH time zone,
	updated_at timestamp WITH time zone
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION jsonb_values (jsonb)
	RETURNS text
	AS $func$
	SELECT
		string_agg(value, '|')
	FROM
		jsonb_each_text($1)
$func$
LANGUAGE sql
IMMUTABLE;
-- +migrate StatementEnd


-- +migrate Up notransaction
CREATE INDEX CONCURRENTLY IF NOT EXISTS audit_log_created_at_idx ON public.audit_log (created_at);
CREATE INDEX CONCURRENTLY audit_log_log_idx ON public.audit_log USING GIN (jsonb_values (log) gin_trgm_ops);


-- +migrate Down
DROP TABLE IF EXISTS public.audit_log;
DROP INDEX IF EXISTS audit_log_created_at_idx;
DROP INDEX IF EXISTS audit_log_log_idx;
DROP FUNCTION IF EXISTS jsonb_values (jsonb);