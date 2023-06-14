CREATE TABLE IF NOT EXISTS url (
    id bigserial PRIMARY KEY,
    short_url text NOT NULL,
    long_url text NOT NULL
);
