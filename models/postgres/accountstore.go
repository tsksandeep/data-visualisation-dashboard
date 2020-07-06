package postgres

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/ctrlrsf/logdna"
	"github.com/pkg/errors"

	"know/db"
	"know/models"
)

type accountStore struct {
	db  *db.DB
	log *log.Client
}

//NewAccountStore initiates a new instance of ConfigStore
func NewAccountStore(db *db.DB, logDNAClient *log.Client) (models.AccountStore, error) {
	if db == nil {
		return nil, errors.New("account store new instance creation failed: invalid database")
	}

	return &accountStore{
		db:  db,
		log: logDNAClient,
	}, nil
}

func (as *accountStore) Save(account *models.Account) error {

	as.log.Log(time.Now(), fmt.Sprintf("adding new account info: %s %s %s", account.Email, account.FirstName, account.LastName))

	if account == nil {
		return models.ErrAddAccount
	}

	if account.Email == "" || account.FirstName == "" || account.LastName == "" || account.Password == "" {
		return models.ErrAddAccount
	}

	sqlStmt := `INSERT INTO account_info (email, firstName, lastName, password) VALUES ($1, $2, $3, $4)`

	_, err := as.db.Exec(sqlStmt, account.Email, account.FirstName, account.LastName, account.Password)
	if err != nil {
		return errors.Wrap(err, "add account record failed: sql statement exec failed")
	}

	return nil
}

func (as *accountStore) Delete(email string) error {

	as.log.Log(time.Now(), fmt.Sprintf("deleting account %s", email))

	if email == "" {
		return models.ErrDeleteAccount
	}

	sqlStmt := `DELETE FROM account_info WHERE email = $1`
	result, err := as.db.Exec(sqlStmt, email)
	if err != nil {
		return errors.Wrap(err, "account deletion failed, sql statement exec failed")
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return models.ErrNoAccountFound
	}

	return nil
}

func (as *accountStore) Get(email string) (*models.Account, error) {

	as.log.Log(time.Now(), fmt.Sprintf("getting account %s", email))

	if email == "" {
		return nil, models.ErrGetAccount
	}

	query := `SELECT email, firstName, lastName, password FROM account_info WHERE email = $1`

	var account models.Account

	row := as.db.QueryRow(query, email)
	err := row.Scan(&account.Email, &account.FirstName, &account.LastName, &account.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoAccountFound
		}

		return nil, errors.Wrap(err, "get account failed due to sql query failure")
	}

	return &account, nil
}

func (as *accountStore) GetAll() ([]models.Account, error) {

	as.log.Log(time.Now(), "getting all account info")
	query := `SELECT email, firstName, lastName, password FROM account_info`

	rows, err := as.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "get account records failed due to sql query failure")
	}

	defer rows.Close()

	var accounts = []models.Account{}
	var account models.Account

	for rows.Next() {
		err := rows.Scan(&account.Email, &account.FirstName, &account.LastName, &account.Password)
		if err != nil {
			return nil, errors.Wrap(err, "get all accounts failed, failed to scan exec result")
		}

		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		return nil, models.ErrNoAccountFound
	}

	return accounts, nil
}
