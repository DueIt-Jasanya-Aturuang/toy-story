CREATE TABLE m_payment_methods
(
    id          VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id  VARCHAR(64),
    name        VARCHAR(128),
    description TEXT,
    image       VARCHAR(255) default 'default-icon.png',
    created_at  DECIMAL     NOT NULL,
    created_by  VARCHAR(64),
    updated_at  DECIMAL     NOT NULL,
    updated_by  VARCHAR(64),
    deleted_at  DECIMAL,
    deleted_by  VARCHAR(64),
    CONSTRAINT fk_m_profile
        FOREIGN KEY (profile_id)
            REFERENCES m_profiles(id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);