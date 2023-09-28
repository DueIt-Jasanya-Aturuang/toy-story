CREATE TABLE m_default_payment_method
(
    id          VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    name        VARCHAR(128),
    description TEXT,
    image       VARCHAR(255) default 'default-icon.png',
    created_at  DECIMAL     NOT NULL,
    created_by  VARCHAR(64),
    updated_at  DECIMAL     NOT NULL,
    updated_by  VARCHAR(64),
    deleted_at  DECIMAL,
    deleted_by  VARCHAR(64)
);

INSERT INTO m_default_payment_method (id, name, description, image, created_at, created_by, updated_at)
VALUES ('0bb8498a-fbcb-49e3-86f0-e8f82bc6c722', 'bca', 'bca online', '/files/payment-images/public/file.png', 1684768516, NULL, 1684768516)
--
-- 0bb8498a-fbcb-49e3-86f0-e8f82bc6c722
--
-- 69f72ea2-1ec5-49a1-ae1b-62cc92a2d949
--
-- a7c94c23-9b26-4c62-ac94-44e062678c1c
--
-- 06c794b2-fe27-48d2-bb8d-1c21d48e4612
--
-- 741e82cd-3229-40ab-8f0c-000f553e15be
--
-- 07a5c7b0-d1bb-4d17-8ab4-8504c32b5e2a
--
-- b9391828-3c37-4ed1-b6a4-78e609d77731
--
-- b184906c-5912-47cc-b4e1-35ba375a3050
--
-- ca3d2b25-20b4-4b19-896b-afd940b824db
--
-- 48adc390-3d30-44a1-a1e0-57a34c50505a