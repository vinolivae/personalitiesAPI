-- O docker compose será responsável por executar as instruções, na chave "volumes"

create table personalities(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO personalities(name, history) VALUES
('João bebe água', 'Era memro do Partido Liberal e envolvido com a política local, tendo ideias conservadoras. João Bebe-Água também foi presidente da Câmara dos Vereadores de São Cristóvão e líder de uma Irmandade de Amparo dos Homens Pardos.	De acordo com a lenda, ele liderou uma rebelião de 400 homens contra a transferência da capital para Santo Antônio de Aracaju.	João Bebe-Água morreu pobre e desacreditado, ainda sonhando com São Cristóvão como capital. Seu apelido “Bebe-Água” é uma referência a seu suposto gosto por bebidas alcoólicas, mas ele teve uma carreira política ativa e desempenhou cargos importantes, incluindo juiz de paz.'),
('Zé Peixe', 'Por muitos anos, Zé Peixe atuou como prático conduzindo embarcações que entravam e saíam de Aracaju, pelo Rio Sergipe. O inusitado, em sua tarefa, se devia ao fato de não necessitar de embarcação de apoio para transportá-lo após conduzir o navio até fora da barra, ele saltava e voltava à terra nadando. Segundo ele, o trajeto a nado poderia durar de 1 hora e meia até 3 horas, dependendo das condições do vento.');