-- +migrate Up
ALTER TABLE public.s_users ADD COLUMN notifications JSONB;

-- +migrate Down
ALTER TABLE public.s_users DROP COLUMN notifications;

