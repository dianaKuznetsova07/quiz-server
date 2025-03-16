create table quiz_question(
    id              bigserial       primary key,
    quiz_id         bigint not null,
    title           text   not null,
    question_type   text   not null,
    options         text[],
    created_at      timestamptz     not null
);