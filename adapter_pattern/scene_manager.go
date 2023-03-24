package main

import (
	"fmt"
)

// Scene defines the interface of all scene
type Scene interface {
	// Enter the specific scene
	Enter()
}

// NormalScene is the normal scene
type NormalScene struct{}

func (ns *NormalScene) Enter() {
	fmt.Println("Enter Normal Scene")
}

// BossScene is the boss scene
type BossScene struct{}

func (bs *BossScene) Start() {
	fmt.Println("Enter Boss Scene")
}

// BossSceneAdapter is the adapter for BossScene to adapt Scene interface
type BossSceneAdapter struct {
	BossScene *BossScene
}

func (bsa *BossSceneAdapter) Enter() {
	bsa.BossScene.Start()
}

// SceneManager is the scene manager, it can switch to any scene by adapter which implements Scene
type SceneManager struct {
	CurrentScene Scene
}

func (sm *SceneManager) ChangeScene(s Scene) {
	sm.CurrentScene = s
	sm.CurrentScene.Enter()
}

func main() {
	sm := &SceneManager{}

	// enter normal scene
	ns := &NormalScene{}
	sm.ChangeScene(ns)

	// eneter boss scene by adapter
	bs := &BossScene{}
	sm.ChangeScene(&BossSceneAdapter{
		BossScene: bs,
	})
}
