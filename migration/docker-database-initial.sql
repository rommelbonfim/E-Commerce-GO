CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    preco DECIMAL(10, 2) NOT NULL,
    quantidade INT NOT NULL,
    categoria VARCHAR(255),
    imagem_url VARCHAR(255),
    disponivel BOOLEAN DEFAULT true
);

INSERT INTO products (nome, descricao, preco, quantidade, categoria, imagem_url, disponivel) VALUES
('Produto 1', 'Descrição do Produto 1', 100.00, 10, 'Eletrônicos', 'http://example.com/produto1.jpg', true),
('Produto 2', 'Descrição do Produto 2', 250.50, 5, 'Vestuário', 'http://example.com/produto2.jpg', false);
