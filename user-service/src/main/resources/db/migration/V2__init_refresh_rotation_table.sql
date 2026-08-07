
-- Table
CREATE TABLE refresh_tokens (
  id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  user_id          UUID  NOT NULL,
  token            varchar(255)  NOT NULL,
  is_revoked       BOOLEAN,
  expired_at       timestamptz  NOT NULL DEFAULT now(),
  created_at       timestamptz  NOT NULL DEFAULT now()
);

ALTER TABLE refresh_tokens
ADD CONSTRAINT fk_refresh_tokens_user
FOREIGN KEY (user_id)
REFERENCES accounts(id)
ON DELETE CASCADE;