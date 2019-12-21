package postgresql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectDB(t *testing.T) {
	//t.Run("Valid connection", func(t *testing.T) {
	//	connStr := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s",
	//		os.Getenv("DB_DRIVER"),
	//		os.Getenv("DB_USER"),
	//		os.Getenv("DB_PASSWORD"),
	//		os.Getenv("DB_HOST"),
	//		os.Getenv("DB_NAME"),
	//		os.Getenv("DB_SSL_MODE"),
	//	)
	//
	//	_, err := ConnectDB(connStr)
	//	assert.NoError(t, err, "should be no error")
	//})

	t.Run("Invalid connection", func(t *testing.T) {
		_, err := ConnectDB("postgres://invalid:@invalid/invalid")
		assert.Error(t, err, "should be an error")
	})
}
