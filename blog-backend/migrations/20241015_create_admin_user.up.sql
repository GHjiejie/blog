-- admin user, admin/Console1@
INSERT INTO users(username, password, role) VALUES('admin', '12345', 0) ON DUPLICATE KEY UPDATE password = '12345';