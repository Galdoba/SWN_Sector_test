package main

type FullName struct {
	firstName string
	surName   string
	female    bool
	test      bool
}

func NewFullName(isfemale bool, names ...string) *FullName {
	fn := FullName{}
	fn.female = isfemale
	parts := names
	site := ""
	for i := range parts {
		if i == 0 {
			site = parts[0]
		}
		if i == 1 {
			fn.firstName = parts[1]
			if fn.firstName == "" {
				if !fn.female {
					fn.firstName = GiveName(site, "Male", -1)
				} else {
					fn.firstName = GiveName(site, "Female", -1)
				}
			}
		}
		if i == 2 {
			fn.surName = parts[2]
			if fn.surName == "" {
				fn.surName = GiveName(site, "Sur", -1)
			}
		}
	}
	return &fn
}

func (fn *FullName) setFirstName(fName string) {
	fn.firstName = fName
}

func (fn *FullName) setSurName(sName string) {
	fn.surName = sName
}

func (fn *FullName) switchGender() {
	if fn.female {
		fn.female = false
	} else {
		fn.female = true
	}
}

func (fn *FullName) FirstName() string {
	return fn.firstName
}
func (fn *FullName) SurName() string {
	return fn.surName
}
func (fn *FullName) FullName() string {
	return fn.firstName + " " + fn.surName
}

func arabicMaleNames() []string {
	names := []string{
		"Aamir",
		"Ayub",
		"Binyamin",
		"Efraim",
		"Ibrahim",
		"Ilyas",
		"Ismail",
		"Jibril",
		"Jumanah",
		"Kazi",
		"Lut",
		"Matta",
		"Mohammed",
		"Mubarak",
		"Mustafa",
		"Nazir",
		"Rahim",
		"Reza",
		"Sharif",
		"Taimur",
		"Usman",
		"Yakub",
		"Yusuf",
		"Zakariya",
		"Zubair",
	}
	return names
}

func arabicFemaleNames() []string {
	names := []string{
		"Aisha",
		"Alimah",
		"Badia",
		"Bisharah",
		"Chanda",
		"Daliya",
		"Fatimah",
		"Ghania",
		"Halah",
		"Kaylah",
		"Khayrah",
		"Layla",
		"Mina",
		"Munisa",
		"Mysha",
		"Naimah",
		"Nissa",
		"Nura",
		"Parveen",
		"Rana",
		"Shalha",
		"Suhira",
		"Tahirah",
		"Yasmin",
		"Zulehka",
	}
	return names
}

func arabicSurNames() []string {
	names := []string{
		"Abdel",
		"Awad",
		"Dahhak",
		"Essa",
		"Hanna",
		"Harbi",
		"Hassan",
		"Isa",
		"Kasim",
		"Katib",
		"Khalil",
		"Malik",
		"Mansoor",
		"Mazin",
		"Musa",
		"Najeeb",
		"Namari",
		"Naser",
		"Rahman",
		"Rasheed",
		"Saleh",
		"Salim",
		"Shadi",
		"Sulaiman",
		"Tabari",
	}
	return names
}

func arabicPlaceNames() []string {
	names := []string{
		"Adan",
		"Ahsa",
		"Andalus",
		"Asmara",
		"Asqlan",
		"Baqubah",
		"Basit",
		"Baysan",
		"Baytlahm",
		"Bursaid",
		"Dahilah",
		"Darasalam",
		"Dawhah",
		"Ganin",
		"Gebal",
		"Gibuti",
		"Giddah",
		"Harmah",
		"Hartum",
		"Hibah",
		"Hims",
		"Hubar",
		"Karbala",
		"Kut",
		"Lacant",
		"Magrit",
		"Masqat",
		"Misr",
		"Muruni",
		"Qabis",
		"Qina",
		"Rabat",
		"Ramlah",
		"Riyadh",
		"Sabtah",
		"Salalah",
		"Sana",
		"Sinqit",
		"Suqutrah",
		"Sur",
		"Tabuk",
		"Tangah",
		"Tarifah",
		"Tarrakunah",
		"Tisit",
		"Uman",
		"Urdunn",
		"Wasqah",
		"Yaburah",
		"Yaman",
	}
	return names
}

func chineseMaleNames() []string {
	names := []string{
		"Aiguo",
		"Bohai",
		"Chao",
		"Dai",
		"Dawei",
		"Duyi",
		"Fa",
		"Fu",
		"Gui",
		"Hong",
		"Jianyu",
		"Kang",
		"Li",
		"Niu",
		"Peng",
		"Quan",
		"Ru",
		"Shen",
		"Shi",
		"Song",
		"Tao",
		"Xue",
		"Yi",
		"Yuan",
		"Zian",
	}
	return names
}

func chineseFemaleNames() []string {
	names := []string{
		"Biyu",
		"Changying",
		"Daiyu",
		"Huidai",
		"Huiliang",
		"Jia",
		"Jingfei",
		"Lan",
		"Liling",
		"Liu",
		"Meili",
		"Niu",
		"Peizhi",
		"Qiao",
		"Qing",
		"Ruolan",
		"Shu",
		"Suyin",
		"Ting",
		"Xia",
		"Xiaowen",
		"Xiulan",
		"Ya",
		"Ying",
		"Zhilan",
	}
	return names
}

func chineseSurNames() []string {
	names := []string{
		"Bai",
		"Cao",
		"Chen",
		"Cui",
		"Ding",
		"Du",
		"Fang",
		"Fu",
		"Guo",
		"Han",
		"Hao",
		"Huang",
		"Lei",
		"Li",
		"Liang",
		"Liu",
		"Long",
		"Song",
		"Tan",
		"Tang",
		"Wang",
		"Wu",
		"Xing",
		"Yang",
		"Zhang",
	}
	return names
}

func chinesePlaceNames() []string {
	names := []string{
		"Andong",
		"Anqing",
		"Anshan",
		"Chaoyang",
		"Chaozhou",
		"Chifeng",
		"Dalian",
		"Dunhuang",
		"Fengjia",
		"Fengtian",
		"Fuliang",
		"Fushun",
		"Gansu",
		"Ganzhou",
		"Guizhou",
		"Hotan",
		"Hunan",
		"Jinan",
		"Jingdezhen",
		"Jinxi",
		"Jinzhou",
		"Kunming",
		"Liaoning",
		"Linyi",
		"Lushun",
		"Luzhou",
		"Ningxia",
		"Pingxiang",
		"Pizhou",
		"Qidong",
		"Qingdao",
		"Qinghai",
		"Rehe",
		"Shanxi",
		"Taiyuan",
		"Tengzhou",
		"Urumqi",
		"Weifang",
		"Wugang",
		"Wuxi",
		"Xiamen",
		"Xian",
		"Xikang",
		"Xining",
		"Xinjiang",
		"Yidu",
		"Yingkou",
		"Yuxi",
		"Zigong",
		"Zoige",
	}
	return names
}

func englishMaleNames() []string {
	names := []string{
		"Adam",
		"Albert",
		"Alfred",
		"Allan",
		"Archibald",
		"Arthur",
		"Basil",
		"Charles",
		"Colin",
		"Donald",
		"Douglas",
		"Edgar",
		"Edmund",
		"Edward",
		"George",
		"Harold",
		"Henry",
		"Ian",
		"James",
		"John",
		"Lewis",
		"Oliver",
		"Philip",
		"Richard",
		"William",
	}
	return names
}

func englishFemaleNames() []string {
	names := []string{
		"Abigail",
		"Anne",
		"Beatrice",
		"Blanche",
		"Catherine",
		"Charlotte",
		"Claire",
		"Eleanor",
		"Elizabeth",
		"Emily",
		"Emma",
		"Georgia",
		"Harriet",
		"Joan",
		"Judy",
		"Julia",
		"Lucy",
		"Lydia",
		"Margaret",
		"Mary",
		"Molly",
		"Nora",
		"Rosie",
		"Sarah",
		"Victoria",
	}
	return names
}

func englishSurNames() []string {
	names := []string{
		"Barker",
		"Brown",
		"Butler",
		"Carter",
		"Chapman",
		"Collins",
		"Cook",
		"Davies",
		"Gray",
		"Green",
		"Harris",
		"Jackson",
		"Jones",
		"Lloyd",
		"Miller",
		"Roberts",
		"Smith",
		"Taylor",
		"Thomas",
		"Turner",
		"Watson",
		"White",
		"Williams",
		"Wood",
		"Young",
	}
	return names
}

func englishPlaceNames() []string {
	names := []string{
		"Aldington",
		"Appleton",
		"Ashdon",
		"Berwick",
		"Bramford",
		"Brimstage",
		"Carden",
		"Churchill",
		"Clifton",
		"Colby",
		"Copford",
		"Cromer",
		"Davenham",
		"Dersingham",
		"Doverdale",
		"Elsted",
		"Ferring",
		"Gissing",
		"Heydon",
		"Holt",
		"Hunston",
		"Hutton",
		"Inkberrow",
		"Inworth",
		"Isfield",
		"Kedington",
		"Latchford",
		"Leigh",
		"Leighton",
		"Maresfield",
		"Markshall",
		"Netherpool",
		"Newton",
		"Oxton",
		"Preston",
		"Ridley",
		"Rochford",
		"Seaford",
		"Selsey",
		"Stanton",
		"Stockham",
		"Stoke",
		"Sutton",
		"Thakeham",
		"Thetford",
		"Thorndon",
		"Ulting",
		"Upton",
		"Westhorpe",
		"Worcester",
	}
	return names
}

func greekMaleNames() []string {
	names := []string{
		"Alexander",
		"Alexius",
		"Anastasius",
		"Christodoulos",
		"Christos",
		"Damian",
		"Dimitris",
		"Dysmas",
		"Elias",
		"Giorgos",
		"Ioannis",
		"Konstantinos",
		"Lambros",
		"Leonidas",
		"Marcos",
		"Miltiades",
		"Nestor",
		"Nikos",
		"Orestes",
		"Petros",
		"Simon",
		"Stavros",
		"Theodore",
		"Vassilios",
		"Yannis",
	}
	return names
}

func greekFemaleNames() []string {
	names := []string{
		"Alexandra",
		"Amalia",
		"Callisto",
		"Charis",
		"Chloe",
		"Dorothea",
		"Elena",
		"Eudoxia",
		"Giada",
		"Helena",
		"Ioanna",
		"Lydia",
		"Melania",
		"Melissa",
		"Nika",
		"Nikolina",
		"Olympias",
		"Philippa",
		"Phoebe",
		"Sophia",
		"Theodora",
		"Valentina",
		"Valeria",
		"Yianna",
		"Zoe",
	}
	return names
}

func greekSurNames() []string {
	names := []string{
		"Andreas",
		"Argyros",
		"Dimitriou",
		"Floros",
		"Gavras",
		"Ioannidis",
		"Katsaros",
		"Kyrkos",
		"Leventis",
		"Makris",
		"Metaxas",
		"Nikolaidis",
		"Pallis",
		"Pappas",
		"Petrou",
		"Raptis",
		"Simonides",
		"Spiros",
		"Stavros",
		"Stephanidis",
		"Stratigos",
		"Terzis",
		"Theodorou",
		"Vasiliadis",
		"Yannakakis",
	}
	return names
}

func greekPlaceNames() []string {
	names := []string{
		"Adramyttion",
		"Ainos",
		"Alikarnassos",
		"Avydos",
		"Dakia",
		"Dardanos",
		"Dekapoli",
		"Dodoni",
		"Efesos",
		"Efstratios",
		"Elefsina",
		"Ellada",
		"Epidavros",
		"Erymanthos",
		"Evripos",
		"Gavdos",
		"Gytheio",
		"Ikaria",
		"Ilios",
		"Illyria",
		"Iraia",
		"Irakleio",
		"Isminos",
		"Ithaki",
		"Kadmeia",
		"Kallisto",
		"Katerini",
		"Kithairon",
		"Kydonia",
		"Lakonia",
		"Leros",
		"Lesvos",
		"Limnos",
		"Lykia",
		"Megara",
		"Messene",
		"Milos",
		"Nikaia",
		"Orontis",
		"Parnasos",
		"Petro",
		"Samos",
		"Syros",
		"Thapsos",
		"Thessalia",
		"Thira",
		"Thiva",
		"Varvara",
		"Voiotia",
		"Vyvlos",
	}
	return names
}

func indianMaleNames() []string {
	names := []string{
		"Amrit",
		"Ashok",
		"Chand",
		"Dinesh",
		"Gobind",
		"Harinder",
		"Jagdish",
		"Johar",
		"Kurien",
		"Lakshman",
		"Madhav",
		"Mahinder",
		"Mohal",
		"Narinder",
		"Nikhil",
		"Omrao",
		"Prasad",
		"Pratap",
		"Ranjit",
		"Sanjay",
		"Shankar",
		"Thakur",
		"Vijay",
		"Vipul",
		"Yash",
	}
	return names
}

func indianFemaleNames() []string {
	names := []string{
		"Amala",
		"Asha",
		"Chandra",
		"Devika",
		"Esha",
		"Gita",
		"Indira",
		"Indrani",
		"Jaya",
		"Jayanti",
		"Kiri",
		"Lalita",
		"Malati",
		"Mira",
		"Mohana",
		"Neela",
		"Nita",
		"Rajani",
		"Sarala",
		"Sarika",
		"Sheela",
		"Sunita",
		"Trishna",
		"Usha",
		"Vasanta",
	}
	return names
}

func indianSurNames() []string {
	names := []string{
		"Achari",
		"Banerjee",
		"Bhatnagar",
		"Bose",
		"Chauhan",
		"Chopra",
		"Das",
		"Dutta",
		"Gupta",
		"Johar",
		"Kapoor",
		"Mahajan",
		"Malhotra",
		"Mehra",
		"Nehru",
		"Patil",
		"Rao",
		"Saxena",
		"Shah",
		"Sharma",
		"Singh",
		"Trivedi",
		"Venkatesan",
		"Verma",
		"Yadav",
	}
	return names
}

func indianPlaceNames() []string {
	names := []string{
		"Ahmedabad",
		"Alipurduar",
		"Alubari",
		"Anjanadri",
		"Ankleshwar",
		"Balarika",
		"Bhanuja",
		"Bhilwada",
		"Brahmaghosa",
		"Bulandshahar",
		"Candrama",
		"Chalisgaon",
		"Chandragiri",
		"Charbagh",
		"Chayanka",
		"Chittorgarh",
		"Dayabasti",
		"Dikpala",
		"Ekanga",
		"Gandhidham",
		"Gollaprolu",
		"Grahisa",
		"Guwahati",
		"Haridasva",
		"Indraprastha",
		"Jaisalmer",
		"Jharonda",
		"Kadambur",
		"Kalasipalyam",
		"Karnataka",
		"Kutchuhery",
		"Lalgola",
		"Mainaguri",
		"Nainital",
		"Nandidurg",
		"Narayanadri",
		"Panipat",
		"Panjagutta",
		"Pathankot",
		"Pathardih",
		"Porbandar",
		"Rajasthan",
		"Renigunta",
		"Sewagram",
		"Shakurbasti",
		"Siliguri",
		"Sonepat",
		"Teliwara",
		"Tinpahar",
		"Villivakkam",
	}
	return names
}

func japaneseMaleNames() []string {
	names := []string{
		"Akira",
		"Daisuke",
		"Fukashi",
		"Goro",
		"Hiro",
		"Hiroya",
		"Hotaka",
		"Katsu",
		"Katsuto",
		"Keishuu",
		"Kyuuto",
		"Mikiya",
		"Mitsunobu",
		"Mitsuru",
		"Naruhiko",
		"Nobu",
		"Shigeo",
		"Shigeto",
		"Shou",
		"Shuji",
		"Takaharu",
		"Teruaki",
		"Tetsushi",
		"Tsukasa",
		"Yasuharu",
	}
	return names
}

func japaneseFemaleNames() []string {
	names := []string{
		"Aemi",
		"Airi",
		"Ako",
		"Ayu",
		"Chikaze",
		"Eriko",
		"Hina",
		"Kaori",
		"Keiko",
		"Kyouka",
		"Mayumi",
		"Miho",
		"Namiko",
		"Natsu",
		"Nobuko",
		"Rei",
		"Ririsa",
		"Sakimi",
		"Shihoko",
		"Shika",
		"Tsukiko",
		"Tsuzune",
		"Yoriko",
		"Yorimi",
		"Yoshiko",
	}
	return names
}

func japaneseSurNames() []string {
	names := []string{
		"Abe",
		"Arakaki",
		"Endo",
		"Fujiwara",
		"Goto",
		"Ito",
		"Kikuchi",
		"Kinjo",
		"Kobayashi",
		"Koga",
		"Komatsu",
		"Maeda",
		"Nakamura",
		"Narita",
		"Ochi",
		"Oshiro",
		"Saito",
		"Sakamoto",
		"Sato",
		"Suzuki",
		"Takahashi",
		"Tanaka",
		"Watanabe",
		"Yamamoto",
		"Yamasaki",
	}
	return names
}

func japanesePlaceNames() []string {
	names := []string{
		"Bando",
		"Chikuma",
		"Chikusei",
		"Chino",
		"Hitachi",
		"Hitachinaka",
		"Hitachiomiya",
		"Hitachiota",
		"Iida",
		"Iiyama",
		"Ina",
		"Inashiki",
		"Ishioka",
		"Itako",
		"Kamisu",
		"Kasama",
		"Kashima",
		"Kasumigaura",
		"Kitaibaraki",
		"Kiyose",
		"Koga",
		"Komagane",
		"Komoro",
		"Matsumoto",
		"Mito",
		"Mitsukaido",
		"Moriya",
		"Nagano",
		"Naka",
		"Nakano",
		"Ogi",
		"Okaya",
		"Omachi",
		"Ryugasaki",
		"Saku",
		"Settsu",
		"Shimotsuma",
		"Shiojiri",
		"Suwa",
		"Suzaka",
		"Takahagi",
		"Takeo",
		"Tomi",
		"Toride",
		"Tsuchiura",
		"Tsukuba",
		"Ueda",
		"Ushiku",
		"Yoshikawa",
		"Yuki",
	}
	return names
}

func latinMaleNames() []string {
	names := []string{
		"Agrippa",
		"Appius",
		"Aulus",
		"Caeso",
		"Decimus",
		"Faustus",
		"Gaius",
		"Gnaeus",
		"Hostus",
		"Lucius",
		"Mamercus",
		"Manius",
		"Marcus",
		"Mettius",
		"Nonus",
		"Numerius",
		"Opiter",
		"Paulus",
		"Proculus",
		"Publius",
		"Quintus",
		"Servius",
		"Tiberius",
		"Titus",
		"Volescus",
	}
	return names
}

func latinFemaleNames() []string {
	names := []string{
		"Appia",
		"Aula",
		"Caesula",
		"Decima",
		"Fausta",
		"Gaia",
		"Gnaea",
		"Hosta",
		"Lucia",
		"Maio",
		"Marcia",
		"Maxima",
		"Mettia",
		"Nona",
		"Numeria",
		"Octavia",
		"Postuma",
		"Prima",
		"Procula",
		"Septima",
		"Servia",
		"Tertia",
		"Tiberia",
		"Titia",
		"Vibia",
	}
	return names
}

func latinSurNames() []string {
	names := []string{
		"Antius",
		"Aurius",
		"Barbatius",
		"Calidius",
		"Cornelius",
		"Decius",
		"Fabius",
		"Flavius",
		"Galerius",
		"Horatius",
		"Julius",
		"Juventius",
		"Licinius",
		"Marius",
		"Minicius",
		"Nerius",
		"Octavius",
		"Pompeius",
		"Quinctius",
		"Rutilius",
		"Sextius",
		"Titius",
		"Ulpius",
		"Valerius",
		"Vitellius",
	}
	return names
}

func latinPlaceNames() []string {
	names := []string{
		"Abilia",
		"Alsium",
		"Aquileia",
		"Argentoratum",
		"Ascrivium",
		"Asculum",
		"Attalia",
		"Barium",
		"Batavorum",
		"Belum",
		"Bobbium",
		"Brigantium",
		"Burgodunum",
		"Camulodunum",
		"Clausentum",
		"Corduba",
		"Coriovallum",
		"Durucobrivis",
		"Eboracum",
		"Emona",
		"Florentia",
		"Lactodurum",
		"Lentia",
		"Lindum",
		"Londinium",
		"Lucus",
		"Lugdunum",
		"Mediolanum",
		"Novaesium",
		"Patavium",
		"Pistoria",
		"Pompeii",
		"Raurica",
		"Rigomagus",
		"Roma",
		"Salernum",
		"Salona",
		"Segovia",
		"Sirmium",
		"Spalatum",
		"Tarraco",
		"Treverorum",
		"Verulamium",
		"Vesontio",
		"Vetera",
		"Vindelicorum",
		"Vindobona",
		"Vinovia",
		"Viroconium",
		"Volubilis",
	}
	return names
}

func negerianMaleNames() []string {
	names := []string{
		"Adesegun",
		"Akintola",
		"Amabere",
		"Arikawe",
		"Asagwara",
		"Chidubem",
		"Chinedu",
		"Chiwetei",
		"Damilola",
		"Esangbedo",
		"Ezenwoye",
		"Folarin",
		"Genechi",
		"Idowu",
		"Kelechi",
		"Ketanndu",
		"Melubari",
		"Nkanta",
		"Obafemi",
		"Olatunde",
		"Olumide",
		"Tombari",
		"Udofia",
		"Uyoata",
		"Uzochi",
	}
	return names
}

func negerianFemaleNames() []string {
	names := []string{
		"Abike",
		"Adesuwa",
		"Adunola",
		"Anguli",
		"Arewa",
		"Asari",
		"Bisola",
		"Chioma",
		"Eduwa",
		"Emilohi",
		"Fehintola",
		"Folasade",
		"Mahparah",
		"Minika",
		"Nkolika",
		"Nkoyo",
		"Nuanae",
		"Obioma",
		"Olafemi",
		"Shanumi",
		"Sominabo",
		"Suliat",
		"Tariere",
		"Temedire",
		"Yemisi",
	}
	return names
}

func negerianSurNames() []string {
	names := []string{
		"Adegboye",
		"Adeniyi",
		"Adeyeku",
		"Adunola",
		"Agbaje",
		"Akpan",
		"Akpehi",
		"Aliki",
		"Asuni",
		"Babangida",
		"Ekim",
		"Ezeiruaku",
		"Fabiola",
		"Fasola",
		"Nwokolo",
		"Nzeocha",
		"Ojo",
		"Okonkwo",
		"Okoye",
		"Olaniyan",
		"Olawale",
		"Olumese",
		"Onajobi",
		"Soyinka",
		"Yamusa",
	}
	return names
}

func negerianPlaceNames() []string {
	names := []string{
		"Abadan",
		"Ador",
		"Agatu",
		"Akamkpa",
		"Akpabuyo",
		"Ala",
		"Askira",
		"Bakassi",
		"Bama",
		"Bayo",
		"Bekwara",
		"Biase",
		"Boki",
		"Buruku",
		"Calabar",
		"Chibok",
		"Damboa",
		"Dikwa",
		"Etung",
		"Gboko",
		"Gubio",
		"Guzamala",
		"Gwoza",
		"Hawul",
		"Ikom",
		"Jere",
		"Kalabalge",
		"Katsina",
		"Knoduga",
		"Konshishatse",
		"Kukawa",
		"Kwande",
		"Kwayakusar",
		"Logo",
		"Mafa",
		"Makurdi",
		"Nganzai",
		"Obanliku",
		"Obi",
		"Obubra",
		"Obudu",
		"Odukpani",
		"Ogbadibo",
		"Ohimini",
		"Okpokwu",
		"Otukpo",
		"Shani",
		"Ugep",
		"Vandeikya",
		"Yala",
	}
	return names
}

func russianMaleNames() []string {
	names := []string{
		"Aleksandr",
		"Andrei",
		"Arkady",
		"Boris",
		"Dmitri",
		"Dominik",
		"Grigory",
		"Igor",
		"Ilya",
		"Ivan",
		"Kiril",
		"Konstantin",
		"Leonid",
		"Nikolai",
		"Oleg",
		"Pavel",
		"Petr",
		"Sergei",
		"Stepan",
		"Valentin",
		"Vasily",
		"Viktor",
		"Yakov",
		"Yegor",
		"Yuri",
	}
	return names
}

func russianFemaleNames() []string {
	names := []string{
		"Aleksandra",
		"Anastasia",
		"Anja",
		"Catarina",
		"Devora",
		"Dima",
		"Ekaterina",
		"Eva",
		"Irina",
		"Karolina",
		"Katlina",
		"Kira",
		"Ludmilla",
		"Mara",
		"Nadezdha",
		"Nastassia",
		"Natalya",
		"Oksana",
		"Olena",
		"Olga",
		"Sofia",
		"Svetlana",
		"Tatyana",
		"Vilma",
		"Yelena",
	}
	return names
}

func russianSurNames() []string {
	names := []string{
		"Abelev",
		"Bobrikov",
		"Chemerkin",
		"Gogunov",
		"Gurov",
		"Iltchenko",
		"Kavelin",
		"Komarov",
		"Korovin",
		"Kurnikov",
		"Lebedev",
		"Litvak",
		"Mekhdiev",
		"Muraviov",
		"Nikitin",
		"Ortov",
		"Peshkov",
		"Romasko",
		"Shvedov",
		"Sikorski",
		"Stolypin",
		"Turov",
		"Volokh",
		"Zaitsev",
		"Zhukov",
	}
	return names
}

func russianPlaceNames() []string {
	names := []string{
		"Amur",
		"Arkhangelsk",
		"Astrakhan",
		"Belgorod",
		"Bryansk",
		"Chelyabinsk",
		"Chita",
		"Gorki",
		"Irkutsk",
		"Ivanovo",
		"Kaliningrad",
		"Kaluga",
		"Kamchatka",
		"Kemerovo",
		"Kirov",
		"Kostroma",
		"Kurgan",
		"Kursk",
		"Leningrad",
		"Lipetsk",
		"Magadan",
		"Moscow",
		"Murmansk",
		"Novgorod",
		"Novosibirsk",
		"Omsk",
		"Orenburg",
		"Oryol",
		"Penza",
		"Perm",
		"Pskov",
		"Rostov",
		"Ryazan",
		"Sakhalin",
		"Samara",
		"Saratov",
		"Smolensk",
		"Sverdlovsk",
		"Tambov",
		"Tomsk",
		"Tula",
		"Tver",
		"Tyumen",
		"Ulyanovsk",
		"Vladimir",
		"Volgograd",
		"Vologda",
		"Voronezh",
		"Vyborg",
		"Yaroslavl",
	}
	return names
}

func spanishMaleNames() []string {
	names := []string{
		"Alejandro",
		"Alonso",
		"Amelio",
		"Armando",
		"Bernardo",
		"Carlos",
		"Cesar",
		"Diego",
		"Emilio",
		"Estevan",
		"Felipe",
		"Francisco",
		"Guillermo",
		"Javier",
		"Jose",
		"Juan",
		"Julio",
		"Luis",
		"Pedro",
		"Raul",
		"Ricardo",
		"Salvador",
		"Santiago",
		"Valeriano",
		"Vicente",
	}
	return names
}

func spanishFemaleNames() []string {
	names := []string{
		"Adalina",
		"Aleta",
		"Ana",
		"Ascencion",
		"Beatriz",
		"Carmela",
		"Celia",
		"Dolores",
		"Elena",
		"Emelina",
		"Felipa",
		"Inez",
		"Isabel",
		"Jacinta",
		"Lucia",
		"Lupe",
		"Maria",
		"Marta",
		"Nina",
		"Paloma",
		"Rafaela",
		"Soledad",
		"Teresa",
		"Valencia",
		"Zenaida",
	}
	return names
}

func spanishSurNames() []string {
	names := []string{
		"Arellano",
		"Arispana",
		"Borrego",
		"Carderas",
		"Carranzo",
		"Cordova",
		"Enciso",
		"Espejo",
		"Gavilan",
		"Guerra",
		"Guillen",
		"Huertas",
		"Illan",
		"Jurado",
		"Moretta",
		"Motolinia",
		"Pancorbo",
		"Paredes",
		"Quesada",
		"Roma",
		"Rubiera",
		"Santoro",
		"Torrillas",
		"Vera",
		"Vivero",
	}
	return names
}

func spanishPlaceNames() []string {
	names := []string{
		"Aguascebas",
		"Alcazar",
		"Barranquete",
		"Bravatas",
		"Cabezudos",
		"Calderon",
		"Cantera",
		"Castillo",
		"Delgadas",
		"Donablanca",
		"Encinetas",
		"Estrella",
		"Faustino",
		"Fuentebravia",
		"Gafarillos",
		"Gironda",
		"Higueros",
		"Huelago",
		"Humilladero",
		"Illora",
		"Isabela",
		"Izbor",
		"Jandilla",
		"Jinetes",
		"Limones",
		"Loreto",
		"Lujar",
		"Marbela",
		"Matagorda",
		"Nacimiento",
		"Niguelas",
		"Ogijares",
		"Ortegicar",
		"Pampanico",
		"Pelado",
		"Quesada",
		"Quintera",
		"Riguelo",
		"Ruescas",
		"Salteras",
		"Santopitar",
		"Taberno",
		"Torres",
		"Umbrete",
		"Valdecazorla",
		"Velez",
		"Vistahermosa",
		"Yeguas",
		"Zahora",
		"Zumeta",
	}
	return names
}

func randomSite() string {
	r := roll1dX(10, -1)
	site := []string{
		"arabic",
		"chinese",
		"english",
		"greek",
		"indian",
		"japanese",
		"latin",
		"negerian",
		"russian",
		"spanish",
	}
	return site[r]
}

func getPlace(site string, index int, f func(i int) string) string {
	//funcName := "Place"
	//getFirstName(female bool, ethnicity string, index int)
	return "test"
}

// func someFunction1(a, b int) int {
// 	return a + b
// }

// func someFunction2(a, b int) int {
// 	return a - b
// }

// func someOtherFunction(a, b int, f func(int, int) int) int {
// 	return f(a, b)
// }

func GiveName(site string, nameType string, index int) string {
	registry := namesRegistry()
	if nameType != "Male" && nameType != "Female" && nameType != "Sur" && nameType != "Place" {
		return "Error: 'nameType' is wrong"
	}
	if _, ok := registry[site+nameType+"Names"]; !ok {
		site = randomSite()
	}
	namesList := registry[site+nameType+"Names"]()
	if index < 0 || index >= len(namesList) {
		index = roll1dX(len(namesList), -1)
	}

	return registry[site+nameType+"Names"]()[index]
}

func namesRegistry() map[string]func() []string {
	registry := map[string]func() []string{
		"arabicMaleNames":     arabicMaleNames,
		"arabicFemaleNames":   arabicFemaleNames,
		"arabicSurNames":      arabicSurNames,
		"arabicPlaceNames":    arabicPlaceNames,
		"chineseMaleNames":    chineseMaleNames,
		"chineseFemaleNames":  chineseFemaleNames,
		"chineseSurNames":     chineseSurNames,
		"chinesePlaceNames":   chinesePlaceNames,
		"englishMaleNames":    englishMaleNames,
		"englishFemaleNames":  englishFemaleNames,
		"englishSurNames":     englishSurNames,
		"englishPlaceNames":   englishPlaceNames,
		"greekMaleNames":      greekMaleNames,
		"greekFemaleNames":    greekFemaleNames,
		"greekSurNames":       greekSurNames,
		"greekPlaceNames":     greekPlaceNames,
		"indianMaleNames":     indianMaleNames,
		"indianFemaleNames":   indianFemaleNames,
		"indianSurNames":      indianSurNames,
		"indianPlaceNames":    indianPlaceNames,
		"japaneseMaleNames":   japaneseMaleNames,
		"japaneseFemaleNames": japaneseFemaleNames,
		"japaneseSurNames":    japaneseSurNames,
		"japanesePlaceNames":  japanesePlaceNames,
		"latinMaleNames":      latinMaleNames,
		"latinFemaleNames":    latinFemaleNames,
		"latinSurNames":       latinSurNames,
		"latinPlaceNames":     latinPlaceNames,
		"negerianMaleNames":   negerianMaleNames,
		"negerianFemaleNames": negerianFemaleNames,
		"negerianSurNames":    negerianSurNames,
		"negerianPlaceNames":  negerianPlaceNames,
		"russianMaleNames":    russianMaleNames,
		"russianFemaleNames":  russianFemaleNames,
		"russianSurNames":     russianSurNames,
		"russianPlaceNames":   russianPlaceNames,
		"spanishMaleNames":    spanishMaleNames,
		"spanishFemaleNames":  spanishFemaleNames,
		"spanishSurNames":     spanishSurNames,
		"spanishPlaceNames":   spanishPlaceNames,
	}
	return registry
}

//getFirstName(nameType string, ethnicity string, index int) string {
