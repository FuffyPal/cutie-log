package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"gitlab.com/fluffypal/cutie-log/internal/i18n"
)

type ProcessStat struct {
	Name string
	CPU  float64
}

var dbConn *sql.DB

func initDB() *sql.DB {
	d, err := sql.Open("sqlite", "./cutie-log.db")
	if err != nil {
		log.Fatal(err)
	}

	d.Exec(`CREATE TABLE IF NOT EXISTS cpu_logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        usage REAL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`)

	d.Exec(`CREATE TABLE IF NOT EXISTS process_logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        usage REAL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`)

	return d
}

func main() {
	dbConn = initDB()
	defer dbConn.Close()

	for {
		fmt.Printf("\n---  %s ---\n", i18n.GetT("menu_title"))
		fmt.Printf("1 - %s\n", i18n.GetT("menu_monitor"))
		fmt.Printf("2 - %s\n", i18n.GetT("menu_report"))
		fmt.Printf("0 - %s\n", i18n.GetT("exit"))
		fmt.Print(i18n.GetT("choice"))

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			runLiveMonitor()
		case "2":
			showMiniReport()
		case "0":
			fmt.Println(i18n.GetT("exit_message"))
			return
		default:
			fmt.Println("Geçersiz / Invalid!")
		}
	}
}

func runLiveMonitor() {
	fmt.Printf("\n %s...\n", i18n.GetT("monitor_start"))

	for i := 0; i < 5; i++ {

		stats, err := getProcessStats()
		if err != nil {
			fmt.Println(i18n.GetT("error"), err)
			continue
		}

		sort.Slice(stats, func(i, j int) bool { return stats[i].CPU > stats[j].CPU })

		var totalCPU float64
		for _, s := range stats {
			totalCPU += s.CPU

			dbConn.Exec("INSERT INTO process_logs (name, usage) VALUES (?, ?)", s.Name, s.CPU)
		}
		dbConn.Exec("INSERT INTO cpu_logs (usage) VALUES (?)", totalCPU)

		fmt.Printf(" [%s] %s: %%%.2f\n", time.Now().Format("15:04:05"), i18n.GetT("total_cpu"), totalCPU)
		time.Sleep(5 * time.Second)
	}
}

func showMiniReport() {
	fmt.Printf("\n---  %s ---\n", i18n.GetT("report_title"))

	var maxCPU float64
	var totalRecords int
	var firstRecord string

	dbConn.QueryRow("SELECT IFNULL(MAX(usage), 0), COUNT(*) FROM cpu_logs").Scan(&maxCPU, &totalRecords)
	dbConn.QueryRow("SELECT IFNULL(MIN(timestamp), 'N/A') FROM cpu_logs").Scan(&firstRecord)

	fmt.Printf(" %s: %%%.2f\n", i18n.GetT("max_cpu_ever"), maxCPU)
	fmt.Printf(" %s %s\n", i18n.GetT("archive_start"), firstRecord)
	fmt.Printf(" %s %d\n", i18n.GetT("total_records"), totalRecords)

	fmt.Printf("\n %s:\n", i18n.GetT("top_apps_ever"))
	rows, err := dbConn.Query(`
		SELECT name, AVG(usage) as avg_usage, MAX(usage) as max_usage
		FROM process_logs 
		GROUP BY name 
		ORDER BY avg_usage DESC 
		LIMIT 5`)

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var name string
			var avg, maxVal float64
			rows.Scan(&name, &avg, &maxVal)
			fmt.Printf("   - %-20s : %s %%%.2f (%s: %%%.2f)\n", name, i18n.GetT("average"), avg, i18n.GetT("peak"), maxVal)
		}
	}

	fmt.Printf("\n%s", i18n.GetT("back_menu"))
	var dummy string
	fmt.Scanln(&dummy)
}
