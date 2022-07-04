-- name: CreateUser :one
INSERT INTO users(
    id, email, username
) VALUES (@ID, @email, @username) RETURNING created_at;

-- name: userExistsByEmail :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = @email);

-- name: userExistsByUsername :one
SELECT EXISTS (SELECT 1 FROM users WHERE username LIKE @username);
