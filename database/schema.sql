CREATE TABLE categories (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  slug text NOT NULL,
  author integer,
  isActive boolean
)
