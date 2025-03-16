create table user_quiz_answer(
    id                      bigserial       primary key,
    user_quiz_id            bigint          not null,
    quiz_question_id        bigint          not null,
    option_answer           text,
    text_answer             text
);