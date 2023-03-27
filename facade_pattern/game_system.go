package main

// UserSystem is the system for operations of user
type UserSystem struct{}

// Login checks user's auth
func (us *UserSystem) Login(username, password string) bool {
	return true
}

// Register is used to register a new user
func (us *UserSystem) Register(username, password string) bool {
	return true
}

// BackpackSystem is the system to handle user's backpack's item
type BackpackSystem struct{}

// GetItems returns specific item
func (bs *BackpackSystem) GetItems(userId int) []string {
	return nil
}

// AddItem adds a new item to user's backpack
func (bs *BackpackSystem) AddItem(userId int, itemId string) bool {
	return true
}

// TaskSystem is the system to handle user's tasks
type TaskSystem struct{}

// GetTasks return all task of user
func (ts *TaskSystem) GetTasks(userId int) []string {
	return nil
}

// FinishTask ends the specific task
func (ts *TaskSystem) FinishTask(userId int, taskId string) bool {
	return true
}

// InstanceSystem is the instance system
type InstanceSystem struct{}

// EnterInstance used to enter the instance
func (is *InstanceSystem) EnterInstance(userId int, instanceId string) bool {
	return true
}

// ExitInstance used the exit the instance
func (is *InstanceSystem) ExitInstance(userId int) bool {
	return true
}
