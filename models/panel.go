package models

// Panel tanımı - il, ilçe, mahalle eklendi
type Panel struct {
	City           string  `json:"city"`
	District       string  `json:"district"`
	Neighborhood   string  `json:"neighborhood"`
	Brand          string  `json:"brand"`
	MaxPower       float64 `json:"max_power"`
	Pattern        []float64
	LastPatternDay int
}

// Panel listesi
var PanelList = []Panel{
	// Ankara
	{City: "Ankara", District: "Cankaya", Neighborhood: "Bahcelievler", Brand: "ABC", MaxPower: 1700, LastPatternDay: -1},
	{City: "Ankara", District: "Cankaya", Neighborhood: "Kizilay", Brand: "DEF", MaxPower: 1620, LastPatternDay: -1},
	{City: "Ankara", District: "Cankaya", Neighborhood: "Ayranci", Brand: "GHI", MaxPower: 1570, LastPatternDay: -1},
	{City: "Ankara", District: "Cankaya", Neighborhood: "Mebusevleri", Brand: "JKL", MaxPower: 1580, LastPatternDay: -1},
	{City: "Ankara", District: "Cankaya", Neighborhood: "Birlik", Brand: "MNO", MaxPower: 1490, LastPatternDay: -1},

	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Demetevler", Brand: "DEF", MaxPower: 1280, LastPatternDay: -1},
	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Ragip Tuzun", Brand: "ABC", MaxPower: 1285, LastPatternDay: -1},
	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Ata", Brand: "GHI", MaxPower: 1290, LastPatternDay: -1},
	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Serhat", Brand: "JKL", MaxPower: 1295, LastPatternDay: -1},
	{City: "Ankara", District: "Yenimahalle", Neighborhood: "Karsiyaka", Brand: "MNO", MaxPower: 1300, LastPatternDay: -1},

	{City: "Ankara", District: "Mamak", Neighborhood: "Ege", Brand: "GHI", MaxPower: 1320, LastPatternDay: -1},
	{City: "Ankara", District: "Mamak", Neighborhood: "Akdere", Brand: "DEF", MaxPower: 1325, LastPatternDay: -1},
	{City: "Ankara", District: "Mamak", Neighborhood: "Kutludugun", Brand: "ABC", MaxPower: 1330, LastPatternDay: -1},
	{City: "Ankara", District: "Mamak", Neighborhood: "Fahri Koruturk", Brand: "JKL", MaxPower: 1335, LastPatternDay: -1},
	{City: "Ankara", District: "Mamak", Neighborhood: "Gulveren", Brand: "MNO", MaxPower: 1340, LastPatternDay: -1},

	{City: "Ankara", District: "Kecioren", Neighborhood: "Etlik", Brand: "JKL", MaxPower: 1270, LastPatternDay: -1},
	{City: "Ankara", District: "Kecioren", Neighborhood: "Baglum", Brand: "DEF", MaxPower: 1275, LastPatternDay: -1},
	{City: "Ankara", District: "Kecioren", Neighborhood: "Sanatoryum", Brand: "ABC", MaxPower: 1280, LastPatternDay: -1},
	{City: "Ankara", District: "Kecioren", Neighborhood: "Osmangazi", Brand: "GHI", MaxPower: 1285, LastPatternDay: -1},
	{City: "Ankara", District: "Kecioren", Neighborhood: "Yukari Ayvali", Brand: "MNO", MaxPower: 1290, LastPatternDay: -1},

	{City: "Ankara", District: "Altindag", Neighborhood: "Aydinlikevler", Brand: "MNO", MaxPower: 1300, LastPatternDay: -1},
	{City: "Ankara", District: "Altindag", Neighborhood: "Ulubey", Brand: "DEF", MaxPower: 1305, LastPatternDay: -1},
	{City: "Ankara", District: "Altindag", Neighborhood: "Karapurcek", Brand: "ABC", MaxPower: 1310, LastPatternDay: -1},
	{City: "Ankara", District: "Altindag", Neighborhood: "Gunesevler", Brand: "JKL", MaxPower: 1315, LastPatternDay: -1},
	{City: "Ankara", District: "Altindag", Neighborhood: "Feridun Celik", Brand: "STU", MaxPower: 1320, LastPatternDay: -1},

	{City: "Ankara", District: "Sincan", Neighborhood: "Fatih", Brand: "PQR", MaxPower: 1310, LastPatternDay: -1},
	{City: "Ankara", District: "Sincan", Neighborhood: "Yenikent", Brand: "DEF", MaxPower: 1315, LastPatternDay: -1},
	{City: "Ankara", District: "Sincan", Neighborhood: "Mustafa Kemal", Brand: "JKL", MaxPower: 1320, LastPatternDay: -1},
	{City: "Ankara", District: "Sincan", Neighborhood: "Menderes", Brand: "ABC", MaxPower: 1325, LastPatternDay: -1},
	{City: "Ankara", District: "Sincan", Neighborhood: "Malazgirt", Brand: "MNO", MaxPower: 1330, LastPatternDay: -1},

	{City: "Ankara", District: "Golbasi", Neighborhood: "Safak", Brand: "STU", MaxPower: 1240, LastPatternDay: -1},
	{City: "Ankara", District: "Golbasi", Neighborhood: "Karagedik", Brand: "DEF", MaxPower: 1245, LastPatternDay: -1},
	{City: "Ankara", District: "Golbasi", Neighborhood: "Ogulbey", Brand: "PQR", MaxPower: 1250, LastPatternDay: -1},
	{City: "Ankara", District: "Golbasi", Neighborhood: "Karaali", Brand: "JKL", MaxPower: 1255, LastPatternDay: -1},
	{City: "Ankara", District: "Golbasi", Neighborhood: "Gorba", Brand: "MNO", MaxPower: 1260, LastPatternDay: -1},

	{City: "Ankara", District: "Etimesgut", Neighborhood: "Eryaman", Brand: "VWX", MaxPower: 1290, LastPatternDay: -1},
	{City: "Ankara", District: "Etimesgut", Neighborhood: "Elvankent", Brand: "DEF", MaxPower: 1295, LastPatternDay: -1},
	{City: "Ankara", District: "Etimesgut", Neighborhood: "Yavuz Selim", Brand: "ABC", MaxPower: 1300, LastPatternDay: -1},
	{City: "Ankara", District: "Etimesgut", Neighborhood: "Topcu", Brand: "JKL", MaxPower: 1305, LastPatternDay: -1},
	{City: "Ankara", District: "Etimesgut", Neighborhood: "Seyrantepe", Brand: "GHI", MaxPower: 1310, LastPatternDay: -1},

	// Istanbul
	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Moda", Brand: "DEF", MaxPower: 1400, LastPatternDay: -1},
	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Fenerbahce", Brand: "ABC", MaxPower: 1405, LastPatternDay: -1},
	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Goztepe", Brand: "JKL", MaxPower: 1410, LastPatternDay: -1},
	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Acibadem", Brand: "MNO", MaxPower: 1415, LastPatternDay: -1},
	{City: "Istanbul", District: "Kadikoy", Neighborhood: "Sogutlucesme", Brand: "PQR", MaxPower: 1420, LastPatternDay: -1},

	{City: "Istanbul", District: "Besiktas", Neighborhood: "Levent", Brand: "ABC", MaxPower: 1450, LastPatternDay: -1},
	{City: "Istanbul", District: "Besiktas", Neighborhood: "Etiler", Brand: "DEF", MaxPower: 1455, LastPatternDay: -1},
	{City: "Istanbul", District: "Besiktas", Neighborhood: "Gayrettepe", Brand: "GHI", MaxPower: 1460, LastPatternDay: -1},
	{City: "Istanbul", District: "Besiktas", Neighborhood: "Ortakoy", Brand: "JKL", MaxPower: 1465, LastPatternDay: -1},
	{City: "Istanbul", District: "Besiktas", Neighborhood: "Arnavutkoy", Brand: "MNO", MaxPower: 1470, LastPatternDay: -1},

	{City: "Istanbul", District: "Sisli", Neighborhood: "Nisantasi", Brand: "GHI", MaxPower: 1410, LastPatternDay: -1},
	{City: "Istanbul", District: "Sisli", Neighborhood: "Mecidiyekoy", Brand: "DEF", MaxPower: 1415, LastPatternDay: -1},
	{City: "Istanbul", District: "Sisli", Neighborhood: "Bomonti", Brand: "ABC", MaxPower: 1420, LastPatternDay: -1},
	{City: "Istanbul", District: "Sisli", Neighborhood: "Halaskargazi", Brand: "JKL", MaxPower: 1425, LastPatternDay: -1},
	{City: "Istanbul", District: "Sisli", Neighborhood: "Ferikoy", Brand: "MNO", MaxPower: 1430, LastPatternDay: -1},

	{City: "Istanbul", District: "Uskudar", Neighborhood: "Cengelkoy", Brand: "JKL", MaxPower: 1430, LastPatternDay: -1},
	{City: "Istanbul", District: "Uskudar", Neighborhood: "Beylerbeyi", Brand: "DEF", MaxPower: 1435, LastPatternDay: -1},
	{City: "Istanbul", District: "Uskudar", Neighborhood: "Altunizade", Brand: "ABC", MaxPower: 1440, LastPatternDay: -1},
	{City: "Istanbul", District: "Uskudar", Neighborhood: "Kisikli", Brand: "GHI", MaxPower: 1445, LastPatternDay: -1},
	{City: "Istanbul", District: "Uskudar", Neighborhood: "Salacak", Brand: "MNO", MaxPower: 1450, LastPatternDay: -1},

	{City: "Istanbul", District: "Fatih", Neighborhood: "Aksaray", Brand: "MNO", MaxPower: 1390, LastPatternDay: -1},
	{City: "Istanbul", District: "Fatih", Neighborhood: "Fener", Brand: "DEF", MaxPower: 1395, LastPatternDay: -1},
	{City: "Istanbul", District: "Fatih", Neighborhood: "Balat", Brand: "ABC", MaxPower: 1400, LastPatternDay: -1},
	{City: "Istanbul", District: "Fatih", Neighborhood: "Capa", Brand: "JKL", MaxPower: 1405, LastPatternDay: -1},
	{City: "Istanbul", District: "Fatih", Neighborhood: "Samatya", Brand: "STU", MaxPower: 1410, LastPatternDay: -1},

	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Yesilkoy", Brand: "PQR", MaxPower: 1460, LastPatternDay: -1},
	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Atakoy", Brand: "DEF", MaxPower: 1465, LastPatternDay: -1},
	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Zeytinlik", Brand: "ABC", MaxPower: 1470, LastPatternDay: -1},
	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Cevizlik", Brand: "JKL", MaxPower: 1475, LastPatternDay: -1},
	{City: "Istanbul", District: "Bakirkoy", Neighborhood: "Sakirpasa", Brand: "MNO", MaxPower: 1480, LastPatternDay: -1},

	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Cihangir", Brand: "STU", MaxPower: 1420, LastPatternDay: -1},
	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Galata", Brand: "DEF", MaxPower: 1425, LastPatternDay: -1},
	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Karakoy", Brand: "ABC", MaxPower: 1430, LastPatternDay: -1},
	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Asmalimescit", Brand: "JKL", MaxPower: 1435, LastPatternDay: -1},
	{City: "Istanbul", District: "Beyoglu", Neighborhood: "Sishane", Brand: "MNO", MaxPower: 1440, LastPatternDay: -1},

	{City: "Istanbul", District: "Maltepe", Neighborhood: "Kucukyali", Brand: "VWX", MaxPower: 1440, LastPatternDay: -1},
	{City: "Istanbul", District: "Maltepe", Neighborhood: "Fındıklı", Brand: "DEF", MaxPower: 1445, LastPatternDay: -1},
	{City: "Istanbul", District: "Maltepe", Neighborhood: "Altaycesme", Brand: "ABC", MaxPower: 1450, LastPatternDay: -1},
	{City: "Istanbul", District: "Maltepe", Neighborhood: "Cevizli", Brand: "JKL", MaxPower: 1455, LastPatternDay: -1},
	{City: "Istanbul", District: "Maltepe", Neighborhood: "Baglarbasi", Brand: "MNO", MaxPower: 1460, LastPatternDay: -1},

	// Izmir
	{City: "Izmir", District: "Konak", Neighborhood: "Alsancak", Brand: "GHI", MaxPower: 1600, LastPatternDay: -1},
	{City: "Izmir", District: "Konak", Neighborhood: "Guzelyali", Brand: "ABC", MaxPower: 1605, LastPatternDay: -1},
	{City: "Izmir", District: "Konak", Neighborhood: "Karsiyaka", Brand: "DEF", MaxPower: 1610, LastPatternDay: -1},
	{City: "Izmir", District: "Konak", Neighborhood: "Hatay", Brand: "MNO", MaxPower: 1615, LastPatternDay: -1},
	{City: "Izmir", District: "Konak", Neighborhood: "Basmane", Brand: "PQR", MaxPower: 1620, LastPatternDay: -1},

	{City: "Izmir", District: "Bornova", Neighborhood: "Kazimdirik", Brand: "ABC", MaxPower: 1620, LastPatternDay: -1},
	{City: "Izmir", District: "Bornova", Neighborhood: "Evka3", Brand: "DEF", MaxPower: 1625, LastPatternDay: -1},
	{City: "Izmir", District: "Bornova", Neighborhood: "Mevlana", Brand: "GHI", MaxPower: 1630, LastPatternDay: -1},
	{City: "Izmir", District: "Bornova", Neighborhood: "Naldoken", Brand: "JKL", MaxPower: 1635, LastPatternDay: -1},
	{City: "Izmir", District: "Bornova", Neighborhood: "Ataturk", Brand: "MNO", MaxPower: 1640, LastPatternDay: -1},

	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Bostanli", Brand: "DEF", MaxPower: 1580, LastPatternDay: -1},
	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Alaybey", Brand: "ABC", MaxPower: 1585, LastPatternDay: -1},
	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Demirkoy", Brand: "GHI", MaxPower: 1590, LastPatternDay: -1},
	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Sahil", Brand: "JKL", MaxPower: 1595, LastPatternDay: -1},
	{City: "Izmir", District: "Karsiyaka", Neighborhood: "Dedebaşı", Brand: "MNO", MaxPower: 1600, LastPatternDay: -1},

	{City: "Izmir", District: "Balcova", Neighborhood: "Teleferik", Brand: "JKL", MaxPower: 1550, LastPatternDay: -1},
	{City: "Izmir", District: "Balcova", Neighborhood: "Caferaga", Brand: "DEF", MaxPower: 1555, LastPatternDay: -1},
	{City: "Izmir", District: "Balcova", Neighborhood: "Koruturk", Brand: "ABC", MaxPower: 1560, LastPatternDay: -1},
	{City: "Izmir", District: "Balcova", Neighborhood: "Onur", Brand: "MNO", MaxPower: 1565, LastPatternDay: -1},
	{City: "Izmir", District: "Balcova", Neighborhood: "Fevzipasa", Brand: "PQR", MaxPower: 1570, LastPatternDay: -1},

	{City: "Izmir", District: "Gaziemir", Neighborhood: "Sarnic", Brand: "PQR", MaxPower: 1570, LastPatternDay: -1},
	{City: "Izmir", District: "Gaziemir", Neighborhood: "Akcam", Brand: "DEF", MaxPower: 1575, LastPatternDay: -1},
	{City: "Izmir", District: "Gaziemir", Neighborhood: "Atifbey", Brand: "ABC", MaxPower: 1580, LastPatternDay: -1},
	{City: "Izmir", District: "Gaziemir", Neighborhood: "Beyazevler", Brand: "JKL", MaxPower: 1585, LastPatternDay: -1},
	{City: "Izmir", District: "Gaziemir", Neighborhood: "Dokuz Eylul", Brand: "MNO", MaxPower: 1590, LastPatternDay: -1},

	// Bursa
	{City: "Bursa", District: "Osmangazi", Neighborhood: "Heykel", Brand: "JKL", MaxPower: 1500, LastPatternDay: -1},
	{City: "Bursa", District: "Osmangazi", Neighborhood: "Altiparmak", Brand: "DEF", MaxPower: 1505, LastPatternDay: -1},
	{City: "Bursa", District: "Osmangazi", Neighborhood: "Cekirge", Brand: "ABC", MaxPower: 1510, LastPatternDay: -1},
	{City: "Bursa", District: "Osmangazi", Neighborhood: "Cinarcik", Brand: "MNO", MaxPower: 1515, LastPatternDay: -1},
	{City: "Bursa", District: "Osmangazi", Neighborhood: "Kukurtlu", Brand: "PQR", MaxPower: 1520, LastPatternDay: -1},

	{City: "Bursa", District: "Nilufer", Neighborhood: "Gorukle", Brand: "ABC", MaxPower: 1520, LastPatternDay: -1},
	{City: "Bursa", District: "Nilufer", Neighborhood: "Ihsaniye", Brand: "DEF", MaxPower: 1525, LastPatternDay: -1},
	{City: "Bursa", District: "Nilufer", Neighborhood: "Odunluk", Brand: "GHI", MaxPower: 1530, LastPatternDay: -1},
	{City: "Bursa", District: "Nilufer", Neighborhood: "Ataevler", Brand: "JKL", MaxPower: 1535, LastPatternDay: -1},
	{City: "Bursa", District: "Nilufer", Neighborhood: "Fethiye", Brand: "MNO", MaxPower: 1540, LastPatternDay: -1},

	{City: "Bursa", District: "Yildirim", Neighborhood: "Ertugrulgazi", Brand: "DEF", MaxPower: 1510, LastPatternDay: -1},
	{City: "Bursa", District: "Yildirim", Neighborhood: "Baris", Brand: "ABC", MaxPower: 1515, LastPatternDay: -1},
	{City: "Bursa", District: "Yildirim", Neighborhood: "Yavuzselim", Brand: "GHI", MaxPower: 1520, LastPatternDay: -1},
	{City: "Bursa", District: "Yildirim", Neighborhood: "Davutdede", Brand: "JKL", MaxPower: 1525, LastPatternDay: -1},
	{City: "Bursa", District: "Yildirim", Neighborhood: "Hacivat", Brand: "MNO", MaxPower: 1530, LastPatternDay: -1},

	{City: "Bursa", District: "Inegol", Neighborhood: "Alanyurt", Brand: "GHI", MaxPower: 1490, LastPatternDay: -1},
	{City: "Bursa", District: "Inegol", Neighborhood: "Yenice", Brand: "DEF", MaxPower: 1495, LastPatternDay: -1},
	{City: "Bursa", District: "Inegol", Neighborhood: "Mesudiye", Brand: "ABC", MaxPower: 1500, LastPatternDay: -1},
	{City: "Bursa", District: "Inegol", Neighborhood: "Osmaniye", Brand: "JKL", MaxPower: 1505, LastPatternDay: -1},
	{City: "Bursa", District: "Inegol", Neighborhood: "Hamzabey", Brand: "MNO", MaxPower: 1510, LastPatternDay: -1},

	{City: "Bursa", District: "Gemlik", Neighborhood: "Kumla", Brand: "MNO", MaxPower: 1505, LastPatternDay: -1},
	{City: "Bursa", District: "Gemlik", Neighborhood: "Umuttepe", Brand: "DEF", MaxPower: 1510, LastPatternDay: -1},
	{City: "Bursa", District: "Gemlik", Neighborhood: "Cinarli", Brand: "ABC", MaxPower: 1515, LastPatternDay: -1},
	{City: "Bursa", District: "Gemlik", Neighborhood: "Fevziye", Brand: "JKL", MaxPower: 1520, LastPatternDay: -1},
	{City: "Bursa", District: "Gemlik", Neighborhood: "Balikpazari", Brand: "STU", MaxPower: 1525, LastPatternDay: -1},

	// Adana
	{City: "Adana", District: "Seyhan", Neighborhood: "Resatbey", Brand: "MNO", MaxPower: 1350, LastPatternDay: -1},
	{City: "Adana", District: "Seyhan", Neighborhood: "Barajyolu", Brand: "JKL", MaxPower: 1355, LastPatternDay: -1},
	{City: "Adana", District: "Seyhan", Neighborhood: "Cemalpaşa", Brand: "DEF", MaxPower: 1360, LastPatternDay: -1},
	{City: "Adana", District: "Seyhan", Neighborhood: "Kurtuluş", Brand: "ABC", MaxPower: 1365, LastPatternDay: -1},
	{City: "Adana", District: "Seyhan", Neighborhood: "Gazipaşa", Brand: "GHI", MaxPower: 1370, LastPatternDay: -1},

	{City: "Adana", District: "Cukurova", Neighborhood: "Turgut Ozal", Brand: "ABC", MaxPower: 1370, LastPatternDay: -1},
	{City: "Adana", District: "Cukurova", Neighborhood: "Guzelyali", Brand: "JKL", MaxPower: 1375, LastPatternDay: -1},
	{City: "Adana", District: "Cukurova", Neighborhood: "Huzurevleri", Brand: "DEF", MaxPower: 1380, LastPatternDay: -1},
	{City: "Adana", District: "Cukurova", Neighborhood: "Yurt", Brand: "PQR", MaxPower: 1385, LastPatternDay: -1},
	{City: "Adana", District: "Cukurova", Neighborhood: "Toros", Brand: "MNO", MaxPower: 1390, LastPatternDay: -1},

	{City: "Adana", District: "Yuregir", Neighborhood: "Serinevler", Brand: "DEF", MaxPower: 1340, LastPatternDay: -1},
	{City: "Adana", District: "Yuregir", Neighborhood: "Kazım Karabekir", Brand: "STU", MaxPower: 1345, LastPatternDay: -1},
	{City: "Adana", District: "Yuregir", Neighborhood: "Akincilar", Brand: "GHI", MaxPower: 1350, LastPatternDay: -1},
	{City: "Adana", District: "Yuregir", Neighborhood: "Sinanpasa", Brand: "MNO", MaxPower: 1355, LastPatternDay: -1},
	{City: "Adana", District: "Yuregir", Neighborhood: "Yavuzlar", Brand: "JKL", MaxPower: 1360, LastPatternDay: -1},

	{City: "Adana", District: "Ceyhan", Neighborhood: "Buyukmangit", Brand: "GHI", MaxPower: 1360, LastPatternDay: -1},
	{City: "Adana", District: "Ceyhan", Neighborhood: "Inonu", Brand: "JKL", MaxPower: 1365, LastPatternDay: -1},
	{City: "Adana", District: "Ceyhan", Neighborhood: "Namik Kemal", Brand: "DEF", MaxPower: 1370, LastPatternDay: -1},
	{City: "Adana", District: "Ceyhan", Neighborhood: "Kurtpinar", Brand: "MNO", MaxPower: 1375, LastPatternDay: -1},
	{City: "Adana", District: "Ceyhan", Neighborhood: "Kucukmangit", Brand: "ABC", MaxPower: 1380, LastPatternDay: -1},

	{City: "Adana", District: "Kozan", Neighborhood: "Sevkiye", Brand: "JKL", MaxPower: 1330, LastPatternDay: -1},
	{City: "Adana", District: "Kozan", Neighborhood: "Tufanpaşa", Brand: "DEF", MaxPower: 1335, LastPatternDay: -1},
	{City: "Adana", District: "Kozan", Neighborhood: "Varsaklar", Brand: "MNO", MaxPower: 1340, LastPatternDay: -1},
	{City: "Adana", District: "Kozan", Neighborhood: "Bağtepe", Brand: "PQR", MaxPower: 1345, LastPatternDay: -1},
	{City: "Adana", District: "Kozan", Neighborhood: "Mahmutlu", Brand: "GHI", MaxPower: 1350, LastPatternDay: -1},

	// Antalya
	{City: "Antalya", District: "Muratpasa", Neighborhood: "Lara", Brand: "PQR", MaxPower: 1450, LastPatternDay: -1},
	{City: "Antalya", District: "Muratpasa", Neighborhood: "Fener", Brand: "DEF", MaxPower: 1455, LastPatternDay: -1},
	{City: "Antalya", District: "Muratpasa", Neighborhood: "Guzeloba", Brand: "STU", MaxPower: 1460, LastPatternDay: -1},
	{City: "Antalya", District: "Muratpasa", Neighborhood: "Meltem", Brand: "JKL", MaxPower: 1465, LastPatternDay: -1},
	{City: "Antalya", District: "Muratpasa", Neighborhood: "Yildiz", Brand: "GHI", MaxPower: 1470, LastPatternDay: -1},

	{City: "Antalya", District: "Kepez", Neighborhood: "Varsak", Brand: "ABC", MaxPower: 1430, LastPatternDay: -1},
	{City: "Antalya", District: "Kepez", Neighborhood: "Duzlercami", Brand: "DEF", MaxPower: 1435, LastPatternDay: -1},
	{City: "Antalya", District: "Kepez", Neighborhood: "Gultepe", Brand: "GHI", MaxPower: 1440, LastPatternDay: -1},
	{City: "Antalya", District: "Kepez", Neighborhood: "Altinova", Brand: "PQR", MaxPower: 1445, LastPatternDay: -1},
	{City: "Antalya", District: "Kepez", Neighborhood: "Odabaşı", Brand: "MNO", MaxPower: 1450, LastPatternDay: -1},

	{City: "Antalya", District: "Konyaalti", Neighborhood: "Uncali", Brand: "DEF", MaxPower: 1420, LastPatternDay: -1},
	{City: "Antalya", District: "Konyaalti", Neighborhood: "Hurma", Brand: "ABC", MaxPower: 1425, LastPatternDay: -1},
	{City: "Antalya", District: "Konyaalti", Neighborhood: "Liman", Brand: "GHI", MaxPower: 1430, LastPatternDay: -1},
	{City: "Antalya", District: "Konyaalti", Neighborhood: "Molla Yusuf", Brand: "JKL", MaxPower: 1435, LastPatternDay: -1},
	{City: "Antalya", District: "Konyaalti", Neighborhood: "Altinkum", Brand: "MNO", MaxPower: 1440, LastPatternDay: -1},

	{City: "Antalya", District: "Alanya", Neighborhood: "Mahmutlar", Brand: "GHI", MaxPower: 1460, LastPatternDay: -1},
	{City: "Antalya", District: "Alanya", Neighborhood: "Kestel", Brand: "DEF", MaxPower: 1465, LastPatternDay: -1},
	{City: "Antalya", District: "Alanya", Neighborhood: "Tosmur", Brand: "STU", MaxPower: 1470, LastPatternDay: -1},
	{City: "Antalya", District: "Alanya", Neighborhood: "Oba", Brand: "JKL", MaxPower: 1475, LastPatternDay: -1},
	{City: "Antalya", District: "Alanya", Neighborhood: "Payallar", Brand: "ABC", MaxPower: 1480, LastPatternDay: -1},

	{City: "Antalya", District: "Manavgat", Neighborhood: "Side", Brand: "JKL", MaxPower: 1440, LastPatternDay: -1},
	{City: "Antalya", District: "Manavgat", Neighborhood: "Ilci", Brand: "DEF", MaxPower: 1445, LastPatternDay: -1},
	{City: "Antalya", District: "Manavgat", Neighborhood: "Kavakli", Brand: "GHI", MaxPower: 1450, LastPatternDay: -1},
	{City: "Antalya", District: "Manavgat", Neighborhood: "Sorgun", Brand: "MNO", MaxPower: 1455, LastPatternDay: -1},
	{City: "Antalya", District: "Manavgat", Neighborhood: "Aydinkent", Brand: "ABC", MaxPower: 1460, LastPatternDay: -1},

	// Konya
	{City: "Konya", District: "Meram", Neighborhood: "Yazir", Brand: "STU", MaxPower: 1550, LastPatternDay: -1},
	{City: "Konya", District: "Meram", Neighborhood: "Havzan", Brand: "DEF", MaxPower: 1555, LastPatternDay: -1},
	{City: "Konya", District: "Meram", Neighborhood: "Seyh Sadrettin", Brand: "GHI", MaxPower: 1560, LastPatternDay: -1},
	{City: "Konya", District: "Meram", Neighborhood: "Aksinne", Brand: "MNO", MaxPower: 1565, LastPatternDay: -1},
	{City: "Konya", District: "Meram", Neighborhood: "Kosova", Brand: "ABC", MaxPower: 1570, LastPatternDay: -1},

	{City: "Konya", District: "Selcuklu", Neighborhood: "Bosna Hersek", Brand: "ABC", MaxPower: 1580, LastPatternDay: -1},
	{City: "Konya", District: "Selcuklu", Neighborhood: "Sancak", Brand: "DEF", MaxPower: 1585, LastPatternDay: -1},
	{City: "Konya", District: "Selcuklu", Neighborhood: "Yazir", Brand: "JKL", MaxPower: 1590, LastPatternDay: -1},
	{City: "Konya", District: "Selcuklu", Neighborhood: "Beyhekim", Brand: "GHI", MaxPower: 1595, LastPatternDay: -1},
	{City: "Konya", District: "Selcuklu", Neighborhood: "Kosova", Brand: "MNO", MaxPower: 1600, LastPatternDay: -1},

	{City: "Konya", District: "Karatay", Neighborhood: "Aziziye", Brand: "DEF", MaxPower: 1540, LastPatternDay: -1},
	{City: "Konya", District: "Karatay", Neighborhood: "Haciveyiszade", Brand: "STU", MaxPower: 1545, LastPatternDay: -1},
	{City: "Konya", District: "Karatay", Neighborhood: "Akabe", Brand: "MNO", MaxPower: 1550, LastPatternDay: -1},
	{City: "Konya", District: "Karatay", Neighborhood: "Fevzi Cakmak", Brand: "JKL", MaxPower: 1555, LastPatternDay: -1},
	{City: "Konya", District: "Karatay", Neighborhood: "Samanpazari", Brand: "ABC", MaxPower: 1560, LastPatternDay: -1},

	{City: "Konya", District: "Beysehir", Neighborhood: "Huyuk", Brand: "GHI", MaxPower: 1520, LastPatternDay: -1},
	{City: "Konya", District: "Beysehir", Neighborhood: "Avsar", Brand: "DEF", MaxPower: 1525, LastPatternDay: -1},
	{City: "Konya", District: "Beysehir", Neighborhood: "Kusbaba", Brand: "ABC", MaxPower: 1530, LastPatternDay: -1},
	{City: "Konya", District: "Beysehir", Neighborhood: "Esence", Brand: "MNO", MaxPower: 1535, LastPatternDay: -1},
	{City: "Konya", District: "Beysehir", Neighborhood: "Yesilova", Brand: "JKL", MaxPower: 1540, LastPatternDay: -1},

	{City: "Konya", District: "Seydisehir", Neighborhood: "Kavak", Brand: "JKL", MaxPower: 1560, LastPatternDay: -1},
	{City: "Konya", District: "Seydisehir", Neighborhood: "Bostandere", Brand: "DEF", MaxPower: 1565, LastPatternDay: -1},
	{City: "Konya", District: "Seydisehir", Neighborhood: "Kizilkas", Brand: "GHI", MaxPower: 1570, LastPatternDay: -1},
	{City: "Konya", District: "Seydisehir", Neighborhood: "Ulukapı", Brand: "MNO", MaxPower: 1575, LastPatternDay: -1},
	{City: "Konya", District: "Seydisehir", Neighborhood: "Pinarbasi", Brand: "ABC", MaxPower: 1580, LastPatternDay: -1},

	// Trabzon
	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Merkez", Brand: "VWX", MaxPower: 1300, LastPatternDay: -1},
	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Kalkinma", Brand: "DEF", MaxPower: 1305, LastPatternDay: -1},
	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Yildizli", Brand: "JKL", MaxPower: 1310, LastPatternDay: -1},
	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Boztepe", Brand: "ABC", MaxPower: 1315, LastPatternDay: -1},
	{City: "Trabzon", District: "Ortahisar", Neighborhood: "Toklu", Brand: "PQR", MaxPower: 1320, LastPatternDay: -1},

	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Sogutlu", Brand: "ABC", MaxPower: 1330, LastPatternDay: -1},
	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Yildizli", Brand: "DEF", MaxPower: 1335, LastPatternDay: -1},
	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Darikli", Brand: "GHI", MaxPower: 1340, LastPatternDay: -1},
	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Caglayan", Brand: "JKL", MaxPower: 1345, LastPatternDay: -1},
	{City: "Trabzon", District: "Akcaabat", Neighborhood: "Yaylacik", Brand: "MNO", MaxPower: 1350, LastPatternDay: -1},

	{City: "Trabzon", District: "Yomra", Neighborhood: "Kasustu", Brand: "DEF", MaxPower: 1310, LastPatternDay: -1},
	{City: "Trabzon", District: "Yomra", Neighborhood: "Sancak", Brand: "STU", MaxPower: 1315, LastPatternDay: -1},
	{City: "Trabzon", District: "Yomra", Neighborhood: "Maden", Brand: "GHI", MaxPower: 1320, LastPatternDay: -1},
	{City: "Trabzon", District: "Yomra", Neighborhood: "Oymali", Brand: "MNO", MaxPower: 1325, LastPatternDay: -1},
	{City: "Trabzon", District: "Yomra", Neighborhood: "Caglayan", Brand: "ABC", MaxPower: 1330, LastPatternDay: -1},

	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Carsi", Brand: "GHI", MaxPower: 1320, LastPatternDay: -1},
	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Yalikoy", Brand: "DEF", MaxPower: 1325, LastPatternDay: -1},
	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Kemaliye", Brand: "ABC", MaxPower: 1330, LastPatternDay: -1},
	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Kucukdere", Brand: "JKL", MaxPower: 1335, LastPatternDay: -1},
	{City: "Trabzon", District: "Vakfikebir", Neighborhood: "Bostanci", Brand: "MNO", MaxPower: 1340, LastPatternDay: -1},

	{City: "Trabzon", District: "Arakli", Neighborhood: "Yaliboyu", Brand: "JKL", MaxPower: 1290, LastPatternDay: -1},
	{City: "Trabzon", District: "Arakli", Neighborhood: "Merkez", Brand: "DEF", MaxPower: 1295, LastPatternDay: -1},
	{City: "Trabzon", District: "Arakli", Neighborhood: "Kucukdere", Brand: "GHI", MaxPower: 1300, LastPatternDay: -1},
	{City: "Trabzon", District: "Arakli", Neighborhood: "Ayvadere", Brand: "MNO", MaxPower: 1305, LastPatternDay: -1},
	{City: "Trabzon", District: "Arakli", Neighborhood: "Pinaralti", Brand: "ABC", MaxPower: 1310, LastPatternDay: -1},

	// Mugla
	{City: "Mugla", District: "Bodrum", Neighborhood: "Yalikavak", Brand: "XYZ", MaxPower: 1200, LastPatternDay: -1},
	{City: "Mugla", District: "Bodrum", Neighborhood: "Bitez", Brand: "DEF", MaxPower: 1205, LastPatternDay: -1},
	{City: "Mugla", District: "Bodrum", Neighborhood: "Turgutreis", Brand: "GHI", MaxPower: 1210, LastPatternDay: -1},
	{City: "Mugla", District: "Bodrum", Neighborhood: "Gumbet", Brand: "JKL", MaxPower: 1215, LastPatternDay: -1},
	{City: "Mugla", District: "Bodrum", Neighborhood: "Gumusluk", Brand: "MNO", MaxPower: 1220, LastPatternDay: -1},

	{City: "Mugla", District: "Fethiye", Neighborhood: "Calis", Brand: "ABC", MaxPower: 1100, LastPatternDay: -1},
	{City: "Mugla", District: "Fethiye", Neighborhood: "Akarca", Brand: "DEF", MaxPower: 1105, LastPatternDay: -1},
	{City: "Mugla", District: "Fethiye", Neighborhood: "Patlangic", Brand: "JKL", MaxPower: 1110, LastPatternDay: -1},
	{City: "Mugla", District: "Fethiye", Neighborhood: "Tasyaka", Brand: "PQR", MaxPower: 1115, LastPatternDay: -1},
	{City: "Mugla", District: "Fethiye", Neighborhood: "Karagozler", Brand: "MNO", MaxPower: 1120, LastPatternDay: -1},

	{City: "Mugla", District: "Mentese", Neighborhood: "Orhaniye", Brand: "DEF", MaxPower: 1230, LastPatternDay: -1},
	{City: "Mugla", District: "Mentese", Neighborhood: "Emirbeyazıt", Brand: "ABC", MaxPower: 1235, LastPatternDay: -1},
	{City: "Mugla", District: "Mentese", Neighborhood: "Yerkesik", Brand: "JKL", MaxPower: 1240, LastPatternDay: -1},
	{City: "Mugla", District: "Mentese", Neighborhood: "Muslihittin", Brand: "GHI", MaxPower: 1245, LastPatternDay: -1},
	{City: "Mugla", District: "Mentese", Neighborhood: "Kotekli", Brand: "STU", MaxPower: 1250, LastPatternDay: -1},

	{City: "Mugla", District: "Milas", Neighborhood: "Gulluk", Brand: "GHI", MaxPower: 1250, LastPatternDay: -1},
	{City: "Mugla", District: "Milas", Neighborhood: "Selimiye", Brand: "DEF", MaxPower: 1255, LastPatternDay: -1},
	{City: "Mugla", District: "Milas", Neighborhood: "Ovakisla", Brand: "ABC", MaxPower: 1260, LastPatternDay: -1},
	{City: "Mugla", District: "Milas", Neighborhood: "Bafa", Brand: "MNO", MaxPower: 1265, LastPatternDay: -1},
	{City: "Mugla", District: "Milas", Neighborhood: "Kafaca", Brand: "PQR", MaxPower: 1270, LastPatternDay: -1},

	{City: "Mugla", District: "Ortaca", Neighborhood: "Dalyan", Brand: "JKL", MaxPower: 1210, LastPatternDay: -1},
	{City: "Mugla", District: "Ortaca", Neighborhood: "Terzialiler", Brand: "DEF", MaxPower: 1215, LastPatternDay: -1},
	{City: "Mugla", District: "Ortaca", Neighborhood: "Eskikoy", Brand: "ABC", MaxPower: 1220, LastPatternDay: -1},
	{City: "Mugla", District: "Ortaca", Neighborhood: "Okcular", Brand: "GHI", MaxPower: 1225, LastPatternDay: -1},
	{City: "Mugla", District: "Ortaca", Neighborhood: "Kemaliye", Brand: "MNO", MaxPower: 1230, LastPatternDay: -1},

	// Duzce
	{City: "Duzce", District: "Merkez", Neighborhood: "Serefiye", Brand: "DEF", MaxPower: 1150, LastPatternDay: -1},
	{City: "Duzce", District: "Merkez", Neighborhood: "Camikebir", Brand: "ABC", MaxPower: 1155, LastPatternDay: -1},
	{City: "Duzce", District: "Merkez", Neighborhood: "Kultur", Brand: "GHI", MaxPower: 1160, LastPatternDay: -1},
	{City: "Duzce", District: "Merkez", Neighborhood: "Hamidiye", Brand: "MNO", MaxPower: 1165, LastPatternDay: -1},
	{City: "Duzce", District: "Merkez", Neighborhood: "Aktoprak", Brand: "PQR", MaxPower: 1170, LastPatternDay: -1},

	{City: "Duzce", District: "Akcakoca", Neighborhood: "Osmaniye", Brand: "ABC", MaxPower: 1160, LastPatternDay: -1},
	{City: "Duzce", District: "Akcakoca", Neighborhood: "Yali", Brand: "DEF", MaxPower: 1165, LastPatternDay: -1},
	{City: "Duzce", District: "Akcakoca", Neighborhood: "Cumhuriyet", Brand: "JKL", MaxPower: 1170, LastPatternDay: -1},
	{City: "Duzce", District: "Akcakoca", Neighborhood: "Uskudar", Brand: "GHI", MaxPower: 1175, LastPatternDay: -1},
	{City: "Duzce", District: "Akcakoca", Neighborhood: "Orhangazi", Brand: "MNO", MaxPower: 1180, LastPatternDay: -1},

	{City: "Duzce", District: "Cumayeri", Neighborhood: "Cilimli", Brand: "GHI", MaxPower: 1130, LastPatternDay: -1},
	{City: "Duzce", District: "Cumayeri", Neighborhood: "Yusuflar", Brand: "DEF", MaxPower: 1135, LastPatternDay: -1},
	{City: "Duzce", District: "Cumayeri", Neighborhood: "Hamascik", Brand: "JKL", MaxPower: 1140, LastPatternDay: -1},
	{City: "Duzce", District: "Cumayeri", Neighborhood: "Esentepe", Brand: "ABC", MaxPower: 1145, LastPatternDay: -1},
	{City: "Duzce", District: "Cumayeri", Neighborhood: "Doganci", Brand: "MNO", MaxPower: 1150, LastPatternDay: -1},

	{City: "Duzce", District: "Golyaka", Neighborhood: "Yazlik", Brand: "JKL", MaxPower: 1140, LastPatternDay: -1},
	{City: "Duzce", District: "Golyaka", Neighborhood: "Cumhuriyet", Brand: "DEF", MaxPower: 1145, LastPatternDay: -1},
	{City: "Duzce", District: "Golyaka", Neighborhood: "Yazlik Mah.", Brand: "GHI", MaxPower: 1150, LastPatternDay: -1},
	{City: "Duzce", District: "Golyaka", Neighborhood: "Asar", Brand: "ABC", MaxPower: 1155, LastPatternDay: -1},
	{City: "Duzce", District: "Golyaka", Neighborhood: "Aksu", Brand: "MNO", MaxPower: 1160, LastPatternDay: -1},

	{City: "Duzce", District: "Kaynasli", Neighborhood: "Camlica", Brand: "MNO", MaxPower: 1170, LastPatternDay: -1},
	{City: "Duzce", District: "Kaynasli", Neighborhood: "Yesiltepe", Brand: "DEF", MaxPower: 1175, LastPatternDay: -1},
	{City: "Duzce", District: "Kaynasli", Neighborhood: "Beyaztarla", Brand: "JKL", MaxPower: 1180, LastPatternDay: -1},
	{City: "Duzce", District: "Kaynasli", Neighborhood: "Kumru", Brand: "GHI", MaxPower: 1185, LastPatternDay: -1},
	{City: "Duzce", District: "Kaynasli", Neighborhood: "Cay", Brand: "ABC", MaxPower: 1190, LastPatternDay: -1},
}
