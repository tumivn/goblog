CREATE TYPE content_type AS ENUM ('blog_post', 'course_chapter', 'course_mcq', 'news', 'faq', 'tutorial', 'forum_thread', 'question');

CREATE TABLE profile(
    id              uuid DEFAULT uuid_generate_v4(),
    username        varchar(50) NOT NULL UNIQUE,
    name            varchar(50),
    email           varchar(255) NOT NULL UNIQUE,
    bio             text,
    highlight       varchar(255),
    expertise       varchar(255),
    website         varchar(255),
    birthday        date,
    phone           varchar(20),
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE tag(
    id              uuid DEFAULT uuid_generate_v4(),
    name            varchar(50) NOT NULL,
    slug            varchar(50) NOT NULL UNIQUE,
    image           varchar(255),
    description     varchar(255),
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE article(
    id              uuid DEFAULT uuid_generate_v4(),
    title           varchar(200) NOT NULL,
    slug            varchar(200) NOT NULL UNIQUE,
    type            content_type,
    image           varchar(255),
    abstraction     varchar(500) NOT NULL,
    body            text,
    approved        bool,
    rate_count      int,
    rate_total      bigint,
    like_count      int,
    published       bool,
    allow_comment   bool,
    allow_rating    bool,
    published_at    time,
    profile_id      uuid,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE article_revision(
    id              uuid DEFAULT uuid_generate_v4(),
    article_id      uuid,
    title           varchar(200) NOT NULL,
    abstraction     varchar(500) NOT NULL,
    body            text,
    version         int,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
)
