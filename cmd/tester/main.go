package tester

type ShopResponse struct {
	ShopID   string `json:"shopID" validate:"required" example:"63ec5ccff19921e0414902b7"` // 샵 ID
	ShopName string `json:"shopName" validate:"required" example:"스토어 쟈넬"`
}

type ShopsResponse []ShopResponse

type ListShopsResponse struct {
	Shops ShopsResponse `json:"shops"`
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
}

func main() {

}
