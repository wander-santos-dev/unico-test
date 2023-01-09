CREATE TABLE IF NOT EXISTS feiras(
id SERIAL PRIMARY KEY,
long VARCHAR(50) NOT NULL,
lat VARCHAR(50)  NOT NULL,
setcens VARCHAR(50)  NOT NULL,
areap VARCHAR(50)  NOT NULL,
coddist VARCHAR(10)  NOT NULL,
distrito VARCHAR(100)  NOT NULL,
codsubpref VARCHAR(10)  NOT NULL,
subprefe VARCHAR(100)  NOT NULL,
regiao5 VARCHAR(20)  NOT NULL,
regiao8 VARCHAR(20)  NOT NULL,
nome_feira VARCHAR(100)  NOT NULL,
registro VARCHAR(20) NOT NULL,
logradouro VARCHAR(120) NOT NULL,
numero VARCHAR(50),
bairro VARCHAR(100) NOT NULL,
referencia TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);