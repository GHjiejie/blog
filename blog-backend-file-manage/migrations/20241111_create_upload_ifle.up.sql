
CREATE TABLE IF NOT EXISTS files (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    file_type VARCHAR(255) NOT NULL, 
    tag VARCHAR(50) NOT NULL, 
    file_content LONGBLOB NULL, 
    file_path VARCHAR(500) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    uploader_id BIGINT NOT NULL,
    INDEX idx_tag (tag(50)),
    INDEX idx_file_type (file_type(255))
);


