package todo

import (
	todov1 "github.com/pivaldi/mmw-contracts/gen/go/todo/v1"
)

// The Topic Name (Routing Key)
const (
	TopicUserTasksDeleted  = "todo.userTasks.deleted.v1"
	TopicUserTaskDeleted   = "todo.userTask.deleted.v1"
	TopicUserTaskCreated   = "todo.userTasks.created.v1"
	TopicUserTaskUpdated   = "todo.userTasks.updated.v1"
	TopicUserTaskReopened  = "todo.userTasks.reopened.v1"
	TopicUserTaskCompleted = "todo.userTasks.completed.v1"
)

var AllEvents = []string{
	TopicUserTasksDeleted,
	TopicUserTaskCreated,
	TopicUserTaskUpdated,
	TopicUserTaskReopened,
	TopicUserTaskCompleted,
}

type UserTasksDeletedEvent = todov1.UserTasksDeletedEvent
