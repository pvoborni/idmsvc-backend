
-- File created by: ./bin/db-tool new domains_org_index
BEGIN;

CREATE INDEX idx_domains_org_id ON domains (org_id);

COMMIT;
