package waitformore

type QueueRepository interface {
	GetUserView(userKey string) (UserView, *RepositoryError)
	CreateNewQueue(queueTitle string) (QueueDefinition, *RepositoryError)
	GetQueueDefinition(adminKey string) (QueueDefinition, *RepositoryError)
	PullByUserKey(userKey string) (QueueDefinition, *RepositoryError)
	CallByAdminKey(adminKey string) (QueueDefinition, *RepositoryError)
	GetStatus() (QueueingStatus, *RepositoryError)
}
