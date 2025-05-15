# 📡 nginx_monitoring – Gerçek Zamanlı NGINX Trafik İzleme

**nginx_monitoring**, NGINX web sunucularını gerçek zamanlı olarak izlemek amacıyla geliştirilmiş, yüksek performanslı bir izleme aracıdır. Proje, Go diliyle yazılmış olup, Prometheus ile veri toplayarak sistem trafiğini ölçer ve Docker kullanılarak kolayca dağıtılabilir hale getirilmiştir.

---

## 🎯 Projenin Amacı

- 🌐 NGINX üzerinde gerçekleşen HTTP trafiğini anlık olarak izlemek
- 📈 Prometheus exporter olarak çalışarak NGINX loglarını metriklere dönüştürmek
- 🔧 Kurulumu kolaylaştırmak için tüm bileşenleri Docker üzerinden yönetmek
- 📊 Grafana gibi araçlarla görsel analiz yapılmasını sağlamak (isteğe bağlı)

---

## ⚙️ Kullanılan Teknolojiler

| Teknoloji      | Açıklama                                   |
|----------------|--------------------------------------------|
| **Go (Golang)**| Performans odaklı sistem servisi           |
| **NGINX**      | İzlenen HTTP sunucusu                      |
| **Prometheus** | Metrikleri toplayan zaman serisi veritabanı|
| **Docker**     | Ortamları kapsüllenmiş çalıştırma altyapısı|
| **Grafana***   | (İsteğe bağlı) grafiksel analiz aracı      |

---
