package waitformore

import (
	"fmt"
	"sync"
)

type RamQueueRepository struct {
	queues []*QueueDefinition
	mutex  sync.Mutex
}

func (r *RamQueueRepository) CreateNewQueue(queueTitle string) (QueueDefinition, *RepositoryError) {
	definition := QueueDefinition{
		UserKey:    GenerateUserKey(),
		AdminKey:   GenerateAdminKey(),
		QueueTitle: queueTitle,
		QueueState: QueueState{
			LastPulled: 0,
			LastCalled: 0,
		},
	}
	r.mutex.Lock()
	r.queues = append(r.queues, &definition)
	r.mutex.Unlock()

	return definition, nil
}

func (r *RamQueueRepository) GetUserView(userKey string) (UserView, *RepositoryError) {

	for _, queueDefinition := range r.queues {
		if userKey == queueDefinition.UserKey {
			return UserView{
				UserKey:    queueDefinition.UserKey,
				QueueTitle: queueDefinition.QueueTitle,
				QueueState: queueDefinition.QueueState,
			}, nil
		}
	}
	return UserView{}, &RepositoryError{
		HttpStatus: 404,
		ErrorText:  fmt.Sprintf("not found by userKey %s", userKey),
	}
}

func (r *RamQueueRepository) GetQueueDefinition(adminKey string) (QueueDefinition, *RepositoryError) {

	for _, queueDefinition := range r.queues {
		if adminKey == queueDefinition.AdminKey {
			return *queueDefinition, nil
		}
	}
	return QueueDefinition{}, &RepositoryError{
		HttpStatus: 404,
		ErrorText:  fmt.Sprintf("not found by userKey %s", adminKey),
	}
}

func (r *RamQueueRepository) PullByUserKey(userKey string) (QueueDefinition, *RepositoryError) {

	var queueDefinition *QueueDefinition
	for _, queueDefinitionToCheck := range r.queues {
		if userKey == queueDefinitionToCheck.UserKey {
			queueDefinition = queueDefinitionToCheck
		}
	}
	if queueDefinition == nil {
		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("not found by userKey %s", userKey),
		}
	}

	var reposErr *RepositoryError

	// synch
	r.mutex.Lock()
	pulled, err := queueDefinition.QueueState.Pull()
	if err == nil {
		queueDefinition.QueueState.LastPulled = pulled
	} else {
		reposErr = &RepositoryError{
			// todo uli better code?
			HttpStatus: 406,
			ErrorText:  fmt.Sprintf("nothing left to pull %s", userKey),
		}
	}
	r.mutex.Unlock()

	return *queueDefinition, reposErr
}

func (r *RamQueueRepository) CallByAdminKey(adminKey string) (QueueDefinition, *RepositoryError) {

	var queueDefinition *QueueDefinition
	for _, queueDefinitionToCheck := range r.queues {
		if adminKey == queueDefinitionToCheck.AdminKey {
			queueDefinition = queueDefinitionToCheck
		}
	}
	if queueDefinition == nil {
		return QueueDefinition{}, &RepositoryError{
			HttpStatus: 404,
			ErrorText:  fmt.Sprintf("not found by adminKey %s", adminKey),
		}
	}

	var reposErr *RepositoryError

	// synch
	r.mutex.Lock()
	called, err := queueDefinition.QueueState.Call()
	if err == nil {
		queueDefinition.QueueState.LastCalled = called
	} else {
		reposErr = &RepositoryError{
			// todo uli better code?
			HttpStatus: 406,
			ErrorText:  fmt.Sprintf("nothing left to call %s", adminKey),
		}
	}
	r.mutex.Unlock()

	return *queueDefinition, reposErr
}
