-- name: CreateUser :one
INSERT INTO "user" (
    id, email, username
) VALUES (@ID, LOWER(@email), @username) RETURNING created_at;

-- name: UserByEmail :one
SELECT * FROM "user" WHERE email = @email;

-- name: UserByUsername :one
SELECT * FROM "user" WHERE username = @username;

-- name: UserExistsByEmail :one
SELECT EXISTS (SELECT 1 FROM "user" WHERE LOWER(email) = LOWER(@email));

-- name: UserExistsByUsername :one
SELECT EXISTS (SELECT 1 FROM "user" WHERE LOWER(username) = LOWER(@username));
