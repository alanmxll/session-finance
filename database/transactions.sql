CREATE TABLE transactions(
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  amount DECIMAL,
  type SMALLINT,
  installment SMALLINT,
  created_at TIMESTAMP
);

INSERT INTO transactions(title, amount, type, installment, created_at)
VALUES ('Freela', '100.0', 0, 1, '2022-01-03 19:23:00');

SELECT * FROM transactions;