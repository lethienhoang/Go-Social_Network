-- name: CreateUser :one
INSERT INTO users(
    id, email, username
) VALUES (@ID, @email, @username) RETURNING created_at;

-- name: UserByEmail :one
SELECT 
    *
FROM users WHERE email = LOWER(@email);

-- name: UserByUsername :one
SELECT 
    *
FROM users WHERE username = LOWER(@username);

-- name: UserExistsByEmail :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = @email);

-- name: UserExistsByUsername :one
SELECT EXISTS (SELECT 1 FROM users WHERE username LIKE @username);
