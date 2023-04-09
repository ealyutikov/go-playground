-- +goose Up
-- +goose StatementBegin
insert into users
values ('1ff5b90d-4430-4e15-92a7-6f7cb482fceb', 'user_1', true),
       ('03f4baf2-9ecf-4041-907a-4154cd415874', 'user_2', true),
       ('e40e8348-0800-4679-9585-24224977fa83', 'user_3', true),
       ('db0cc29e-bb64-48cc-ba62-525acafa9443', 'user_4', true),
       ('4441dfae-b963-4cce-a014-681fd0556583', 'user_5', true);
-- +goose StatementEnd
