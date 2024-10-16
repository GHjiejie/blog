-- target user role policy
INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES
        ('g', 'admin', 'user-admin-policy', '');

INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES
        ('g', 'user-admin-policy', 'admin-policy', '')

-- admin all api
INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
        -- admin user policy
        ('p', 'admin-policy', '/v1/*', '*');        