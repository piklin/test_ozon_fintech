CREATE TABLE IF NOT EXISTS short_urls
(
  id          SERIAL          NOT NULL UNIQUE,
  short_url   VARCHAR(255)    NOT NULL,
  full_url    VARCHAR(255)    NOT NULL
);
