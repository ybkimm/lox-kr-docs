package parser

import (
	_i0 "github.com/dcaiafa/lox/internal/ast"
)

var _rules = []int32{
	0, 1, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8,
	8, 8, 9, 10, 10, 10, 10, 11, 11, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 24,
	24, 25, 25, 25, 25, 26, 26, 27, 28, 28, 29, 29, 29, 29,
	30, 31, 32, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38,
	39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45,
	46, 46, 47, 47, 48, 48, 49, 49, 50, 50, 51, 51, 52, 52,
	53, 53, 54, 54, 55, 55, 56, 56, 57, 57, 58, 58, 59, 59,
}

var _termCounts = []int32{
	1, 2, 1, 1, 1, 3, 1, 1, 5, 2, 1, 2, 1, 1,
	1, 1, 6, 1, 1, 1, 1, 4, 4, 3, 1, 1, 1, 1,
	1, 1, 1, 6, 5, 4, 5, 3, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 3, 3, 1, 4, 1, 1, 1, 1, 1, 1,
	1, 4, 1, 4, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0,
	2, 1, 1, 0, 3, 1, 2, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 0, 2, 1, 1, 0, 2, 1, 1, 0, 2, 1,
	2, 1, 3, 1, 2, 1, 1, 0, 1, 0, 2, 1, 1, 0,
}

var _actions = []int32{
	155, 166, 169, 178, 185, 194, 197, 200, 203, 210, 217, 224, 227, 234,
	243, 256, 275, 282, 295, 298, 301, 314, 327, 334, 347, 350, 361, 364,
	367, 370, 391, 412, 433, 452, 471, 478, 497, 518, 537, 558, 561, 350,
	566, 581, 612, 350, 643, 646, 649, 682, 713, 724, 739, 770, 795, 820,
	823, 828, 833, 838, 851, 870, 881, 886, 889, 713, 894, 905, 908, 919,
	922, 933, 936, 947, 958, 969, 980, 350, 991, 1012, 350, 1017, 1042, 1067,
	1092, 1117, 1142, 1167, 1170, 1175, 1180, 1185, 1190, 1219, 1248, 1251, 1280, 1309,
	1314, 1339, 1356, 1373, 1386, 1389, 1410, 1413, 1444, 1475, 1482, 1489, 1496, 1503,
	1508, 1511, 1516, 1527, 1542, 870, 1555, 1558, 1561, 1566, 1571, 1588, 1605, 1622,
	1639, 1656, 1673, 1690, 1699, 1712, 1715, 1728, 1749, 1770, 1803, 1810, 1813, 1816,
	1819, 1822, 1825, 1828, 1847, 1858, 1869, 1874, 1690, 1887, 1890, 1893, 1896, 1901,
	1906, 10, 0, -61, 1, 1, 16, -61, 39, 2, 15, -61, 2, 0,
	-2, 8, 0, -63, 16, -63, 39, -63, 15, -63, 6, 0, -65, 16,
	6, 15, 7, 8, 0, -60, 16, -60, 39, 13, 15, -60, 2, 0,
	2147483647, 2, 39, 15, 2, 39, 14, 6, 0, -4, 16, -4, 15, -4,
	6, 0, -3, 16, -3, 15, -3, 6, 0, -67, 16, -67, 15, -67,
	2, 0, -1, 6, 0, -64, 16, 6, 15, 7, 8, 0, -62, 16,
	-62, 39, -62, 15, -62, 12, 0, -69, 32, -73, 16, -69, 39, 17,
	15, -69, 17, 18, 18, 0, -83, 30, 24, 20, 25, 32, 26, 16,
	-83, 19, 27, 21, 28, 39, 29, 15, -83, 6, 0, -66, 16, -66,
	15, -66, 12, 0, -7, 32, -7, 16, -7, 39, -7, 15, -7, 17,
	-7, 2, 32, -72, 2, 32, 39, 12, 0, -6, 32, -6, 16, -6,
	39, -6, 15, -6, 17, -6, 12, 0, -71, 32, -71, 16, -71, 39,
	-71, 15, -71, 17, -71, 6, 0, -5, 16, -5, 15, -5, 12, 0,
	-68, 32, -73, 16, -68, 39, 17, 15, -68, 17, 18, 2, 32, 56,
	10, 32, 43, 34, 44, 35, -107, 8, 45, 7, 46, 2, 3, 41,
	2, 32, 55, 2, 32, 40, 20, 6, -30, 0, -30, 30, -30, 20,
	-30, 32, -30, 16, -30, 19, -30, 21, -30, 39, -30, 15, -30, 20,
	6, -29, 0, -29, 30, -29, 20, -29, 32, -29, 16, -29, 19, -29,
	21, -29, 39, -29, 15, -29, 20, 6, -27, 0, -27, 30, -27, 20,
	-27, 32, -27, 16, -27, 19, -27, 21, -27, 39, -27, 15, -27, 18,
	0, -25, 30, -25, 20, -25, 32, -25, 16, -25, 19, -25, 21, -25,
	39, -25, 15, -25, 18, 0, -85, 30, -85, 20, -85, 32, -85, 16,
	-85, 19, -85, 21, -85, 39, -85, 15, -85, 6, 0, -23, 16, -23,
	15, -23, 18, 0, -82, 30, 24, 20, 25, 32, 26, 16, -82, 19,
	27, 21, 28, 39, 29, 15, -82, 20, 6, -28, 0, -28, 30, -28,
	20, -28, 32, -28, 16, -28, 19, -28, 21, -28, 39, -28, 15, -28,
	18, 0, -24, 30, -24, 20, -24, 32, -24, 16, -24, 19, -24, 21,
	-24, 39, -24, 15, -24, 20, 6, -26, 0, -26, 30, -26, 20, -26,
	32, -26, 16, -26, 19, -26, 21, -26, 39, -26, 15, -26, 2, 3,
	61, 4, 39, 62, 5, -87, 14, 9, -37, 18, -37, 28, -37, 39,
	-37, 4, 80, 23, -37, 22, -37, 30, 9, -44, 18, -44, 28, -44,
	32, -44, 34, -44, 39, -44, 35, -44, 13, -44, 8, -44, 4, -44,
	23, -44, 22, -44, 7, -44, 12, -44, 11, -44, 30, 9, -43, 18,
	-43, 28, -43, 32, -43, 34, -43, 39, -43, 35, -43, 13, -43, 8,
	-43, 4, -43, 23, -43, 22, -43, 7, -43, 12, -43, 11, -43, 2,
	35, -106, 2, 35, 89, 32, 9, -48, 18, -48, 28, -48, 32, -48,
	34, -48, 39, -48, 35, -48, 13, -48, 8, -48, 4, -48, 23, -48,
	22, -48, 10, 88, 7, -48, 12, -48, 11, -48, 30, 9, -45, 18,
	-45, 28, -45, 32, -45, 34, -45, 39, -45, 35, -45, 13, -45, 8,
	-45, 4, -45, 23, -45, 22, -45, 7, -45, 12, -45, 11, -45, 10,
	18, 66, 28, 67, 39, -95, 23, 68, 22, 69, 14, 9, -101, 18,
	-101, 28, -101, 39, -101, 4, -101, 23, -101, 22, -101, 30, 9, -105,
	18, -105, 28, -105, 32, -105, 34, -105, 39, -105, 35, -105, 13, 82,
	8, -105, 4, -105, 23, -105, 22, -105, 7, -105, 12, 83, 11, 84,
	24, 9, -103, 18, -103, 28, -103, 32, -103, 34, -103, 39, -103, 35,
	-103, 8, -103, 4, -103, 23, -103, 22, -103, 7, -103, 24, 9, -38,
	18, -38, 28, -38, 32, 43, 34, 44, 39, -38, 35, -107, 8, 45,
	4, -38, 23, -38, 22, -38, 7, 46, 2, 3, 77, 4, 32, -36,
	39, -36, 4, 32, -99, 39, -99, 4, 32, 56, 39, 78, 12, 0,
	-70, 32, -70, 16, -70, 39, -70, 15, -70, 17, -70, 18, 0, -84,
	30, -84, 20, -84, 32, -84, 16, -84, 19, -84, 21, -84, 39, -84,
	15, -84, 10, 29, 91, 24, 92, 32, 93, 26, 94, 34, 95, 4,
	39, -89, 5, -89, 2, 5, 101, 4, 39, 113, 5, -86, 10, 18,
	-56, 28, -56, 39, -56, 23, -56, 22, -56, 2, 8, 112, 10, 18,
	-58, 28, -58, 39, -58, 23, -58, 22, -58, 2, 8, 111, 10, 18,
	-97, 28, -97, 39, -97, 23, -97, 22, -97, 2, 39, 103, 10, 18,
	66, 28, 67, 39, -94, 23, 68, 22, 69, 10, 18, -52, 28, -52,
	39, -52, 23, -52, 22, -52, 10, 18, -55, 28, -55, 39, -55, 23,
	-55, 22, -55, 10, 18, -54, 28, -54, 39, -54, 23, -54, 22, -54,
	10, 18, -53, 28, -53, 39, -53, 23, -53, 22, -53, 20, 6, -35,
	0, -35, 30, -35, 20, -35, 32, -35, 16, -35, 19, -35, 21, -35,
	39, -35, 15, -35, 4, 32, -98, 39, -98, 24, 9, -102, 18, -102,
	28, -102, 32, -102, 34, -102, 39, -102, 35, -102, 8, -102, 4, -102,
	23, -102, 22, -102, 7, -102, 24, 9, -42, 18, -42, 28, -42, 32,
	-42, 34, -42, 39, -42, 35, -42, 8, -42, 4, -42, 23, -42, 22,
	-42, 7, -42, 24, 9, -41, 18, -41, 28, -41, 32, -41, 34, -41,
	39, -41, 35, -41, 8, -41, 4, -41, 23, -41, 22, -41, 7, -41,
	24, 9, -40, 18, -40, 28, -40, 32, -40, 34, -40, 39, -40, 35,
	-40, 8, -40, 4, -40, 23, -40, 22, -40, 7, -40, 24, 9, -104,
	18, -104, 28, -104, 32, -104, 34, -104, 39, -104, 35, -104, 8, -104,
	4, -104, 23, -104, 22, -104, 7, -104, 24, 9, -39, 18, -39, 28,
	-39, 32, -39, 34, -39, 39, -39, 35, -39, 8, -39, 4, -39, 23,
	-39, 22, -39, 7, -39, 2, 9, 105, 4, 35, -107, 7, 46, 4,
	38, 107, 37, 108, 4, 39, 116, 4, 117, 4, 39, -10, 4, -10,
	28, 2, -14, 9, -14, 24, -14, 32, -14, 25, -14, 26, -14, 34,
	-14, 39, -14, 13, -14, 4, -14, 27, -14, 12, -14, 14, -14, 11,
	-14, 28, 2, -12, 9, -12, 24, -12, 32, -12, 25, -12, 26, -12,
	34, -12, 39, -12, 13, -12, 4, -12, 27, -12, 12, -12, 14, -12,
	11, -12, 2, 8, 129, 28, 2, -13, 9, -13, 24, -13, 32, -13,
	25, -13, 26, -13, 34, -13, 39, -13, 13, -13, 4, -13, 27, -13,
	12, -13, 14, -13, 11, -13, 28, 2, -15, 9, -15, 24, -15, 32,
	-15, 25, -15, 26, -15, 34, -15, 39, -15, 13, -15, 4, -15, 27,
	-15, 12, -15, 14, -15, 11, -15, 4, 39, -75, 4, -75, 24, 24,
	-81, 32, -81, 25, -81, 26, -81, 34, -81, 39, -81, 13, 123, 4,
	-81, 27, -81, 12, 124, 14, 125, 11, 126, 16, 24, -77, 32, -77,
	25, -77, 26, -77, 34, -77, 39, -77, 4, -77, 27, -77, 16, 24,
	92, 32, 93, 25, 118, 26, 94, 34, 95, 39, -79, 4, -79, 27,
	119, 12, 6, -91, 30, 24, 20, 25, 32, 26, 19, 27, 39, 29,
	2, 39, 133, 20, 6, -33, 0, -33, 30, -33, 20, -33, 32, -33,
	16, -33, 19, -33, 21, -33, 39, -33, 15, -33, 2, 39, 134, 30,
	9, -46, 18, -46, 28, -46, 32, -46, 34, -46, 39, -46, 35, -46,
	13, -46, 8, -46, 4, -46, 23, -46, 22, -46, 7, -46, 12, -46,
	11, -46, 30, 9, -47, 18, -47, 28, -47, 32, -47, 34, -47, 39,
	-47, 35, -47, 13, -47, 8, -47, 4, -47, 23, -47, 22, -47, 7,
	-47, 12, -47, 11, -47, 6, 36, -50, 38, -50, 37, -50, 6, 36,
	-51, 38, -51, 37, -51, 6, 36, -109, 38, -109, 37, -109, 6, 36,
	135, 38, 107, 37, 108, 4, 9, -111, 32, 137, 2, 32, 139, 4,
	39, -88, 5, -88, 10, 18, -96, 28, -96, 39, -96, 23, -96, 22,
	-96, 14, 9, -100, 18, -100, 28, -100, 39, -100, 4, -100, 23, -100,
	22, -100, 12, 0, -8, 32, -8, 16, -8, 39, -8, 15, -8, 17,
	-8, 2, 8, 141, 2, 8, 142, 4, 39, -78, 4, -78, 4, 39,
	-9, 4, -9, 16, 24, -76, 32, -76, 25, -76, 26, -76, 34, -76,
	39, -76, 4, -76, 27, -76, 16, 24, -19, 32, -19, 25, -19, 26,
	-19, 34, -19, 39, -19, 4, -19, 27, -19, 16, 24, -17, 32, -17,
	25, -17, 26, -17, 34, -17, 39, -17, 4, -17, 27, -17, 16, 24,
	-18, 32, -18, 25, -18, 26, -18, 34, -18, 39, -18, 4, -18, 27,
	-18, 16, 24, -20, 32, -20, 25, -20, 26, -20, 34, -20, 39, -20,
	4, -20, 27, -20, 16, 24, -80, 32, -80, 25, -80, 26, -80, 34,
	-80, 39, -80, 4, -80, 27, -80, 16, 24, -11, 32, -11, 25, -11,
	26, -11, 34, -11, 39, -11, 4, -11, 27, -11, 8, 24, 92, 32,
	93, 26, 94, 34, 95, 12, 6, -93, 30, -93, 20, -93, 32, -93,
	19, -93, 39, -93, 2, 6, 143, 12, 6, -90, 30, 24, 20, 25,
	32, 26, 19, 27, 39, 29, 20, 6, -32, 0, -32, 30, -32, 20,
	-32, 32, -32, 16, -32, 19, -32, 21, -32, 39, -32, 15, -32, 20,
	6, -34, 0, -34, 30, -34, 20, -34, 32, -34, 16, -34, 19, -34,
	21, -34, 39, -34, 15, -34, 32, 9, -49, 18, -49, 28, -49, 32,
	-49, 34, -49, 39, -49, 35, -49, 13, -49, 8, -49, 4, -49, 23,
	-49, 22, -49, 10, -49, 7, -49, 12, -49, 11, -49, 6, 36, -108,
	38, -108, 37, -108, 2, 9, -110, 2, 9, 144, 2, 9, 145, 2,
	2, 148, 2, 33, 149, 2, 33, 150, 18, 0, -31, 30, -31, 20,
	-31, 32, -31, 16, -31, 19, -31, 21, -31, 39, -31, 15, -31, 10,
	18, -57, 28, -57, 39, -57, 23, -57, 22, -57, 10, 18, -59, 28,
	-59, 39, -59, 23, -59, 22, -59, 4, 39, -74, 4, -74, 12, 6,
	-92, 30, -92, 20, -92, 32, -92, 19, -92, 39, -92, 2, 9, 152,
	2, 9, 153, 2, 9, 154, 4, 39, -21, 4, -21, 4, 39, -22,
	4, -22, 28, 2, -16, 9, -16, 24, -16, 32, -16, 25, -16, 26,
	-16, 34, -16, 39, -16, 13, -16, 4, -16, 27, -16, 12, -16, 14,
	-16, 11, -16,
}

var _goto = []int32{
	155, 162, 162, 163, 162, 162, 162, 162, 162, 162, 162, 162, 174, 162,
	181, 192, 162, 162, 162, 162, 162, 162, 162, 211, 218, 223, 162, 162,
	162, 162, 162, 162, 162, 162, 162, 242, 162, 162, 162, 162, 257, 262,
	162, 162, 162, 281, 162, 162, 162, 162, 300, 162, 315, 162, 320, 162,
	162, 162, 331, 162, 162, 334, 162, 162, 162, 347, 162, 162, 162, 162,
	162, 162, 362, 162, 162, 162, 162, 373, 162, 162, 392, 162, 162, 162,
	162, 162, 162, 162, 407, 412, 162, 162, 162, 162, 162, 162, 162, 162,
	417, 162, 422, 433, 162, 162, 162, 162, 162, 162, 162, 162, 448, 451,
	162, 162, 162, 162, 162, 454, 162, 162, 162, 162, 162, 162, 162, 162,
	162, 162, 162, 465, 162, 162, 470, 162, 162, 162, 162, 162, 162, 162,
	162, 162, 162, 162, 162, 162, 162, 162, 481, 162, 162, 162, 162, 162,
	162, 6, 34, 3, 35, 4, 1, 5, 0, 10, 12, 8, 3, 9,
	2, 10, 36, 11, 37, 12, 6, 12, 8, 3, 9, 2, 16, 10,
	40, 19, 5, 20, 4, 21, 38, 22, 39, 23, 18, 19, 30, 17,
	31, 14, 32, 13, 33, 45, 34, 46, 35, 18, 36, 15, 37, 16,
	38, 6, 40, 19, 5, 20, 4, 59, 4, 20, 57, 53, 58, 18,
	54, 42, 57, 47, 27, 48, 26, 49, 21, 50, 22, 51, 25, 52,
	23, 53, 55, 54, 14, 19, 30, 17, 31, 14, 32, 13, 60, 18,
	36, 15, 37, 16, 38, 4, 47, 63, 48, 64, 18, 54, 42, 57,
	47, 27, 48, 26, 49, 21, 65, 22, 51, 25, 52, 23, 53, 55,
	54, 18, 54, 42, 57, 47, 27, 48, 26, 49, 21, 87, 22, 51,
	25, 52, 23, 53, 55, 54, 14, 29, 70, 51, 71, 52, 72, 30,
	73, 33, 74, 32, 75, 31, 76, 4, 24, 85, 56, 86, 10, 57,
	47, 27, 48, 26, 49, 25, 52, 23, 81, 2, 20, 79, 12, 41,
	90, 9, 96, 6, 97, 8, 98, 7, 99, 42, 100, 14, 29, 70,
	51, 102, 52, 72, 30, 73, 33, 74, 32, 75, 31, 76, 10, 29,
	114, 30, 73, 33, 74, 32, 75, 31, 76, 18, 54, 42, 57, 47,
	27, 48, 26, 49, 21, 104, 22, 51, 25, 52, 23, 53, 55, 54,
	14, 57, 47, 27, 48, 26, 49, 22, 115, 25, 52, 23, 53, 55,
	54, 4, 57, 47, 27, 106, 4, 28, 109, 58, 110, 4, 10, 127,
	44, 128, 10, 9, 96, 11, 120, 43, 121, 8, 98, 7, 122, 14,
	19, 30, 17, 31, 14, 130, 49, 131, 50, 132, 18, 36, 16, 38,
	2, 28, 136, 2, 59, 138, 10, 9, 96, 6, 146, 8, 98, 7,
	99, 42, 100, 4, 9, 96, 8, 140, 10, 19, 30, 17, 31, 14,
	147, 18, 36, 16, 38, 4, 9, 96, 8, 151,
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
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[[][]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 2:
		return p.on_spec__error(
			_cast[Error](p._stack.Peek(0).Sym),
		)
	case 3:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 4:
		return p.on_section(
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 5:
		return p.on_parser_section(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 6:
		return p.on_parser_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 7:
		return p.on_parser_statement__nl(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 8:
		return p.on_parser_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]*_i0.ParserProd](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 9:
		return p.on_parser_prod(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym),
		)
	case 10:
		return p.on_parser_prod__empty(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 11:
		return p.on_parser_term_card(
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[_i0.ParserTermType](p._stack.Peek(0).Sym),
		)
	case 12:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 13:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 14:
		return p.on_parser_term__token(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 15:
		return p.on_parser_term__list(
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 16:
		return p.on_parser_list(
			_cast[Token](p._stack.Peek(5).Sym),
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 17:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 18:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 19:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 20:
		return p.on_parser_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 21:
		return p.on_parser_qualif(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 22:
		return p.on_parser_qualif(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 23:
		return p.on_lexer_section(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 24:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 25:
		return p.on_lexer_statement(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 26:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 27:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 28:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 29:
		return p.on_lexer_rule(
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 30:
		return p.on_lexer_rule__nl(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 31:
		return p.on_mode(
			_cast[Token](p._stack.Peek(5).Sym),
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[[]Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 32:
		return p.on_token_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 33:
		return p.on_frag_rule(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(2).Sym),
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 34:
		return p.on_macro_rule(
			_cast[Token](p._stack.Peek(4).Sym),
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 35:
		return p.on_external_rule(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]*_i0.ExternalName](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 36:
		return p.on_external_name(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 37:
		return p.on_lexer_expr(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 38:
		return p.on_lexer_factor(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 39:
		return p.on_lexer_term_card(
			_cast[_i0.LexerTerm](p._stack.Peek(1).Sym),
			_cast[_i0.Card](p._stack.Peek(0).Sym),
		)
	case 40:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 41:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 42:
		return p.on_lexer_card(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 43:
		return p.on_lexer_term__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 44:
		return p.on_lexer_term__tok(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 45:
		return p.on_lexer_term__char_class_expr(
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 46:
		return p.on_lexer_term__expr(
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerExpr](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 47:
		return p.on_char_class_expr__binary(
			_cast[_i0.CharClassExpr](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[_i0.CharClassExpr](p._stack.Peek(0).Sym),
		)
	case 48:
		return p.on_char_class_expr__char_class(
			_cast[*_i0.CharClass](p._stack.Peek(0).Sym),
		)
	case 49:
		return p.on_char_class(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 50:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 51:
		return p.on_char_class_item(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 52:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 53:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 54:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 55:
		return p.on_action(
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 56:
		return p.on_action_discard(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 57:
		return p.on_action_push_mode(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 58:
		return p.on_action_pop_mode(
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 59:
		return p.on_action_emit(
			_cast[Token](p._stack.Peek(3).Sym),
			_cast[Token](p._stack.Peek(2).Sym),
			_cast[Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 60: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 61: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 62:
		{ // OneOrMoreF
			l := _cast[[]Token](p._stack.Peek(1).Sym)
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 63:
		{ // OneOrMoreF
			var l []Token
			e := _cast[Token](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 64: // ZeroOrMore
		return _cast[[][]_i0.Statement](p._stack.Peek(0).Sym)
	case 65: // ZeroOrMore
		{
			var zero [][]_i0.Statement
			return zero
		}
	case 66: // OneOrMore
		return append(
			_cast[[][]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 67: // OneOrMore
		return [][]_i0.Statement{
			_cast[[]_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 68: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 69: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 70: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 71: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 72: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 73: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 74: // List
		return append(
			_cast[[]*_i0.ParserProd](p._stack.Peek(2).Sym),
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		)
	case 75: // List
		return []*_i0.ParserProd{
			_cast[*_i0.ParserProd](p._stack.Peek(0).Sym),
		}
	case 76: // OneOrMore
		return append(
			_cast[[]*_i0.ParserTerm](p._stack.Peek(1).Sym),
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		)
	case 77: // OneOrMore
		return []*_i0.ParserTerm{
			_cast[*_i0.ParserTerm](p._stack.Peek(0).Sym),
		}
	case 78: // ZeroOrOne
		return _cast[*_i0.ProdQualifier](p._stack.Peek(0).Sym)
	case 79: // ZeroOrOne
		{
			var zero *_i0.ProdQualifier
			return zero
		}
	case 80: // ZeroOrOne
		return _cast[_i0.ParserTermType](p._stack.Peek(0).Sym)
	case 81: // ZeroOrOne
		{
			var zero _i0.ParserTermType
			return zero
		}
	case 82: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 83: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 84:
		{ // OneOrMoreF
			l := _cast[[]_i0.Statement](p._stack.Peek(1).Sym)
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 85:
		{ // OneOrMoreF
			var l []_i0.Statement
			e := _cast[_i0.Statement](p._stack.Peek(0).Sym)
			if !e.Discard() {
				l = append(l, e)
			}
			return l
		}
	case 86: // ZeroOrMore
		return _cast[[]Token](p._stack.Peek(0).Sym)
	case 87: // ZeroOrMore
		{
			var zero []Token
			return zero
		}
	case 88: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 89: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 90: // ZeroOrMore
		return _cast[[]_i0.Statement](p._stack.Peek(0).Sym)
	case 91: // ZeroOrMore
		{
			var zero []_i0.Statement
			return zero
		}
	case 92: // OneOrMore
		return append(
			_cast[[]_i0.Statement](p._stack.Peek(1).Sym),
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		)
	case 93: // OneOrMore
		return []_i0.Statement{
			_cast[_i0.Statement](p._stack.Peek(0).Sym),
		}
	case 94: // ZeroOrMore
		return _cast[[]_i0.Action](p._stack.Peek(0).Sym)
	case 95: // ZeroOrMore
		{
			var zero []_i0.Action
			return zero
		}
	case 96: // OneOrMore
		return append(
			_cast[[]_i0.Action](p._stack.Peek(1).Sym),
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		)
	case 97: // OneOrMore
		return []_i0.Action{
			_cast[_i0.Action](p._stack.Peek(0).Sym),
		}
	case 98: // OneOrMore
		return append(
			_cast[[]*_i0.ExternalName](p._stack.Peek(1).Sym),
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		)
	case 99: // OneOrMore
		return []*_i0.ExternalName{
			_cast[*_i0.ExternalName](p._stack.Peek(0).Sym),
		}
	case 100: // List
		return append(
			_cast[[]*_i0.LexerFactor](p._stack.Peek(2).Sym),
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		)
	case 101: // List
		return []*_i0.LexerFactor{
			_cast[*_i0.LexerFactor](p._stack.Peek(0).Sym),
		}
	case 102: // OneOrMore
		return append(
			_cast[[]*_i0.LexerTermCard](p._stack.Peek(1).Sym),
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		)
	case 103: // OneOrMore
		return []*_i0.LexerTermCard{
			_cast[*_i0.LexerTermCard](p._stack.Peek(0).Sym),
		}
	case 104: // ZeroOrOne
		return _cast[_i0.Card](p._stack.Peek(0).Sym)
	case 105: // ZeroOrOne
		{
			var zero _i0.Card
			return zero
		}
	case 106: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 107: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	case 108: // OneOrMore
		return append(
			_cast[[]Token](p._stack.Peek(1).Sym),
			_cast[Token](p._stack.Peek(0).Sym),
		)
	case 109: // OneOrMore
		return []Token{
			_cast[Token](p._stack.Peek(0).Sym),
		}
	case 110: // ZeroOrOne
		return _cast[Token](p._stack.Peek(0).Sym)
	case 111: // ZeroOrOne
		{
			var zero Token
			return zero
		}
	default:
		panic("unreachable")
	}
}
