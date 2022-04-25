package adventofcode

import "testing"

var lines []string = []string{
	"452,244 -> 452,303",
	"958,109 -> 958,639",
	"809,31 -> 778,31",
	"927,139 -> 917,139",
	"56,298 -> 273,298",
	"959,53 -> 959,739",
	"80,255 -> 434,255",
	"750,500 -> 73,500",
	"528,920 -> 428,820",
	"455,552 -> 289,718",
	"800,363 -> 800,25",
	"688,745 -> 688,696",
	"364,41 -> 364,351",
	"693,642 -> 693,67",
	"564,143 -> 755,143",
	"622,682 -> 622,869",
	"311,515 -> 311,120",
	"210,255 -> 210,363",
	"66,550 -> 540,550",
	"243,485 -> 381,623",
	"321,712 -> 680,712",
	"549,501 -> 70,22",
	"357,893 -> 823,893",
	"797,882 -> 536,621",
	"389,724 -> 389,329",
	"859,75 -> 859,591",
	"195,618 -> 161,618",
	"508,668 -> 974,202",
	"543,871 -> 889,525",
	"99,659 -> 371,387",
	"602,553 -> 547,553",
	"98,930 -> 98,683",
	"741,522 -> 895,522",
	"143,636 -> 683,636",
	"56,751 -> 56,555",
	"696,117 -> 784,117",
	"625,827 -> 277,827",
	"908,633 -> 213,633",
	"482,290 -> 256,290",
	"717,566 -> 717,923",
	"258,34 -> 728,34",
	"862,54 -> 862,869",
	"57,983 -> 987,53",
	"749,712 -> 715,678",
	"342,291 -> 147,291",
	"48,308 -> 48,600",
	"44,749 -> 823,749",
	"671,533 -> 944,533",
	"848,804 -> 315,804",
	"960,310 -> 980,310",
	"730,426 -> 730,330",
	"85,633 -> 690,28",
	"432,791 -> 692,531",
	"217,128 -> 217,763",
	"205,275 -> 781,851",
	"207,329 -> 201,329",
	"338,724 -> 338,783",
	"434,89 -> 501,89",
	"812,298 -> 812,480",
	"299,190 -> 752,190",
	"441,301 -> 441,872",
	"643,913 -> 643,21",
	"518,204 -> 518,442",
	"730,601 -> 930,601",
	"352,268 -> 766,268",
	"444,55 -> 444,372",
	"240,914 -> 656,498",
	"507,858 -> 507,581",
	"246,805 -> 805,805",
	"955,960 -> 381,386",
	"828,128 -> 818,128",
	"218,775 -> 557,436",
	"896,529 -> 173,529",
	"523,338 -> 581,338",
	"972,203 -> 770,203",
	"515,635 -> 728,848",
	"852,380 -> 852,469",
	"963,829 -> 688,829",
	"320,161 -> 234,161",
	"517,531 -> 919,933",
	"971,75 -> 578,75",
	"244,666 -> 244,308",
	"640,311 -> 164,787",
	"373,311 -> 373,611",
	"50,784 -> 653,181",
	"770,93 -> 96,767",
	"30,19 -> 887,876",
	"964,860 -> 447,343",
	"436,645 -> 436,813",
	"10,11 -> 982,983",
	"907,638 -> 907,979",
	"453,864 -> 453,246",
	"578,132 -> 757,132",
	"55,456 -> 548,456",
	"127,626 -> 493,626",
	"816,14 -> 748,14",
	"938,873 -> 938,478",
	"315,374 -> 902,961",
	"137,386 -> 137,478",
	"121,954 -> 702,373",
	"92,874 -> 714,252",
	"521,660 -> 34,660",
	"677,401 -> 531,401",
	"55,278 -> 331,278",
	"400,263 -> 426,237",
	"236,974 -> 236,545",
	"579,30 -> 579,198",
	"886,257 -> 320,823",
	"53,67 -> 922,936",
	"374,773 -> 470,773",
	"241,115 -> 241,61",
	"206,252 -> 519,565",
	"388,617 -> 388,778",
	"423,315 -> 978,315",
	"616,759 -> 174,317",
	"770,511 -> 770,487",
	"118,258 -> 309,258",
	"320,908 -> 320,433",
	"71,358 -> 50,337",
	"37,461 -> 175,323",
	"374,758 -> 374,96",
	"268,139 -> 268,785",
	"72,484 -> 364,484",
	"937,286 -> 574,286",
	"554,602 -> 554,475",
	"677,863 -> 245,863",
	"975,33 -> 54,954",
	"557,960 -> 560,957",
	"989,892 -> 175,78",
	"447,47 -> 830,430",
	"407,718 -> 634,945",
	"418,304 -> 418,103",
	"890,278 -> 890,915",
	"449,619 -> 243,825",
	"798,661 -> 137,661",
	"796,251 -> 67,980",
	"122,88 -> 122,518",
	"505,407 -> 695,217",
	"447,961 -> 447,524",
	"608,708 -> 390,708",
	"110,636 -> 878,636",
	"376,517 -> 376,409",
	"400,213 -> 767,213",
	"369,220 -> 369,203",
	"141,934 -> 921,154",
	"808,243 -> 808,983",
	"529,923 -> 529,438",
	"399,324 -> 350,324",
	"632,83 -> 276,83",
	"254,632 -> 812,74",
	"981,594 -> 453,594",
	"30,893 -> 472,893",
	"22,355 -> 614,355",
	"519,732 -> 519,911",
	"392,425 -> 392,710",
	"987,953 -> 987,82",
	"519,108 -> 370,108",
	"52,897 -> 52,449",
	"17,83 -> 913,979",
	"612,305 -> 118,799",
	"225,365 -> 225,639",
	"878,928 -> 45,95",
	"238,279 -> 908,279",
	"352,27 -> 352,572",
	"72,11 -> 347,286",
	"332,195 -> 332,189",
	"273,726 -> 204,726",
	"401,817 -> 498,817",
	"118,170 -> 118,675",
	"924,868 -> 528,472",
	"721,948 -> 886,948",
	"735,379 -> 331,379",
	"256,298 -> 256,187",
	"600,886 -> 314,886",
	"937,275 -> 937,385",
	"840,777 -> 92,29",
	"844,96 -> 844,691",
	"833,597 -> 833,762",
	"475,138 -> 475,207",
	"32,984 -> 675,984",
	"101,528 -> 101,497",
	"884,399 -> 302,981",
	"178,124 -> 670,616",
	"363,84 -> 527,84",
	"744,784 -> 336,784",
	"588,963 -> 588,543",
	"298,525 -> 626,525",
	"610,496 -> 265,496",
	"648,294 -> 460,106",
	"859,646 -> 859,867",
	"469,48 -> 987,566",
	"273,689 -> 103,859",
	"556,745 -> 440,745",
	"44,923 -> 941,26",
	"569,777 -> 781,565",
	"791,408 -> 276,923",
	"289,185 -> 289,608",
	"866,783 -> 760,783",
	"302,153 -> 420,35",
	"45,545 -> 479,979",
	"309,54 -> 850,54",
	"36,48 -> 973,985",
	"698,671 -> 129,102",
	"778,339 -> 587,339",
	"114,566 -> 114,621",
	"747,450 -> 583,286",
	"558,698 -> 558,912",
	"60,43 -> 987,970",
	"770,243 -> 770,490",
	"58,811 -> 58,131",
	"839,364 -> 839,614",
	"833,633 -> 833,208",
	"685,441 -> 600,526",
	"919,894 -> 769,894",
	"478,799 -> 601,922",
	"957,957 -> 237,957",
	"472,169 -> 472,96",
	"269,185 -> 269,117",
	"582,187 -> 582,460",
	"243,506 -> 722,506",
	"702,191 -> 125,768",
	"542,63 -> 542,398",
	"379,222 -> 920,763",
	"930,658 -> 930,724",
	"524,33 -> 524,739",
	"987,949 -> 50,12",
	"976,71 -> 852,71",
	"46,59 -> 921,934",
	"906,526 -> 906,623",
	"621,455 -> 621,590",
	"954,55 -> 954,499",
	"740,52 -> 740,448",
	"185,959 -> 937,207",
	"563,349 -> 261,651",
	"852,938 -> 114,938",
	"972,22 -> 16,978",
	"249,142 -> 199,142",
	"282,469 -> 282,831",
	"667,686 -> 685,686",
	"901,95 -> 901,578",
	"186,903 -> 336,903",
	"662,248 -> 99,811",
	"485,248 -> 833,248",
	"12,154 -> 734,154",
	"87,270 -> 760,270",
	"926,191 -> 799,191",
	"796,793 -> 451,448",
	"443,292 -> 625,292",
	"944,62 -> 944,859",
	"988,489 -> 374,489",
	"693,279 -> 693,55",
	"893,617 -> 893,247",
	"168,905 -> 910,905",
	"850,429 -> 773,429",
	"943,13 -> 364,13",
	"409,880 -> 409,806",
	"620,646 -> 567,699",
	"194,949 -> 194,660",
	"775,682 -> 775,456",
	"741,161 -> 581,161",
	"889,91 -> 42,938",
	"725,495 -> 273,43",
	"211,22 -> 897,22",
	"314,604 -> 562,604",
	"299,193 -> 914,193",
	"240,388 -> 612,16",
	"367,752 -> 367,114",
	"330,34 -> 931,34",
	"809,909 -> 809,757",
	"912,140 -> 224,140",
	"545,804 -> 503,804",
	"848,758 -> 848,638",
	"917,226 -> 538,226",
	"43,34 -> 43,97",
	"47,54 -> 47,435",
	"32,52 -> 843,863",
	"200,361 -> 398,163",
	"466,388 -> 705,388",
	"928,127 -> 104,951",
	"869,940 -> 869,787",
	"697,563 -> 167,563",
	"754,828 -> 754,141",
	"661,100 -> 968,100",
	"582,44 -> 582,182",
	"876,550 -> 876,268",
	"737,240 -> 27,950",
	"664,730 -> 633,730",
	"244,805 -> 228,805",
	"226,605 -> 211,605",
	"911,524 -> 548,524",
	"814,256 -> 889,181",
	"432,944 -> 432,235",
	"445,25 -> 445,499",
	"233,619 -> 91,477",
	"151,248 -> 880,977",
	"267,149 -> 317,149",
	"968,346 -> 968,371",
	"665,617 -> 665,770",
	"518,625 -> 518,579",
	"451,458 -> 628,635",
	"764,442 -> 255,442",
	"240,363 -> 680,803",
	"443,153 -> 863,153",
	"674,880 -> 973,581",
	"876,236 -> 876,817",
	"569,335 -> 569,907",
	"74,400 -> 74,679",
	"495,804 -> 827,472",
	"182,76 -> 832,726",
	"376,370 -> 986,370",
	"986,233 -> 541,233",
	"386,379 -> 386,837",
	"624,637 -> 454,637",
	"844,866 -> 179,201",
	"392,89 -> 627,89",
	"559,326 -> 525,326",
	"681,188 -> 418,188",
	"796,386 -> 796,637",
	"35,72 -> 952,989",
	"735,873 -> 407,873",
	"756,238 -> 328,238",
	"872,284 -> 394,284",
	"157,873 -> 243,873",
	"621,657 -> 913,657",
	"13,42 -> 13,598",
	"831,725 -> 831,270",
	"944,51 -> 87,908",
	"409,820 -> 913,316",
	"250,323 -> 622,323",
	"435,397 -> 435,409",
	"691,198 -> 691,952",
	"302,352 -> 872,352",
	"370,790 -> 362,790",
	"114,74 -> 926,74",
	"561,273 -> 561,169",
	"240,658 -> 240,584",
	"677,357 -> 557,477",
	"480,796 -> 544,732",
	"981,568 -> 319,568",
	"63,928 -> 957,34",
	"715,145 -> 81,779",
	"918,554 -> 292,554",
	"943,550 -> 943,23",
	"425,477 -> 277,329",
	"428,718 -> 428,612",
	"213,647 -> 244,616",
	"549,610 -> 549,39",
	"433,330 -> 433,922",
	"903,921 -> 148,166",
	"357,706 -> 913,150",
	"990,254 -> 990,368",
	"937,47 -> 937,34",
	"868,86 -> 260,694",
	"919,905 -> 140,126",
	"256,918 -> 980,194",
	"880,346 -> 331,895",
	"706,392 -> 706,328",
	"283,986 -> 758,511",
	"967,106 -> 123,106",
	"486,92 -> 486,47",
	"298,338 -> 888,928",
	"665,767 -> 987,767",
	"139,898 -> 139,94",
	"257,808 -> 257,388",
	"865,631 -> 865,624",
	"168,913 -> 168,984",
	"848,936 -> 197,285",
	"195,249 -> 764,249",
	"117,900 -> 249,900",
	"813,130 -> 10,933",
	"693,81 -> 693,554",
	"595,448 -> 987,840",
	"293,140 -> 904,751",
	"198,289 -> 234,289",
	"714,671 -> 989,396",
	"161,600 -> 161,465",
	"458,955 -> 807,955",
	"485,124 -> 485,835",
	"352,911 -> 691,911",
	"896,949 -> 103,156",
	"618,171 -> 618,847",
	"563,887 -> 44,887",
	"293,801 -> 293,492",
	"903,714 -> 651,714",
	"959,18 -> 265,712",
	"28,134 -> 851,957",
	"888,374 -> 647,374",
	"184,566 -> 184,411",
	"977,244 -> 926,193",
	"864,189 -> 427,626",
	"687,262 -> 700,275",
	"86,82 -> 891,887",
	"517,779 -> 517,599",
	"850,59 -> 256,653",
	"595,352 -> 595,142",
	"62,191 -> 618,747",
	"56,395 -> 124,463",
	"902,138 -> 418,138",
	"603,131 -> 185,549",
	"446,546 -> 446,397",
	"278,788 -> 916,788",
	"153,19 -> 707,573",
	"580,468 -> 502,390",
	"278,794 -> 447,963",
	"260,495 -> 315,550",
	"648,392 -> 726,392",
	"232,457 -> 633,858",
	"475,697 -> 740,697",
	"526,186 -> 526,588",
	"868,287 -> 868,668",
	"88,57 -> 88,984",
	"539,963 -> 539,859",
	"499,214 -> 963,678",
	"105,16 -> 937,848",
	"566,158 -> 566,26",
	"675,915 -> 675,555",
	"702,123 -> 360,123",
	"37,47 -> 120,130",
	"826,787 -> 79,787",
	"232,33 -> 833,634",
	"455,775 -> 455,919",
	"300,488 -> 300,92",
	"36,84 -> 36,514",
	"700,916 -> 700,391",
	"710,585 -> 710,14",
	"989,938 -> 120,69",
	"437,514 -> 437,277",
	"400,358 -> 400,623",
	"482,908 -> 802,908",
	"869,842 -> 53,26",
	"531,560 -> 180,560",
	"18,312 -> 18,223",
	"541,888 -> 541,438",
	"631,751 -> 631,238",
	"820,208 -> 192,836",
	"976,453 -> 561,38",
	"118,544 -> 318,744",
	"283,818 -> 283,138",
	"876,623 -> 327,623",
	"11,855 -> 283,855",
	"438,460 -> 145,167",
	"788,379 -> 250,379",
	"376,130 -> 890,130",
	"616,747 -> 538,825",
	"906,403 -> 906,306",
	"710,371 -> 710,200",
	"871,917 -> 629,917",
	"751,103 -> 345,103",
	"374,608 -> 641,608",
	"481,902 -> 198,902",
	"512,686 -> 590,764",
	"187,134 -> 949,896",
	"628,58 -> 170,516",
	"54,370 -> 23,401",
	"391,718 -> 523,718",
	"617,637 -> 448,637",
	"108,208 -> 233,208",
	"57,478 -> 633,478",
	"122,693 -> 734,693",
	"980,537 -> 980,639",
	"134,448 -> 134,439",
	"81,812 -> 515,812",
	"483,133 -> 215,133",
	"68,250 -> 68,476",
	"204,101 -> 204,934",
	"525,626 -> 924,227",
	"783,519 -> 373,929",
	"813,166 -> 813,837",
	"12,988 -> 988,12",
	"49,180 -> 46,180",
	"782,585 -> 782,194",
	"798,98 -> 459,437",
	"904,83 -> 333,654",
	"404,417 -> 404,819",
	"289,228 -> 431,86",
	"281,876 -> 964,193",
	"66,15 -> 66,570",
	"134,159 -> 784,809",
	"356,684 -> 442,598",
	"211,624 -> 572,624",
	"19,985 -> 986,18",
	"978,429 -> 978,791",
	"181,955 -> 832,304",
	"566,859 -> 566,281",
	"391,232 -> 535,232",
	"798,848 -> 17,67",
	"530,203 -> 530,932",
	"502,562 -> 287,562",
	"710,499 -> 451,758",
	"886,420 -> 163,420",
	"858,289 -> 632,515",
	"939,441 -> 939,280",
	"293,220 -> 444,220",
	"363,676 -> 159,676",
	"983,24 -> 46,961",
	"725,585 -> 154,14",
	"236,453 -> 467,453",
	"806,395 -> 806,557",
	"149,583 -> 149,200",
	"858,704 -> 858,272",
}

func TestDay5(t *testing.T) {
	expectedRes := 18423
	res := day5(lines)
	if res != expectedRes {
		t.Errorf("expected answer is %d, what we get is %d", expectedRes, res)
	}
}
