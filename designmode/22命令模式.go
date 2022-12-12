package main

import "fmt"

// 命令的接受者  医生
type Doctor struct{}

func (d *Doctor) treatEyes() {
	fmt.Println("treat eyes...")
}

func (d *Doctor) treatNose() {
	fmt.Println("treat nose...")
}

// 抽象的命令
type Command interface {
	Treat()
}

// 具体的命令 治疗眼睛的病单
type CommandTreatEyes struct {
	doctor *Doctor
}

func (cmd *CommandTreatEyes) Treat() {
	cmd.doctor.treatEyes()
}

// 具体的命令 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 命令的掉用者 护士
type Nurse struct {
	CmdList []Command
}

// 发送命令的方法 发送病单
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

// 业务逻辑层 病人
func main() {

	//医生
	doctor := new(Doctor)

	//病人 通过填写病单来看病 看什么病指定什么病单，还要指定一个医生
	cmdEye := CommandTreatEyes{doctor}
	cmdNose := CommandTreatNose{doctor}

	//护士 通过收集病人填写的病单传递给医生诊断治疗
	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}
