# build service

```
go fmt trademe/... && go build -o ./bin/trademe ./cmd
```

# run service
```
./bin/trademe test1
```
test1 result will show as default but subcommand can also support `test2` or `test3` or `test4` or `test5`

# run service
```
 go test -timeout 5000ms -cover trademe/...
 ```

 ```
 2022/11/26 03:41:10 properties initialised successfully with length 291
2022/11/26 03:41:10 result from test1: {
			"length": 194,
			"properties": [
				{
					"Id": "20",
					"StreetAddress": "1 Edgewood PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 415000
				},
				{
					"Id": "21",
					"StreetAddress": "1 HUNTER CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 385000
				},
				{
					"Id": "74",
					"StreetAddress": "1 Stonebrook DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 690000
				},
				{
					"Id": "60",
					"StreetAddress": "1 SYCAMORE PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 780000
				},
				{
					"Id": "85",
					"StreetAddress": "1 Frederick ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 505000
				},
				{
					"Id": "5",
					"StreetAddress": "1 Kapuka LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 209000
				},
				{
					"Id": "23",
					"StreetAddress": "1 RIMU LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 525000
				},
				{
					"Id": "27",
					"StreetAddress": "1 EELY POINT RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 620000
				},
				{
					"Id": "75",
					"StreetAddress": "1 Hyland ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 560000
				},
				{
					"Id": "77",
					"StreetAddress": "1 Homestead CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1810000
				},
				{
					"Id": "83",
					"StreetAddress": "1 Heuchan LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 670000
				},
				{
					"Id": "44",
					"StreetAddress": "1 Islington PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 240000
				},
				{
					"Id": "46",
					"StreetAddress": "1 Weatherall CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 475000
				},
				{
					"Id": "81",
					"StreetAddress": "1 MILL END",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 650000
				},
				{
					"Id": "18",
					"StreetAddress": "1 Centre CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 355000
				},
				{
					"Id": "62",
					"StreetAddress": "1 Alpha CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 560000
				},
				{
					"Id": "93",
					"StreetAddress": "1 Finch ST",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/16",
					"Value": 600000
				},
				{
					"Id": "30",
					"StreetAddress": "1 Manuka CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 560000
				},
				{
					"Id": "52",
					"StreetAddress": "1 Pearce PL",
					"Town": "Wanaka",
					"ValuationDate": "1/01/16",
					"Value": 735000
				},
				{
					"Id": "11",
					"StreetAddress": "1 WAIMANA PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 840000
				},
				{
					"Id": "13",
					"StreetAddress": "1 Toms WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 2290000
				},
				{
					"Id": "11",
					"StreetAddress": "1 WAIMANA PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 830000
				},
				{
					"Id": "22",
					"StreetAddress": "1 KOWHAI DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 530000
				},
				{
					"Id": "31",
					"StreetAddress": "1 WILEY RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 470000
				},
				{
					"Id": "63",
					"StreetAddress": "1 Coromandel ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 580000
				},
				{
					"Id": "84",
					"StreetAddress": "1 SARGOOD DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 505000
				},
				{
					"Id": "88",
					"StreetAddress": "1 Hikuwai DR",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/16",
					"Value": 340000
				},
				{
					"Id": "89",
					"StreetAddress": "1 Harrier LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1060000
				},
				{
					"Id": "33",
					"StreetAddress": "1 Briar Bank  DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1310000
				},
				{
					"Id": "52",
					"StreetAddress": "1 Pearce PL",
					"Town": "Wanaka",
					"ValuationDate": "1/01/15",
					"Value": 725000
				},
				{
					"Id": "70",
					"StreetAddress": "1 OAKWOOD PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 630000
				},
				{
					"Id": "37",
					"StreetAddress": "1 ROB ROY LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 440000
				},
				{
					"Id": "70",
					"StreetAddress": "1 OAKWOOD PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 640000
				},
				{
					"Id": "28",
					"StreetAddress": "1 LINDSAY PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 445000
				},
				{
					"Id": "32",
					"StreetAddress": "1 Baker GR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 880000
				},
				{
					"Id": "74",
					"StreetAddress": "1 Stonebrook DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 700000
				},
				{
					"Id": "28",
					"StreetAddress": "1 LINDSAY PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 435000
				},
				{
					"Id": "55",
					"StreetAddress": "1 Hogan LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 445000
				},
				{
					"Id": "77",
					"StreetAddress": "1 Homestead CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1800000
				},
				{
					"Id": "91",
					"StreetAddress": "1 Sherwin AVE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/15",
					"Value": 490000
				},
				{
					"Id": "1",
					"StreetAddress": "1 Northburn RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 340000
				},
				{
					"Id": "85",
					"StreetAddress": "1 Frederick ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 515000
				},
				{
					"Id": "65",
					"StreetAddress": "1 Maggies Way",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 425000
				},
				{
					"Id": "3",
					"StreetAddress": "1 Mount Linton AVE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 840000
				},
				{
					"Id": "4",
					"StreetAddress": "1 Kamahi ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 215000
				},
				{
					"Id": "44",
					"StreetAddress": "1 Islington PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 250000
				},
				{
					"Id": "47",
					"StreetAddress": "1 Terranova PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 950000
				},
				{
					"Id": "76",
					"StreetAddress": "1 TAPLEY PADDOCK",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 770000
				},
				{
					"Id": "60",
					"StreetAddress": "1 SYCAMORE PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 770000
				},
				{
					"Id": "20",
					"StreetAddress": "1 Edgewood PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 425000
				},
				{
					"Id": "64",
					"StreetAddress": "1 Niger ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 535000
				},
				{
					"Id": "76",
					"StreetAddress": "1 TAPLEY PADDOCK",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 780000
				},
				{
					"Id": "81",
					"StreetAddress": "1 MILL END",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 660000
				},
				{
					"Id": "69",
					"StreetAddress": "1 MEADOWSTONE DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 710000
				},
				{
					"Id": "80",
					"StreetAddress": "1 LARCH PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 630000
				},
				{
					"Id": "34",
					"StreetAddress": "1 LAKESIDE RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 490000
				},
				{
					"Id": "62",
					"StreetAddress": "1 Alpha CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 550000
				},
				{
					"Id": "73",
					"StreetAddress": "1 Lansdown ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 750000
				},
				{
					"Id": "22",
					"StreetAddress": "1 KOWHAI DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 540000
				},
				{
					"Id": "56",
					"StreetAddress": "1 ARDMORE ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1250000
				},
				{
					"Id": "40",
					"StreetAddress": "1 APOLLO PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 445000
				},
				{
					"Id": "48",
					"StreetAddress": "1 Cliff Wilson ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1260000
				},
				{
					"Id": "94",
					"StreetAddress": "1 Poppy LANE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/16",
					"Value": 455000
				},
				{
					"Id": "26",
					"StreetAddress": "1 AUBREY RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1035000
				},
				{
					"Id": "57",
					"StreetAddress": "1 Bullock Creek LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 11175000
				},
				{
					"Id": "83",
					"StreetAddress": "1 Heuchan LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 660000
				},
				{
					"Id": "14",
					"StreetAddress": "1 MULBERRY LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 475000
				},
				{
					"Id": "17",
					"StreetAddress": "1 Clutha PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 760000
				},
				{
					"Id": "6",
					"StreetAddress": "1 Mohua MEWS",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 610000
				},
				{
					"Id": "82",
					"StreetAddress": "1 Bills WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 810000
				},
				{
					"Id": "33",
					"StreetAddress": "1 Briar Bank  DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1320000
				},
				{
					"Id": "57",
					"StreetAddress": "1 Bullock Creek LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 11185000
				},
				{
					"Id": "72",
					"StreetAddress": "1 Jessies CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 380000
				},
				{
					"Id": "3",
					"StreetAddress": "1 Mount Linton AVE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 830000
				},
				{
					"Id": "9",
					"StreetAddress": "1 Penrith Park DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1300000
				},
				{
					"Id": "45",
					"StreetAddress": "1 Hidden Hills DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1330000
				},
				{
					"Id": "10",
					"StreetAddress": "1 ATHERTON PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 710000
				},
				{
					"Id": "19",
					"StreetAddress": "1 Valley CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 850000
				},
				{
					"Id": "88",
					"StreetAddress": "1 Hikuwai DR",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/15",
					"Value": 330000
				},
				{
					"Id": "61",
					"StreetAddress": "1 FAULKS TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 840000
				},
				{
					"Id": "65",
					"StreetAddress": "1 Maggies Way",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 435000
				},
				{
					"Id": "25",
					"StreetAddress": "1 COLLINS ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 570000
				},
				{
					"Id": "49",
					"StreetAddress": "1 TOTARA TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 510000
				},
				{
					"Id": "59",
					"StreetAddress": "1 Primary LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 480000
				},
				{
					"Id": "75",
					"StreetAddress": "1 Hyland ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 550000
				},
				{
					"Id": "86",
					"StreetAddress": "1 Connell TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 650000
				},
				{
					"Id": "66",
					"StreetAddress": "1 Hollyhock LANE",
					"Town": "QUEENSTOWN",
					"ValuationDate": "1/01/16",
					"Value": 1140000
				},
				{
					"Id": "92",
					"StreetAddress": "1 Hardie PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 880000
				},
				{
					"Id": "2",
					"StreetAddress": "1 Mount Ida PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 340000
				},
				{
					"Id": "46",
					"StreetAddress": "1 Weatherall CL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 485000
				},
				{
					"Id": "94",
					"StreetAddress": "1 Poppy LANE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/15",
					"Value": 445000
				},
				{
					"Id": "12",
					"StreetAddress": "1 ROTO PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1530000
				},
				{
					"Id": "24",
					"StreetAddress": "1 CHERRY CT",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 555000
				},
				{
					"Id": "18",
					"StreetAddress": "1 Centre CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 345000
				},
				{
					"Id": "24",
					"StreetAddress": "1 CHERRY CT",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 545000
				},
				{
					"Id": "39",
					"StreetAddress": "1 Fastness CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 690000
				},
				{
					"Id": "51",
					"StreetAddress": "1 Bovett PL",
					"Town": "Wanaka",
					"ValuationDate": "1/01/15",
					"Value": 545000
				},
				{
					"Id": "82",
					"StreetAddress": "1 Bills WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 800000
				},
				{
					"Id": "90",
					"StreetAddress": "1 Ewing PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 830000
				},
				{
					"Id": "93",
					"StreetAddress": "1 Finch ST",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/15",
					"Value": 590000
				},
				{
					"Id": "96",
					"StreetAddress": "1 Balneaves LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 300000
				},
				{
					"Id": "30",
					"StreetAddress": "1 Manuka CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 570000
				},
				{
					"Id": "67",
					"StreetAddress": "1 ELDERBERRY CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1400000
				},
				{
					"Id": "29",
					"StreetAddress": "1 WINDERS ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 810000
				},
				{
					"Id": "43",
					"StreetAddress": "1 Moncrieff  PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 580000
				},
				{
					"Id": "69",
					"StreetAddress": "1 MEADOWSTONE DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 700000
				},
				{
					"Id": "21",
					"StreetAddress": "1 HUNTER CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 395000
				},
				{
					"Id": "58",
					"StreetAddress": "1 DUNMORE ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1350000
				},
				{
					"Id": "61",
					"StreetAddress": "1 FAULKS TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 830000
				},
				{
					"Id": "36",
					"StreetAddress": "1 Allenby PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 700000
				},
				{
					"Id": "16",
					"StreetAddress": "1 Clearview ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1280000
				},
				{
					"Id": "36",
					"StreetAddress": "1 Allenby PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 690000
				},
				{
					"Id": "38",
					"StreetAddress": "1 Ansted PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 640000
				},
				{
					"Id": "50",
					"StreetAddress": "1 Koru WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 620000
				},
				{
					"Id": "54",
					"StreetAddress": "1 Bob Lee PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 660000
				},
				{
					"Id": "41",
					"StreetAddress": "1 AEOLUS PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 430000
				},
				{
					"Id": "42",
					"StreetAddress": "1 Peak View RDGE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1810000
				},
				{
					"Id": "43",
					"StreetAddress": "1 Moncrieff  PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 590000
				},
				{
					"Id": "49",
					"StreetAddress": "1 TOTARA TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 520000
				},
				{
					"Id": "54",
					"StreetAddress": "1 Bob Lee PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 670000
				},
				{
					"Id": "10",
					"StreetAddress": "1 ATHERTON PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 700000
				},
				{
					"Id": "15",
					"StreetAddress": "1 Range View PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 350000
				},
				{
					"Id": "41",
					"StreetAddress": "1 AEOLUS PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 420000
				},
				{
					"Id": "16",
					"StreetAddress": "1 Clearview ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1290000
				},
				{
					"Id": "38",
					"StreetAddress": "1 Ansted PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 650000
				},
				{
					"Id": "34",
					"StreetAddress": "1 LAKESIDE RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 500000
				},
				{
					"Id": "35",
					"StreetAddress": "1 PLANTATION RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 405000
				},
				{
					"Id": "73",
					"StreetAddress": "1 Lansdown ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 760000
				},
				{
					"Id": "5",
					"StreetAddress": "1 Kapuka LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 199000
				},
				{
					"Id": "14",
					"StreetAddress": "1 MULBERRY LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 465000
				},
				{
					"Id": "79",
					"StreetAddress": "1 Sunrise Bay DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 3050000
				},
				{
					"Id": "97",
					"StreetAddress": "1 Mill Green",
					"Town": "Arrowtown",
					"ValuationDate": "1/01/15",
					"Value": 850000
				},
				{
					"Id": "15",
					"StreetAddress": "1 Range View PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 360000
				},
				{
					"Id": "45",
					"StreetAddress": "1 Hidden Hills DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1340000
				},
				{
					"Id": "51",
					"StreetAddress": "1 Bovett PL",
					"Town": "Wanaka",
					"ValuationDate": "1/01/16",
					"Value": 555000
				},
				{
					"Id": "53",
					"StreetAddress": "1 Ironside DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 630000
				},
				{
					"Id": "7",
					"StreetAddress": "1 Kakapo CT",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 480000
				},
				{
					"Id": "64",
					"StreetAddress": "1 Niger ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 525000
				},
				{
					"Id": "67",
					"StreetAddress": "1 ELDERBERRY CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1390000
				},
				{
					"Id": "7",
					"StreetAddress": "1 Kakapo CT",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 490000
				},
				{
					"Id": "9",
					"StreetAddress": "1 Penrith Park DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1310000
				},
				{
					"Id": "58",
					"StreetAddress": "1 DUNMORE ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1360000
				},
				{
					"Id": "59",
					"StreetAddress": "1 Primary LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 490000
				},
				{
					"Id": "63",
					"StreetAddress": "1 Coromandel ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 590000
				},
				{
					"Id": "96",
					"StreetAddress": "1 Balneaves LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 310000
				},
				{
					"Id": "78",
					"StreetAddress": "1 NORMAN TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 670000
				},
				{
					"Id": "80",
					"StreetAddress": "1 LARCH PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 620000
				},
				{
					"Id": "26",
					"StreetAddress": "1 AUBREY RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1045000
				},
				{
					"Id": "29",
					"StreetAddress": "1 WINDERS ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 820000
				},
				{
					"Id": "55",
					"StreetAddress": "1 Hogan LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 455000
				},
				{
					"Id": "97",
					"StreetAddress": "1 Mill Green",
					"Town": "Arrowtown",
					"ValuationDate": "1/01/16",
					"Value": 860000
				},
				{
					"Id": "13",
					"StreetAddress": "1 Toms WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 2280000
				},
				{
					"Id": "68",
					"StreetAddress": "1 Foxglove HTS",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 2570000
				},
				{
					"Id": "72",
					"StreetAddress": "1 Jessies CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 370000
				},
				{
					"Id": "68",
					"StreetAddress": "1 Foxglove HTS",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 2580000
				},
				{
					"Id": "92",
					"StreetAddress": "1 Hardie PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 890000
				},
				{
					"Id": "35",
					"StreetAddress": "1 PLANTATION RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 395000
				},
				{
					"Id": "42",
					"StreetAddress": "1 Peak View RDGE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1800000
				},
				{
					"Id": "87",
					"StreetAddress": "1 Soho ST",
					"Town": "QUEENSTOWN",
					"ValuationDate": "1/01/15",
					"Value": 370000
				},
				{
					"Id": "71",
					"StreetAddress": "1 MEADOWBROOK PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 705000
				},
				{
					"Id": "86",
					"StreetAddress": "1 Connell TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 660000
				},
				{
					"Id": "2",
					"StreetAddress": "1 Mount Ida PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 330000
				},
				{
					"Id": "32",
					"StreetAddress": "1 Baker GR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 870000
				},
				{
					"Id": "78",
					"StreetAddress": "1 NORMAN TCE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 680000
				},
				{
					"Id": "91",
					"StreetAddress": "1 Sherwin AVE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/16",
					"Value": 500000
				},
				{
					"Id": "95",
					"StreetAddress": "1 Warbler LANE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/16",
					"Value": 470000
				},
				{
					"Id": "4",
					"StreetAddress": "1 Kamahi ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 205000
				},
				{
					"Id": "8",
					"StreetAddress": "1 Mt Gold PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1310000
				},
				{
					"Id": "8",
					"StreetAddress": "1 Mt Gold PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 1320000
				},
				{
					"Id": "25",
					"StreetAddress": "1 COLLINS ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 580000
				},
				{
					"Id": "89",
					"StreetAddress": "1 Harrier LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1050000
				},
				{
					"Id": "31",
					"StreetAddress": "1 WILEY RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 480000
				},
				{
					"Id": "84",
					"StreetAddress": "1 SARGOOD DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 515000
				},
				{
					"Id": "1",
					"StreetAddress": "1 Northburn RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 330000
				},
				{
					"Id": "12",
					"StreetAddress": "1 ROTO PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1520000
				},
				{
					"Id": "17",
					"StreetAddress": "1 Clutha PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 750000
				},
				{
					"Id": "37",
					"StreetAddress": "1 ROB ROY LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 430000
				},
				{
					"Id": "40",
					"StreetAddress": "1 APOLLO PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 435000
				},
				{
					"Id": "90",
					"StreetAddress": "1 Ewing PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 840000
				},
				{
					"Id": "27",
					"StreetAddress": "1 EELY POINT RD",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 610000
				},
				{
					"Id": "39",
					"StreetAddress": "1 Fastness CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 700000
				},
				{
					"Id": "47",
					"StreetAddress": "1 Terranova PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 960000
				},
				{
					"Id": "79",
					"StreetAddress": "1 Sunrise Bay DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 3060000
				},
				{
					"Id": "66",
					"StreetAddress": "1 Hollyhock LANE",
					"Town": "QUEENSTOWN",
					"ValuationDate": "1/01/15",
					"Value": 1130000
				},
				{
					"Id": "95",
					"StreetAddress": "1 Warbler LANE",
					"Town": "ALBERT TOWN",
					"ValuationDate": "1/01/15",
					"Value": 460000
				},
				{
					"Id": "6",
					"StreetAddress": "1 Mohua MEWS",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 620000
				},
				{
					"Id": "87",
					"StreetAddress": "1 Soho ST",
					"Town": "QUEENSTOWN",
					"ValuationDate": "1/01/16",
					"Value": 380000
				},
				{
					"Id": "71",
					"StreetAddress": "1 MEADOWBROOK PL",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 695000
				},
				{
					"Id": "50",
					"StreetAddress": "1 Koru WAY",
					"Town": "WANAKA",
					"ValuationDate": "1/01/16",
					"Value": 630000
				},
				{
					"Id": "19",
					"StreetAddress": "1 Valley CRES",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 840000
				},
				{
					"Id": "23",
					"StreetAddress": "1 RIMU LANE",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 515000
				},
				{
					"Id": "48",
					"StreetAddress": "1 Cliff Wilson ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1250000
				},
				{
					"Id": "53",
					"StreetAddress": "1 Ironside DR",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 620000
				},
				{
					"Id": "56",
					"StreetAddress": "1 ARDMORE ST",
					"Town": "WANAKA",
					"ValuationDate": "1/01/15",
					"Value": 1240000
				}
			]
		}

```