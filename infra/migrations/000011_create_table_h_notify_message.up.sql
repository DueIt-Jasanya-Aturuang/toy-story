CREATE TABLE h_notify_message
(
    id         VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    name       VARCHAR(64),
    status     VARCHAR(20),
    title      VARCHAR(64),
    icon       VARCHAR(64),
    message    TEXT,
    created_at DECIMAL     NOT NULL,
    created_by VARCHAR(64),
    updated_at DECIMAL     NOT NULL,
    updated_by VARCHAR(64),
    deleted_at DECIMAL,
    deleted_by VARCHAR(64)
);

INSERT INTO h_notify_message
(id, name, status, title, icon, message, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES ('01HBXAC9X9HWSFKHNWN1HDDPJR', 'DAILY_NOTIFY', 'on', 'notifikasi harian', '/files/icon-images/public/notif.svg',
        'hallo %s ini notif harian kamu', 1684768516, 'admin', 1684768516, NULL,
        NULL, NULL);

INSERT INTO h_notify_message
(id, name, status, title, icon, message, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES ('01HBXAC9X9HWSFKHNWN5GTQ11G', 'MONTHLY_PERIOD', 'on', 'notifikasi bulanan',
        '/files/icon-images/public/notif.svg', 'hallo %s ini notif bulanan kamu', 1684768516, 'admin', 1684768516, NULL,
        NULL, NULL);
