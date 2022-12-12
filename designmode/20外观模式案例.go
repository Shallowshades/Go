package main

import "fmt"

// 电视机
type TV struct{}

func (t *TV) On() {
	fmt.Println("put on TV...")
}

func (t *TV) Off() {
	fmt.Println("Put off TV...")
}

// 音响
type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("put on VoiceBox...")
}

func (v *VoiceBox) Off() {
	fmt.Println("put off VoiceBox...")
}

// 灯光
type Light struct{}

func (v *Light) On() {
	fmt.Println("put on Light...")
}

func (v *Light) Off() {
	fmt.Println("put off Light...")
}

// 游戏机
type Xbox struct{}

func (v *Xbox) On() {
	fmt.Println("put on Xbox...")
}

func (v *Xbox) Off() {
	fmt.Println("put off Xbox...")
}

// 麦克风
type MicroPhone struct{}

func (v *MicroPhone) On() {
	fmt.Println("put on MicroPhone...")
}

func (v *MicroPhone) Off() {
	fmt.Println("put off MicroPhone...")
}

// 投影仪
type Projector struct{}

func (v *Projector) On() {
	fmt.Println("put on Projector...")
}

func (v *Projector) Off() {
	fmt.Println("put off Projector...")
}

// 家庭影院（外观）
type HomePlayerFacade struct {
	tv         TV
	voiceBox   VoiceBox
	light      Light
	xbox       Xbox
	microphone MicroPhone
	projector  Projector
}

// KTV 模式
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("It's about to enter KTV mode...")
	hp.tv.On()
	hp.projector.On()
	hp.microphone.On()
	hp.light.Off()
	hp.voiceBox.On()

}

// 游戏模式
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("It's about to enter game mode...")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func main() {

	homePlayer := new(HomePlayerFacade)
	homePlayer.DoKTV()

	fmt.Println("-----------------")
	homePlayer.DoGame()
}
