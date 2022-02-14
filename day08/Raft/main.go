package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 定义3节点常量
const raftCount = 3

// Leader 声明leader 对象
type Leader struct {
	// 任期
	Term int
	// LeaderId编号
	LeaderId int
}

// Raft 声明raft
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	votedFor int
	// 3个状态
	// 0 follower   1 candidate  2 leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 设置当前节点的领导
	currentLeader int
	// 节点间发信息的通道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳信号的通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

// 0 还没上任  -1 没有编号
var leader = Leader{0, -1}

func main() {
	// 过程：有3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 会产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
		Make(i)
	}
	for {
	}
}

func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// 刚初始化，还没有投票，用-1表示
	rf.votedFor = -1
	// 刚刚初始化的时候所有结点都是follower
	rf.state = 0
	rf.timeout = 0
	// 刚刚初始化目前没有领导id
	rf.currentLeader = -1
	// 节点任期
	rf.setTerm(0)
	// 初始化各种通道
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子，因为我们结点的选举周期是随机的（150ms-300ms），需要随机数，所以设置一个随机数种子备用
	rand.Seed(time.Now().UnixNano())

	// 选举的协程
	go rf.election()

	// 心跳检测的协程
	go rf.sendLeaderHeartBeat()

	return rf
}

// 设置任期，这里封装成函数，方便其他地方调用
func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

func (rf *Raft) election() {
	// 设置标记，是否选出了leader
	var result bool
	result = false
	for {
		// 设置超时时间（150-300）
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()

		select {
		// 这里是延时等待，等待结点的选举周期结束之后，
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为：", rf.state)
		}

		if !result {
			// 选主逻辑
			result = rf.election_one_round(&leader)
		}
	}
}

// 生成指定范围的随机数
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (rf *Raft) election_one_round(leader *Leader) bool {
	// 定义超时
	var timeout int64
	timeout = 100
	// 投票数量
	var vote int
	// 定义是否开始产生心跳信号
	var triggerHeartbeat bool
	// 时间
	last := millisecond()
	// 用于返回值
	success := false

	// 选举前给当前节点变成candidate
	rf.mu.Lock()
	// 修改状态
	rf.becomeCandidate()
	rf.mu.Unlock()

	// 开始选主
	fmt.Println("开始选主")
	for {
		// 遍历所有结点拉选票
		for i := 0; i < raftCount; i++ {
			// 自己就不用拉选票了，跳过自己
			if i != rf.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}

		// 设置投票数量
		vote = 0
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量加1
					vote++
					// 若选票个数，大于节点个数 / 2 ，则成功
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						// 变化成主节点，选主成功了
						// 开始触发心跳信号检测
						triggerHeartbeat = true
						rf.mu.Lock()
						// 变主
						rf.becomeLeader()
						rf.mu.Unlock()
						// 由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}

		// 做最后校验工作
		// 若不超时，且票数大于一半，则选举成功，break
		if timeout+last < millisecond() || (vote > raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1) // 任期要加1
	rf.votedFor = rf.me            // 变为candidate后第一票投给自己
	rf.currentLeader = -1
}

func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// leader节点发送心跳信号
// 顺便完成数据同步，这里先不实现了，后面有现成的
// 看小弟挂没挂
func (rf *Raft) sendLeaderHeartBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

// 用于返回给leader的确认信号
func (rf *Raft) sendAppendEntriesImpl() {
	// 是主就别跑了
	if rf.currentLeader == rf.me {
		// 此时是leader
		// 记录确认信号的节点个数
		var success_cout = 0
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					rf.heartbeatRe <- true
				}()
			}
		}
		// 计算返回确认信号个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_cout++
					if success_cout > raftCount/2 {
						fmt.Println("投票选举成功，心跳信号OK")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}
