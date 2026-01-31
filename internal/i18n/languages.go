package i18n

import (
	"os"
	"strings"
)

var translations = map[string]map[string]string{
	"tr": {
		"menu_title":    "🐾(づ｡◕‿‿◕｡づ) Cutie-Log Veri Merkezi🐾",
		"recent_logs":   "⸜⸝ᵕᴗᵕ⸝⸝ Son Kayıtlar (Seçmek için numara girin):",
		"menu_monitor":  "(⌐■_■) Anlık Takip (İzle & Kaydet)",
		"menu_report":   "(➧ ◕ 𝝈 ◕ ➧) Mini Rapor (Tüm Zamanlar)",
		"monitor_start": "(っ◕‿◕)っ Takip başladı (5 sn aralıkla)",
		"exit":          "(•ㅅ•) Çıkış",
		"choice":        "⸝⸝ᵕᴗᵕ⸝⸝ Seçiminiz: ",
		"report_title":  "( ✿ >◡<) MİNİ RAPOR ÖZETİ",
		"max_cpu_ever":  "(•̀ᴗ•́)و Tüm Zamanların Rekor CPU Yükü",
		"top_apps_ever": "(∩ᄑ_ᄑ)⊃━☆ En Çok Kaynak Tüketen 5 Uygulama",
		"total_cpu":     "(˵•̀ᴗ•́˵) Toplam İzlenen CPU",
		"loading":       "(๑•̀ω•́๑) arka planda çalışıyor...",
		"back_menu":     "(づ ◕‿◕ )づ Menüye dönmek için Enter'a basın...",
		"exit_message":  "づ｡◕‿◕｡づ Patiler dinlenmeye gidiyor... Hoşça kal!",
		"error":         "(╥﹏╥) Hata:",
		"archive_start": "(⏳ᐵ) Arşiv Başlangıcı:",
		"total_records": "(⚗︎) Toplam Kayıt Sayısı:",
		"average":       "(∿) Ortalama",
		"peak":          "(⛰︎) Zirve",
	},
	"en": {
		"menu_title":    "🐾(づ｡◕‿‿◕｡づ) Cutie-Log Data Center🐾",
		"recent_logs":   "⸜⸝ᵕᴗᵕ⸝⸝ Recent Logs (Enter number to select):",
		"menu_monitor":  "(⌐■_■) Real-time Monitoring (Watch & Save)",
		"menu_report":   "(➧ ◕ 𝝈 ◕ ➧) Mini Report (All-Time)",
		"monitor_start": "(っ◕‿◕)っ Monitoring started (5s intervals)",
		"exit":          "(•ㅅ•) Exit",
		"choice":        "⸝⸝ᵕᴗᵕ⸝⸝ Your choice: ",
		"report_title":  "( ✿ >◡<) MINI REPORT SUMMARY",
		"max_cpu_ever":  "(•̀ᴗ•́)و All-Time Record CPU Load",
		"top_apps_ever": "(∩ᄑ_ᄑ)⊃━☆ Top 5 Resource Consuming Apps",
		"total_cpu":     "(˵•̀ᴗ•́˵) Total Monitored CPU",
		"loading":       "(๑•̀ω•́๑) working in background...",
		"back_menu":     "(づ ◕‿◕ )づ Press Enter to return to the menu...",
		"exit_message":  "づ｡◕‿◕｡づ Paws are going to rest... Goodbye!",
		"error":         "(╥﹏╥) Error:",
		"archive_start": "(⏳ᐵ) Archive Start:",
		"total_records": "(⚗︎) Total Records:",
		"average":       "(∿) Average",
		"peak":          "(⛰︎) Peak",
	},
}

const defaultLang = "en"

func GetT(key string) string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = defaultLang
	}

	langCode := strings.Split(lang, "_")[0]
	langCode = strings.ToLower(langCode)

	if dict, ok := translations[langCode]; ok {
		if val, ok := dict[key]; ok {
			return val
		}
	}

	return translations[defaultLang][key]
}
