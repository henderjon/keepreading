package djrp

type day struct {
	Gospel, Nt, Ot, Wisdom string
}

// the Discipleship Journal Bible Reading Plan
var Plan = [][]day{
	[]day{
		{
			Gospel: "Matt 1:1-17",
			Nt:     "Acts 1:1-11",
			Wisdom: "Psalm 1",
			Ot:     "Gen 1-2",
		},
		{
			Gospel: "Matt 1:18-25",
			Nt:     "Acts 1:12-26",
			Wisdom: "Psalm 2",
			Ot:     "Gen 3-4",
		},
		{
			Gospel: "Matt 2:1-12",
			Nt:     "Acts 2:1-21",
			Wisdom: "Psalm 3",
			Ot:     "Gen 5-8",
		},
		{
			Gospel: "Matt 2:13-23",
			Nt:     "Acts 2:22-47",
			Wisdom: "Psalm 4",
			Ot:     "Gen 9-11",
		},
		{
			Gospel: "Matt 3:1-12",
			Nt:     "Acts 3",
			Wisdom: "Psalm 5",
			Ot:     "Gen 12-14",
		},
		{
			Gospel: "Matt 3:13-17",
			Nt:     "Acts 4:1-22",
			Wisdom: "Psalm 6",
			Ot:     "Gen 15-17",
		},
		{
			Gospel: "Matt 4:1-11",
			Nt:     "Acts 4:23-37",
			Wisdom: "Psalm 7",
			Ot:     "Gen 18-20",
		},
		{
			Gospel: "Matt 4:12-17",
			Nt:     "Acts 5:1-16",
			Wisdom: "Psalm 8",
			Ot:     "Gen 21-23",
		},
		{
			Gospel: "Matt 4:18-25",
			Nt:     "Acts 5:17-42",
			Wisdom: "Psalm 9",
			Ot:     "Gen 24",
		},
		{
			Gospel: "Matt 5:1-12",
			Nt:     "Acts 6",
			Wisdom: "Psalm 10",
			Ot:     "Gen 25-26",
		},
		{
			Gospel: "Matt 5:13-20",
			Nt:     "Acts 7:1-38",
			Wisdom: "Psalm 11",
			Ot:     "Gen 27-28",
		},
		{
			Gospel: "Matt 5:21-32",
			Nt:     "Acts 7:39-60",
			Wisdom: "Psalm 12",
			Ot:     "Gen 29-30",
		},
		{
			Gospel: "Matt 5:33-48",
			Nt:     "Acts 8:1-25",
			Wisdom: "Psalm 13",
			Ot:     "Gen 31",
		},
		{
			Gospel: "Matt 6:1-15",
			Nt:     "Acts 8:26-40",
			Wisdom: "Psalm 14",
			Ot:     "Gen 32-33",
		},
		{
			Gospel: "Matt 6:16-24",
			Nt:     "Acts 9:1-19",
			Wisdom: "Psalm 15",
			Ot:     "Gen 34-35",
		},
		{
			Gospel: "Matt 6:25-34",
			Nt:     "Acts 9:20-43",
			Wisdom: "Psalm 16",
			Ot:     "Gen 36",
		},
		{
			Gospel: "Matt 7:1-14",
			Nt:     "Acts 10:1-23",
			Wisdom: "Psalm 17",
			Ot:     "Gen 37-38",
		},
		{
			Gospel: "Matt 7:15-29",
			Nt:     "Acts 10:24-48",
			Wisdom: "Psalm 18:1-24",
			Ot:     "Gen 39-40",
		},
		{
			Gospel: "Matt 8:1-13",
			Nt:     "Acts 11:1-18",
			Wisdom: "Psalm 18:25-50",
			Ot:     "Gen 41",
		},
		{
			Gospel: "Matt 8:14-22",
			Nt:     "Acts 11:19-30",
			Wisdom: "Psalm 19",
			Ot:     "Gen 42-43",
		},
		{
			Gospel: "Matt 8:23-34",
			Nt:     "Acts 12",
			Wisdom: "Psalm 20",
			Ot:     "Gen 44-45",
		},
		{
			Gospel: "Matt 9:1-13",
			Nt:     "Acts 13:1-25",
			Wisdom: "Psalm 21",
			Ot:     "Gen 46-47",
		},
		{
			Gospel: "Matt 9:14-26",
			Nt:     "Acts 13:26-52",
			Wisdom: "Psalm 22:1-11",
			Ot:     "Gen 48",
		},
		{
			Gospel: "Matt 9:27-38",
			Nt:     "Acts 14",
			Wisdom: "Psalm 22:12-31",
			Ot:     "Gen 49",
		},
		{
			Gospel: "Matt 10:1-20",
			Nt:     "Acts 15:1-21",
			Wisdom: "Psalm 23",
			Ot:     "Gen 50",
		},
	},
	[]day{
		{
			Gospel: "Matt 10:21-42",
			Nt:     "Acts 15:22-41",
			Wisdom: "Psalm 24",
			Ot:     "Ex 1-3",
		},
		{
			Gospel: "Matt 11:1-19",
			Nt:     "Acts 16:1-15",
			Wisdom: "Psalm 25",
			Ot:     "Ex 4-6",
		},
		{
			Gospel: "Matt 11:20-30",
			Nt:     "Acts 16:16-40",
			Wisdom: "Psalm 26",
			Ot:     "Ex 7-9",
		},
		{
			Gospel: "Matt 12:1-21",
			Nt:     "Acts 17:1-15",
			Wisdom: "Psalm 27",
			Ot:     "Ex 10-12",
		},
		{
			Gospel: "Matt 12:22-37",
			Nt:     "Acts 17:16-34",
			Wisdom: "Psalm 28",
			Ot:     "Ex 13-15",
		},
		{
			Gospel: "Matt 12:38-50",
			Nt:     "Acts 18:1-17",
			Wisdom: "Psalm 29",
			Ot:     "Ex 16-18",
		},
		{
			Gospel: "Matt 13:1-23",
			Nt:     "Acts 18:18-28",
			Wisdom: "Psalm 30",
			Ot:     "Ex 19-20",
		},
		{
			Gospel: "Matt 13:24-43",
			Nt:     "Acts 19:1-22",
			Wisdom: "Psalm 31",
			Ot:     "Ex 21-23",
		},
		{
			Gospel: "Matt 13:44-58",
			Nt:     "Acts 19:23-41",
			Wisdom: "Psalm 32",
			Ot:     "Ex 24-26",
		},
		{
			Gospel: "Matt 14:1-21",
			Nt:     "Acts 20:1-12",
			Wisdom: "Psalm 33",
			Ot:     "Ex 27-29",
		},
		{
			Gospel: "Matt 14:22-36",
			Nt:     "Acts 20:13-38",
			Wisdom: "Psalm 34",
			Ot:     "Ex 30-31",
		},
		{
			Gospel: "Matt 15:1-20",
			Nt:     "Acts 21:1-26",
			Wisdom: "Psalm 35",
			Ot:     "Ex 32-33",
		},
		{
			Gospel: "Matt 15:21-39",
			Nt:     "Acts 21:27-40",
			Wisdom: "Psalm 36",
			Ot:     "Ex 34",
		},
		{
			Gospel: "Matt 16:1-12",
			Nt:     "Acts 22",
			Wisdom: "Psalm 37:1-22",
			Ot:     "Ex 35-37",
		},
		{
			Gospel: "Matt 16:13-28",
			Nt:     "Acts 23:1-11",
			Wisdom: "Psalm 37:23-40",
			Ot:     "Ex 38-40",
		},
		{
			Gospel: "Matt 17:1-13",
			Nt:     "Acts 23:12-35",
			Wisdom: "Psalm 38",
			Ot:     "Lev 1-4",
		},
		{
			Gospel: "Matt 17:14-27",
			Nt:     "Acts 24",
			Wisdom: "Psalm 39",
			Ot:     "Lev 5-7",
		},
		{
			Gospel: "Matt 18:1-14",
			Nt:     "Acts 25:1-12",
			Wisdom: "Psalm 40",
			Ot:     "Lev 8-10",
		},
		{
			Gospel: "Matt 18:15-35",
			Nt:     "Acts 25:13-27",
			Wisdom: "Psalm 41",
			Ot:     "Lev 11-13",
		},
		{
			Gospel: "Matt 19:1-15",
			Nt:     "Acts 26:1-18",
			Wisdom: "Psalm 42",
			Ot:     "Lev 14-15",
		},
		{
			Gospel: "Matt 19:16-30",
			Nt:     "Acts 26:19-32",
			Wisdom: "Psalm 43",
			Ot:     "Lev 16-17",
		},
		{
			Gospel: "Matt 20:1-16",
			Nt:     "Acts 27:1-26",
			Wisdom: "Psalm 44",
			Ot:     "Lev 18-20",
		},
		{
			Gospel: "Matt 20:17-34",
			Nt:     "Acts 27:27-44",
			Wisdom: "Psalm 45",
			Ot:     "Lev 21-23",
		},
		{
			Gospel: "Matt 21:1-11",
			Nt:     "Acts 28:1-16",
			Wisdom: "Psalm 46",
			Ot:     "Lev 24-25",
		},
		{
			Gospel: "Matt 21:12-22",
			Nt:     "Acts 28:17-31",
			Wisdom: "Psalm 47",
			Ot:     "Lev 26-27",
		},
	},
	[]day{
		{
			Gospel: "Matt 21:23-32",
			Nt:     "Romans 1:1-17",
			Wisdom: "Psalm 48",
			Ot:     "Num 1-2",
		},
		{
			Gospel: "Matt 21:33-46",
			Nt:     "Romans 1:18-32",
			Wisdom: "Psalm 49",
			Ot:     "Num 3-4",
		},
		{
			Gospel: "Matt 22:1-14",
			Nt:     "Romans 2",
			Wisdom: "Psalm 50",
			Ot:     "Num 5-6",
		},
		{
			Gospel: "Matt 22:15-33",
			Nt:     "Romans 3",
			Wisdom: "Psalm 51",
			Ot:     "Num 7-8",
		},
		{
			Gospel: "Matt 22:34-46",
			Nt:     "Romans 4",
			Wisdom: "Psalm 52",
			Ot:     "Num 9-11",
		},
		{
			Gospel: "Matt 23:1-12",
			Nt:     "Romans 5:1-11",
			Wisdom: "Psalm 53",
			Ot:     "Num 12-14",
		},
		{
			Gospel: "Matt 23:13-24",
			Nt:     "Romans 5:12-21",
			Wisdom: "Psalm 54",
			Ot:     "Num 15-17",
		},
		{
			Gospel: "Matt 23:25-39",
			Nt:     "Romans 6:1-14",
			Wisdom: "Psalm 55",
			Ot:     "Num 18-20",
		},
		{
			Gospel: "Matt 24:1-14",
			Nt:     "Romans 6:15-23",
			Wisdom: "Psalm 56",
			Ot:     "Num 21-22",
		},
		{
			Gospel: "Matt 24:15-35",
			Nt:     "Romans 7:1-12",
			Wisdom: "Psalm 57",
			Ot:     "Num 23-25",
		},
		{
			Gospel: "Matt 24:36-51",
			Nt:     "Romans 7:13-25",
			Wisdom: "Psalm 58",
			Ot:     "Num 26-27",
		},
		{
			Gospel: "Matt 25:1-13",
			Nt:     "Romans 8:1-17",
			Wisdom: "Psalm 59",
			Ot:     "Num 28-30",
		},
		{
			Gospel: "Matt 25:14-30",
			Nt:     "Romans 8:18-39",
			Wisdom: "Psalm 60",
			Ot:     "Num 31-32",
		},
		{
			Gospel: "Matt 25:31-46",
			Nt:     "Romans 9:1-18",
			Wisdom: "Psalm 61",
			Ot:     "Num 33-36",
		},
		{
			Gospel: "Matt 26:1-16",
			Nt:     "Romans 9:19-33",
			Wisdom: "Psalm 62",
			Ot:     "Deut 1-3",
		},
		{
			Gospel: "Matt 26:17-35",
			Nt:     "Romans 10",
			Wisdom: "Psalm 63",
			Ot:     "Deut 4-5",
		},
		{
			Gospel: "Matt 26:36-56",
			Nt:     "Romans 11:1-24",
			Wisdom: "Psalm 64",
			Ot:     "Deut 6-8",
		},
		{
			Gospel: "Matt 26:57-75",
			Nt:     "Romans 11:25-36",
			Wisdom: "Psalm 65",
			Ot:     "Deut 9-12",
		},
		{
			Gospel: "Matt 27:1-10",
			Nt:     "Romans 12:1-8",
			Wisdom: "Psalm 66",
			Ot:     "Deut 13-17",
		},
		{
			Gospel: "Matt 27:11-26",
			Nt:     "Romans 12:9-21",
			Wisdom: "Psalm 67",
			Ot:     "Deut 18-21",
		},
		{
			Gospel: "Matt 27:27-44",
			Nt:     "Romans 13",
			Wisdom: "Psalm 68",
			Ot:     "Deut 22-26",
		},
		{
			Gospel: "Matt 27:45-56",
			Nt:     "Romans 14",
			Wisdom: "Psalm 69:1-18",
			Ot:     "Deut 27-28",
		},
		{
			Gospel: "Matt 27:57-66",
			Nt:     "Romans 15:1-13",
			Wisdom: "Psalm 69:19-36",
			Ot:     "Deut 29-31",
		},
		{
			Gospel: "Matt 28:1-10",
			Nt:     "Romans 15:14-33",
			Wisdom: "Psalm 70",
			Ot:     "Deut 32",
		},
		{
			Gospel: "Matt 28:11-20",
			Nt:     "Romans 16",
			Wisdom: "Psalm 71",
			Ot:     "Deut 33-34",
		},
	},
	[]day{
		{
			Gospel: "Mark 1:1-8",
			Nt:     "1 Cor 1:1-17",
			Wisdom: "Psalm 72",
			Ot:     "Josh 1-2",
		},
		{
			Gospel: "Mark 1:9-20",
			Nt:     "1 Cor 1:18-31",
			Wisdom: "Psalm 73",
			Ot:     "Josh 3-5",
		},
		{
			Gospel: "Mark 1:21-34",
			Nt:     "1 Cor 2",
			Wisdom: "Psalm 74",
			Ot:     "Josh 6-7",
		},
		{
			Gospel: "Mark 1:35-45",
			Nt:     "1 Cor 3",
			Wisdom: "Psalm 75",
			Ot:     "Josh 8-9",
		},
		{
			Gospel: "Mark 2:1-12",
			Nt:     "1 Cor 4",
			Wisdom: "Psalm 76",
			Ot:     "Josh 10-12",
		},
		{
			Gospel: "Mark 2:13-17",
			Nt:     "1 Cor 5",
			Wisdom: "Psalm 77",
			Ot:     "Josh 13-14",
		},
		{
			Gospel: "Mark 2:18-28",
			Nt:     "1 Cor 6:1-11",
			Wisdom: "Psalm 78:1-39",
			Ot:     "Josh 15-17",
		},
		{
			Gospel: "Mark 3:1-19",
			Nt:     "1 Cor 6:12-20",
			Wisdom: "Psalm 78:40-72",
			Ot:     "Josh 18-19",
		},
		{
			Gospel: "Mark 3:20-35",
			Nt:     "1 Cor 7:1-16",
			Wisdom: "Psalm 79",
			Ot:     "Josh 20-21",
		},
		{
			Gospel: "Mark 4:1-20",
			Nt:     "1 Cor 7:17-40",
			Wisdom: "Psalm 80",
			Ot:     "Josh 22-23",
		},
		{
			Gospel: "Mark 4:21-41",
			Nt:     "1 Cor 8",
			Wisdom: "Psalm 81",
			Ot:     "Josh 24",
		},
		{
			Gospel: "Mark 5:1-20",
			Nt:     "1 Cor 9:1-12",
			Wisdom: "Psalm 82",
			Ot:     "Judges 1-3",
		},
		{
			Gospel: "Mark 5:21-43",
			Nt:     "1 Cor 9:13-27",
			Wisdom: "Psalm 83",
			Ot:     "Judges 4-5",
		},
		{
			Gospel: "Mark 6:1-13",
			Nt:     "1 Cor 10:1-13",
			Wisdom: "Psalm 84",
			Ot:     "Judges 6-7",
		},
		{
			Gospel: "Mark 6:14-29",
			Nt:     "1 Cor 10:14-33",
			Wisdom: "Psalm 85",
			Ot:     "Judges 8",
		},
		{
			Gospel: "Mark 6:30-44",
			Nt:     "1 Cor 11:1-16",
			Wisdom: "Psalm 86",
			Ot:     "Judges 9",
		},
		{
			Gospel: "Mark 6:45-56",
			Nt:     "1 Cor 11:17-34",
			Wisdom: "Psalm 87",
			Ot:     "Judges 10-12",
		},
		{
			Gospel: "Mark 7:1-23",
			Nt:     "1 Cor 12:1-13",
			Wisdom: "Psalm 88",
			Ot:     "Judges 13-15",
		},
		{
			Gospel: "Mark 7:24-37",
			Nt:     "1 Cor 12:14-31",
			Wisdom: "Psalm 89:1-18",
			Ot:     "Judges 16",
		},
		{
			Gospel: "Mark 8:1-13",
			Nt:     "1 Cor 13",
			Wisdom: "Psalm 89:19-52",
			Ot:     "Judges 17-18",
		},
		{
			Gospel: "Mark 8:14-21",
			Nt:     "1 Cor 14:1-25",
			Wisdom: "Psalm 90",
			Ot:     "Judges 19",
		},
		{
			Gospel: "Mark 8:22-30",
			Nt:     "1 Cor 14:26-40",
			Wisdom: "Psalm 91",
			Ot:     "Judges 20-21",
		},
		{
			Gospel: "Mark 8:31-38",
			Nt:     "1 Cor 15:1-28",
			Wisdom: "Psalm 92",
			Ot:     "Ruth 1",
		},
		{
			Gospel: "Mark 9:1-13",
			Nt:     "1 Cor 15:29-58",
			Wisdom: "Psalm 93",
			Ot:     "Ruth 2-3",
		},
		{
			Gospel: "Mark 9:14-32",
			Nt:     "1 Cor 16",
			Wisdom: "Psalm 94",
			Ot:     "Ruth 4",
		},
	},
	[]day{
		{
			Gospel: "Mark 9:33-50",
			Nt:     "2 Cor 1:1-11",
			Wisdom: "Psalm 95",
			Ot:     "1 Sam 1-2",
		},
		{
			Gospel: "Mark 10:1-16",
			Nt:     "2 Cor 1:12-24",
			Wisdom: "Psalm 96",
			Ot:     "1 Sam 3-5",
		},
		{
			Gospel: "Mark 10:17-34",
			Nt:     "2 Cor 2",
			Wisdom: "Psalm 97",
			Ot:     "1 Sam 6-8",
		},
		{
			Gospel: "Mark 10:35-52",
			Nt:     "2 Cor 3",
			Wisdom: "Psalm 98",
			Ot:     "1 Sam 9-10",
		},
		{
			Gospel: "Mark 11:1-11",
			Nt:     "2 Cor 4",
			Wisdom: "Psalm 99",
			Ot:     "1 Sam 11-13",
		},
		{
			Gospel: "Mark 11:12-26",
			Nt:     "2 Cor 5",
			Wisdom: "Psalm 100",
			Ot:     "1 Sam 14",
		},
		{
			Gospel: "Mark 11:27-33",
			Nt:     "2 Cor 6",
			Wisdom: "Psalm 101",
			Ot:     "1 Sam 15-16",
		},
		{
			Gospel: "Mark 12:1-12",
			Nt:     "2 Cor 7",
			Wisdom: "Psalm 102",
			Ot:     "1 Sam 17-18",
		},
		{
			Gospel: "Mark 12:13-27",
			Nt:     "2 Cor 8",
			Wisdom: "Psalm 103",
			Ot:     "1 Sam 19-20",
		},
		{
			Gospel: "Mark 12:28-34",
			Nt:     "2 Cor 9",
			Wisdom: "Psalm 104",
			Ot:     "1 Sam 21-23",
		},
		{
			Gospel: "Mark 12:35-44",
			Nt:     "2 Cor 10",
			Wisdom: "Psalm 105",
			Ot:     "1 Sam 24-25",
		},
		{
			Gospel: "Mark 13:1-13",
			Nt:     "2 Cor 11:1-15",
			Wisdom: "Psalm 106:1-23",
			Ot:     "1 Sam 26-28",
		},
		{
			Gospel: "Mark 13:14-31",
			Nt:     "2 Cor 11:16-33",
			Wisdom: "Psalm 106:24-48",
			Ot:     "1 Sam 29-31",
		},
		{
			Gospel: "Mark 13:32-37",
			Nt:     "2 Cor 12:1-10",
			Wisdom: "Psalm 107",
			Ot:     "2 Sam 1-2",
		},
		{
			Gospel: "Mark 14:1-11",
			Nt:     "2 Cor 12:11-21",
			Wisdom: "Psalm 108",
			Ot:     "2 Sam 3-4",
		},
		{
			Gospel: "Mark 14:12-31",
			Nt:     "2 Cor 13",
			Wisdom: "Psalm 109",
			Ot:     "2 Sam 5-7",
		},
		{
			Gospel: "Mark 14:32-42",
			Nt:     "Gal 1",
			Wisdom: "Psalm 110",
			Ot:     "2 Sam 8-10",
		},
		{
			Gospel: "Mark 14:43-52",
			Nt:     "Gal 2",
			Wisdom: "Psalm 111",
			Ot:     "2 Sam 11-12",
		},
		{
			Gospel: "Mark 14:53-65",
			Nt:     "Gal 3:1-14",
			Wisdom: "Psalm 112",
			Ot:     "2 Sam 13",
		},
		{
			Gospel: "Mark 14:66-72",
			Nt:     "Gal 3:15-29",
			Wisdom: "Psalm 113",
			Ot:     "2 Sam 14-15",
		},
		{
			Gospel: "Mark 15:1-15",
			Nt:     "Gal 4:1-20",
			Wisdom: "Psalm 114",
			Ot:     "2 Sam 16-17",
		},
		{
			Gospel: "Mark 15:16-32",
			Nt:     "Gal 4:21-31",
			Wisdom: "Psalm 115",
			Ot:     "2 Sam 18-19",
		},
		{
			Gospel: "Mark 15:33-41",
			Nt:     "Gal 5:1-12",
			Wisdom: "Psalm 116",
			Ot:     "2 Sam 20-21",
		},
		{
			Gospel: "Mark 15:42-47",
			Nt:     "Gal 5:13-26",
			Wisdom: "Psalm 117",
			Ot:     "2 Sam 22",
		},
		{
			Gospel: "Mark 16",
			Nt:     "Gal 6",
			Wisdom: "Psalm 118",
			Ot:     "2 Sam 23-24",
		},
	},
	[]day{
		{
			Gospel: "Luke 1:1-25",
			Nt:     "Eph 1:1-14",
			Wisdom: "Psalm 119:1-8",
			Ot:     "1 Kings 1",
		},
		{
			Gospel: "Luke 1:26-38",
			Nt:     "Eph 1:15-23",
			Wisdom: "Psalm 119:9-16",
			Ot:     "1 Kings 2-3",
		},
		{
			Gospel: "Luke 1:39-56",
			Nt:     "Eph 2:1-10",
			Wisdom: "Psalm 119:17-24",
			Ot:     "1 Kings 4-5",
		},
		{
			Gospel: "Luke 1:57-66",
			Nt:     "Eph 2:11-22",
			Wisdom: "Psalm 119:25-32",
			Ot:     "1 Kings 6-7",
		},
		{
			Gospel: "Luke 1:67-80",
			Nt:     "Eph 3:1-13",
			Wisdom: "Psalm 119:33-40",
			Ot:     "1 Kings 8",
		},
		{
			Gospel: "Luke 2:1-20",
			Nt:     "Eph 3:14-21",
			Wisdom: "Psalm 119:41-48",
			Ot:     "1 Kings 9-10",
		},
		{
			Gospel: "Luke 2:21-40",
			Nt:     "Eph 4:1-16",
			Wisdom: "Psalm 119:49-56",
			Ot:     "1 Kings 11",
		},
		{
			Gospel: "Luke 2:41-52",
			Nt:     "Eph 4:17-24",
			Wisdom: "Psalm 119:57-64",
			Ot:     "1 Kings 12",
		},
		{
			Gospel: "Luke 3:1-20",
			Nt:     "Eph 4:25-32",
			Wisdom: "Psalm 119:65-72",
			Ot:     "1 Kings 13-14",
		},
		{
			Gospel: "Luke 3:21-38",
			Nt:     "Eph 5:1-21",
			Wisdom: "Psalm 119:73-80",
			Ot:     "1 Kings 15-16",
		},
		{
			Gospel: "Luke 4:1-12",
			Nt:     "Eph 5:22-33",
			Wisdom: "Psalm 119:81-88",
			Ot:     "1 Kings 17-18",
		},
		{
			Gospel: "Luke 4:13-30",
			Nt:     "Eph 6:1-9",
			Wisdom: "Psalm 119:89-96",
			Ot:     "1 Kings 19-20",
		},
		{
			Gospel: "Luke 4:31-37",
			Nt:     "Eph 6:10-24",
			Wisdom: "Psalm 119:97-104",
			Ot:     "1 Kings 21-22",
		},
		{
			Gospel: "Luke 4:38-44",
			Nt:     "Phil 1:1-11",
			Wisdom: "Psalm 119:105-112",
			Ot:     "2 Kings 1-3",
		},
		{
			Gospel: "Luke 5:1-11",
			Nt:     "Phil 1:12-20",
			Wisdom: "Psalm 119:113-120",
			Ot:     "2 Kings 4-5",
		},
		{
			Gospel: "Luke 5:12-16",
			Nt:     "Phil 1:21-30",
			Wisdom: "Psalm 119:121-128",
			Ot:     "2 Kings 6-7",
		},
		{
			Gospel: "Luke 5:17-26",
			Nt:     "Phil 2:1-11",
			Wisdom: "Psalm 119:129-136",
			Ot:     "2 Kings 8-9",
		},
		{
			Gospel: "Luke 5:27-32",
			Nt:     "Phil 2:12-18",
			Wisdom: "Psalm 119:137-144",
			Ot:     "2 Kings 10-11",
		},
		{
			Gospel: "Luke 5:33-39",
			Nt:     "Phil 2:19-30",
			Wisdom: "Psalm 119:145-152",
			Ot:     "2 Kings 12-13",
		},
		{
			Gospel: "Luke 6:1-16",
			Nt:     "Phil 3:1-9",
			Wisdom: "Psalm 119:153-160",
			Ot:     "2 Kings 14-15",
		},
		{
			Gospel: "Luke 6:17-26",
			Nt:     "Phil 3:10-14",
			Wisdom: "Psalm 119:161-168",
			Ot:     "2 Kings 16-17",
		},
		{
			Gospel: "Luke 6:27-36",
			Nt:     "Phil 3:15-21",
			Wisdom: "Psalm 119:169-176",
			Ot:     "2 Kings 18-19",
		},
		{
			Gospel: "Luke 6:37-42",
			Nt:     "Phil 4:1-7",
			Wisdom: "Psalm 120",
			Ot:     "2 Kings 20-21",
		},
		{
			Gospel: "Luke 6:43-49",
			Nt:     "Phil 4:8-13",
			Wisdom: "Psalm 121",
			Ot:     "2 Kings 22-23",
		},
		{
			Gospel: "Luke 7:1-10",
			Nt:     "Phil 4:14-23",
			Wisdom: "Psalm 122",
			Ot:     "2 Kings 24-25",
		},
	},
	[]day{
		{
			Gospel: "Luke 7:11-17",
			Nt:     "Col 1:1-14",
			Wisdom: "Psalm 123-124",
			Ot:     "1 Chron 1-2",
		},
		{
			Gospel: "Luke 7:18-35",
			Nt:     "Col 1:15-29",
			Wisdom: "Psalm 125",
			Ot:     "1 Chron 3-4",
		},
		{
			Gospel: "Luke 7:36-50",
			Nt:     "Col 2:1-7",
			Wisdom: "Psalm 126",
			Ot:     "1 Chron 5-6",
		},
		{
			Gospel: "Luke 8:1-15",
			Nt:     "Col 2:8-15",
			Wisdom: "Psalm 127",
			Ot:     "1 Chron 7-9",
		},
		{
			Gospel: "Luke 8:16-25",
			Nt:     "Col 2:16-23",
			Wisdom: "Psalm 128",
			Ot:     "1 Chron 10-11",
		},
		{
			Gospel: "Luke 8:26-39",
			Nt:     "Col 3:1-14",
			Wisdom: "Psalm 129",
			Ot:     "1 Chron 12-14",
		},
		{
			Gospel: "Luke 8:40-56",
			Nt:     "Col 3:15-25",
			Wisdom: "Psalm 130-131",
			Ot:     "1 Chron 15-16",
		},
		{
			Gospel: "Luke 9:1-17",
			Nt:     "Col 4:1-9",
			Wisdom: "Psalm 132",
			Ot:     "1 Chron 17-19",
		},
		{
			Gospel: "Luke 9:18-27",
			Nt:     "Col 4:10-18",
			Wisdom: "Psalm 133-134",
			Ot:     "1 Chron 20-22",
		},
		{
			Gospel: "Luke 9:28-36",
			Nt:     "1 Thess 1",
			Wisdom: "Psalm 135",
			Ot:     "1 Chron 23-25",
		},
		{
			Gospel: "Luke 9:37-50",
			Nt:     "1 Thess 2:1-9",
			Wisdom: "Psalm 136",
			Ot:     "1 Chron 26-28",
		},
		{
			Gospel: "Luke 9:51-62",
			Nt:     "1 Thess 2:10-20",
			Wisdom: "Psalm 137",
			Ot:     "1 Chron 29",
		},
		{
			Gospel: "Luke 10:1-16",
			Nt:     "1 Thess 3:1-6",
			Wisdom: "Psalm 138",
			Ot:     "2 Chron 1-2",
		},
		{
			Gospel: "Luke 10:17-24",
			Nt:     "1 Thess 3:7-13",
			Wisdom: "Psalm 139",
			Ot:     "2 Chron 3-5",
		},
		{
			Gospel: "Luke 10:25-37",
			Nt:     "1 Thess 4:1-10",
			Wisdom: "Psalm 140",
			Ot:     "2 Chron 6-7",
		},
		{
			Gospel: "Luke 10:38-42",
			Nt:     "1 Thess 4:11-18",
			Wisdom: "Psalm 141",
			Ot:     "2 Chron 8-9",
		},
		{
			Gospel: "Luke 11:1-13",
			Nt:     "1 Thess 5:1-11",
			Wisdom: "Psalm 142",
			Ot:     "2 Chron 10-12",
		},
		{
			Gospel: "Luke 11:14-28",
			Nt:     "1 Thess 5:12-28",
			Wisdom: "Psalm 143",
			Ot:     "2 Chron 13-16",
		},
		{
			Gospel: "Luke 11:29-36",
			Nt:     "2 Thess 1:1-7",
			Wisdom: "Psalm 144",
			Ot:     "2 Chron 17-19",
		},
		{
			Gospel: "Luke 11:37-54",
			Nt:     "2 Thess 1:8-12",
			Wisdom: "Psalm 145",
			Ot:     "2 Chron 20-21",
		},
		{
			Gospel: "Luke 12:1-12",
			Nt:     "2 Thess 2:1-12",
			Wisdom: "Psalm 146",
			Ot:     "2 Chron 22-24",
		},
		{
			Gospel: "Luke 12:13-21",
			Nt:     "2 Thess 2:13-17",
			Wisdom: "Psalm 147",
			Ot:     "2 Chron 25-27",
		},
		{
			Gospel: "Luke 12:22-34",
			Nt:     "2 Thess 3:1-5",
			Wisdom: "Psalm 148",
			Ot:     "2 Chron 28-29",
		},
		{
			Gospel: "Luke 12:35-48",
			Nt:     "2 Thess 3:6-13",
			Wisdom: "Psalm 149",
			Ot:     "2 Chron 30-33",
		},
		{
			Gospel: "Luke 12:49-59",
			Nt:     "2 Thess 3:14-18",
			Wisdom: "Psalm 150",
			Ot:     "2 Chron 34-36",
		},
	},
	[]day{
		{
			Gospel: "Luke 13:1-9",
			Nt:     "1 Tim 1:1-11",
			Wisdom: "Prov 1",
			Ot:     "Ezra 1-2",
		},
		{
			Gospel: "Luke 13:10-21",
			Nt:     "1 Tim 1:12-20",
			Wisdom: "Prov 2",
			Ot:     "Ezra 3",
		},
		{
			Gospel: "Luke 13:22-35",
			Nt:     "1 Tim 2",
			Wisdom: "Prov 3",
			Ot:     "Ezra 4-5",
		},
		{
			Gospel: "Luke 14:1-14",
			Nt:     "1 Tim 3:1-10",
			Wisdom: "Prov 4",
			Ot:     "Ezra 6",
		},
		{
			Gospel: "Luke 14:15-24",
			Nt:     "1 Tim 3:11-16",
			Wisdom: "Prov 5",
			Ot:     "Ezra 7",
		},
		{
			Gospel: "Luke 14:25-35",
			Nt:     "1 Tim 4",
			Wisdom: "Prov 6",
			Ot:     "Ezra 8",
		},
		{
			Gospel: "Luke 15:1-10",
			Nt:     "1 Tim 5:1-15",
			Wisdom: "Prov 7",
			Ot:     "Ezra 9",
		},
		{
			Gospel: "Luke 15:11-32",
			Nt:     "1 Tim 5:16-25",
			Wisdom: "Prov 8",
			Ot:     "Ezra 10",
		},
		{
			Gospel: "Luke 16:1-9",
			Nt:     "1 Tim 6:1-10",
			Wisdom: "Prov 9",
			Ot:     "Neh 1-2",
		},
		{
			Gospel: "Luke 16:10-18",
			Nt:     "1 Tim 6:11-21",
			Wisdom: "Prov 10:1-16",
			Ot:     "Neh 3",
		},
		{
			Gospel: "Luke 16:19-31",
			Nt:     "2 Tim 1:1-7",
			Wisdom: "Prov 10:17-32",
			Ot:     "Neh 4-5",
		},
		{
			Gospel: "Luke 17:1-10",
			Nt:     "2 Tim 1:8-18",
			Wisdom: "Prov 11:1-15",
			Ot:     "Neh 6",
		},
		{
			Gospel: "Luke 17:11-19",
			Nt:     "2 Tim 2:1-13",
			Wisdom: "Prov 11:16-31",
			Ot:     "Neh 7",
		},
		{
			Gospel: "Luke 17:20-37",
			Nt:     "2 Tim 2:14-26",
			Wisdom: "Prov 12:1-14",
			Ot:     "Neh 8",
		},
		{
			Gospel: "Luke 18:1-8",
			Nt:     "2 Tim 3:1-9",
			Wisdom: "Prov 12:15-28",
			Ot:     "Neh 9",
		},
		{
			Gospel: "Luke 18:9-17",
			Nt:     "2 Tim 3:10-17",
			Wisdom: "Prov 13:1-12",
			Ot:     "Neh 10",
		},
		{
			Gospel: "Luke 18:18-30",
			Nt:     "2 Tim 4",
			Wisdom: "Prov 13:13-25",
			Ot:     "Neh 11",
		},
		{
			Gospel: "Luke 18:31-43",
			Nt:     "Titus 1:1-9",
			Wisdom: "Prov 14:1-18",
			Ot:     "Neh 12",
		},
		{
			Gospel: "Luke 19:1-10",
			Nt:     "Titus 1:10-16",
			Wisdom: "Prov 14:19-35",
			Ot:     "Neh 13",
		},
		{
			Gospel: "Luke 19:11-27",
			Nt:     "Titus 2:1-10",
			Wisdom: "Prov 15:1-17",
			Ot:     "Est 1",
		},
		{
			Gospel: "Luke 19:28-38",
			Nt:     "Titus 2:11-15",
			Wisdom: "Prov 15:18-33",
			Ot:     "Est 2",
		},
		{
			Gospel: "Luke 19:39-48",
			Nt:     "Titus 3:1-8",
			Wisdom: "Prov 16:1-16",
			Ot:     "Est 3-4",
		},
		{
			Gospel: "Luke 20:1-8",
			Nt:     "Titus 3:9-15",
			Wisdom: "Prov 16:17-33",
			Ot:     "Est 5-6",
		},
		{
			Gospel: "Luke 20:9-19",
			Nt:     "Phile 1-11",
			Wisdom: "Prov 17:1-14",
			Ot:     "Est 7-8",
		},
		{
			Gospel: "Luke 20:20-26",
			Nt:     "Phile 12-25",
			Wisdom: "Prov 17:15-28",
			Ot:     "Est 9-10",
		},
	},
	[]day{
		{
			Gospel: "Luke 20:27-40",
			Nt:     "Heb 1:1-9",
			Wisdom: "Prov 18",
			Ot:     "Isa 1-2",
		},
		{
			Gospel: "Luke 20:41-47",
			Nt:     "Heb 1:10-14",
			Wisdom: "Prov 19:1-14",
			Ot:     "Isa 3-5",
		},
		{
			Gospel: "Luke 21:1-19",
			Nt:     "Heb 2:1-9",
			Wisdom: "Prov 19:15-29",
			Ot:     "Isa 6-8",
		},
		{
			Gospel: "Luke 21:20-28",
			Nt:     "Heb 2:10-18",
			Wisdom: "Prov 20:1-15",
			Ot:     "Isa 9-10",
		},
		{
			Gospel: "Luke 21:29-38",
			Nt:     "Heb 3",
			Wisdom: "Prov 20:16-30",
			Ot:     "Isa 11-13",
		},
		{
			Gospel: "Luke 22:1-13",
			Nt:     "Heb 4:1-11",
			Wisdom: "Prov 21:1-16",
			Ot:     "Isa 14-16",
		},
		{
			Gospel: "Luke 22:14-23",
			Nt:     "Heb 4:12-16",
			Wisdom: "Prov 21:17-31",
			Ot:     "Isa 17-20",
		},
		{
			Gospel: "Luke 22:24-30",
			Nt:     "Heb 5",
			Wisdom: "Prov 22:1-16",
			Ot:     "Isa 21-23",
		},
		{
			Gospel: "Luke 22:31-38",
			Nt:     "Heb 6:1-12",
			Wisdom: "Prov 22:17-29",
			Ot:     "Isa 24-26",
		},
		{
			Gospel: "Luke 22:39:46",
			Nt:     "Heb 6:13-20",
			Wisdom: "Prov 23:1-18",
			Ot:     "Isa 27-28",
		},
		{
			Gospel: "Luke 22:47-53",
			Nt:     "Heb 7:1-10",
			Wisdom: "Prov 23:19-35",
			Ot:     "Isa 29-30",
		},
		{
			Gospel: "Luke 22:54-62",
			Nt:     "Heb 7:11-28",
			Wisdom: "Prov 24:1-22",
			Ot:     "Isa 31-33",
		},
		{
			Gospel: "Luke 22:63-71",
			Nt:     "Heb 8:1-6",
			Wisdom: "Prov 24:23-34",
			Ot:     "Isa 34-36",
		},
		{
			Gospel: "Luke 23:1-12",
			Nt:     "Heb 8:7-13",
			Wisdom: "Prov 25:1-14",
			Ot:     "Isa 37-39",
		},
		{
			Gospel: "Luke 23:13-25",
			Nt:     "Heb 9:1-10",
			Wisdom: "Prov 25:15-28",
			Ot:     "Isa 40-41",
		},
		{
			Gospel: "Luke 23:26-31",
			Nt:     "Heb 9:11-28",
			Wisdom: "Prov 26:1-16",
			Ot:     "Isa 42-43",
		},
		{
			Gospel: "Luke 23:32-37",
			Nt:     "Heb 10:1-18",
			Wisdom: "Prov 26:17-28",
			Ot:     "Isa 44-45",
		},
		{
			Gospel: "Luke 23:38-43",
			Nt:     "Heb 10:19-39",
			Wisdom: "Prov 27:1-14",
			Ot:     "Isa 46-48",
		},
		{
			Gospel: "Luke 23:44-49",
			Nt:     "Heb 11:1-16",
			Wisdom: "Prov 27:15-27",
			Ot:     "Isa 49-50",
		},
		{
			Gospel: "Luke 23:50-56",
			Nt:     "Heb 11:17-31",
			Wisdom: "Prov 28:1-14",
			Ot:     "Isa 51-53",
		},
		{
			Gospel: "Luke 24:1-12",
			Nt:     "Heb 11:32-40",
			Wisdom: "Prov 28:15-28",
			Ot:     "Isa 54-55",
		},
		{
			Gospel: "Luke 24:13-27",
			Nt:     "Heb 12:1-13",
			Wisdom: "Prov 29:1-14",
			Ot:     "Isa 56-58",
		},
		{
			Gospel: "Luke 24:28-35",
			Nt:     "Heb 12:14-29",
			Wisdom: "Prov 29:15-27",
			Ot:     "Isa 59-61",
		},
		{
			Gospel: "Luke 24:36-44",
			Nt:     "Heb 13:1-8",
			Wisdom: "Prov 30",
			Ot:     "Isa 62-64",
		},
		{
			Gospel: "Luke 24:45-53",
			Nt:     "Heb 13:9-25",
			Wisdom: "Prov 31",
			Ot:     "Isa 65-66",
		},
	},
	[]day{
		{
			Gospel: "John 1:1-18",
			Nt:     "James 1:1-11",
			Wisdom: "Ecc 1",
			Ot:     "Jer 1-2",
		},
		{
			Gospel: "John 1:19-28",
			Nt:     "James 1:12-18",
			Wisdom: "Ecc 2:1-16",
			Ot:     "Jer 3-4",
		},
		{
			Gospel: "John 1:29-34",
			Nt:     "James 1:19-27",
			Wisdom: "Ecc 2:17-26",
			Ot:     "Jer 5-6",
		},
		{
			Gospel: "John 1:35-42",
			Nt:     "James 2:1-13",
			Wisdom: "Ecc 3:1-15",
			Ot:     "Jer 7-9",
		},
		{
			Gospel: "John 1:43-51",
			Nt:     "James 2:14-26",
			Wisdom: "Ecc 3:16-22",
			Ot:     "Jer 10-11",
		},
		{
			Gospel: "John 2:1-11",
			Nt:     "James 3:1-12",
			Wisdom: "Ecc 4",
			Ot:     "Jer 12-13",
		},
		{
			Gospel: "John 2:12-25",
			Nt:     "James 3:13-18",
			Wisdom: "Ecc 5",
			Ot:     "Jer 14-15",
		},
		{
			Gospel: "John 3:1-15",
			Nt:     "James 4:1-10",
			Wisdom: "Ecc 6",
			Ot:     "Jer 16-18",
		},
		{
			Gospel: "John 3:16-21",
			Nt:     "James 4:11-17",
			Wisdom: "Ecc 7:1-14",
			Ot:     "Jer 19-22",
		},
		{
			Gospel: "John 3:22-36",
			Nt:     "James 5:1-6",
			Wisdom: "Ecc 7:15-29",
			Ot:     "Jer 23-25",
		},
		{
			Gospel: "John 4:1-14",
			Nt:     "James 5:7-12",
			Wisdom: "Ecc 8",
			Ot:     "Jer 26-29",
		},
		{
			Gospel: "John 4:15-26",
			Nt:     "James 5:13-20",
			Wisdom: "Ecc 9",
			Ot:     "Jer 30-31",
		},
		{
			Gospel: "John 4:27-42",
			Nt:     "1 Peter 1:1-9",
			Wisdom: "Ecc 10",
			Ot:     "Jer 32-34",
		},
		{
			Gospel: "John 4:43-54",
			Nt:     "1 Peter 1:10-16",
			Wisdom: "Ecc 11",
			Ot:     "Jer 35-38",
		},
		{
			Gospel: "John 5:1-15",
			Nt:     "1 Peter 1:17-25",
			Wisdom: "Ecc 12",
			Ot:     "Jer 39-43",
		},
		{
			Gospel: "John 5:16-30",
			Nt:     "1 Peter 2:1-8",
			Wisdom: "Song 1",
			Ot:     "Jer 44-46",
		},
		{
			Gospel: "John 5:31-47",
			Nt:     "1 Peter 2:9-17",
			Wisdom: "Song 2",
			Ot:     "Jer 47-48",
		},
		{
			Gospel: "John 6:1-15",
			Nt:     "1 Peter 2:18-25",
			Wisdom: "Song 3",
			Ot:     "Jer 49",
		},
		{
			Gospel: "John 6:16-24",
			Nt:     "1 Peter 3:1-7",
			Wisdom: "Song 4:1-7",
			Ot:     "Jer 50",
		},
		{
			Gospel: "John 6:25-40",
			Nt:     "1 Peter 3:8-12",
			Wisdom: "Song 4:8-16",
			Ot:     "Jer 51",
		},
		{
			Gospel: "John 6:41-59",
			Nt:     "1 Peter 3:13-22",
			Wisdom: "Song 5",
			Ot:     "Jer 52",
		},
		{
			Gospel: "John 6:60-71",
			Nt:     "1 Peter 4:1-11",
			Wisdom: "Song 6",
			Ot:     "Lam 1",
		},
		{
			Gospel: "John 7:1-13",
			Nt:     "1 Peter 4:12-19",
			Wisdom: "Song 7",
			Ot:     "Lam 2",
		},
		{
			Gospel: "John 7:14-24",
			Nt:     "1 Peter 5:1-7",
			Wisdom: "Song 8:1-7",
			Ot:     "Lam 3",
		},
		{
			Gospel: "John 7:25-36",
			Nt:     "1 Peter 5:8-14",
			Wisdom: "Song 8:8-14",
			Ot:     "Lam 4-5",
		},
	},
	[]day{
		{
			Gospel: "John 7:37-44",
			Nt:     "2 Peter 1:1-11",
			Wisdom: "Job 1",
			Ot:     "Ez 1-3",
		},
		{
			Gospel: "John 7:45-53",
			Nt:     "2 Peter 1:12-21",
			Wisdom: "Job 2",
			Ot:     "Ez 4-8",
		},
		{
			Gospel: "John 8:1-11",
			Nt:     "2 Peter 2:1-9",
			Wisdom: "Job 3",
			Ot:     "Ez 9-12",
		},
		{
			Gospel: "John 8:12-20",
			Nt:     "2 Peter 2:10-16",
			Wisdom: "Job 4",
			Ot:     "Ez 13-15",
		},
		{
			Gospel: "John 8:21-30",
			Nt:     "2 Peter 2:17-22",
			Wisdom: "Job 5",
			Ot:     "Ez 16",
		},
		{
			Gospel: "John 8:31-47",
			Nt:     "2 Peter 3:1-9",
			Wisdom: "Job 6",
			Ot:     "Ez 17-19",
		},
		{
			Gospel: "John 8:48-59",
			Nt:     "2 Peter 3:10-18",
			Wisdom: "Job 7",
			Ot:     "Ez 20-21",
		},
		{
			Gospel: "John 9:1-12",
			Nt:     "1 John 1:1-4",
			Wisdom: "Job 8",
			Ot:     "Ez 22-23",
		},
		{
			Gospel: "John 9:13-25",
			Nt:     "1 John 1:5-10",
			Wisdom: "Job 9:1-20",
			Ot:     "Ez 24-26",
		},
		{
			Gospel: "John 9:26-41",
			Nt:     "1 John 2:1-11",
			Wisdom: "Job 9:21-35",
			Ot:     "Ez 27-28",
		},
		{
			Gospel: "John 10:1-10",
			Nt:     "1 John 2:12-17",
			Wisdom: "Job 10",
			Ot:     "Ez 29-30",
		},
		{
			Gospel: "John 10:11-21",
			Nt:     "1 John 2:18-23",
			Wisdom: "Job 11",
			Ot:     "Ez 31-32",
		},
		{
			Gospel: "John 10:22-42",
			Nt:     "1 John 2:24-29",
			Wisdom: "Job 12",
			Ot:     "Ez 33-34",
		},
		{
			Gospel: "John 11:1-16",
			Nt:     "1 John 3:1-10",
			Wisdom: "Job 13",
			Ot:     "Ez 35-37",
		},
		{
			Gospel: "John 11:17-37",
			Nt:     "1 John 3:11-18",
			Wisdom: "Job 14",
			Ot:     "Ez 38-39",
		},
		{
			Gospel: "John 11:38-44",
			Nt:     "1 John 3:19-24",
			Wisdom: "Job 15:1-16",
			Ot:     "Ez 40-41",
		},
		{
			Gospel: "John 11:45-57",
			Nt:     "1 John 4:1-6",
			Wisdom: "Job 15:17-35",
			Ot:     "Ez 42-44",
		},
		{
			Gospel: "John 12:1-11",
			Nt:     "1 John 4:7-21",
			Wisdom: "Job 16",
			Ot:     "Ez 45-47",
		},
		{
			Gospel: "John 12:12-19",
			Nt:     "1 John 5:1-12",
			Wisdom: "Job 17",
			Ot:     "Ez 48",
		},
		{
			Gospel: "John 12:20-36",
			Nt:     "1 John 5:13-21",
			Wisdom: "Job 18",
			Ot:     "Dan 1-2",
		},
		{
			Gospel: "John 12:37-50",
			Nt:     "2 John 1-13",
			Wisdom: "Job 19",
			Ot:     "Dan 3-4",
		},
		{
			Gospel: "John 13:1-11",
			Nt:     "3 John 1-14",
			Wisdom: "Job 20",
			Ot:     "Dan 5-6",
		},
		{
			Gospel: "John 13:12-17",
			Nt:     "Jude 1-7",
			Wisdom: "Job 21:1-21",
			Ot:     "Dan 7-8",
		},
		{
			Gospel: "John 13:18-30",
			Nt:     "Jude 8-16",
			Wisdom: "Job 21:22-34",
			Ot:     "Dan 9",
		},
		{
			Gospel: "John 13:31-38",
			Nt:     "Jude 17-25",
			Wisdom: "Job 22",
			Ot:     "Dan 10-12",
		},
	},
	[]day{
		{
			Gospel: "John 14:1-14",
			Nt:     "Rev 1:1-8",
			Wisdom: "Job 23",
			Ot:     "Hos 1-3",
		},
		{
			Gospel: "John 14:15-21",
			Nt:     "Rev 1:9-20",
			Wisdom: "Job 24",
			Ot:     "Hos 4-6",
		},
		{
			Gospel: "John 14:22-31",
			Nt:     "Rev 2:1-17",
			Wisdom: "Job 25-26",
			Ot:     "Hos 7-8",
		},
		{
			Gospel: "John 15:1-8",
			Nt:     "Rev 2:18-29",
			Wisdom: "Job 27",
			Ot:     "Hos 9-12",
		},
		{
			Gospel: "John 15:9-17",
			Nt:     "Rev 3:1-13",
			Wisdom: "Job 28",
			Ot:     "Hos 13-14",
		},
		{
			Gospel: "John 15:18-27",
			Nt:     "Rev 3:14-22",
			Wisdom: "Job 29",
			Ot:     "Joel 1",
		},
		{
			Gospel: "John 16:1-11",
			Nt:     "Rev 4",
			Wisdom: "Job 30",
			Ot:     "Joel 2-3",
		},
		{
			Gospel: "John 16:12-24",
			Nt:     "Rev 5",
			Wisdom: "Job 31:1-23",
			Ot:     "Amos 1-2",
		},
		{
			Gospel: "John 16:25-33",
			Nt:     "Rev 6",
			Wisdom: "Job 31:24-40",
			Ot:     "Amos 3-4",
		},
		{
			Gospel: "John 17:1-5",
			Nt:     "Rev 7",
			Wisdom: "Job 32",
			Ot:     "Amos 5-6",
		},
		{
			Gospel: "John 17:6-19",
			Nt:     "Rev 8",
			Wisdom: "Job 33:1-11",
			Ot:     "Amos 7-9",
		},
		{
			Gospel: "John 17:20-26",
			Nt:     "Rev 9",
			Wisdom: "Job 33:12-33",
			Ot:     "Obadiah 1-21",
		},
		{
			Gospel: "John 18:1-18",
			Nt:     "Rev 10",
			Wisdom: "Job 34:1-20",
			Ot:     "Jonah 1-4",
		},
		{
			Gospel: "John 18:19-27",
			Nt:     "Rev 11",
			Wisdom: "Job 34:21-37",
			Ot:     "Micah 1-3",
		},
		{
			Gospel: "John 18:28-40",
			Nt:     "Rev 12",
			Wisdom: "Job 35",
			Ot:     "Micah 4-5",
		},
		{
			Gospel: "John 19:1-16",
			Nt:     "Rev 13",
			Wisdom: "Job 36:1-15",
			Ot:     "Micah 6-7",
		},
		{
			Gospel: "John 19:17-27",
			Nt:     "Rev 14",
			Wisdom: "Job 36:16-33",
			Ot:     "Nahum 1-3",
		},
		{
			Gospel: "John 19:28-37",
			Nt:     "Rev 15",
			Wisdom: "Job 37",
			Ot:     "Hab 1-3",
		},
		{
			Gospel: "John 19:38-42",
			Nt:     "Rev 16",
			Wisdom: "Job 38:1-21",
			Ot:     "Zeph 1-2",
		},
		{
			Gospel: "John 20:1-9",
			Nt:     "Rev 17",
			Wisdom: "Job 38:22-41",
			Ot:     "Zeph 3",
		},
		{
			Gospel: "John 20:10-18",
			Nt:     "Rev 18",
			Wisdom: "Job 39",
			Ot:     "Hag 1-2",
		},
		{
			Gospel: "John 20:19-23",
			Nt:     "Rev 19",
			Wisdom: "Job 40",
			Ot:     "Zech 1-5",
		},
		{
			Gospel: "John 20:24-31",
			Nt:     "Rev 20",
			Wisdom: "Job 41:1-11",
			Ot:     "Zech 6-9",
		},
		{
			Gospel: "John 21:1-14",
			Nt:     "Rev 21",
			Wisdom: "Job 41:12-34",
			Ot:     "Zech 10-14",
		},
		{
			Gospel: "John 21:15-25",
			Nt:     "Rev 22",
			Wisdom: "Job 42",
			Ot:     "Mal 1-4",
		},
	},
}