# 🐾 Cutie-Log 🐾

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-003B57?style=for-the-badge&logo=sqlite&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux-brightgreen?style=for-the-badge)

---

### 🌐 Language Selection / Dil Seçimi
- [🇬🇧 English Version](#-english)
- [🇹🇷 Türkçe Versiyon](#-türkçe)

---

## 🇬🇧 English

**Cutie-Log** is a lightweight, cross-platform system resource monitor that doesn't just watch your CPU—it remembers it! Using a native Go implementation, it tracks process loads and stores them in a safe, CGO-free SQLite database.

### ✨ Features
* **Dual-Core Heart:** Specialized collectors for both Windows (NT API) and Linux (`/proc` filesystem).
* **Safe Storage:** Powered by `glebarez/go-sqlite`, meaning zero C dependencies and 100% portability.
* **Clean Architecture:** Modular internal structure for easy translations and expansions.
* **GitLab CI/CD Ready:** Automatically builds and tests for both OS targets on every push.

### 🚀 Quick Start
1.  **Build for your OS:**
    ```bash
        go build -o cutie-log ./src
            ```
            2.  **Run it:**
                ```bash
                    ./cutie-log
                        ```

                        ---

                        ## 🇹🇷 Türkçe

                        **Cutie-Log**, sistem kaynaklarını izleyen ama bunu yaparken yormayan, minnoş ama disiplinli bir performans takip aracıdır. CPU yükünü takip eder ve tüm verileri "saf Go" (CGO-free) SQLite veritabanına titizlikle kaydeder.

                        ### ✨ Özellikler
                        * **Çift Motorlu Yapı:** Windows (NT API) ve Linux (`/proc`) için özel olarak optimize edilmiş veri toplayıcılar.
                        * **Güvenli Kayıt:** `glebarez/go-sqlite` altyapısı sayesinde hiçbir C kütüphanesine ihtiyaç duymadan, her yerde çalışabilen taşınabilir yapı.
                        * **Temiz Kod:** Modüler iç yapı sayesinde kolayca çeviri eklenebilir ve geliştirilebilir.
                        * **GitLab CI/CD Entegre:** Her `push` yaptığında hem Linux hem Windows için otomatik olarak test edilir ve derlenir.

                        ### 🚀 Hızlı Başlangıç
                        1.  **Kendi sistemin için derle:**
                            ```bash
                                go build -o cutie-log ./src
                                    ```
                                    2.  **Çalıştır:**
                                        ```bash
                                            ./cutie-log
                                                ```

                                                ---

                                                > _Developed with 🐾 by fluffypal_
                                                