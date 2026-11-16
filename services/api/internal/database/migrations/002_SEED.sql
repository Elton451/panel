-- +goose Up
INSERT INTO users (
  id,
  firstName,
  lastName
) VALUES ( 1, 'Admin', 'Orderer' );

-- +goose Down
DELETE FROM users
  WHERE id = 1;


