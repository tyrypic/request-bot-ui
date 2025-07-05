-- 1) Тип статуса: создаём, только если его ещё нет
DO $$
BEGIN
  -- проверяем, есть ли уже тип с именем auth_status
  IF NOT EXISTS (
    SELECT 1
      FROM pg_type
     WHERE typname = 'auth_status'
  ) THEN
    CREATE TYPE auth_status AS ENUM (
      'pending',
      'approved',
      'rejected'
    );
  END IF;
END
$$;

-- 2) Таблица users
CREATE TABLE IF NOT EXISTS users (
  id           BIGSERIAL    PRIMARY KEY,
  telegram_id  BIGINT       NOT NULL UNIQUE,
  username     TEXT,
  first_name   TEXT,
  last_name    TEXT,
  real_name    TEXT,
  email        TEXT,
  age          INT          CHECK(age > 0),
  city         TEXT,
  is_admin     BOOLEAN      NOT NULL DEFAULT FALSE,
  created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
  approved_at  TIMESTAMP
);

-- 3) Таблица auth_requests
CREATE TABLE IF NOT EXISTS auth_requests (
  id           BIGSERIAL    PRIMARY KEY,
  user_id      BIGINT       NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  status       auth_status  NOT NULL DEFAULT 'pending',
  created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
  resolved_at  TIMESTAMP,
  resolved_by  BIGINT       REFERENCES users(id),
  UNIQUE (user_id)
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_users_telegram_id ON users(telegram_id);
CREATE INDEX IF NOT EXISTS idx_requests_user_id  ON auth_requests(user_id);
CREATE INDEX IF NOT EXISTS idx_requests_status   ON auth_requests(status);
