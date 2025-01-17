package tester

import "github.com/myyrakle/swag/pkg/array"

type ShopResponse struct {
	ShopID   string `json:"shopID" validate:"required" example:"63ec5ccff19921e0414902b7"` // 샵 ID
	ShopName string `json:"shopName" validate:"required" example:"스토어 쟈넬"`
}

type ShopsResponseV1 []ShopResponse

type ShopsResponseV2 array.Array[ShopResponse]

type ShopsResponseV3 struct {
	array.Array[ShopResponse]
}

type ListShopsResponse struct {
	Shops   ShopsResponseV1 `json:"shopsV1"`
	ShopsV2 ShopsResponseV2 `json:"shopsV2"`
	ShopsV3 ShopsResponseV3 `json:"shopsV3"`
}

// ListShops
// @Tags Shops
// @Summary 샵 목록 조회 ✅
// @Description 샵 목록을 조회하는 API 입니다.
// @Produce json
// @Success  200 {object} ListShopsResponse "샵 목록 응답 값"
// @Security BearerAuth
// @Router /shops [get]
func ListShops() {
	foo := array.Array[ShopResponse]{}
	if foo.IsNotEmpty() {
		// do something
	}
}

func main() {

}
