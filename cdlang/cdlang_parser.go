// Code generated from CDLang.g4 by ANTLR 4.9. DO NOT EDIT.

package cdlang // CDLang
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 354,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 65, 10, 2, 12, 2,
	14, 2, 68, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 5, 5, 77, 10,
	5, 3, 5, 3, 5, 7, 5, 81, 10, 5, 12, 5, 14, 5, 84, 11, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 96, 10, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 104, 10, 6, 3, 6, 3, 6, 3, 7, 5, 7, 109,
	10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 118, 10, 7, 12,
	7, 14, 7, 121, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7,
	8, 131, 10, 8, 12, 8, 14, 8, 134, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8,
	140, 10, 8, 12, 8, 14, 8, 143, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 156, 10, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 174, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 186, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 208, 10,
	20, 12, 20, 14, 20, 211, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 217,
	10, 20, 12, 20, 14, 20, 220, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	226, 10, 20, 12, 20, 14, 20, 229, 11, 20, 3, 20, 5, 20, 232, 10, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 240, 10, 21, 12, 21, 14,
	21, 243, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 249, 10, 22, 3, 22,
	3, 22, 3, 22, 7, 22, 254, 10, 22, 12, 22, 14, 22, 257, 11, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 5, 23, 264, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 5, 24, 273, 10, 24, 3, 25, 5, 25, 276, 10, 25, 3,
	25, 3, 25, 3, 25, 7, 25, 281, 10, 25, 12, 25, 14, 25, 284, 11, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 291, 10, 26, 3, 26, 3, 26, 7, 26, 295,
	10, 26, 12, 26, 14, 26, 298, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	6, 27, 305, 10, 27, 13, 27, 14, 27, 306, 3, 27, 3, 27, 3, 28, 3, 28, 6,
	28, 313, 10, 28, 13, 28, 14, 28, 314, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28,
	321, 10, 28, 3, 28, 5, 28, 324, 10, 28, 3, 29, 3, 29, 7, 29, 328, 10, 29,
	12, 29, 14, 29, 331, 11, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 5, 30, 341, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 5, 30, 352, 10, 30, 3, 30, 2, 2, 31, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 2, 2, 2, 375, 2, 66, 3, 2, 2, 2, 4, 69,
	3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 76, 3, 2, 2, 2, 10, 87, 3, 2, 2, 2, 12,
	108, 3, 2, 2, 2, 14, 124, 3, 2, 2, 2, 16, 146, 3, 2, 2, 2, 18, 155, 3,
	2, 2, 2, 20, 157, 3, 2, 2, 2, 22, 159, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2,
	26, 175, 3, 2, 2, 2, 28, 185, 3, 2, 2, 2, 30, 187, 3, 2, 2, 2, 32, 193,
	3, 2, 2, 2, 34, 196, 3, 2, 2, 2, 36, 200, 3, 2, 2, 2, 38, 231, 3, 2, 2,
	2, 40, 233, 3, 2, 2, 2, 42, 248, 3, 2, 2, 2, 44, 263, 3, 2, 2, 2, 46, 265,
	3, 2, 2, 2, 48, 275, 3, 2, 2, 2, 50, 285, 3, 2, 2, 2, 52, 301, 3, 2, 2,
	2, 54, 310, 3, 2, 2, 2, 56, 329, 3, 2, 2, 2, 58, 351, 3, 2, 2, 2, 60, 65,
	5, 6, 4, 2, 61, 65, 5, 10, 6, 2, 62, 65, 5, 12, 7, 2, 63, 65, 5, 14, 8,
	2, 64, 60, 3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 63,
	3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 3, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 38, 2, 2, 70, 5, 3, 2,
	2, 2, 71, 72, 5, 4, 3, 2, 72, 73, 7, 3, 2, 2, 73, 74, 5, 42, 22, 2, 74,
	7, 3, 2, 2, 2, 75, 77, 7, 4, 2, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2,
	2, 77, 82, 3, 2, 2, 2, 78, 79, 7, 38, 2, 2, 79, 81, 7, 4, 2, 2, 80, 78,
	3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 86, 7, 38, 2, 2, 86, 9, 3,
	2, 2, 2, 87, 88, 7, 5, 2, 2, 88, 89, 5, 8, 5, 2, 89, 90, 7, 6, 2, 2, 90,
	91, 7, 38, 2, 2, 91, 95, 7, 7, 2, 2, 92, 93, 5, 8, 5, 2, 93, 94, 7, 6,
	2, 2, 94, 96, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97,
	3, 2, 2, 2, 97, 98, 7, 38, 2, 2, 98, 99, 7, 8, 2, 2, 99, 103, 7, 9, 2,
	2, 100, 101, 5, 8, 5, 2, 101, 102, 7, 6, 2, 2, 102, 104, 3, 2, 2, 2, 103,
	100, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106,
	7, 38, 2, 2, 106, 11, 3, 2, 2, 2, 107, 109, 7, 10, 2, 2, 108, 107, 3, 2,
	2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 7, 11, 2, 2,
	111, 112, 7, 38, 2, 2, 112, 113, 7, 7, 2, 2, 113, 114, 7, 8, 2, 2, 114,
	115, 5, 22, 12, 2, 115, 119, 7, 12, 2, 2, 116, 118, 5, 24, 13, 2, 117,
	116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120,
	3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 13,
	2, 2, 123, 13, 3, 2, 2, 2, 124, 125, 7, 11, 2, 2, 125, 126, 7, 38, 2, 2,
	126, 127, 7, 7, 2, 2, 127, 132, 5, 16, 9, 2, 128, 129, 7, 14, 2, 2, 129,
	131, 5, 16, 9, 2, 130, 128, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130,
	3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2,
	2, 2, 135, 136, 7, 8, 2, 2, 136, 137, 5, 22, 12, 2, 137, 141, 7, 12, 2,
	2, 138, 140, 5, 24, 13, 2, 139, 138, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2,
	141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 144, 3, 2, 2, 2, 143,
	141, 3, 2, 2, 2, 144, 145, 7, 13, 2, 2, 145, 15, 3, 2, 2, 2, 146, 147,
	7, 38, 2, 2, 147, 148, 7, 15, 2, 2, 148, 149, 7, 38, 2, 2, 149, 17, 3,
	2, 2, 2, 150, 151, 7, 16, 2, 2, 151, 156, 7, 38, 2, 2, 152, 156, 7, 17,
	2, 2, 153, 154, 7, 18, 2, 2, 154, 156, 7, 38, 2, 2, 155, 150, 3, 2, 2,
	2, 155, 152, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 19, 3, 2, 2, 2, 157,
	158, 7, 42, 2, 2, 158, 21, 3, 2, 2, 2, 159, 160, 7, 19, 2, 2, 160, 161,
	7, 43, 2, 2, 161, 162, 7, 6, 2, 2, 162, 163, 7, 43, 2, 2, 163, 164, 7,
	6, 2, 2, 164, 165, 7, 43, 2, 2, 165, 23, 3, 2, 2, 2, 166, 174, 5, 34, 18,
	2, 167, 174, 5, 36, 19, 2, 168, 174, 5, 38, 20, 2, 169, 174, 5, 40, 21,
	2, 170, 174, 5, 30, 16, 2, 171, 174, 5, 32, 17, 2, 172, 174, 5, 26, 14,
	2, 173, 166, 3, 2, 2, 2, 173, 167, 3, 2, 2, 2, 173, 168, 3, 2, 2, 2, 173,
	169, 3, 2, 2, 2, 173, 170, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 172,
	3, 2, 2, 2, 174, 25, 3, 2, 2, 2, 175, 176, 5, 42, 22, 2, 176, 177, 7, 20,
	2, 2, 177, 178, 7, 12, 2, 2, 178, 179, 5, 28, 15, 2, 179, 180, 7, 13, 2,
	2, 180, 27, 3, 2, 2, 2, 181, 182, 7, 21, 2, 2, 182, 183, 7, 14, 2, 2, 183,
	186, 5, 42, 22, 2, 184, 186, 7, 22, 2, 2, 185, 181, 3, 2, 2, 2, 185, 184,
	3, 2, 2, 2, 186, 29, 3, 2, 2, 2, 187, 188, 7, 38, 2, 2, 188, 189, 7, 3,
	2, 2, 189, 190, 7, 38, 2, 2, 190, 191, 7, 6, 2, 2, 191, 192, 7, 38, 2,
	2, 192, 31, 3, 2, 2, 2, 193, 194, 7, 23, 2, 2, 194, 195, 7, 38, 2, 2, 195,
	33, 3, 2, 2, 2, 196, 197, 7, 38, 2, 2, 197, 198, 7, 24, 2, 2, 198, 199,
	5, 42, 22, 2, 199, 35, 3, 2, 2, 2, 200, 201, 7, 38, 2, 2, 201, 202, 7,
	20, 2, 2, 202, 203, 5, 42, 22, 2, 203, 37, 3, 2, 2, 2, 204, 205, 7, 25,
	2, 2, 205, 209, 7, 12, 2, 2, 206, 208, 5, 24, 13, 2, 207, 206, 3, 2, 2,
	2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210,
	212, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 232, 7, 13, 2, 2, 213, 214,
	7, 26, 2, 2, 214, 218, 7, 12, 2, 2, 215, 217, 5, 24, 13, 2, 216, 215, 3,
	2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2,
	2, 219, 221, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 232, 7, 13, 2, 2, 222,
	223, 7, 27, 2, 2, 223, 227, 7, 12, 2, 2, 224, 226, 5, 36, 19, 2, 225, 224,
	3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2,
	2, 2, 228, 230, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 232, 7, 13, 2, 2,
	231, 204, 3, 2, 2, 2, 231, 213, 3, 2, 2, 2, 231, 222, 3, 2, 2, 2, 232,
	39, 3, 2, 2, 2, 233, 234, 7, 28, 2, 2, 234, 235, 7, 38, 2, 2, 235, 236,
	7, 7, 2, 2, 236, 241, 5, 18, 10, 2, 237, 238, 7, 14, 2, 2, 238, 240, 5,
	18, 10, 2, 239, 237, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2,
	2, 2, 241, 242, 3, 2, 2, 2, 242, 244, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2,
	244, 245, 7, 8, 2, 2, 245, 41, 3, 2, 2, 2, 246, 247, 7, 38, 2, 2, 247,
	249, 7, 6, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250,
	3, 2, 2, 2, 250, 251, 7, 38, 2, 2, 251, 255, 7, 12, 2, 2, 252, 254, 5,
	44, 23, 2, 253, 252, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2,
	2, 2, 255, 256, 3, 2, 2, 2, 256, 258, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2,
	258, 259, 7, 13, 2, 2, 259, 43, 3, 2, 2, 2, 260, 264, 5, 46, 24, 2, 261,
	264, 5, 50, 26, 2, 262, 264, 5, 52, 27, 2, 263, 260, 3, 2, 2, 2, 263, 261,
	3, 2, 2, 2, 263, 262, 3, 2, 2, 2, 264, 45, 3, 2, 2, 2, 265, 266, 7, 38,
	2, 2, 266, 272, 7, 15, 2, 2, 267, 273, 7, 43, 2, 2, 268, 273, 7, 42, 2,
	2, 269, 273, 5, 48, 25, 2, 270, 273, 5, 18, 10, 2, 271, 273, 5, 56, 29,
	2, 272, 267, 3, 2, 2, 2, 272, 268, 3, 2, 2, 2, 272, 269, 3, 2, 2, 2, 272,
	270, 3, 2, 2, 2, 272, 271, 3, 2, 2, 2, 273, 47, 3, 2, 2, 2, 274, 276, 7,
	4, 2, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 3, 2, 2,
	2, 277, 282, 7, 38, 2, 2, 278, 279, 7, 4, 2, 2, 279, 281, 7, 38, 2, 2,
	280, 278, 3, 2, 2, 2, 281, 284, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 49, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285, 290, 7,
	38, 2, 2, 286, 287, 7, 29, 2, 2, 287, 288, 5, 8, 5, 2, 288, 289, 7, 30,
	2, 2, 289, 291, 3, 2, 2, 2, 290, 286, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2,
	291, 292, 3, 2, 2, 2, 292, 296, 7, 12, 2, 2, 293, 295, 5, 44, 23, 2, 294,
	293, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 296, 297,
	3, 2, 2, 2, 297, 299, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 299, 300, 7, 13,
	2, 2, 300, 51, 3, 2, 2, 2, 301, 302, 7, 38, 2, 2, 302, 304, 7, 31, 2, 2,
	303, 305, 5, 54, 28, 2, 304, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306,
	304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309,
	7, 32, 2, 2, 309, 53, 3, 2, 2, 2, 310, 312, 7, 12, 2, 2, 311, 313, 5, 44,
	23, 2, 312, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2,
	314, 315, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 320, 7, 13, 2, 2, 317,
	321, 7, 33, 2, 2, 318, 321, 7, 34, 2, 2, 319, 321, 7, 35, 2, 2, 320, 317,
	3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 320, 321, 3, 2,
	2, 2, 321, 323, 3, 2, 2, 2, 322, 324, 7, 14, 2, 2, 323, 322, 3, 2, 2, 2,
	323, 324, 3, 2, 2, 2, 324, 55, 3, 2, 2, 2, 325, 326, 7, 36, 2, 2, 326,
	328, 5, 58, 30, 2, 327, 325, 3, 2, 2, 2, 328, 331, 3, 2, 2, 2, 329, 327,
	3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 57, 3, 2, 2, 2, 331, 329, 3, 2,
	2, 2, 332, 352, 7, 38, 2, 2, 333, 334, 7, 38, 2, 2, 334, 335, 7, 31, 2,
	2, 335, 336, 7, 38, 2, 2, 336, 340, 7, 37, 2, 2, 337, 341, 7, 33, 2, 2,
	338, 341, 5, 18, 10, 2, 339, 341, 7, 38, 2, 2, 340, 337, 3, 2, 2, 2, 340,
	338, 3, 2, 2, 2, 340, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 352,
	7, 32, 2, 2, 343, 344, 7, 38, 2, 2, 344, 345, 7, 31, 2, 2, 345, 346, 5,
	18, 10, 2, 346, 347, 7, 3, 2, 2, 347, 348, 7, 38, 2, 2, 348, 349, 7, 32,
	2, 2, 349, 352, 3, 2, 2, 2, 350, 352, 5, 18, 10, 2, 351, 332, 3, 2, 2,
	2, 351, 333, 3, 2, 2, 2, 351, 343, 3, 2, 2, 2, 351, 350, 3, 2, 2, 2, 352,
	59, 3, 2, 2, 2, 35, 64, 66, 76, 82, 95, 103, 108, 119, 132, 141, 155, 173,
	185, 209, 218, 227, 231, 241, 248, 255, 263, 272, 275, 282, 290, 296, 306,
	314, 320, 323, 329, 340, 351,
}
var literalNames = []string{
	"", "':='", "'::'", "'rpc'", "'.'", "'('", "')'", "'=>'", "'disabled'",
	"'scenario'", "'{'", "'}'", "','", "':'", "'$'", "'_'", "'#'", "'version'",
	"'>>'", "'OK'", "'ERROR'", "'close'", "'<<'", "'AtLeastOnce'", "'ZeroOrMore'",
	"'AnyOrder'", "'execute'", "'<'", "'>'", "'['", "']'", "'*'", "'+'", "'?'",
	"'/'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"ID", "WS", "LINE_COMMENT", "BLOCK_COMMENT", "STRING", "NUMBER",
}

var ruleNames = []string{
	"contract", "constName", "constProto", "domainName", "mapping", "scenario",
	"subScenario", "variableDeclaration", "variable", "constant", "version",
	"instruction", "call", "callResponse", "openStream", "closeStream", "send",
	"receive", "group", "execute", "protobuf", "protobufField", "protobufFieldSimple",
	"enumer", "protobufFieldGroup", "protobufFieldRepeated", "protobufFieldRepeatedRow",
	"path", "pathElement",
}

type CDLangParser struct {
	*antlr.BaseParser
}

// NewCDLangParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CDLangParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCDLangParser(input antlr.TokenStream) *CDLangParser {
	this := new(CDLangParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CDLang.g4"

	return this
}

// CDLangParser tokens.
const (
	CDLangParserEOF           = antlr.TokenEOF
	CDLangParserT__0          = 1
	CDLangParserT__1          = 2
	CDLangParserT__2          = 3
	CDLangParserT__3          = 4
	CDLangParserT__4          = 5
	CDLangParserT__5          = 6
	CDLangParserT__6          = 7
	CDLangParserT__7          = 8
	CDLangParserT__8          = 9
	CDLangParserT__9          = 10
	CDLangParserT__10         = 11
	CDLangParserT__11         = 12
	CDLangParserT__12         = 13
	CDLangParserT__13         = 14
	CDLangParserT__14         = 15
	CDLangParserT__15         = 16
	CDLangParserT__16         = 17
	CDLangParserT__17         = 18
	CDLangParserT__18         = 19
	CDLangParserT__19         = 20
	CDLangParserT__20         = 21
	CDLangParserT__21         = 22
	CDLangParserT__22         = 23
	CDLangParserT__23         = 24
	CDLangParserT__24         = 25
	CDLangParserT__25         = 26
	CDLangParserT__26         = 27
	CDLangParserT__27         = 28
	CDLangParserT__28         = 29
	CDLangParserT__29         = 30
	CDLangParserT__30         = 31
	CDLangParserT__31         = 32
	CDLangParserT__32         = 33
	CDLangParserT__33         = 34
	CDLangParserT__34         = 35
	CDLangParserID            = 36
	CDLangParserWS            = 37
	CDLangParserLINE_COMMENT  = 38
	CDLangParserBLOCK_COMMENT = 39
	CDLangParserSTRING        = 40
	CDLangParserNUMBER        = 41
)

// CDLangParser rules.
const (
	CDLangParserRULE_contract                 = 0
	CDLangParserRULE_constName                = 1
	CDLangParserRULE_constProto               = 2
	CDLangParserRULE_domainName               = 3
	CDLangParserRULE_mapping                  = 4
	CDLangParserRULE_scenario                 = 5
	CDLangParserRULE_subScenario              = 6
	CDLangParserRULE_variableDeclaration      = 7
	CDLangParserRULE_variable                 = 8
	CDLangParserRULE_constant                 = 9
	CDLangParserRULE_version                  = 10
	CDLangParserRULE_instruction              = 11
	CDLangParserRULE_call                     = 12
	CDLangParserRULE_callResponse             = 13
	CDLangParserRULE_openStream               = 14
	CDLangParserRULE_closeStream              = 15
	CDLangParserRULE_send                     = 16
	CDLangParserRULE_receive                  = 17
	CDLangParserRULE_group                    = 18
	CDLangParserRULE_execute                  = 19
	CDLangParserRULE_protobuf                 = 20
	CDLangParserRULE_protobufField            = 21
	CDLangParserRULE_protobufFieldSimple      = 22
	CDLangParserRULE_enumer                   = 23
	CDLangParserRULE_protobufFieldGroup       = 24
	CDLangParserRULE_protobufFieldRepeated    = 25
	CDLangParserRULE_protobufFieldRepeatedRow = 26
	CDLangParserRULE_path                     = 27
	CDLangParserRULE_pathElement              = 28
)

// IContractContext is an interface to support dynamic dispatch.
type IContractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContractContext differentiates from other interfaces.
	IsContractContext()
}

type ContractContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContractContext() *ContractContext {
	var p = new(ContractContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_contract
	return p
}

func (*ContractContext) IsContractContext() {}

func NewContractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContractContext {
	var p = new(ContractContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_contract

	return p
}

func (s *ContractContext) GetParser() antlr.Parser { return s.parser }

func (s *ContractContext) AllConstProto() []IConstProtoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstProtoContext)(nil)).Elem())
	var tst = make([]IConstProtoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstProtoContext)
		}
	}

	return tst
}

func (s *ContractContext) ConstProto(i int) IConstProtoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstProtoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstProtoContext)
}

func (s *ContractContext) AllMapping() []IMappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMappingContext)(nil)).Elem())
	var tst = make([]IMappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMappingContext)
		}
	}

	return tst
}

func (s *ContractContext) Mapping(i int) IMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMappingContext)
}

func (s *ContractContext) AllScenario() []IScenarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScenarioContext)(nil)).Elem())
	var tst = make([]IScenarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScenarioContext)
		}
	}

	return tst
}

func (s *ContractContext) Scenario(i int) IScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScenarioContext)
}

func (s *ContractContext) AllSubScenario() []ISubScenarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubScenarioContext)(nil)).Elem())
	var tst = make([]ISubScenarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubScenarioContext)
		}
	}

	return tst
}

func (s *ContractContext) SubScenario(i int) ISubScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubScenarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubScenarioContext)
}

func (s *ContractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContractContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterContract(s)
	}
}

func (s *ContractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitContract(s)
	}
}

func (s *ContractContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitContract(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Contract() (localctx IContractContext) {
	localctx = NewContractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CDLangParserRULE_contract)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CDLangParserT__2)|(1<<CDLangParserT__7)|(1<<CDLangParserT__8))) != 0) || _la == CDLangParserID {
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(58)
				p.ConstProto()
			}

		case 2:
			{
				p.SetState(59)
				p.Mapping()
			}

		case 3:
			{
				p.SetState(60)
				p.Scenario()
			}

		case 4:
			{
				p.SetState(61)
				p.SubScenario()
			}

		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstNameContext is an interface to support dynamic dispatch.
type IConstNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// IsConstNameContext differentiates from other interfaces.
	IsConstNameContext()
}

type ConstNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      antlr.Token
}

func NewEmptyConstNameContext() *ConstNameContext {
	var p = new(ConstNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_constName
	return p
}

func (*ConstNameContext) IsConstNameContext() {}

func NewConstNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstNameContext {
	var p = new(ConstNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_constName

	return p
}

func (s *ConstNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstNameContext) GetN() antlr.Token { return s.n }

func (s *ConstNameContext) SetN(v antlr.Token) { s.n = v }

func (s *ConstNameContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ConstNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterConstName(s)
	}
}

func (s *ConstNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitConstName(s)
	}
}

func (s *ConstNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitConstName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ConstName() (localctx IConstNameContext) {
	localctx = NewConstNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CDLangParserRULE_constName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)

		var _m = p.Match(CDLangParserID)

		localctx.(*ConstNameContext).n = _m
	}

	return localctx
}

// IConstProtoContext is an interface to support dynamic dispatch.
type IConstProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n rule contexts.
	GetN() IConstNameContext

	// GetP returns the p rule contexts.
	GetP() IProtobufContext

	// SetN sets the n rule contexts.
	SetN(IConstNameContext)

	// SetP sets the p rule contexts.
	SetP(IProtobufContext)

	// IsConstProtoContext differentiates from other interfaces.
	IsConstProtoContext()
}

type ConstProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      IConstNameContext
	p      IProtobufContext
}

func NewEmptyConstProtoContext() *ConstProtoContext {
	var p = new(ConstProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_constProto
	return p
}

func (*ConstProtoContext) IsConstProtoContext() {}

func NewConstProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstProtoContext {
	var p = new(ConstProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_constProto

	return p
}

func (s *ConstProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstProtoContext) GetN() IConstNameContext { return s.n }

func (s *ConstProtoContext) GetP() IProtobufContext { return s.p }

func (s *ConstProtoContext) SetN(v IConstNameContext) { s.n = v }

func (s *ConstProtoContext) SetP(v IProtobufContext) { s.p = v }

func (s *ConstProtoContext) ConstName() IConstNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstNameContext)
}

func (s *ConstProtoContext) Protobuf() IProtobufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufContext)
}

func (s *ConstProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterConstProto(s)
	}
}

func (s *ConstProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitConstProto(s)
	}
}

func (s *ConstProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitConstProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ConstProto() (localctx IConstProtoContext) {
	localctx = NewConstProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CDLangParserRULE_constProto)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)

		var _x = p.ConstName()

		localctx.(*ConstProtoContext).n = _x
	}
	{
		p.SetState(70)
		p.Match(CDLangParserT__0)
	}
	{
		p.SetState(71)

		var _x = p.Protobuf()

		localctx.(*ConstProtoContext).p = _x
	}

	return localctx
}

// IDomainNameContext is an interface to support dynamic dispatch.
type IDomainNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainNameContext differentiates from other interfaces.
	IsDomainNameContext()
}

type DomainNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainNameContext() *DomainNameContext {
	var p = new(DomainNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_domainName
	return p
}

func (*DomainNameContext) IsDomainNameContext() {}

func NewDomainNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainNameContext {
	var p = new(DomainNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_domainName

	return p
}

func (s *DomainNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *DomainNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *DomainNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterDomainName(s)
	}
}

func (s *DomainNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitDomainName(s)
	}
}

func (s *DomainNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitDomainName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) DomainName() (localctx IDomainNameContext) {
	localctx = NewDomainNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CDLangParserRULE_domainName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDLangParserT__1 {
		{
			p.SetState(73)
			p.Match(CDLangParserT__1)
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(76)
				p.Match(CDLangParserID)
			}
			{
				p.SetState(77)
				p.Match(CDLangParserT__1)
			}

		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(83)
		p.Match(CDLangParserID)
	}

	return localctx
}

// IMappingContext is an interface to support dynamic dispatch.
type IMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f token.
	GetF() antlr.Token

	// GetReq returns the req token.
	GetReq() antlr.Token

	// GetResp returns the resp token.
	GetResp() antlr.Token

	// SetF sets the f token.
	SetF(antlr.Token)

	// SetReq sets the req token.
	SetReq(antlr.Token)

	// SetResp sets the resp token.
	SetResp(antlr.Token)

	// GetFd returns the fd rule contexts.
	GetFd() IDomainNameContext

	// GetReqd returns the reqd rule contexts.
	GetReqd() IDomainNameContext

	// GetRespd returns the respd rule contexts.
	GetRespd() IDomainNameContext

	// SetFd sets the fd rule contexts.
	SetFd(IDomainNameContext)

	// SetReqd sets the reqd rule contexts.
	SetReqd(IDomainNameContext)

	// SetRespd sets the respd rule contexts.
	SetRespd(IDomainNameContext)

	// IsMappingContext differentiates from other interfaces.
	IsMappingContext()
}

type MappingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fd     IDomainNameContext
	f      antlr.Token
	reqd   IDomainNameContext
	req    antlr.Token
	respd  IDomainNameContext
	resp   antlr.Token
}

func NewEmptyMappingContext() *MappingContext {
	var p = new(MappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_mapping
	return p
}

func (*MappingContext) IsMappingContext() {}

func NewMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingContext {
	var p = new(MappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_mapping

	return p
}

func (s *MappingContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingContext) GetF() antlr.Token { return s.f }

func (s *MappingContext) GetReq() antlr.Token { return s.req }

func (s *MappingContext) GetResp() antlr.Token { return s.resp }

func (s *MappingContext) SetF(v antlr.Token) { s.f = v }

func (s *MappingContext) SetReq(v antlr.Token) { s.req = v }

func (s *MappingContext) SetResp(v antlr.Token) { s.resp = v }

func (s *MappingContext) GetFd() IDomainNameContext { return s.fd }

func (s *MappingContext) GetReqd() IDomainNameContext { return s.reqd }

func (s *MappingContext) GetRespd() IDomainNameContext { return s.respd }

func (s *MappingContext) SetFd(v IDomainNameContext) { s.fd = v }

func (s *MappingContext) SetReqd(v IDomainNameContext) { s.reqd = v }

func (s *MappingContext) SetRespd(v IDomainNameContext) { s.respd = v }

func (s *MappingContext) AllDomainName() []IDomainNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDomainNameContext)(nil)).Elem())
	var tst = make([]IDomainNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDomainNameContext)
		}
	}

	return tst
}

func (s *MappingContext) DomainName(i int) IDomainNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDomainNameContext)
}

func (s *MappingContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *MappingContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *MappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterMapping(s)
	}
}

func (s *MappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitMapping(s)
	}
}

func (s *MappingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitMapping(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Mapping() (localctx IMappingContext) {
	localctx = NewMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CDLangParserRULE_mapping)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(CDLangParserT__2)
	}
	{
		p.SetState(86)

		var _x = p.DomainName()

		localctx.(*MappingContext).fd = _x
	}
	{
		p.SetState(87)
		p.Match(CDLangParserT__3)
	}
	{
		p.SetState(88)

		var _m = p.Match(CDLangParserID)

		localctx.(*MappingContext).f = _m
	}
	{
		p.SetState(89)
		p.Match(CDLangParserT__4)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(90)

			var _x = p.DomainName()

			localctx.(*MappingContext).reqd = _x
		}
		{
			p.SetState(91)
			p.Match(CDLangParserT__3)
		}

	}
	{
		p.SetState(95)

		var _m = p.Match(CDLangParserID)

		localctx.(*MappingContext).req = _m
	}
	{
		p.SetState(96)
		p.Match(CDLangParserT__5)
	}
	{
		p.SetState(97)
		p.Match(CDLangParserT__6)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(98)

			var _x = p.DomainName()

			localctx.(*MappingContext).respd = _x
		}
		{
			p.SetState(99)
			p.Match(CDLangParserT__3)
		}

	}
	{
		p.SetState(103)

		var _m = p.Match(CDLangParserID)

		localctx.(*MappingContext).resp = _m
	}

	return localctx
}

// IScenarioContext is an interface to support dynamic dispatch.
type IScenarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDisabled returns the disabled token.
	GetDisabled() antlr.Token

	// GetSos returns the sos token.
	GetSos() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// GetEos returns the eos token.
	GetEos() antlr.Token

	// SetDisabled sets the disabled token.
	SetDisabled(antlr.Token)

	// SetSos sets the sos token.
	SetSos(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetEos sets the eos token.
	SetEos(antlr.Token)

	// GetVer returns the ver rule contexts.
	GetVer() IVersionContext

	// SetVer sets the ver rule contexts.
	SetVer(IVersionContext)

	// IsScenarioContext differentiates from other interfaces.
	IsScenarioContext()
}

type ScenarioContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	disabled antlr.Token
	sos      antlr.Token
	name     antlr.Token
	ver      IVersionContext
	eos      antlr.Token
}

func NewEmptyScenarioContext() *ScenarioContext {
	var p = new(ScenarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_scenario
	return p
}

func (*ScenarioContext) IsScenarioContext() {}

func NewScenarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScenarioContext {
	var p = new(ScenarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_scenario

	return p
}

func (s *ScenarioContext) GetParser() antlr.Parser { return s.parser }

func (s *ScenarioContext) GetDisabled() antlr.Token { return s.disabled }

func (s *ScenarioContext) GetSos() antlr.Token { return s.sos }

func (s *ScenarioContext) GetName() antlr.Token { return s.name }

func (s *ScenarioContext) GetEos() antlr.Token { return s.eos }

func (s *ScenarioContext) SetDisabled(v antlr.Token) { s.disabled = v }

func (s *ScenarioContext) SetSos(v antlr.Token) { s.sos = v }

func (s *ScenarioContext) SetName(v antlr.Token) { s.name = v }

func (s *ScenarioContext) SetEos(v antlr.Token) { s.eos = v }

func (s *ScenarioContext) GetVer() IVersionContext { return s.ver }

func (s *ScenarioContext) SetVer(v IVersionContext) { s.ver = v }

func (s *ScenarioContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ScenarioContext) Version() IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *ScenarioContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ScenarioContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ScenarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScenarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScenarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterScenario(s)
	}
}

func (s *ScenarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitScenario(s)
	}
}

func (s *ScenarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitScenario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Scenario() (localctx IScenarioContext) {
	localctx = NewScenarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CDLangParserRULE_scenario)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDLangParserT__7 {
		{
			p.SetState(105)

			var _m = p.Match(CDLangParserT__7)

			localctx.(*ScenarioContext).disabled = _m
		}

	}
	{
		p.SetState(108)

		var _m = p.Match(CDLangParserT__8)

		localctx.(*ScenarioContext).sos = _m
	}
	{
		p.SetState(109)

		var _m = p.Match(CDLangParserID)

		localctx.(*ScenarioContext).name = _m
	}
	{
		p.SetState(110)
		p.Match(CDLangParserT__4)
	}
	{
		p.SetState(111)
		p.Match(CDLangParserT__5)
	}
	{
		p.SetState(112)

		var _x = p.Version()

		localctx.(*ScenarioContext).ver = _x
	}
	{
		p.SetState(113)
		p.Match(CDLangParserT__9)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CDLangParserT__20-21))|(1<<(CDLangParserT__22-21))|(1<<(CDLangParserT__23-21))|(1<<(CDLangParserT__24-21))|(1<<(CDLangParserT__25-21))|(1<<(CDLangParserID-21)))) != 0 {
		{
			p.SetState(114)
			p.Instruction()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)

		var _m = p.Match(CDLangParserT__10)

		localctx.(*ScenarioContext).eos = _m
	}

	return localctx
}

// ISubScenarioContext is an interface to support dynamic dispatch.
type ISubScenarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSos returns the sos token.
	GetSos() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// GetEos returns the eos token.
	GetEos() antlr.Token

	// SetSos sets the sos token.
	SetSos(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetEos sets the eos token.
	SetEos(antlr.Token)

	// GetVer returns the ver rule contexts.
	GetVer() IVersionContext

	// SetVer sets the ver rule contexts.
	SetVer(IVersionContext)

	// IsSubScenarioContext differentiates from other interfaces.
	IsSubScenarioContext()
}

type SubScenarioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sos    antlr.Token
	name   antlr.Token
	ver    IVersionContext
	eos    antlr.Token
}

func NewEmptySubScenarioContext() *SubScenarioContext {
	var p = new(SubScenarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_subScenario
	return p
}

func (*SubScenarioContext) IsSubScenarioContext() {}

func NewSubScenarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubScenarioContext {
	var p = new(SubScenarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_subScenario

	return p
}

func (s *SubScenarioContext) GetParser() antlr.Parser { return s.parser }

func (s *SubScenarioContext) GetSos() antlr.Token { return s.sos }

func (s *SubScenarioContext) GetName() antlr.Token { return s.name }

func (s *SubScenarioContext) GetEos() antlr.Token { return s.eos }

func (s *SubScenarioContext) SetSos(v antlr.Token) { s.sos = v }

func (s *SubScenarioContext) SetName(v antlr.Token) { s.name = v }

func (s *SubScenarioContext) SetEos(v antlr.Token) { s.eos = v }

func (s *SubScenarioContext) GetVer() IVersionContext { return s.ver }

func (s *SubScenarioContext) SetVer(v IVersionContext) { s.ver = v }

func (s *SubScenarioContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *SubScenarioContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *SubScenarioContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *SubScenarioContext) Version() IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *SubScenarioContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *SubScenarioContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *SubScenarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubScenarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubScenarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterSubScenario(s)
	}
}

func (s *SubScenarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitSubScenario(s)
	}
}

func (s *SubScenarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitSubScenario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) SubScenario() (localctx ISubScenarioContext) {
	localctx = NewSubScenarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CDLangParserRULE_subScenario)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)

		var _m = p.Match(CDLangParserT__8)

		localctx.(*SubScenarioContext).sos = _m
	}
	{
		p.SetState(123)

		var _m = p.Match(CDLangParserID)

		localctx.(*SubScenarioContext).name = _m
	}
	{
		p.SetState(124)
		p.Match(CDLangParserT__4)
	}
	{
		p.SetState(125)
		p.VariableDeclaration()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserT__11 {
		{
			p.SetState(126)
			p.Match(CDLangParserT__11)
		}
		{
			p.SetState(127)
			p.VariableDeclaration()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(CDLangParserT__5)
	}
	{
		p.SetState(134)

		var _x = p.Version()

		localctx.(*SubScenarioContext).ver = _x
	}
	{
		p.SetState(135)
		p.Match(CDLangParserT__9)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CDLangParserT__20-21))|(1<<(CDLangParserT__22-21))|(1<<(CDLangParserT__23-21))|(1<<(CDLangParserT__24-21))|(1<<(CDLangParserT__25-21))|(1<<(CDLangParserID-21)))) != 0 {
		{
			p.SetState(136)
			p.Instruction()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)

		var _m = p.Match(CDLangParserT__10)

		localctx.(*SubScenarioContext).eos = _m
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetType_name returns the type_name token.
	GetType_name() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetType_name sets the type_name token.
	SetType_name(antlr.Token)

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	name      antlr.Token
	type_name antlr.Token
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) GetName() antlr.Token { return s.name }

func (s *VariableDeclarationContext) GetType_name() antlr.Token { return s.type_name }

func (s *VariableDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableDeclarationContext) SetType_name(v antlr.Token) { s.type_name = v }

func (s *VariableDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *VariableDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CDLangParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)

		var _m = p.Match(CDLangParserID)

		localctx.(*VariableDeclarationContext).name = _m
	}
	{
		p.SetState(145)
		p.Match(CDLangParserT__12)
	}
	{
		p.SetState(146)

		var _m = p.Match(CDLangParserID)

		localctx.(*VariableDeclarationContext).type_name = _m
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStr_name returns the str_name token.
	GetStr_name() antlr.Token

	// GetSkiped returns the skiped token.
	GetSkiped() antlr.Token

	// GetNum_name returns the num_name token.
	GetNum_name() antlr.Token

	// SetStr_name sets the str_name token.
	SetStr_name(antlr.Token)

	// SetSkiped sets the skiped token.
	SetSkiped(antlr.Token)

	// SetNum_name sets the num_name token.
	SetNum_name(antlr.Token)

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	str_name antlr.Token
	skiped   antlr.Token
	num_name antlr.Token
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) GetStr_name() antlr.Token { return s.str_name }

func (s *VariableContext) GetSkiped() antlr.Token { return s.skiped }

func (s *VariableContext) GetNum_name() antlr.Token { return s.num_name }

func (s *VariableContext) SetStr_name(v antlr.Token) { s.str_name = v }

func (s *VariableContext) SetSkiped(v antlr.Token) { s.skiped = v }

func (s *VariableContext) SetNum_name(v antlr.Token) { s.num_name = v }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CDLangParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CDLangParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(CDLangParserT__13)
		}
		{
			p.SetState(149)

			var _m = p.Match(CDLangParserID)

			localctx.(*VariableContext).str_name = _m
		}

	case CDLangParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)

			var _m = p.Match(CDLangParserT__14)

			localctx.(*VariableContext).skiped = _m
		}

	case CDLangParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(CDLangParserT__15)
		}
		{
			p.SetState(152)

			var _m = p.Match(CDLangParserID)

			localctx.(*VariableContext).num_name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStr returns the str token.
	GetStr() antlr.Token

	// SetStr sets the str token.
	SetStr(antlr.Token)

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	str    antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetStr() antlr.Token { return s.str }

func (s *ConstantContext) SetStr(v antlr.Token) { s.str = v }

func (s *ConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(CDLangParserSTRING, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CDLangParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)

		var _m = p.Match(CDLangParserSTRING)

		localctx.(*ConstantContext).str = _m
	}

	return localctx
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMajor returns the major token.
	GetMajor() antlr.Token

	// GetMinor returns the minor token.
	GetMinor() antlr.Token

	// GetPatch returns the patch token.
	GetPatch() antlr.Token

	// SetMajor sets the major token.
	SetMajor(antlr.Token)

	// SetMinor sets the minor token.
	SetMinor(antlr.Token)

	// SetPatch sets the patch token.
	SetPatch(antlr.Token)

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	major  antlr.Token
	minor  antlr.Token
	patch  antlr.Token
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_version
	return p
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) GetMajor() antlr.Token { return s.major }

func (s *VersionContext) GetMinor() antlr.Token { return s.minor }

func (s *VersionContext) GetPatch() antlr.Token { return s.patch }

func (s *VersionContext) SetMajor(v antlr.Token) { s.major = v }

func (s *VersionContext) SetMinor(v antlr.Token) { s.minor = v }

func (s *VersionContext) SetPatch(v antlr.Token) { s.patch = v }

func (s *VersionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserNUMBER)
}

func (s *VersionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserNUMBER, i)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CDLangParserRULE_version)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(CDLangParserT__16)
	}
	{
		p.SetState(158)

		var _m = p.Match(CDLangParserNUMBER)

		localctx.(*VersionContext).major = _m
	}
	{
		p.SetState(159)
		p.Match(CDLangParserT__3)
	}
	{
		p.SetState(160)

		var _m = p.Match(CDLangParserNUMBER)

		localctx.(*VersionContext).minor = _m
	}
	{
		p.SetState(161)
		p.Match(CDLangParserT__3)
	}
	{
		p.SetState(162)

		var _m = p.Match(CDLangParserNUMBER)

		localctx.(*VersionContext).patch = _m
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Send() ISendContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISendContext)
}

func (s *InstructionContext) Receive() IReceiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReceiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReceiveContext)
}

func (s *InstructionContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *InstructionContext) Execute() IExecuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *InstructionContext) OpenStream() IOpenStreamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenStreamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenStreamContext)
}

func (s *InstructionContext) CloseStream() ICloseStreamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICloseStreamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICloseStreamContext)
}

func (s *InstructionContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (s *InstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CDLangParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Send()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Receive()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Group()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(167)
			p.Execute()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(168)
			p.OpenStream()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(169)
			p.CloseStream()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(170)
			p.Call()
		}

	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReq returns the req rule contexts.
	GetReq() IProtobufContext

	// SetReq sets the req rule contexts.
	SetReq(IProtobufContext)

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	req    IProtobufContext
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) GetReq() IProtobufContext { return s.req }

func (s *CallContext) SetReq(v IProtobufContext) { s.req = v }

func (s *CallContext) CallResponse() ICallResponseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallResponseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallResponseContext)
}

func (s *CallContext) Protobuf() IProtobufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CDLangParserRULE_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)

		var _x = p.Protobuf()

		localctx.(*CallContext).req = _x
	}
	{
		p.SetState(174)
		p.Match(CDLangParserT__17)
	}
	{
		p.SetState(175)
		p.Match(CDLangParserT__9)
	}
	{
		p.SetState(176)
		p.CallResponse()
	}
	{
		p.SetState(177)
		p.Match(CDLangParserT__10)
	}

	return localctx
}

// ICallResponseContext is an interface to support dynamic dispatch.
type ICallResponseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOk returns the ok token.
	GetOk() antlr.Token

	// GetErr returns the err token.
	GetErr() antlr.Token

	// SetOk sets the ok token.
	SetOk(antlr.Token)

	// SetErr sets the err token.
	SetErr(antlr.Token)

	// GetResp returns the resp rule contexts.
	GetResp() IProtobufContext

	// SetResp sets the resp rule contexts.
	SetResp(IProtobufContext)

	// IsCallResponseContext differentiates from other interfaces.
	IsCallResponseContext()
}

type CallResponseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ok     antlr.Token
	resp   IProtobufContext
	err    antlr.Token
}

func NewEmptyCallResponseContext() *CallResponseContext {
	var p = new(CallResponseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_callResponse
	return p
}

func (*CallResponseContext) IsCallResponseContext() {}

func NewCallResponseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallResponseContext {
	var p = new(CallResponseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_callResponse

	return p
}

func (s *CallResponseContext) GetParser() antlr.Parser { return s.parser }

func (s *CallResponseContext) GetOk() antlr.Token { return s.ok }

func (s *CallResponseContext) GetErr() antlr.Token { return s.err }

func (s *CallResponseContext) SetOk(v antlr.Token) { s.ok = v }

func (s *CallResponseContext) SetErr(v antlr.Token) { s.err = v }

func (s *CallResponseContext) GetResp() IProtobufContext { return s.resp }

func (s *CallResponseContext) SetResp(v IProtobufContext) { s.resp = v }

func (s *CallResponseContext) Protobuf() IProtobufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufContext)
}

func (s *CallResponseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallResponseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallResponseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterCallResponse(s)
	}
}

func (s *CallResponseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitCallResponse(s)
	}
}

func (s *CallResponseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitCallResponse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) CallResponse() (localctx ICallResponseContext) {
	localctx = NewCallResponseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CDLangParserRULE_callResponse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CDLangParserT__18:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)

			var _m = p.Match(CDLangParserT__18)

			localctx.(*CallResponseContext).ok = _m
		}
		{
			p.SetState(180)
			p.Match(CDLangParserT__11)
		}
		{
			p.SetState(181)

			var _x = p.Protobuf()

			localctx.(*CallResponseContext).resp = _x
		}

	case CDLangParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)

			var _m = p.Match(CDLangParserT__19)

			localctx.(*CallResponseContext).err = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOpenStreamContext is an interface to support dynamic dispatch.
type IOpenStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStream returns the stream token.
	GetStream() antlr.Token

	// GetDomain returns the domain token.
	GetDomain() antlr.Token

	// GetMethod returns the method token.
	GetMethod() antlr.Token

	// SetStream sets the stream token.
	SetStream(antlr.Token)

	// SetDomain sets the domain token.
	SetDomain(antlr.Token)

	// SetMethod sets the method token.
	SetMethod(antlr.Token)

	// IsOpenStreamContext differentiates from other interfaces.
	IsOpenStreamContext()
}

type OpenStreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stream antlr.Token
	domain antlr.Token
	method antlr.Token
}

func NewEmptyOpenStreamContext() *OpenStreamContext {
	var p = new(OpenStreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_openStream
	return p
}

func (*OpenStreamContext) IsOpenStreamContext() {}

func NewOpenStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenStreamContext {
	var p = new(OpenStreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_openStream

	return p
}

func (s *OpenStreamContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenStreamContext) GetStream() antlr.Token { return s.stream }

func (s *OpenStreamContext) GetDomain() antlr.Token { return s.domain }

func (s *OpenStreamContext) GetMethod() antlr.Token { return s.method }

func (s *OpenStreamContext) SetStream(v antlr.Token) { s.stream = v }

func (s *OpenStreamContext) SetDomain(v antlr.Token) { s.domain = v }

func (s *OpenStreamContext) SetMethod(v antlr.Token) { s.method = v }

func (s *OpenStreamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *OpenStreamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *OpenStreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenStreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenStreamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterOpenStream(s)
	}
}

func (s *OpenStreamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitOpenStream(s)
	}
}

func (s *OpenStreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitOpenStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) OpenStream() (localctx IOpenStreamContext) {
	localctx = NewOpenStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CDLangParserRULE_openStream)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)

		var _m = p.Match(CDLangParserID)

		localctx.(*OpenStreamContext).stream = _m
	}
	{
		p.SetState(186)
		p.Match(CDLangParserT__0)
	}
	{
		p.SetState(187)

		var _m = p.Match(CDLangParserID)

		localctx.(*OpenStreamContext).domain = _m
	}
	{
		p.SetState(188)
		p.Match(CDLangParserT__3)
	}
	{
		p.SetState(189)

		var _m = p.Match(CDLangParserID)

		localctx.(*OpenStreamContext).method = _m
	}

	return localctx
}

// ICloseStreamContext is an interface to support dynamic dispatch.
type ICloseStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStream returns the stream token.
	GetStream() antlr.Token

	// SetStream sets the stream token.
	SetStream(antlr.Token)

	// IsCloseStreamContext differentiates from other interfaces.
	IsCloseStreamContext()
}

type CloseStreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	stream antlr.Token
}

func NewEmptyCloseStreamContext() *CloseStreamContext {
	var p = new(CloseStreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_closeStream
	return p
}

func (*CloseStreamContext) IsCloseStreamContext() {}

func NewCloseStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseStreamContext {
	var p = new(CloseStreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_closeStream

	return p
}

func (s *CloseStreamContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseStreamContext) GetStream() antlr.Token { return s.stream }

func (s *CloseStreamContext) SetStream(v antlr.Token) { s.stream = v }

func (s *CloseStreamContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *CloseStreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseStreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloseStreamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterCloseStream(s)
	}
}

func (s *CloseStreamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitCloseStream(s)
	}
}

func (s *CloseStreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitCloseStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) CloseStream() (localctx ICloseStreamContext) {
	localctx = NewCloseStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CDLangParserRULE_closeStream)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(CDLangParserT__20)
	}
	{
		p.SetState(192)

		var _m = p.Match(CDLangParserID)

		localctx.(*CloseStreamContext).stream = _m
	}

	return localctx
}

// ISendContext is an interface to support dynamic dispatch.
type ISendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCh returns the ch token.
	GetCh() antlr.Token

	// SetCh sets the ch token.
	SetCh(antlr.Token)

	// IsSendContext differentiates from other interfaces.
	IsSendContext()
}

type SendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ch     antlr.Token
}

func NewEmptySendContext() *SendContext {
	var p = new(SendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_send
	return p
}

func (*SendContext) IsSendContext() {}

func NewSendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendContext {
	var p = new(SendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_send

	return p
}

func (s *SendContext) GetParser() antlr.Parser { return s.parser }

func (s *SendContext) GetCh() antlr.Token { return s.ch }

func (s *SendContext) SetCh(v antlr.Token) { s.ch = v }

func (s *SendContext) Protobuf() IProtobufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufContext)
}

func (s *SendContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *SendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterSend(s)
	}
}

func (s *SendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitSend(s)
	}
}

func (s *SendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitSend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Send() (localctx ISendContext) {
	localctx = NewSendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CDLangParserRULE_send)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)

		var _m = p.Match(CDLangParserID)

		localctx.(*SendContext).ch = _m
	}
	{
		p.SetState(195)
		p.Match(CDLangParserT__21)
	}
	{
		p.SetState(196)
		p.Protobuf()
	}

	return localctx
}

// IReceiveContext is an interface to support dynamic dispatch.
type IReceiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCh returns the ch token.
	GetCh() antlr.Token

	// SetCh sets the ch token.
	SetCh(antlr.Token)

	// IsReceiveContext differentiates from other interfaces.
	IsReceiveContext()
}

type ReceiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ch     antlr.Token
}

func NewEmptyReceiveContext() *ReceiveContext {
	var p = new(ReceiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_receive
	return p
}

func (*ReceiveContext) IsReceiveContext() {}

func NewReceiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiveContext {
	var p = new(ReceiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_receive

	return p
}

func (s *ReceiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiveContext) GetCh() antlr.Token { return s.ch }

func (s *ReceiveContext) SetCh(v antlr.Token) { s.ch = v }

func (s *ReceiveContext) Protobuf() IProtobufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufContext)
}

func (s *ReceiveContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ReceiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterReceive(s)
	}
}

func (s *ReceiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitReceive(s)
	}
}

func (s *ReceiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitReceive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Receive() (localctx IReceiveContext) {
	localctx = NewReceiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CDLangParserRULE_receive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)

		var _m = p.Match(CDLangParserID)

		localctx.(*ReceiveContext).ch = _m
	}
	{
		p.SetState(199)
		p.Match(CDLangParserT__17)
	}
	{
		p.SetState(200)
		p.Protobuf()
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAt_least_once returns the at_least_once token.
	GetAt_least_once() antlr.Token

	// GetZero_or_more returns the zero_or_more token.
	GetZero_or_more() antlr.Token

	// GetAny_order returns the any_order token.
	GetAny_order() antlr.Token

	// SetAt_least_once sets the at_least_once token.
	SetAt_least_once(antlr.Token)

	// SetZero_or_more sets the zero_or_more token.
	SetZero_or_more(antlr.Token)

	// SetAny_order sets the any_order token.
	SetAny_order(antlr.Token)

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	at_least_once antlr.Token
	zero_or_more  antlr.Token
	any_order     antlr.Token
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) GetAt_least_once() antlr.Token { return s.at_least_once }

func (s *GroupContext) GetZero_or_more() antlr.Token { return s.zero_or_more }

func (s *GroupContext) GetAny_order() antlr.Token { return s.any_order }

func (s *GroupContext) SetAt_least_once(v antlr.Token) { s.at_least_once = v }

func (s *GroupContext) SetZero_or_more(v antlr.Token) { s.zero_or_more = v }

func (s *GroupContext) SetAny_order(v antlr.Token) { s.any_order = v }

func (s *GroupContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *GroupContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *GroupContext) AllReceive() []IReceiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReceiveContext)(nil)).Elem())
	var tst = make([]IReceiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReceiveContext)
		}
	}

	return tst
}

func (s *GroupContext) Receive(i int) IReceiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReceiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReceiveContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CDLangParserRULE_group)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CDLangParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)

			var _m = p.Match(CDLangParserT__22)

			localctx.(*GroupContext).at_least_once = _m
		}
		{
			p.SetState(203)
			p.Match(CDLangParserT__9)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CDLangParserT__20-21))|(1<<(CDLangParserT__22-21))|(1<<(CDLangParserT__23-21))|(1<<(CDLangParserT__24-21))|(1<<(CDLangParserT__25-21))|(1<<(CDLangParserID-21)))) != 0 {
			{
				p.SetState(204)
				p.Instruction()
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(210)
			p.Match(CDLangParserT__10)
		}

	case CDLangParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)

			var _m = p.Match(CDLangParserT__23)

			localctx.(*GroupContext).zero_or_more = _m
		}
		{
			p.SetState(212)
			p.Match(CDLangParserT__9)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CDLangParserT__20-21))|(1<<(CDLangParserT__22-21))|(1<<(CDLangParserT__23-21))|(1<<(CDLangParserT__24-21))|(1<<(CDLangParserT__25-21))|(1<<(CDLangParserID-21)))) != 0 {
			{
				p.SetState(213)
				p.Instruction()
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(CDLangParserT__10)
		}

	case CDLangParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)

			var _m = p.Match(CDLangParserT__24)

			localctx.(*GroupContext).any_order = _m
		}
		{
			p.SetState(221)
			p.Match(CDLangParserT__9)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CDLangParserID {
			{
				p.SetState(222)
				p.Receive()
			}

			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(228)
			p.Match(CDLangParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_execute
	return p
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) GetName() antlr.Token { return s.name }

func (s *ExecuteContext) SetName(v antlr.Token) { s.name = v }

func (s *ExecuteContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *ExecuteContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExecuteContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (s *ExecuteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitExecute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CDLangParserRULE_execute)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(CDLangParserT__25)
	}
	{
		p.SetState(232)

		var _m = p.Match(CDLangParserID)

		localctx.(*ExecuteContext).name = _m
	}
	{
		p.SetState(233)
		p.Match(CDLangParserT__4)
	}
	{
		p.SetState(234)
		p.Variable()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserT__11 {
		{
			p.SetState(235)
			p.Match(CDLangParserT__11)
		}
		{
			p.SetState(236)
			p.Variable()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
		p.Match(CDLangParserT__5)
	}

	return localctx
}

// IProtobufContext is an interface to support dynamic dispatch.
type IProtobufContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDomain returns the domain token.
	GetDomain() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetDomain sets the domain token.
	SetDomain(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsProtobufContext differentiates from other interfaces.
	IsProtobufContext()
}

type ProtobufContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	domain antlr.Token
	name   antlr.Token
}

func NewEmptyProtobufContext() *ProtobufContext {
	var p = new(ProtobufContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobuf
	return p
}

func (*ProtobufContext) IsProtobufContext() {}

func NewProtobufContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufContext {
	var p = new(ProtobufContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobuf

	return p
}

func (s *ProtobufContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufContext) GetDomain() antlr.Token { return s.domain }

func (s *ProtobufContext) GetName() antlr.Token { return s.name }

func (s *ProtobufContext) SetDomain(v antlr.Token) { s.domain = v }

func (s *ProtobufContext) SetName(v antlr.Token) { s.name = v }

func (s *ProtobufContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *ProtobufContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *ProtobufContext) AllProtobufField() []IProtobufFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem())
	var tst = make([]IProtobufFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProtobufFieldContext)
		}
	}

	return tst
}

func (s *ProtobufContext) ProtobufField(i int) IProtobufFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldContext)
}

func (s *ProtobufContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobuf(s)
	}
}

func (s *ProtobufContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobuf(s)
	}
}

func (s *ProtobufContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobuf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Protobuf() (localctx IProtobufContext) {
	localctx = NewProtobufContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CDLangParserRULE_protobuf)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)

			var _m = p.Match(CDLangParserID)

			localctx.(*ProtobufContext).domain = _m
		}
		{
			p.SetState(245)
			p.Match(CDLangParserT__3)
		}

	}
	{
		p.SetState(248)

		var _m = p.Match(CDLangParserID)

		localctx.(*ProtobufContext).name = _m
	}
	{
		p.SetState(249)
		p.Match(CDLangParserT__9)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserID {
		{
			p.SetState(250)
			p.ProtobufField()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
		p.Match(CDLangParserT__10)
	}

	return localctx
}

// IProtobufFieldContext is an interface to support dynamic dispatch.
type IProtobufFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtobufFieldContext differentiates from other interfaces.
	IsProtobufFieldContext()
}

type ProtobufFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtobufFieldContext() *ProtobufFieldContext {
	var p = new(ProtobufFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobufField
	return p
}

func (*ProtobufFieldContext) IsProtobufFieldContext() {}

func NewProtobufFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufFieldContext {
	var p = new(ProtobufFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobufField

	return p
}

func (s *ProtobufFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufFieldContext) ProtobufFieldSimple() IProtobufFieldSimpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldSimpleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldSimpleContext)
}

func (s *ProtobufFieldContext) ProtobufFieldGroup() IProtobufFieldGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldGroupContext)
}

func (s *ProtobufFieldContext) ProtobufFieldRepeated() IProtobufFieldRepeatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldRepeatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldRepeatedContext)
}

func (s *ProtobufFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobufField(s)
	}
}

func (s *ProtobufFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobufField(s)
	}
}

func (s *ProtobufFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobufField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ProtobufField() (localctx IProtobufFieldContext) {
	localctx = NewProtobufFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CDLangParserRULE_protobufField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.ProtobufFieldSimple()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.ProtobufFieldGroup()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.ProtobufFieldRepeated()
		}

	}

	return localctx
}

// IProtobufFieldSimpleContext is an interface to support dynamic dispatch.
type IProtobufFieldSimpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetVal_number returns the val_number token.
	GetVal_number() antlr.Token

	// GetVal_string returns the val_string token.
	GetVal_string() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetVal_number sets the val_number token.
	SetVal_number(antlr.Token)

	// SetVal_string sets the val_string token.
	SetVal_string(antlr.Token)

	// GetVal_enum returns the val_enum rule contexts.
	GetVal_enum() IEnumerContext

	// GetVal_var returns the val_var rule contexts.
	GetVal_var() IVariableContext

	// GetVal_path returns the val_path rule contexts.
	GetVal_path() IPathContext

	// SetVal_enum sets the val_enum rule contexts.
	SetVal_enum(IEnumerContext)

	// SetVal_var sets the val_var rule contexts.
	SetVal_var(IVariableContext)

	// SetVal_path sets the val_path rule contexts.
	SetVal_path(IPathContext)

	// IsProtobufFieldSimpleContext differentiates from other interfaces.
	IsProtobufFieldSimpleContext()
}

type ProtobufFieldSimpleContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       antlr.Token
	val_number antlr.Token
	val_string antlr.Token
	val_enum   IEnumerContext
	val_var    IVariableContext
	val_path   IPathContext
}

func NewEmptyProtobufFieldSimpleContext() *ProtobufFieldSimpleContext {
	var p = new(ProtobufFieldSimpleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobufFieldSimple
	return p
}

func (*ProtobufFieldSimpleContext) IsProtobufFieldSimpleContext() {}

func NewProtobufFieldSimpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufFieldSimpleContext {
	var p = new(ProtobufFieldSimpleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobufFieldSimple

	return p
}

func (s *ProtobufFieldSimpleContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufFieldSimpleContext) GetName() antlr.Token { return s.name }

func (s *ProtobufFieldSimpleContext) GetVal_number() antlr.Token { return s.val_number }

func (s *ProtobufFieldSimpleContext) GetVal_string() antlr.Token { return s.val_string }

func (s *ProtobufFieldSimpleContext) SetName(v antlr.Token) { s.name = v }

func (s *ProtobufFieldSimpleContext) SetVal_number(v antlr.Token) { s.val_number = v }

func (s *ProtobufFieldSimpleContext) SetVal_string(v antlr.Token) { s.val_string = v }

func (s *ProtobufFieldSimpleContext) GetVal_enum() IEnumerContext { return s.val_enum }

func (s *ProtobufFieldSimpleContext) GetVal_var() IVariableContext { return s.val_var }

func (s *ProtobufFieldSimpleContext) GetVal_path() IPathContext { return s.val_path }

func (s *ProtobufFieldSimpleContext) SetVal_enum(v IEnumerContext) { s.val_enum = v }

func (s *ProtobufFieldSimpleContext) SetVal_var(v IVariableContext) { s.val_var = v }

func (s *ProtobufFieldSimpleContext) SetVal_path(v IPathContext) { s.val_path = v }

func (s *ProtobufFieldSimpleContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ProtobufFieldSimpleContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CDLangParserNUMBER, 0)
}

func (s *ProtobufFieldSimpleContext) STRING() antlr.TerminalNode {
	return s.GetToken(CDLangParserSTRING, 0)
}

func (s *ProtobufFieldSimpleContext) Enumer() IEnumerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumerContext)
}

func (s *ProtobufFieldSimpleContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ProtobufFieldSimpleContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ProtobufFieldSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufFieldSimpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufFieldSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobufFieldSimple(s)
	}
}

func (s *ProtobufFieldSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobufFieldSimple(s)
	}
}

func (s *ProtobufFieldSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobufFieldSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ProtobufFieldSimple() (localctx IProtobufFieldSimpleContext) {
	localctx = NewProtobufFieldSimpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CDLangParserRULE_protobufFieldSimple)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)

		var _m = p.Match(CDLangParserID)

		localctx.(*ProtobufFieldSimpleContext).name = _m
	}
	{
		p.SetState(264)
		p.Match(CDLangParserT__12)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(265)

			var _m = p.Match(CDLangParserNUMBER)

			localctx.(*ProtobufFieldSimpleContext).val_number = _m
		}

	case 2:
		{
			p.SetState(266)

			var _m = p.Match(CDLangParserSTRING)

			localctx.(*ProtobufFieldSimpleContext).val_string = _m
		}

	case 3:
		{
			p.SetState(267)

			var _x = p.Enumer()

			localctx.(*ProtobufFieldSimpleContext).val_enum = _x
		}

	case 4:
		{
			p.SetState(268)

			var _x = p.Variable()

			localctx.(*ProtobufFieldSimpleContext).val_var = _x
		}

	case 5:
		{
			p.SetState(269)

			var _x = p.Path()

			localctx.(*ProtobufFieldSimpleContext).val_path = _x
		}

	}

	return localctx
}

// IEnumerContext is an interface to support dynamic dispatch.
type IEnumerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumerContext differentiates from other interfaces.
	IsEnumerContext()
}

type EnumerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumerContext() *EnumerContext {
	var p = new(EnumerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_enumer
	return p
}

func (*EnumerContext) IsEnumerContext() {}

func NewEnumerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumerContext {
	var p = new(EnumerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_enumer

	return p
}

func (s *EnumerContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumerContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *EnumerContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *EnumerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterEnumer(s)
	}
}

func (s *EnumerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitEnumer(s)
	}
}

func (s *EnumerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitEnumer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Enumer() (localctx IEnumerContext) {
	localctx = NewEnumerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CDLangParserRULE_enumer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDLangParserT__1 {
		{
			p.SetState(272)
			p.Match(CDLangParserT__1)
		}

	}
	{
		p.SetState(275)
		p.Match(CDLangParserID)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserT__1 {
		{
			p.SetState(276)
			p.Match(CDLangParserT__1)
		}
		{
			p.SetState(277)
			p.Match(CDLangParserID)
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProtobufFieldGroupContext is an interface to support dynamic dispatch.
type IProtobufFieldGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetCast returns the cast rule contexts.
	GetCast() IDomainNameContext

	// SetCast sets the cast rule contexts.
	SetCast(IDomainNameContext)

	// IsProtobufFieldGroupContext differentiates from other interfaces.
	IsProtobufFieldGroupContext()
}

type ProtobufFieldGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	cast   IDomainNameContext
}

func NewEmptyProtobufFieldGroupContext() *ProtobufFieldGroupContext {
	var p = new(ProtobufFieldGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobufFieldGroup
	return p
}

func (*ProtobufFieldGroupContext) IsProtobufFieldGroupContext() {}

func NewProtobufFieldGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufFieldGroupContext {
	var p = new(ProtobufFieldGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobufFieldGroup

	return p
}

func (s *ProtobufFieldGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufFieldGroupContext) GetName() antlr.Token { return s.name }

func (s *ProtobufFieldGroupContext) SetName(v antlr.Token) { s.name = v }

func (s *ProtobufFieldGroupContext) GetCast() IDomainNameContext { return s.cast }

func (s *ProtobufFieldGroupContext) SetCast(v IDomainNameContext) { s.cast = v }

func (s *ProtobufFieldGroupContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ProtobufFieldGroupContext) AllProtobufField() []IProtobufFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem())
	var tst = make([]IProtobufFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProtobufFieldContext)
		}
	}

	return tst
}

func (s *ProtobufFieldGroupContext) ProtobufField(i int) IProtobufFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldContext)
}

func (s *ProtobufFieldGroupContext) DomainName() IDomainNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainNameContext)
}

func (s *ProtobufFieldGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufFieldGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufFieldGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobufFieldGroup(s)
	}
}

func (s *ProtobufFieldGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobufFieldGroup(s)
	}
}

func (s *ProtobufFieldGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobufFieldGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ProtobufFieldGroup() (localctx IProtobufFieldGroupContext) {
	localctx = NewProtobufFieldGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CDLangParserRULE_protobufFieldGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)

		var _m = p.Match(CDLangParserID)

		localctx.(*ProtobufFieldGroupContext).name = _m
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDLangParserT__26 {
		{
			p.SetState(284)
			p.Match(CDLangParserT__26)
		}
		{
			p.SetState(285)

			var _x = p.DomainName()

			localctx.(*ProtobufFieldGroupContext).cast = _x
		}
		{
			p.SetState(286)
			p.Match(CDLangParserT__27)
		}

	}
	{
		p.SetState(290)
		p.Match(CDLangParserT__9)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserID {
		{
			p.SetState(291)
			p.ProtobufField()
		}

		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(297)
		p.Match(CDLangParserT__10)
	}

	return localctx
}

// IProtobufFieldRepeatedContext is an interface to support dynamic dispatch.
type IProtobufFieldRepeatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsProtobufFieldRepeatedContext differentiates from other interfaces.
	IsProtobufFieldRepeatedContext()
}

type ProtobufFieldRepeatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyProtobufFieldRepeatedContext() *ProtobufFieldRepeatedContext {
	var p = new(ProtobufFieldRepeatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobufFieldRepeated
	return p
}

func (*ProtobufFieldRepeatedContext) IsProtobufFieldRepeatedContext() {}

func NewProtobufFieldRepeatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufFieldRepeatedContext {
	var p = new(ProtobufFieldRepeatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobufFieldRepeated

	return p
}

func (s *ProtobufFieldRepeatedContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufFieldRepeatedContext) GetName() antlr.Token { return s.name }

func (s *ProtobufFieldRepeatedContext) SetName(v antlr.Token) { s.name = v }

func (s *ProtobufFieldRepeatedContext) ID() antlr.TerminalNode {
	return s.GetToken(CDLangParserID, 0)
}

func (s *ProtobufFieldRepeatedContext) AllProtobufFieldRepeatedRow() []IProtobufFieldRepeatedRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProtobufFieldRepeatedRowContext)(nil)).Elem())
	var tst = make([]IProtobufFieldRepeatedRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProtobufFieldRepeatedRowContext)
		}
	}

	return tst
}

func (s *ProtobufFieldRepeatedContext) ProtobufFieldRepeatedRow(i int) IProtobufFieldRepeatedRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldRepeatedRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldRepeatedRowContext)
}

func (s *ProtobufFieldRepeatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufFieldRepeatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufFieldRepeatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobufFieldRepeated(s)
	}
}

func (s *ProtobufFieldRepeatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobufFieldRepeated(s)
	}
}

func (s *ProtobufFieldRepeatedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobufFieldRepeated(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ProtobufFieldRepeated() (localctx IProtobufFieldRepeatedContext) {
	localctx = NewProtobufFieldRepeatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CDLangParserRULE_protobufFieldRepeated)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)

		var _m = p.Match(CDLangParserID)

		localctx.(*ProtobufFieldRepeatedContext).name = _m
	}
	{
		p.SetState(300)
		p.Match(CDLangParserT__28)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CDLangParserT__9 {
		{
			p.SetState(301)
			p.ProtobufFieldRepeatedRow()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(306)
		p.Match(CDLangParserT__29)
	}

	return localctx
}

// IProtobufFieldRepeatedRowContext is an interface to support dynamic dispatch.
type IProtobufFieldRepeatedRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetZero_or_more returns the zero_or_more token.
	GetZero_or_more() antlr.Token

	// GetOne_or_more returns the one_or_more token.
	GetOne_or_more() antlr.Token

	// GetZero_or_one returns the zero_or_one token.
	GetZero_or_one() antlr.Token

	// SetZero_or_more sets the zero_or_more token.
	SetZero_or_more(antlr.Token)

	// SetOne_or_more sets the one_or_more token.
	SetOne_or_more(antlr.Token)

	// SetZero_or_one sets the zero_or_one token.
	SetZero_or_one(antlr.Token)

	// IsProtobufFieldRepeatedRowContext differentiates from other interfaces.
	IsProtobufFieldRepeatedRowContext()
}

type ProtobufFieldRepeatedRowContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	zero_or_more antlr.Token
	one_or_more  antlr.Token
	zero_or_one  antlr.Token
}

func NewEmptyProtobufFieldRepeatedRowContext() *ProtobufFieldRepeatedRowContext {
	var p = new(ProtobufFieldRepeatedRowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_protobufFieldRepeatedRow
	return p
}

func (*ProtobufFieldRepeatedRowContext) IsProtobufFieldRepeatedRowContext() {}

func NewProtobufFieldRepeatedRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobufFieldRepeatedRowContext {
	var p = new(ProtobufFieldRepeatedRowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_protobufFieldRepeatedRow

	return p
}

func (s *ProtobufFieldRepeatedRowContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobufFieldRepeatedRowContext) GetZero_or_more() antlr.Token { return s.zero_or_more }

func (s *ProtobufFieldRepeatedRowContext) GetOne_or_more() antlr.Token { return s.one_or_more }

func (s *ProtobufFieldRepeatedRowContext) GetZero_or_one() antlr.Token { return s.zero_or_one }

func (s *ProtobufFieldRepeatedRowContext) SetZero_or_more(v antlr.Token) { s.zero_or_more = v }

func (s *ProtobufFieldRepeatedRowContext) SetOne_or_more(v antlr.Token) { s.one_or_more = v }

func (s *ProtobufFieldRepeatedRowContext) SetZero_or_one(v antlr.Token) { s.zero_or_one = v }

func (s *ProtobufFieldRepeatedRowContext) AllProtobufField() []IProtobufFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem())
	var tst = make([]IProtobufFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProtobufFieldContext)
		}
	}

	return tst
}

func (s *ProtobufFieldRepeatedRowContext) ProtobufField(i int) IProtobufFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtobufFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProtobufFieldContext)
}

func (s *ProtobufFieldRepeatedRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobufFieldRepeatedRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobufFieldRepeatedRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterProtobufFieldRepeatedRow(s)
	}
}

func (s *ProtobufFieldRepeatedRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitProtobufFieldRepeatedRow(s)
	}
}

func (s *ProtobufFieldRepeatedRowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitProtobufFieldRepeatedRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) ProtobufFieldRepeatedRow() (localctx IProtobufFieldRepeatedRowContext) {
	localctx = NewProtobufFieldRepeatedRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CDLangParserRULE_protobufFieldRepeatedRow)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(CDLangParserT__9)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CDLangParserID {
		{
			p.SetState(309)
			p.ProtobufField()
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(314)
		p.Match(CDLangParserT__10)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CDLangParserT__30:
		{
			p.SetState(315)

			var _m = p.Match(CDLangParserT__30)

			localctx.(*ProtobufFieldRepeatedRowContext).zero_or_more = _m
		}

	case CDLangParserT__31:
		{
			p.SetState(316)

			var _m = p.Match(CDLangParserT__31)

			localctx.(*ProtobufFieldRepeatedRowContext).one_or_more = _m
		}

	case CDLangParserT__32:
		{
			p.SetState(317)

			var _m = p.Match(CDLangParserT__32)

			localctx.(*ProtobufFieldRepeatedRowContext).zero_or_one = _m
		}

	case CDLangParserT__9, CDLangParserT__11, CDLangParserT__29:

	default:
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CDLangParserT__11 {
		{
			p.SetState(320)
			p.Match(CDLangParserT__11)
		}

	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPathElement() []IPathElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPathElementContext)(nil)).Elem())
	var tst = make([]IPathElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPathElementContext)
		}
	}

	return tst
}

func (s *PathContext) PathElement(i int) IPathElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPathElementContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CDLangParserRULE_path)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CDLangParserT__33 {
		{
			p.SetState(323)
			p.Match(CDLangParserT__33)
		}
		{
			p.SetState(324)
			p.PathElement()
		}

		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPathElementContext is an interface to support dynamic dispatch.
type IPathElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetA returns the a token.
	GetA() antlr.Token

	// GetE returns the e token.
	GetE() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetA sets the a token.
	SetA(antlr.Token)

	// SetE sets the e token.
	SetE(antlr.Token)

	// GetRd_var returns the rd_var rule contexts.
	GetRd_var() IVariableContext

	// GetWr_var returns the wr_var rule contexts.
	GetWr_var() IVariableContext

	// GetParam returns the param rule contexts.
	GetParam() IVariableContext

	// SetRd_var sets the rd_var rule contexts.
	SetRd_var(IVariableContext)

	// SetWr_var sets the wr_var rule contexts.
	SetWr_var(IVariableContext)

	// SetParam sets the param rule contexts.
	SetParam(IVariableContext)

	// IsPathElementContext differentiates from other interfaces.
	IsPathElementContext()
}

type PathElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	key    antlr.Token
	a      antlr.Token
	rd_var IVariableContext
	e      antlr.Token
	wr_var IVariableContext
	param  IVariableContext
}

func NewEmptyPathElementContext() *PathElementContext {
	var p = new(PathElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CDLangParserRULE_pathElement
	return p
}

func (*PathElementContext) IsPathElementContext() {}

func NewPathElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathElementContext {
	var p = new(PathElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CDLangParserRULE_pathElement

	return p
}

func (s *PathElementContext) GetParser() antlr.Parser { return s.parser }

func (s *PathElementContext) GetName() antlr.Token { return s.name }

func (s *PathElementContext) GetKey() antlr.Token { return s.key }

func (s *PathElementContext) GetA() antlr.Token { return s.a }

func (s *PathElementContext) GetE() antlr.Token { return s.e }

func (s *PathElementContext) SetName(v antlr.Token) { s.name = v }

func (s *PathElementContext) SetKey(v antlr.Token) { s.key = v }

func (s *PathElementContext) SetA(v antlr.Token) { s.a = v }

func (s *PathElementContext) SetE(v antlr.Token) { s.e = v }

func (s *PathElementContext) GetRd_var() IVariableContext { return s.rd_var }

func (s *PathElementContext) GetWr_var() IVariableContext { return s.wr_var }

func (s *PathElementContext) GetParam() IVariableContext { return s.param }

func (s *PathElementContext) SetRd_var(v IVariableContext) { s.rd_var = v }

func (s *PathElementContext) SetWr_var(v IVariableContext) { s.wr_var = v }

func (s *PathElementContext) SetParam(v IVariableContext) { s.param = v }

func (s *PathElementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(CDLangParserID)
}

func (s *PathElementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(CDLangParserID, i)
}

func (s *PathElementContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PathElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.EnterPathElement(s)
	}
}

func (s *PathElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CDLangListener); ok {
		listenerT.ExitPathElement(s)
	}
}

func (s *PathElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CDLangVisitor:
		return t.VisitPathElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CDLangParser) PathElement() (localctx IPathElementContext) {
	localctx = NewPathElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CDLangParserRULE_pathElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)

			var _m = p.Match(CDLangParserID)

			localctx.(*PathElementContext).name = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)

			var _m = p.Match(CDLangParserID)

			localctx.(*PathElementContext).name = _m
		}
		{
			p.SetState(332)
			p.Match(CDLangParserT__28)
		}
		{
			p.SetState(333)

			var _m = p.Match(CDLangParserID)

			localctx.(*PathElementContext).key = _m
		}
		{
			p.SetState(334)
			p.Match(CDLangParserT__34)
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CDLangParserT__30:
			{
				p.SetState(335)

				var _m = p.Match(CDLangParserT__30)

				localctx.(*PathElementContext).a = _m
			}

		case CDLangParserT__13, CDLangParserT__14, CDLangParserT__15:
			{
				p.SetState(336)

				var _x = p.Variable()

				localctx.(*PathElementContext).rd_var = _x
			}

		case CDLangParserID:
			{
				p.SetState(337)

				var _m = p.Match(CDLangParserID)

				localctx.(*PathElementContext).e = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(340)
			p.Match(CDLangParserT__29)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)

			var _m = p.Match(CDLangParserID)

			localctx.(*PathElementContext).name = _m
		}
		{
			p.SetState(342)
			p.Match(CDLangParserT__28)
		}
		{
			p.SetState(343)

			var _x = p.Variable()

			localctx.(*PathElementContext).wr_var = _x
		}
		{
			p.SetState(344)
			p.Match(CDLangParserT__0)
		}
		{
			p.SetState(345)

			var _m = p.Match(CDLangParserID)

			localctx.(*PathElementContext).key = _m
		}
		{
			p.SetState(346)
			p.Match(CDLangParserT__29)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(348)

			var _x = p.Variable()

			localctx.(*PathElementContext).param = _x
		}

	}

	return localctx
}
