package main

import (
	"math/rand"
	"time"
	"fmt"
)

func fillColors(colorMap *map[int][]string) {
	color := *colorMap
	color[1] = []string{"White", "#FFFFFF"}
	color[2] = []string{"Black", "#000000"}
	color[3] = []string{"Red", "#FF0000"}
	color[4] = []string{"Lime", "#00FF00"}
	color[5] = []string{"Blue", "#0000FF"}
	color[6] = []string{"Yellow", "#FFFF00"}
	color[7] = []string{"Cyan", "#00FFFF"}
	color[8] = []string{"Magenta", "#FF00FF"}
	color[9] = []string{"Silver", "#C0C0C0"}
	color[10] = []string{"Gray", "#808080"}
	color[11] = []string{"Maroon", "#800000"}
	color[12] = []string{"Olive", "#808000"}
	color[13] = []string{"Green", "#008000"}
	color[14] = []string{"Purple", "#800080"}
	color[15] = []string{"Teal", "#008080"}
	color[16] = []string{"Navy", "#000080"}
	color[17] = []string{"Orange", "#FFA500"}
	color[18] = []string{"Gold", "#FFD700"}
	color[19] = []string{"Pink", "#FFC0CB"}
	color[20] = []string{"Brown", "#A52A2A"}
	color[21] = []string{"Beige", "#F5F5DC"}
	color[22] = []string{"Lavender", "#E6E6FA"}
	color[23] = []string{"Turquoise", "#40E0D0"}
	color[24] = []string{"Indigo", "#4B0082"}
	color[25] = []string{"Violet", "#EE82EE"}
	color[26] = []string{"Khaki", "#F0E68C"}
	color[27] = []string{"Coral", "#FF7F50"}
	color[28] = []string{"Salmon", "#FA8072"}
	color[29] = []string{"Chocolate", "#D2691E"}
	color[30] = []string{"Crimson", "#DC143C"}
	color[31] = []string{"Mint", "#98FF98"}
	color[32] = []string{"Emerald", "#50C878"}
	color[33] = []string{"Ruby", "#E0115F"}
	color[34] = []string{"Ivory", "#FFFFF0"}
	color[35] = []string{"Teal", "#008080"}
	color[36] = []string{"SkyBlue", "#87CEEB"}
	color[37] = []string{"RoyalBlue", "#4169E1"}
	color[38] = []string{"SlateBlue", "#6A5ACD"}
	color[39] = []string{"SeaGreen", "#2E8B57"}
	color[40] = []string{"ForestGreen", "#228B22"}
	color[41] = []string{"LawnGreen", "#7CFC00"}
	color[42] = []string{"SpringGreen", "#00FF7F"}
	color[43] = []string{"Aqua", "#00FFFF"}
	color[44] = []string{"MidnightBlue", "#191970"}
	color[45] = []string{"CornflowerBlue", "#6495ED"}
	color[46] = []string{"DarkOrchid", "#9932CC"}
	color[47] = []string{"FireBrick", "#B22222"}
	color[48] = []string{"Tomato", "#FF6347"}
	color[49] = []string{"SlateGray", "#708090"}
}

func GenerateRandomColor() []string {
	rand.Seed(time.Now().UnixNano())
	colorMap := make(map[int][]string)
	fillColors(&colorMap)
	op:=colorMap[rand.Intn(50)]
	fmt.Println("the color is ",op)
	return op
}
