package main

import (
	"fmt"
)

type Process struct {
	ID             int
	BurstTime      int
	RemainingTime  int
	WaitingTime    int
	TurnaroundTime int
}

type RoundRobin struct {
	processes []Process
	quantum   int
}

func NewRoundRobin(processes []Process, quantum int) *RoundRobin {
	copiedProcesses := make([]Process, len(processes))
	copy(copiedProcesses, processes)

	for i := range copiedProcesses {
		copiedProcesses[i].RemainingTime = copiedProcesses[i].BurstTime
		copiedProcesses[i].WaitingTime = 0
		copiedProcesses[i].TurnaroundTime = 0
	}

	return &RoundRobin{
		processes: copiedProcesses,
		quantum:   quantum,
	}
}

func (rr *RoundRobin) Execute() ([]Process, float64, float64) {
	queue := make([]int, 0)

	for i := range rr.processes {
		queue = append(queue, i)
	}

	currentTime := 0
	completionTimes := make([]int, len(rr.processes))

	for len(queue) > 0 {
		processIndex := queue[0]
		queue = queue[1:]

		process := &rr.processes[processIndex]

		executionTime := process.RemainingTime
		if executionTime > rr.quantum {
			executionTime = rr.quantum
		}

		currentTime += executionTime
		process.RemainingTime -= executionTime

		if process.RemainingTime > 0 {
			queue = append(queue, processIndex)
		} else {
			completionTimes[processIndex] = currentTime
		}
	}

	totalWaitingTime := 0
	totalTurnaroundTime := 0

	for i := range rr.processes {
		rr.processes[i].TurnaroundTime = completionTimes[i]
		rr.processes[i].WaitingTime = rr.processes[i].TurnaroundTime - rr.processes[i].BurstTime
		totalWaitingTime += rr.processes[i].WaitingTime
		totalTurnaroundTime += rr.processes[i].TurnaroundTime
	}

	avgWaitingTime := float64(totalWaitingTime) / float64(len(rr.processes))
	avgTurnaroundTime := float64(totalTurnaroundTime) / float64(len(rr.processes))

	return rr.processes, avgWaitingTime, avgTurnaroundTime
}

func PrintResults(processes []Process, avgWaitingTime, avgTurnaroundTime float64) {
	fmt.Println("\n╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║          RESULTADO DO ALGORITMO ROUND ROBIN              ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")

	fmt.Printf("%-6s %-12s %-14s %-15s %-10s\n", "ID", "Burst Time", "Waiting Time", "Turnaround Time", "Completion")
	fmt.Println("─────────────────────────────────────────────────────────────")

	for _, p := range processes {
		fmt.Printf("%-6d %-12d %-14d %-15d %-10d\n",
			p.ID, p.BurstTime, p.WaitingTime, p.TurnaroundTime, p.TurnaroundTime)
	}

	fmt.Println("─────────────────────────────────────────────────────────────")
	fmt.Printf("\nTempo médio de espera: %.2f ms\n", avgWaitingTime)
	fmt.Printf("Tempo médio de retorno: %.2f ms\n", avgTurnaroundTime)
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("      ROUND ROBIN - ALGORITMO DE ESCALONAMENTO")
	fmt.Println("═══════════════════════════════════════════════════════════")

	fmt.Println("\n📋 EXEMPLO 1: Escalonamento Round Robin")
	fmt.Println("Quantum: 4ms")

	processes1 := []Process{
		{ID: 1, BurstTime: 8},
		{ID: 2, BurstTime: 4},
		{ID: 3, BurstTime: 2},
		{ID: 4, BurstTime: 1},
	}

	rr1 := NewRoundRobin(processes1, 4)

	result1, avgWait1, avgTurnaround1 := rr1.Execute()
	PrintResults(result1, avgWait1, avgTurnaround1)

	fmt.Println("\n\n📋 EXEMPLO 2: Comparação com Quantum = 2ms")

	processes2 := []Process{
		{ID: 1, BurstTime: 8},
		{ID: 2, BurstTime: 4},
		{ID: 3, BurstTime: 2},
		{ID: 4, BurstTime: 1},
	}

	rr2 := NewRoundRobin(processes2, 2)
	result2, avgWait2, avgTurnaround2 := rr2.Execute()
	PrintResults(result2, avgWait2, avgTurnaround2)

	fmt.Println("\n\n╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║              COMPARAÇÃO DE QUANTUMS                       ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Printf("\nQuantum = 4ms: Espera média = %.2f ms\n", avgWait1)
	fmt.Printf("Quantum = 2ms: Espera média = %.2f ms\n", avgWait2)
}
