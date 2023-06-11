-- Populate Assets
INSERT INTO ak.assets (id, user_id, asset_id, asset_type, name, description)
SELECT
    md5(random()::text || clock_timestamp()::text)::uuid,
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.id ELSE u2.id END,
    c.id,
    'chart',
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END || '-chart-' || row_number() OVER (PARTITION BY 'chart' ORDER BY c.id),
    'Faved chart by ' || CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END
FROM ak.charts c
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'johndoe') u1
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'emilysmith') u2
LIMIT 6;

INSERT INTO ak.assets (id, user_id, asset_id, asset_type, name, description)
SELECT
    md5(random()::text || clock_timestamp()::text)::uuid,
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.id ELSE u2.id END,
    i.id,
    'insight',
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END || '-insight-' || row_number() OVER (PARTITION BY 'insight' ORDER BY i.id),
    'Faved insight by ' || CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END
FROM ak.insights i
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'johndoe') u1
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'emilysmith') u2
LIMIT 4;

INSERT INTO ak.assets (id, user_id, asset_id, asset_type, name, description)
SELECT
    md5(random()::text || clock_timestamp()::text)::uuid,
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.id ELSE u2.id END,
    a.id,
    'audience',
    CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END || '-audience-' || row_number() OVER (PARTITION BY 'audience' ORDER BY a.id),
    'Faved audience by ' || CASE WHEN (row_number() OVER ()) % 2 = 1 THEN u1.username ELSE u2.username END
FROM ak.audiences a
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'johndoe') u1
         CROSS JOIN LATERAL (SELECT id, username FROM ak.users WHERE username = 'emilysmith') u2
LIMIT 4;
