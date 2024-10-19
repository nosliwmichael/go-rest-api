package queries

const (
	SelectUserById = `
		SELECT
			ID,
			USERNAME,
			FIRST_NAME,
			LAST_NAME,
			AGE
		FROM
			user
		WHERE
			ID = ?
	`
	InsertUser = `
		INSERT INTO user
			(USERNAME,
			FIRST_NAME,
			LAST_NAME,
			AGE)
		VALUES
			(?, ?, ?, ?)
	`
)
