-- tenants テーブルにサンプルデータを挿入
INSERT INTO tenants (name) VALUES ('Tenant A');
INSERT INTO tenants (name) VALUES ('Tenant B');

-- 最後に挿入された2つのtenant_idを取得
WITH inserted_tenants AS (
  SELECT id FROM tenants ORDER BY id DESC LIMIT 2
)
-- persons テーブルにサンプルデータを挿入
, insert_persons AS (
  SELECT
    tenant.id AS tenant_id,
    generate_series(1, 3) AS series  -- 1から3までの数列を生成
  FROM
    inserted_tenants tenant
)
INSERT INTO persons (tenant_id, name, email, birth_date)
SELECT
  tenant_id,
  'Person ' || series || ' of Tenant ' || tenant_id,
  'person' || series || '_tenant' || tenant_id || '@example.com',
  ('1990-01-01'::date + (series * 365))::date  -- ダミーの誕生日を生成
FROM insert_persons;
