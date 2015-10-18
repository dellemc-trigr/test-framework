package shared

var Tags = struct {
	TASK_TYPE string
	CONTAINER_NAME string
	FILESERVER_IP string
	TARGET_HOST string
	ACCEPTED_HOST string
}{
	TASK_TYPE: "TASK_TYPE",
	CONTAINER_NAME: "CONTAINER_NAME",
	FILESERVER_IP: "FILESERVER_IP",
	TARGET_HOST: "TARGET_HOST",
	ACCEPTED_HOST: "ACCEPTED_HOST",
}

var TaskTypes = struct {
	RUN_CONTAINER string
	CHECKPOINT_CONTAINER string
	RESTORE_CONTAINER string
	TEST_TASK string
	GET_LOGS string
}{
	RUN_CONTAINER: "RUN_CONTAINER",
	CHECKPOINT_CONTAINER: "CHECKPOINT_CONTAINER",
	RESTORE_CONTAINER: "RESTORE_CONTAINER",
	TEST_TASK: "TEST_TASK",
	GET_LOGS: "GET_LOGS",
}