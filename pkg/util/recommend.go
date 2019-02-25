package util

import "math"

//1.平台尽可能多地提供细致的标签
//2.开放用户向物品关联自定义标签权利
//3.需要考虑水军如何过滤
func ItemCF(res_item map[int][]int,base_item map[int]int) map[int]float64{
	LikeScore := make(map[int]float64)
	//计算电影与item关联电影的相似度
	LenBase := len(base_item)
	for key,val := range res_item {
		LenRes := len(val)
		for _,vv := range val{
			if base_item[vv] == 1 {
				LikeScore[key] += 1
			}
		}
		LikeScore[key] = LikeScore[key]/math.Sqrt(float64(LenBase*LenRes))
	}
	return LikeScore
}
