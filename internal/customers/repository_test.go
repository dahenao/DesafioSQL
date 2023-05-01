package customers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	dns := mysql.Config{
		User:   "user1",
		Passwd: "secret_password",
		DBName: "fantasy_products",
		Addr:   "127.0.0.1",
	}
	txdb.Register("txdb", "mysql", dns.FormatDSN())
}

func TestTotalAmoutCustomers(t *testing.T) {

	t.Run("Total Amount", func(t *testing.T) {
		//arrrange
		db, err := sql.Open("txdb", "identifier")
		assert.NoError(t, err)
		defer db.Close()
		rp := NewRepository(db)
		//act

		exp := []*Top5products{
			{Description: "Vinegar - Raspberry", Count: 660},
			{Description: "Flour - Corn, Fine", Count: 521},
			{Description: "Cookie - Oatmeal", Count: 467},
		}

		result, err := rp.GetTop5SoldPrd()

		//asssert
		assert.NoError(t, err)
		assert.Equal(t, exp, result)
	})

}
