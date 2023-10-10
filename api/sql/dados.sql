insert into users (nome, nick, email, senha)
values
("Usuário 1", "user_1", "user1@gmail.com", "ThisScriptExampleAreNotWorkingOnWebProjectExample"), -- user1
("Usuário 2", "user_2", "user2@gmail.com", "ThisScriptExampleAreNotWorkingOnWebProjectExample"), -- user2
("Usuário 3", "user_3", "user3@gmail.com", "ThisScriptExampleAreNotWorkingOnWebProjectExample"); -- user3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);
