CREATE TABLE user_company_moderators (
    user_id INT NOT NULL,
    company_id INT NOT NULL,
    PRIMARY KEY (user_id, company_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

-- Добавление индексов для улучшения производительности
CREATE INDEX idx_user_moderators_id ON user_company_moderators(user_id);
CREATE INDEX idx_company_moderators_id ON user_company_moderators(company_id);