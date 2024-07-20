package parser

import (
	_i1 "github.com/dcaiafa/lox/internal/ast"
	_i0 "github.com/dcaiafa/loxlex/simplelexer"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8,
	8, 8, 9, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14, 14,
	14, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 24, 24,
	25, 25, 25, 25, 26, 26, 27, 28, 28, 29, 29, 29, 29, 30,
	31, 32, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39,
	39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46,
	46, 47, 47, 48, 48, 49, 49, 50, 50, 51, 51, 52, 52, 53,
	53, 54, 54, 55, 55, 56, 56, 57, 57,
}

var _termCounts = []int32{
	1, 2, 1, 1, 1, 3, 1, 1, 5, 2, 1, 2, 1, 1,
	1, 1, 6, 1, 1, 1, 4, 4, 3, 1, 1, 1, 1, 1,
	1, 1, 5, 5, 4, 5, 3, 1, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 3, 3, 1, 4, 1, 1, 1, 1, 1, 1, 1,
	4, 1, 4, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2,
	1, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 1, 0, 2, 1, 1, 0, 2, 1, 2, 1, 3, 1, 2,
	1, 1, 0, 1, 0, 2, 1, 1, 0,
}

var _actions = []int32{
	150, 161, 164, 173, 180, 189, 192, 195, 198, 205, 212, 219, 222, 229,
	238, 251, 270, 277, 290, 293, 296, 309, 322, 329, 342, 345, 356, 359,
	362, 365, 386, 407, 428, 447, 466, 473, 492, 513, 532, 553, 556, 345,
	559, 574, 605, 345, 636, 639, 642, 675, 706, 717, 732, 763, 788, 813,
	816, 821, 826, 831, 844, 863, 874, 706, 887, 898, 901, 912, 915, 926,
	929, 940, 951, 962, 973, 345, 984, 1005, 345, 1010, 1035, 1060, 1085, 1110,
	1135, 1160, 1163, 1168, 1173, 1178, 1183, 1210, 1237, 1240, 1267, 1294, 1299, 1322,
	1339, 1356, 1369, 1372, 1385, 1388, 1409, 1412, 1443, 1474, 1481, 1488, 1495, 1502,
	1507, 1510, 1521, 1536, 863, 1549, 1552, 1555, 1560, 1565, 1582, 1599, 1616, 1633,
	1650, 1667, 1676, 1695, 1716, 1737, 1770, 1777, 1780, 1783, 1786, 1799, 1802, 1805,
	1808, 1819, 1830, 1667, 1835, 1838, 1841, 1844, 1849, 1854, 10, 0, -60, 1,
	1, 17, -60, 40, 2, 16, -60, 2, 0, -2, 8, 0, -62, 17,
	-62, 40, -62, 16, -62, 6, 0, -64, 17, 6, 16, 7, 8, 0,
	-59, 17, -59, 40, 13, 16, -59, 2, 0, 2147483647, 2, 40, 15, 2,
	40, 14, 6, 0, -4, 17, -4, 16, -4, 6, 0, -3, 17, -3,
	16, -3, 6, 0, -66, 17, -66, 16, -66, 2, 0, -1, 6, 0,
	-63, 17, 6, 16, 7, 8, 0, -61, 17, -61, 40, -61, 16, -61,
	12, 0, -68, 33, -72, 17, -68, 40, 17, 16, -68, 18, 18, 18,
	0, -82, 31, 24, 21, 25, 33, 26, 17, -82, 20, 27, 22, 28,
	40, 29, 16, -82, 6, 0, -65, 17, -65, 16, -65, 12, 0, -7,
	33, -7, 17, -7, 40, -7, 16, -7, 18, -7, 2, 33, -71, 2,
	33, 39, 12, 0, -6, 33, -6, 17, -6, 40, -6, 16, -6, 18,
	-6, 12, 0, -70, 33, -70, 17, -70, 40, -70, 16, -70, 18, -70,
	6, 0, -5, 17, -5, 16, -5, 12, 0, -67, 33, -72, 17, -67,
	40, 17, 16, -67, 18, 18, 2, 33, 56, 10, 33, 43, 35, 44,
	36, -102, 8, 45, 7, 46, 2, 3, 41, 2, 33, 55, 2, 33,
	40, 20, 6, -29, 0, -29, 31, -29, 21, -29, 33, -29, 17, -29,
	20, -29, 22, -29, 40, -29, 16, -29, 20, 6, -28, 0, -28, 31,
	-28, 21, -28, 33, -28, 17, -28, 20, -28, 22, -28, 40, -28, 16,
	-28, 20, 6, -26, 0, -26, 31, -26, 21, -26, 33, -26, 17, -26,
	20, -26, 22, -26, 40, -26, 16, -26, 18, 0, -24, 31, -24, 21,
	-24, 33, -24, 17, -24, 20, -24, 22, -24, 40, -24, 16, -24, 18,
	0, -84, 31, -84, 21, -84, 33, -84, 17, -84, 20, -84, 22, -84,
	40, -84, 16, -84, 6, 0, -22, 17, -22, 16, -22, 18, 0, -81,
	31, 24, 21, 25, 33, 26, 17, -81, 20, 27, 22, 28, 40, 29,
	16, -81, 20, 6, -27, 0, -27, 31, -27, 21, -27, 33, -27, 17,
	-27, 20, -27, 22, -27, 40, -27, 16, -27, 18, 0, -23, 31, -23,
	21, -23, 33, -23, 17, -23, 20, -23, 22, -23, 40, -23, 16, -23,
	20, 6, -25, 0, -25, 31, -25, 21, -25, 33, -25, 17, -25, 20,
	-25, 22, -25, 40, -25, 16, -25, 2, 3, 61, 2, 5, 62, 14,
	9, -36, 19, -36, 29, -36, 40, -36, 4, 78, 24, -36, 23, -36,
	30, 9, -43, 19, -43, 29, -43, 33, -43, 35, -43, 40, -43, 36,
	-43, 13, -43, 8, -43, 4, -43, 24, -43, 23, -43, 7, -43, 12,
	-43, 11, -43, 30, 9, -42, 19, -42, 29, -42, 33, -42, 35, -42,
	40, -42, 36, -42, 13, -42, 8, -42, 4, -42, 24, -42, 23, -42,
	7, -42, 12, -42, 11, -42, 2, 36, -101, 2, 36, 87, 32, 9,
	-47, 19, -47, 29, -47, 33, -47, 35, -47, 40, -47, 36, -47, 13,
	-47, 8, -47, 4, -47, 24, -47, 23, -47, 10, 86, 7, -47, 12,
	-47, 11, -47, 30, 9, -44, 19, -44, 29, -44, 33, -44, 35, -44,
	40, -44, 36, -44, 13, -44, 8, -44, 4, -44, 24, -44, 23, -44,
	7, -44, 12, -44, 11, -44, 10, 19, 64, 29, 65, 40, -90, 24,
	66, 23, 67, 14, 9, -96, 19, -96, 29, -96, 40, -96, 4, -96,
	24, -96, 23, -96, 30, 9, -100, 19, -100, 29, -100, 33, -100, 35,
	-100, 40, -100, 36, -100, 13, 80, 8, -100, 4, -100, 24, -100, 23,
	-100, 7, -100, 12, 81, 11, 82, 24, 9, -98, 19, -98, 29, -98,
	33, -98, 35, -98, 40, -98, 36, -98, 8, -98, 4, -98, 24, -98,
	23, -98, 7, -98, 24, 9, -37, 19, -37, 29, -37, 33, 43, 35,
	44, 40, -37, 36, -102, 8, 45, 4, -37, 24, -37, 23, -37, 7,
	46, 2, 3, 75, 4, 33, -35, 40, -35, 4, 33, -94, 40, -94,
	4, 33, 56, 40, 76, 12, 0, -69, 33, -69, 17, -69, 40, -69,
	16, -69, 18, -69, 18, 0, -83, 31, -83, 21, -83, 33, -83, 17,
	-83, 20, -83, 22, -83, 40, -83, 16, -83, 10, 30, 89, 25, 90,
	33, 91, 27, 92, 35, 93, 12, 6, -86, 31, 24, 21, 25, 33,
	26, 20, 27, 40, 29, 10, 19, -55, 29, -55, 40, -55, 24, -55,
	23, -55, 2, 8, 112, 10, 19, -57, 29, -57, 40, -57, 24, -57,
	23, -57, 2, 8, 111, 10, 19, -92, 29, -92, 40, -92, 24, -92,
	23, -92, 2, 40, 103, 10, 19, 64, 29, 65, 40, -89, 24, 66,
	23, 67, 10, 19, -51, 29, -51, 40, -51, 24, -51, 23, -51, 10,
	19, -54, 29, -54, 40, -54, 24, -54, 23, -54, 10, 19, -53, 29,
	-53, 40, -53, 24, -53, 23, -53, 10, 19, -52, 29, -52, 40, -52,
	24, -52, 23, -52, 20, 6, -34, 0, -34, 31, -34, 21, -34, 33,
	-34, 17, -34, 20, -34, 22, -34, 40, -34, 16, -34, 4, 33, -93,
	40, -93, 24, 9, -97, 19, -97, 29, -97, 33, -97, 35, -97, 40,
	-97, 36, -97, 8, -97, 4, -97, 24, -97, 23, -97, 7, -97, 24,
	9, -41, 19, -41, 29, -41, 33, -41, 35, -41, 40, -41, 36, -41,
	8, -41, 4, -41, 24, -41, 23, -41, 7, -41, 24, 9, -40, 19,
	-40, 29, -40, 33, -40, 35, -40, 40, -40, 36, -40, 8, -40, 4,
	-40, 24, -40, 23, -40, 7, -40, 24, 9, -39, 19, -39, 29, -39,
	33, -39, 35, -39, 40, -39, 36, -39, 8, -39, 4, -39, 24, -39,
	23, -39, 7, -39, 24, 9, -99, 19, -99, 29, -99, 33, -99, 35,
	-99, 40, -99, 36, -99, 8, -99, 4, -99, 24, -99, 23, -99, 7,
	-99, 24, 9, -38, 19, -38, 29, -38, 33, -38, 35, -38, 40, -38,
	36, -38, 8, -38, 4, -38, 24, -38, 23, -38, 7, -38, 2, 9,
	105, 4, 36, -102, 7, 46, 4, 39, 107, 38, 108, 4, 40, 115,
	4, 116, 4, 40, -10, 4, -10, 26, 2, -14, 9, -14, 25, -14,
	33, -14, 26, -14, 27, -14, 35, -14, 40, -14, 13, -14, 4, -14,
	28, -14, 12, -14, 11, -14, 26, 2, -12, 9, -12, 25, -12, 33,
	-12, 26, -12, 27, -12, 35, -12, 40, -12, 13, -12, 4, -12, 28,
	-12, 12, -12, 11, -12, 2, 8, 127, 26, 2, -13, 9, -13, 25,
	-13, 33, -13, 26, -13, 27, -13, 35, -13, 40, -13, 13, -13, 4,
	-13, 28, -13, 12, -13, 11, -13, 26, 2, -15, 9, -15, 25, -15,
	33, -15, 26, -15, 27, -15, 35, -15, 40, -15, 13, -15, 4, -15,
	28, -15, 12, -15, 11, -15, 4, 40, -74, 4, -74, 22, 25, -80,
	33, -80, 26, -80, 27, -80, 35, -80, 40, -80, 13, 122, 4, -80,
	28, -80, 12, 123, 11, 124, 16, 25, -76, 33, -76, 26, -76, 27,
	-76, 35, -76, 40, -76, 4, -76, 28, -76, 16, 25, 90, 33, 91,
	26, 117, 27, 92, 35, 93, 40, -78, 4, -78, 28, 118, 12, 6,
	-88, 31, -88, 21, -88, 33, -88, 20, -88, 40, -88, 2, 6, 128,
	12, 6, -85, 31, 24, 21, 25, 33, 26, 20, 27, 40, 29, 2,
	40, 129, 20, 6, -32, 0, -32, 31, -32, 21, -32, 33, -32, 17,
	-32, 20, -32, 22, -32, 40, -32, 16, -32, 2, 40, 130, 30, 9,
	-45, 19, -45, 29, -45, 33, -45, 35, -45, 40, -45, 36, -45, 13,
	-45, 8, -45, 4, -45, 24, -45, 23, -45, 7, -45, 12, -45, 11,
	-45, 30, 9, -46, 19, -46, 29, -46, 33, -46, 35, -46, 40, -46,
	36, -46, 13, -46, 8, -46, 4, -46, 24, -46, 23, -46, 7, -46,
	12, -46, 11, -46, 6, 37, -49, 39, -49, 38, -49, 6, 37, -50,
	39, -50, 38, -50, 6, 37, -104, 39, -104, 38, -104, 6, 37, 131,
	39, 107, 38, 108, 4, 9, -106, 33, 133, 2, 33, 135, 10, 19,
	-91, 29, -91, 40, -91, 24, -91, 23, -91, 14, 9, -95, 19, -95,
	29, -95, 40, -95, 4, -95, 24, -95, 23, -95, 12, 0, -8, 33,
	-8, 17, -8, 40, -8, 16, -8, 18, -8, 2, 8, 138, 2, 8,
	139, 4, 40, -77, 4, -77, 4, 40, -9, 4, -9, 16, 25, -75,
	33, -75, 26, -75, 27, -75, 35, -75, 40, -75, 4, -75, 28, -75,
	16, 25, -18, 33, -18, 26, -18, 27, -18, 35, -18, 40, -18, 4,
	-18, 28, -18, 16, 25, -17, 33, -17, 26, -17, 27, -17, 35, -17,
	40, -17, 4, -17, 28, -17, 16, 25, -19, 33, -19, 26, -19, 27,
	-19, 35, -19, 40, -19, 4, -19, 28, -19, 16, 25, -79, 33, -79,
	26, -79, 27, -79, 35, -79, 40, -79, 4, -79, 28, -79, 16, 25,
	-11, 33, -11, 26, -11, 27, -11, 35, -11, 40, -11, 4, -11, 28,
	-11, 8, 25, 90, 33, 91, 27, 92, 35, 93, 18, 0, -30, 31,
	-30, 21, -30, 33, -30, 17, -30, 20, -30, 22, -30, 40, -30, 16,
	-30, 20, 6, -31, 0, -31, 31, -31, 21, -31, 33, -31, 17, -31,
	20, -31, 22, -31, 40, -31, 16, -31, 20, 6, -33, 0, -33, 31,
	-33, 21, -33, 33, -33, 17, -33, 20, -33, 22, -33, 40, -33, 16,
	-33, 32, 9, -48, 19, -48, 29, -48, 33, -48, 35, -48, 40, -48,
	36, -48, 13, -48, 8, -48, 4, -48, 24, -48, 23, -48, 10, -48,
	7, -48, 12, -48, 11, -48, 6, 37, -103, 39, -103, 38, -103, 2,
	9, -105, 2, 9, 140, 2, 9, 141, 12, 6, -87, 31, -87, 21,
	-87, 33, -87, 20, -87, 40, -87, 2, 2, 143, 2, 34, 144, 2,
	34, 145, 10, 19, -56, 29, -56, 40, -56, 24, -56, 23, -56, 10,
	19, -58, 29, -58, 40, -58, 24, -58, 23, -58, 4, 40, -73, 4,
	-73, 2, 9, 147, 2, 9, 148, 2, 9, 149, 4, 40, -20, 4,
	-20, 4, 40, -21, 4, -21, 26, 2, -16, 9, -16, 25, -16, 33,
	-16, 26, -16, 27, -16, 35, -16, 40, -16, 13, -16, 4, -16, 28,
	-16, 12, -16, 11, -16,
}

var _goto = []int32{
	150, 157, 157, 158, 157, 157, 157, 157, 157, 157, 157, 157, 169, 157,
	176, 187, 157, 157, 157, 157, 157, 157, 157, 206, 213, 218, 157, 157,
	157, 157, 157, 157, 157, 157, 157, 237, 157, 157, 157, 157, 157, 252,
	157, 157, 157, 271, 157, 157, 157, 157, 290, 157, 305, 157, 310, 157,
	157, 157, 321, 157, 157, 324, 337, 352, 157, 157, 157, 157, 157, 157,
	367, 157, 157, 157, 157, 378, 157, 157, 397, 157, 157, 157, 157, 157,
	157, 157, 412, 417, 157, 157, 157, 157, 157, 157, 157, 157, 422, 157,
	427, 157, 157, 438, 157, 157, 157, 157, 157, 157, 157, 157, 449, 452,
	157, 157, 157, 157, 455, 157, 157, 157, 157, 157, 157, 157, 157, 157,
	157, 466, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157,
	157, 157, 157, 471, 157, 157, 157, 157, 157, 157, 6, 34, 3, 35,
	4, 1, 5, 0, 10, 12, 8, 3, 9, 2, 10, 36, 11, 37,
	12, 6, 12, 8, 3, 9, 2, 16, 10, 40, 19, 5, 20, 4,
	21, 38, 22, 39, 23, 18, 19, 30, 17, 31, 14, 32, 13, 33,
	45, 34, 46, 35, 18, 36, 15, 37, 16, 38, 6, 40, 19, 5,
	20, 4, 59, 4, 20, 57, 51, 58, 18, 52, 42, 55, 47, 27,
	48, 26, 49, 21, 50, 22, 51, 25, 52, 23, 53, 53, 54, 14,
	19, 30, 17, 31, 14, 32, 13, 60, 18, 36, 15, 37, 16, 38,
	18, 52, 42, 55, 47, 27, 48, 26, 49, 21, 63, 22, 51, 25,
	52, 23, 53, 53, 54, 18, 52, 42, 55, 47, 27, 48, 26, 49,
	21, 85, 22, 51, 25, 52, 23, 53, 53, 54, 14, 29, 68, 49,
	69, 50, 70, 30, 71, 33, 72, 32, 73, 31, 74, 4, 24, 83,
	54, 84, 10, 55, 47, 27, 48, 26, 49, 25, 52, 23, 79, 2,
	20, 77, 12, 41, 88, 9, 94, 6, 95, 8, 96, 7, 97, 42,
	98, 14, 19, 30, 17, 31, 14, 99, 47, 100, 48, 101, 18, 36,
	16, 38, 14, 29, 68, 49, 102, 50, 70, 30, 71, 33, 72, 32,
	73, 31, 74, 10, 29, 113, 30, 71, 33, 72, 32, 73, 31, 74,
	18, 52, 42, 55, 47, 27, 48, 26, 49, 21, 104, 22, 51, 25,
	52, 23, 53, 53, 54, 14, 55, 47, 27, 48, 26, 49, 22, 114,
	25, 52, 23, 53, 53, 54, 4, 55, 47, 27, 106, 4, 28, 109,
	56, 110, 4, 10, 125, 44, 126, 10, 9, 94, 11, 119, 43, 120,
	8, 96, 7, 121, 10, 19, 30, 17, 31, 14, 136, 18, 36, 16,
	38, 2, 28, 132, 2, 57, 134, 10, 9, 94, 6, 142, 8, 96,
	7, 97, 42, 98, 4, 9, 94, 8, 137, 4, 9, 94, 8, 146,
}

type _Bounds struct {
	Begin Token
	End   Token
	Empty bool
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

type Error struct {
	Token    Token
	Expected []int
}

func _Find(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type _Lexer interface {
	ReadToken() (Token, int)
}

type _item struct {
	State  int32
	Sym    any
	Bounds _Bounds
}

type lox struct {
	_lex   _Lexer
	_stack _Stack[_item]

	_la    int
	_lasym any

	_qla    int
	_qlasym any
}

func (p *parser) parse(lex _Lexer) bool {
	const accept = 2147483647

	p._lex = lex
	p._qla = -1
	p._stack.Push(_item{})

	p._readToken()

	for {
		topState := p._stack.Peek(0).State
		action, ok := _Find(_actions, topState, int32(p._la))
		if !ok {
			if !p._recover() {
				return false
			}
			continue
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			latok, ok := p._lasym.(Token)
			if !ok {
				latok = p._lasym.(Error).Token
			}
			p._stack.Push(_item{
				State: action,
				Sym:   p._lasym,
				Bounds: _Bounds{
					Begin: latok,
					End:   latok,
				},
			})
			p._readToken()
		} else { // reduce
			prod := -action
			termCount := _termCounts[int(prod)]
			rule := _rules[int(prod)]
			res := p._act(prod)

			// Compute reduction token bounds.
			// Trim leading and trailing empty bounds.
			boundSlice := p._stack.PeekSlice(int(termCount))
			for len(boundSlice) > 0 && boundSlice[0].Bounds.Empty {
				boundSlice = boundSlice[1:]
			}
			for len(boundSlice) > 0 && boundSlice[len(boundSlice)-1].Bounds.Empty {
				boundSlice = boundSlice[:len(boundSlice)-1]
			}
			var bounds _Bounds
			if len(boundSlice) > 0 {
				bounds.Begin = boundSlice[0].Bounds.Begin
				bounds.End = boundSlice[len(boundSlice)-1].Bounds.End
			} else {
				bounds.Empty = true
			}
			if !bounds.Empty {
				p._onBounds(res, bounds.Begin, bounds.End)
			}
			p._stack.Pop(int(termCount))
			topState = p._stack.Peek(0).State
			nextState, _ := _Find(_goto, topState, rule)
			p._stack.Push(_item{
				State:  nextState,
				Sym:    res,
				Bounds: bounds,
			})
		}
	}

	return true
}

// recoverLookahead can be called during an error production action (an action
// for a production that has a @error term) to recover the lookahead that was
// possibly lost in the process of reducing the error production.
func (p *parser) recoverLookahead(typ int, tok Token) {
	if p._qla != -1 {
		panic("recovered lookahead already pending")
	}

	p._qla = p._la
	p._qlasym = p._lasym
	p._la = typ
	p._lasym = tok
}

func (p *parser) _readToken() {
	if p._qla != -1 {
		p._la = p._qla
		p._lasym = p._qlasym
		p._qla = -1
		p._qlasym = nil
		return
	}

	p._lasym, p._la = p._lex.ReadToken()
	if p._la == ERROR {
		p._lasym = p._makeError()
	}
}

func (p *parser) _recover() bool {
	errSym, ok := p._lasym.(Error)
	if !ok {
		errSym = p._makeError()
	}

	for p._la == ERROR {
		p._readToken()
	}

	for {
		save := p._stack

		for len(p._stack) >= 1 {
			state := p._stack.Peek(0).State

			for {
				action, ok := _Find(_actions, state, int32(ERROR))
				if !ok {
					break
				}

				if action < 0 {
					prod := -action
					rule := _rules[int(prod)]
					state, _ = _Find(_goto, state, rule)
					continue
				}

				state = action

				_, ok = _Find(_actions, state, int32(p._la))
				if !ok {
					break
				}

				p._qla = p._la
				p._qlasym = p._lasym
				p._la = ERROR
				p._lasym = errSym
				return true
			}

			p._stack.Pop(1)
		}

		if p._la == EOF {
			return false
		}

		p._stack = save
		p._readToken()
	}
}

func (p *parser) _makeError() Error {
	e := Error{
		Token: p._lasym.(Token),
	}

	// Compile list of allowed tokens at this state.
	// See _Find for the format of the _actions table.
	s := p._stack.Peek(0).State
	i := int(_actions[int(s)])
	count := int(_actions[i])
	i++
	end := i + count
	for ; i < end; i += 2 {
		e.Expected = append(e.Expected, int(_actions[i]))
	}

	return e
}

func (p *parser) _act(prod int32) any {
	switch prod {
	case 1:
		return p.on_spec(
			_cast[[]_i0.Token](p._stack.Peek(1).Sym),
			_cast[[][]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_spec__error(
			_cast[Error](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_section(
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_section(
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_parser_section(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_parser_statement(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_parser_statement__nl(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_parser_rule(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]*_i1.ParserProd](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_parser_prod(
			_cast[[]*_i1.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i1.ProdQualifier](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_parser_prod__empty(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_parser_term_card(
			_cast[*_i1.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i1.ParserTermType](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_parser_term__token(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_parser_term__token(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_parser_term__token(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_parser_term__list(
			_cast[*_i1.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_parser_list(
			_cast[_i0.Token](p._stack.Peek(5).Sym),
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[*_i1.ParserTerm](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*_i1.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_parser_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_parser_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_parser_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_parser_qualif(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_parser_qualif(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_lexer_section(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_lexer_statement(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_lexer_rule(
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_lexer_rule__nl(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_mode(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]_i1.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_token_rule(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[*_i1.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i1.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_frag_rule(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[*_i1.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i1.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_macro_rule(
			_cast[_i0.Token](p._stack.Peek(4).Sym),
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*_i1.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 34:
		return p.on_external_rule(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]*_i1.ExternalName](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_external_name(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_lexer_expr(
			_cast[[]*_i1.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_lexer_factor(
			_cast[[]*_i1.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_lexer_term_card(
			_cast[_i1.LexerTerm](p._stack.Peek(1).Sym),
			_cast[_i1.Card](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_lexer_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_lexer_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_lexer_card(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 42:
		return p.on_lexer_term__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 43:
		return p.on_lexer_term__tok(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 44:
		return p.on_lexer_term__char_class_expr(
			_cast[_i1.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 45:
		return p.on_lexer_term__expr(
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[*_i1.LexerExpr](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 46:
		return p.on_char_class_expr__binary(
			_cast[_i1.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i1.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 47:
		return p.on_char_class_expr__char_class(
			_cast[*_i1.CharClass](p._stack.Peek(0).Sym),
		)
	case 48:
		return p.on_char_class(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[[]_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 49:
		return p.on_char_class_item(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_char_class_item(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_action(
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_action(
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		)
	case 53:
		return p.on_action(
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		)
	case 54:
		return p.on_action(
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		)
	case 55:
		return p.on_action_discard(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 56:
		return p.on_action_push_mode(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 57:
		return p.on_action_pop_mode(
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 58:
		return p.on_action_emit(
			_cast[_i0.Token](p._stack.Peek(3).Sym),
			_cast[_i0.Token](p._stack.Peek(2).Sym),
			_cast[_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 59: // ZeroOrMore
		return _cast[[]_i0.Token](p._stack.Peek(0).Sym)
	case 60: // ZeroOrMore
		{
			var zero []_i0.Token
			return zero
		}
	case 61: // OneOrMore
		return append(
			_cast[[]_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 62: // OneOrMore
		return []_i0.Token{
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		}
	case 63: // ZeroOrMore
		return _cast[[][]_i1.Statement](p._stack.Peek(0).Sym)
	case 64: // ZeroOrMore
		{
			var zero [][]_i1.Statement
			return zero
		}
	case 65: // OneOrMore
		return append(
			_cast[[][]_i1.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 66: // OneOrMore
		return [][]_i1.Statement{
			_cast[[]_i1.Statement](p._stack.Peek(0).Sym),
		}
	case 67: // ZeroOrMore
		return _cast[[]_i1.Statement](p._stack.Peek(0).Sym)
	case 68: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 69: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._stack.Peek(1).Sym),
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 70: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		}
	case 71: // ZeroOrOne
		return _cast[_i0.Token](p._stack.Peek(0).Sym)
	case 72: // ZeroOrOne
		{
			var zero _i0.Token
			return zero
		}
	case 73: // List
		return append(
			_cast[[]*_i1.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i1.ParserProd](p._stack.Peek(0).Sym),
		)
	case 74: // List
		return []*_i1.ParserProd{
			_cast[*_i1.ParserProd](p._stack.Peek(0).Sym),
		}
	case 75: // OneOrMore
		return append(
			_cast[[]*_i1.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i1.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 76: // OneOrMore
		return []*_i1.ParserTerm{
			_cast[*_i1.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 77: // ZeroOrOne
		return _cast[*_i1.ProdQualifier](p._stack.Peek(0).Sym)
	case 78: // ZeroOrOne
		{
			var zero *_i1.ProdQualifier
			return zero
		}
	case 79: // ZeroOrOne
		return _cast[_i1.ParserTermType](p._stack.Peek(0).Sym)
	case 80: // ZeroOrOne
		{
			var zero _i1.ParserTermType
			return zero
		}
	case 81: // ZeroOrMore
		return _cast[[]_i1.Statement](p._stack.Peek(0).Sym)
	case 82: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 83: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._stack.Peek(1).Sym),
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 84: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		}
	case 85: // ZeroOrMore
		return _cast[[]_i1.Statement](p._stack.Peek(0).Sym)
	case 86: // ZeroOrMore
		{
			var zero []_i1.Statement
			return zero
		}
	case 87: // OneOrMore
		return append(
			_cast[[]_i1.Statement](p._stack.Peek(1).Sym),
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		)
	case 88: // OneOrMore
		return []_i1.Statement{
			_cast[_i1.Statement](p._stack.Peek(0).Sym),
		}
	case 89: // ZeroOrMore
		return _cast[[]_i1.Action](p._stack.Peek(0).Sym)
	case 90: // ZeroOrMore
		{
			var zero []_i1.Action
			return zero
		}
	case 91: // OneOrMore
		return append(
			_cast[[]_i1.Action](p._stack.Peek(1).Sym),
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		)
	case 92: // OneOrMore
		return []_i1.Action{
			_cast[_i1.Action](p._stack.Peek(0).Sym),
		}
	case 93: // OneOrMore
		return append(
			_cast[[]*_i1.ExternalName](p._stack.Peek(1).Sym),
			_cast[*_i1.ExternalName](p._stack.Peek(0).Sym),
		)
	case 94: // OneOrMore
		return []*_i1.ExternalName{
			_cast[*_i1.ExternalName](p._stack.Peek(0).Sym),
		}
	case 95: // List
		return append(
			_cast[[]*_i1.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i1.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 96: // List
		return []*_i1.LexerFactor{
			_cast[*_i1.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 97: // OneOrMore
		return append(
			_cast[[]*_i1.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i1.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 98: // OneOrMore
		return []*_i1.LexerTermCard{
			_cast[*_i1.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 99: // ZeroOrOne
		return _cast[_i1.Card](p._stack.Peek(0).Sym)
	case 100: // ZeroOrOne
		{
			var zero _i1.Card
			return zero
		}
	case 101: // ZeroOrOne
		return _cast[_i0.Token](p._stack.Peek(0).Sym)
	case 102: // ZeroOrOne
		{
			var zero _i0.Token
			return zero
		}
	case 103: // OneOrMore
		return append(
			_cast[[]_i0.Token](p._stack.Peek(1).Sym),
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		)
	case 104: // OneOrMore
		return []_i0.Token{
			_cast[_i0.Token](p._stack.Peek(0).Sym),
		}
	case 105: // ZeroOrOne
		return _cast[_i0.Token](p._stack.Peek(0).Sym)
	case 106: // ZeroOrOne
		{
			var zero _i0.Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
