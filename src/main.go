package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"sort"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"gitlab.com/fluffypal/cutie-log/internal/i18n"
)

var db *sql.DB

type ProcessStat struct {
	Name string
	CPU  float64
}

func initDB() {
	var err error
	// "sqlite" sürücüsü glebarez kütüphanesi içindir
	db, err = sql.Open("sqlite", "./cutie-log.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, timestamp TEXT, app_name TEXT, cpu_usage REAL)")
	if err != nil {
		log.Fatal(err)
	}
}

func collectAndSave() {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		// İşletim sistemine göre ilgili collector_*.go dosyasındaki fonksiyon çalışır
		currentStats, err := getProcessStats()
		if err != nil {
			continue
		}

		sort.Slice(currentStats, func(i, j int) bool {
			return currentStats[i].CPU > currentStats[j].CPU
		})

		now := time.Now().Format("15:04:05")
		for i := 0; i < 5 && i < len(currentStats); i++ {
			db.Exec("INSERT INTO logs (timestamp, app_name, cpu_usage) VALUES (?, ?, ?)",
				now, currentStats[i].Name, currentStats[i].CPU)
		}
	}
}

func showMenu() {
	for {
		fmt.Printf("\n--- %s ---\n", i18n.GetT("menu_title"))
		rows, _ := db.Query("SELECT MIN(id), timestamp, SUM(cpu_usage) FROM logs GROUP BY timestamp ORDER BY id DESC LIMIT 10")

		var timestamps []string
		i := 1
		fmt.Println(i18n.GetT("recent_logs"))
		for rows.Next() {
			var id int
			var ts string
			var totalCPU float64
			rows.Scan(&id, &ts, &totalCPU)
			fmt.Printf("%d. [%s] | %s: %%%.2f\n", i, ts, i18n.GetT("total_cpu"), totalCPU)
			timestamps = append(timestamps, ts)
			i++
		}
		rows.Close()

		fmt.Printf("0. %s\n", i18n.GetT("exit"))
		fmt.Print(i18n.GetT("choice"))

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		} else if choice > 0 && choice <= len(timestamps) {
			selectedTime := timestamps[choice-1]
			fmt.Printf("\n--- [%s] ---\n", selectedTime)
			detailRows, _ := db.Query("SELECT app_name, cpu_usage FROM logs WHERE timestamp = ?", selectedTime)
			for detailRows.Next() {
				var name string
				var cpu float64
				detailRows.Scan(&name, &cpu)
				fmt.Printf("   App: %-20s | CPU: %%%.2f\n", name, cpu)
			}
			detailRows.Close()
			fmt.Printf("\n%s", i18n.GetT("back_menu"))
			var dummy string
			fmt.Scanln(&dummy) // Enter bekle
		}
	}
}

func main() {
	initDB()
	// Dinamik yükleme mesajı
	fmt.Printf("🐾 Cutie-Log (%d cores) %s\n", runtime.NumCPU(), i18n.GetT("loading"))

	go collectAndSave()
	showMenu()
}
