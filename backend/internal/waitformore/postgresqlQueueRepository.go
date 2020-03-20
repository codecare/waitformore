package waitformore

import (
	"database/sql"
	"fmt"
	"time"
)

/*
CREATE TABLE QueueDefinition (
  id    SERIAL NOT NULL,
  Title   varchar(255) NOT NULL,
  UserKey   varchar(255) NOT NULL,
  AdminKey  varchar(255) NOT NULL,
  LastPulled integer NOT NULL,
  LastCalled integer NOT NULL,
  created timestamp not null,
  updated timestamp not null,
  deleted timestamp null,
  version int not null default 0
);
*/

type QueueDefinitionDbWrapper struct {
	queueDefinition QueueDefinition
	id              int
	version         int
}

type PostgresqlQueueRepository struct {
	DatabaseConnectionPool *sql.DB
}

func (p PostgresqlQueueRepository) GetUserView(userKey string) (UserView, *RepositoryError) {

	row := p.DatabaseConnectionPool.QueryRow("SELECT UserKey, Title, LastPulled, LastCalled FROM QueueDefinition WHERE UserKey = $1 and deleted is null", userKey)

	result := UserView{}
	err := row.Scan(&result.UserKey, &result.QueueTitle, &result.QueueState.LastPulled, &result.QueueState.LastCalled)

	if err == sql.ErrNoRows {
		return UserView{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("queue not found"),
		}
	} else if err != nil {
		fmt.Printf("db error: %v\n", err)

		return UserView{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}

	return result, nil
}

func (p PostgresqlQueueRepository) CreateNewQueue(queueTitle string) (QueueDefinition, *RepositoryError) {

	definition := QueueDefinition{
		UserKey:    GenerateUserKey(),
		AdminKey:   GenerateAdminKey(),
		QueueTitle: queueTitle,
		QueueState: QueueState{
			LastPulled: 0,
			LastCalled: 0,
		},
	}
	now := time.Now()
	_, err := p.DatabaseConnectionPool.Exec("INSERT INTO QueueDefinition (Title, UserKey, AdminKey, LastPulled, LastCalled, Created, Updated, Deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		definition.QueueTitle, definition.UserKey, definition.AdminKey, definition.QueueState.LastPulled, definition.QueueState.LastCalled, now, now, nil)
	if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	return definition, nil
}

func (p PostgresqlQueueRepository) GetQueueDefinition(adminKey string) (QueueDefinition, *RepositoryError) {

	row := p.DatabaseConnectionPool.QueryRow("SELECT UserKey, AdminKey, Title, LastPulled, LastCalled FROM QueueDefinition WHERE adminKey = $1 and deleted is null", adminKey)

	result := QueueDefinition{}
	err := row.Scan(&result.UserKey, &result.AdminKey, &result.QueueTitle, &result.QueueState.LastPulled, &result.QueueState.LastCalled)

	if err == sql.ErrNoRows {
		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("queue not found"),
		}
	} else if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	return result, nil
}

func (p PostgresqlQueueRepository) PullByUserKey(userKey string) (QueueDefinition, *RepositoryError) {

	row := p.DatabaseConnectionPool.QueryRow("SELECT id, version, UserKey, AdminKey, Title, LastPulled, LastCalled FROM QueueDefinition WHERE userKey = $1 and deleted is null", userKey)

	wrapper := QueueDefinitionDbWrapper{}
	err := row.Scan(&wrapper.id, &wrapper.version, &wrapper.queueDefinition.UserKey, &wrapper.queueDefinition.AdminKey, &wrapper.queueDefinition.QueueTitle, &wrapper.queueDefinition.QueueState.LastPulled, &wrapper.queueDefinition.QueueState.LastCalled)

	if err == sql.ErrNoRows {
		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("queue not found"),
		}
	} else if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}

	pulled, err := wrapper.queueDefinition.QueueState.Pull()
	if err == nil {
		wrapper.queueDefinition.QueueState.LastPulled = pulled
	} else {
		return wrapper.queueDefinition, &RepositoryError{

			HttpStatus: 406,
			ErrorText:  fmt.Sprintf("nothing left to pull %s", userKey),
		}
	}

	now := time.Now()
	updateResult, err := p.DatabaseConnectionPool.Exec("update QueueDefinition set LastPulled = $1, Updated=$2, version = version + 1 where id=$3 and version=$4", wrapper.queueDefinition.QueueState.LastPulled, now, wrapper.id, wrapper.version)
	if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	rowsAffected, err := updateResult.RowsAffected()
	if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	if rowsAffected != 1 {
		// we have a concurrent update
		// todo we should retry here up to 3 times
		fmt.Printf("concurrent update error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}

	return wrapper.queueDefinition, nil
}

func (p PostgresqlQueueRepository) CallByAdminKey(adminKey string) (QueueDefinition, *RepositoryError) {

	row := p.DatabaseConnectionPool.QueryRow("SELECT id, version, UserKey, AdminKey, Title, LastPulled, LastCalled FROM QueueDefinition WHERE adminKey = $1 and deleted is null", adminKey)

	wrapper := QueueDefinitionDbWrapper{}
	err := row.Scan(&wrapper.id, &wrapper.version, &wrapper.queueDefinition.UserKey, &wrapper.queueDefinition.AdminKey, &wrapper.queueDefinition.QueueTitle, &wrapper.queueDefinition.QueueState.LastPulled, &wrapper.queueDefinition.QueueState.LastCalled)

	if err == sql.ErrNoRows {
		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("queue not found"),
		}
	} else if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}

	called, err := wrapper.queueDefinition.QueueState.Call()
	if err == nil {
		wrapper.queueDefinition.QueueState.LastCalled = called
	} else {
		return wrapper.queueDefinition, &RepositoryError{

			HttpStatus: 406,
			ErrorText:  fmt.Sprintf("nothing left to call %s", adminKey),
		}
	}

	now := time.Now()
	updateResult, err := p.DatabaseConnectionPool.Exec("update QueueDefinition set LastCalled = $1, Updated=$2, version = version + 1 where id=$3 and version=$4", wrapper.queueDefinition.QueueState.LastCalled, now, wrapper.id, wrapper.version)
	if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	rowsAffected, err := updateResult.RowsAffected()
	if err != nil {
		fmt.Printf("db error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	if rowsAffected != 1 {
		// we have a concurrent update
		// todo we should retry here up to 3 times
		fmt.Printf("concurrent update error: %v\n", err)

		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}

	return wrapper.queueDefinition, nil
}

func (p PostgresqlQueueRepository) GetStatus() (QueueingStatus, *RepositoryError) {

	row := p.DatabaseConnectionPool.QueryRow("select (select count(1) from queuedefinition) as TotalQueues, (select count(1) from queuedefinition where updated > (CURRENT_TIMESTAMP - INTERVAL '1 day')) as UpdatesLastDay")

	result := QueueingStatus{}
	err := row.Scan(&result.TotalQueues, &result.UpdatesLastDay)

	if err == sql.ErrNoRows {
		fmt.Printf("db error: %v\n", err)

		return QueueingStatus{}, &RepositoryError{
			HttpStatus: 500,
			ErrorText:  fmt.Sprintf("internal server error"),
		}
	}
	return result, nil
}
