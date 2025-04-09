CREATE TYPE content_type AS ENUM ('blog_post', 'course_chapter', 'course_mcq', 'news', 'faq', 'tutorial', 'forum_thread', 'question');

CREATE TABLE profile
(
    id         uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username   varchar(50)  NOT NULL UNIQUE,
    name       varchar(50),
    email      varchar(255) NOT NULL UNIQUE,
    bio        text,
    highlight  varchar(255),
    expertise  varchar(255),
    website    varchar(255),
    birthday   date,
    phone      varchar(20),
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE INDEX idx_profile_username ON profile (username);
CREATE INDEX idx_profile_email ON profile (email);

CREATE TABLE tag
(
    id          uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        varchar(50) NOT NULL,
    slug        varchar(50) NOT NULL UNIQUE,
    image       varchar(255),
    description varchar(255),
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL
);

CREATE INDEX idx_tag_slug ON tag (slug);

CREATE TABLE article
(
    id            uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title         varchar(200) NOT NULL,
    slug          varchar(200) NOT NULL UNIQUE,
    type          content_type,
    image         varchar(255),
    abstraction   varchar(500) NOT NULL,
    body          text,
    approved      bool,
    rate_count    int,
    rate_total    bigint,
    like_count    int,
    published     bool,
    allow_comment bool,
    allow_rating  bool,
    published_at  TIMESTAMP,
    profile_id    uuid REFERENCES profile (id),
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);

CREATE INDEX idx_article_slug ON article (slug);
CREATE INDEX idx_article_profile_id ON article (profile_id);

CREATE TABLE article_revision
(
    id          uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    article_id  uuid REFERENCES article (id),
    title       varchar(200) NOT NULL,
    abstraction varchar(500) NOT NULL,
    body        text,
    version     int,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL
);

CREATE INDEX idx_article_revision_article_id ON article_revision (article_id);