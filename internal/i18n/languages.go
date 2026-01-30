package i18n

import (
	"os"
	"strings"
)

var translations = map[string]map[string]string{
	"tr": {
		"menu_title":  "🐾 Cutie-Log Veri Merkezi",
		"recent_logs": "Son Kayıtlar (Seçmek için numara girin):",
		"exit":        "Çıkış",
		"choice":      "Seçiminiz: ",
		"total_cpu":   "Toplam İzlenen CPU",
		"loading":     "arka planda çalışıyor...",
		"back_menu":   "Ana menüye dönmek için Enter'a basın.",
	},
	"en": {
		"menu_title":  "🐾 Cutie-Log Data Center",
		"recent_logs": "Recent Logs (Enter number to select):",
		"exit":        "Exit",
		"choice":      "Your choice: ",
		"total_cpu":   "Total Monitored CPU",
		"loading":     "working in background...",
		"back_menu":   "Press Enter to return to main menu.",
	},
}

const defaultLang = "en"

func GetT(key string) string {
	// Sistem dilini al (NixOS/Docker ortamından)
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = defaultLang
	}

	langCode := strings.Split(lang, "_")[0]
	langCode = strings.ToLower(langCode)

	// Dil var mı kontrol et, yoksa İngilizce'ye dön
	if dict, ok := translations[langCode]; ok {
		if val, ok := dict[key]; ok {
			return val
		}
	}

	// Almanca (de) gibi olmayan bir dil gelirse burası çalışır
	return translations[defaultLang][key]
}
