CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscriptions (
                               id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                               service_name TEXT NOT NULL,
                               price INTEGER NOT NULL,
                               user_id UUID NOT NULL,
                               start_date TIMESTAMP NOT NULL,
                               end_date TIMESTAMP,
                               created_at TIMESTAMP NOT NULL DEFAULT now(),
                               updated_at TIMESTAMP NOT NULL DEFAULT now()
);


