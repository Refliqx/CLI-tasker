package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Refliqx/tasker-project/internal/storage"
	"github.com/Refliqx/tasker-project/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// setup dependency
	repo := storage.NewFileRepository("data/tasks.json")
	service := task.NewService(repo)

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("error: judul task wajib diisi")
			return
		}

		title := os.Args[2]
		description := ""
		var dueDate time.Time

		newTask, err := service.AddTask(title, description, dueDate)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("task berhasil ditambahkan:")
		fmt.Printf("ID: %d | %s\n", newTask.ID, newTask.Title)

	case "list":
		tasks, err := service.ListTask()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("belum ada task")
			return
		}

		for _, t := range tasks {
			status := " "
			if t.IsCompleted {
				status = "✓"
			}

			date := t.CreatedAt.Format("2006-01-02")

			fmt.Printf("[%s] %d. %s (%s)\n", status, t.ID, t.Title, date)

			if t.Description != "" {
				fmt.Printf("    - %s\n", t.Description)
			}
		}

	case "update":
		if len(os.Args) < 3 {
			fmt.Println("error: ID task wajib diisi")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: ID harus angka")
			return
		}

		title := ""
		description := ""
		var dueDate time.Time

		if len(os.Args) >= 4 {
			title = os.Args[3]
		}
		if len(os.Args) >= 5 {
			description = os.Args[4]
		}

		updated, err := service.UpdateTask(id, title, description, dueDate)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("task berhasil diupdate:")
		fmt.Printf("%d. %s\n", updated.ID, updated.Title)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("error: ID task wajib diisi")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: ID harus angka")
			return
		}

		if err := service.DeleteTask(id); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("task berhasil dihapus")

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("error: ID task wajib diisi")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("error: ID harus berupa angka")
			return
		}

		if err := service.MarkDone(id); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("task ditandai selesai")

	default:
		fmt.Println("command tidak dikenal:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Tasker Pro - CLI Task Manager")
	fmt.Println()
	fmt.Println("Perintah:")
	fmt.Println("  add <judul>        menambahkan task baru")
	fmt.Println("  list               menampilkan semua task")
	fmt.Println("  update <id>        merubah task yang telah anda buat")
	fmt.Println("  delete <id>        menghapus task yang telah anda buat")
	fmt.Println("  done <id>          menandai task sebagai selesai")
}
