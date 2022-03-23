-- +migrate Up
CREATE TABLE public.s_users
(
    id                                          bigserial PRIMARY KEY,
    username                                    VARCHAR(30) NOT NULL check (username = lower(username)),
    password_hash                               VARCHAR(64) NOT NULL,
    forgot_password_hash                        VARCHAR(64) DEFAULT NULL::VARCHAR,
    forgot_password_hash_created_at             timestamp with time zone DEFAULT NULL,
    email                                       VARCHAR(64) NOT NULL check (email = lower(email)),
    email_verify_hash                           VARCHAR(64) DEFAULT NULL::VARCHAR,
    email_verify_created_at                     timestamp with time zone DEFAULT NULL,
    first_name                                  VARCHAR(30) DEFAULT NULL::VARCHAR,
    last_name                                   VARCHAR(30) DEFAULT NULL::VARCHAR,
    mobile                                      VARCHAR(30) DEFAULT NULL::VARCHAR,
    mobile_verify_code                          VARCHAR(6)  DEFAULT NULL::VARCHAR,
    mobile_verify_created_at                    timestamp with time zone DEFAULT NULL,
    mobile_verified                             boolean DEFAULT FALSE,
    recovery_questions                          JSONB,
    recovery_questions_set                      boolean DEFAULT FALSE,
    avatar_url                                  VARCHAR(2000) DEFAULT NULL::VARCHAR,
    roles                                       VARCHAR(30)[],
    active                                      boolean DEFAULT FALSE,
    social_login                                boolean DEFAULT FALSE,
    language                                    varchar(2),
    address                                     JSONB,
    country                                     VARCHAR(2)  DEFAULT NULL::VARCHAR,
    customer_id                                  VARCHAR(30)  DEFAULT NULL::VARCHAR,
    authenticator_enabled                       boolean DEFAULT FALSE,
    authenticator_secret                        VARCHAR(64)  DEFAULT NULL::VARCHAR,
    last_login_at                               timestamp with time zone,
    created_at                                  timestamp with time zone,
    updated_at                                  timestamp with time zone,
    deleted_at                                  timestamp with time zone
);

-- +migrate Up notransaction
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS s_users_unique_email_idx ON public.s_users (email);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS s_users_unique_username_idx ON public.s_users (username);
CREATE INDEX CONCURRENTLY IF NOT EXISTS s_users_first_name_idx ON public.s_users (first_name);

CREATE INDEX CONCURRENTLY users_trgm_at_idx ON public.s_users
    USING GIST ((username || ' ' || email || ' ' || coalesce (first_name,'') || ' ' || coalesce (last_name,'')) gist_trgm_ops);


-- +migrate Down
DROP INDEX IF EXISTS s_users_unique_email_idx;
DROP INDEX IF EXISTS s_users_unique_username_idx;
DROP INDEX IF EXISTS s_users_first_name_idx;
DROP INDEX IF EXISTS users_trgm_at_idx;
DROP TABLE IF EXISTS public.s_users;
