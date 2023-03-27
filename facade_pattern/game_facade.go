package main

// GameFacade is the facade for the game system, client can communicate with the game server throught it
type GameFacade struct {
	userSystem     *UserSystem
	backpackSystem *BackpackSystem
	taskSystem     *TaskSystem
	instanceSystem *InstanceSystem
}

func NewGameFacade() *GameFacade {
	return &GameFacade{
		userSystem:     &UserSystem{},
		backpackSystem: &BackpackSystem{},
		taskSystem:     &TaskSystem{},
		instanceSystem: &InstanceSystem{},
	}
}

func (gf *GameFacade) Login(username, password string) bool {
	return gf.userSystem.Login(username, password)
}

func (gf *GameFacade) Register(username, password string) bool {
	return gf.userSystem.Register(username, password)
}

func (gf *GameFacade) GetItems(userId int) []string {
	return gf.backpackSystem.GetItems(userId)
}

func (gf *GameFacade) AddItem(userId int, itemId string) bool {
	return gf.backpackSystem.AddItem(userId, itemId)
}

func (gf *GameFacade) GetTasks(userId int) []string {
	return gf.taskSystem.GetTasks(userId)
}

func (gf *GameFacade) FinishTask(userId int, taskId string) bool {
	return gf.taskSystem.FinishTask(userId, taskId)
}

func (gf *GameFacade) EnterInstance(userId int, instanceId string) bool {
	return gf.instanceSystem.EnterInstance(userId, instanceId)
}

func (gf *GameFacade) ExitInstance(userId int) bool {
	return gf.instanceSystem.ExitInstance(userId)
}
