package facade

import "fmt"

type TV struct{}

func (c *TV) On() {
	fmt.Println("打开电视")
}
func (c *TV) Off() {
	fmt.Println("关闭电视")
}

type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}

func (v *VoiceBox) Off() {
	fmt.Println("关闭 音箱")
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("打开 灯光")
}

func (l *Light) Off() {
	fmt.Println("关闭 灯光")
}

//家庭影院(外观)
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
}

//刚回家
func (hp *HomePlayerFacade) back() {
	fmt.Println("欢迎回家")
	hp.tv.On()
	hp.vb.On()
	hp.light.On()
}

//刚回家
func (hp *HomePlayerFacade) leave() {
	fmt.Println("一路安全")
	hp.tv.Off()
	hp.vb.Off()
	hp.light.Off()
}
