CREATE TABLE IF NOT EXISTS todos (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT false NOT NULL,
    content VARCHAR,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    tenant_id VARCHAR(36) NOT NULL
)