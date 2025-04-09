DROP INDEX IF EXISTS idx_profile_username;
DROP INDEX IF EXISTS idx_profile_email;
DROP INDEX IF EXISTS idx_tag_slug;
DROP INDEX IF EXISTS idx_article_slug;
DROP INDEX IF EXISTS idx_article_profile_id;
DROP INDEX IF EXISTS idx_article_revision_article_id;

DROP TABLE IF EXISTS article_revision;
DROP TABLE IF EXISTS article;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS profile;
DROP TYPE IF EXISTS content_type;