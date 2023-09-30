CREATE TABLE t_income_history
(
    id                       VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id VARCHAR(64),
    income_type_id           VARCHAR(64),
    payment_method_id        VARCHAR(64) NULL,
    payment_name             VARCHAR(64) NULL,
    income_amount            DECIMAL,
    description              VARCHAR(55),
    time_income_history      DATE,
    show_time_income_history VARCHAR(64),
    created_at               DECIMAL     NOT NULL,
    created_by               VARCHAR(64),
    updated_at               DECIMAL     NOT NULL,
    updated_by               VARCHAR(64),
    deleted_at               DECIMAL,
    deleted_by               VARCHAR(64),
    constraint fk_m_income_type
        foreign key (income_type_id)
            references m_income_type (id)
            on delete cascade
            on update cascade,
    constraint fk_m_profile
        foreign key (profile_id)
            references m_profiles (id)
            on delete cascade
            on update cascade,
    constraint fk_m_payment_method
        foreign key (payment_method_id)
            references m_payment_methods (id)
            on delete cascade
            on update cascade
);