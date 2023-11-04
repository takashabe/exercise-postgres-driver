CREATE TABLE tenants (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE persons (
  id SERIAL PRIMARY KEY,
  tenant_id INTEGER NOT NULL REFERENCES tenants(id),
  name VARCHAR(50),
  email VARCHAR(100),
  birth_date DATE,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
ALTER TABLE persons ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_policy ON persons
FOR ALL
USING (tenant_id = current_setting('app.tenant_id')::INTEGER)
