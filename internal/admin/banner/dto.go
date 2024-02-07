package banner

type DataBannerDTO struct {
	Categories []GetBannerDTO `json:"banners"`
	Count      int            `json:"count"`
}

type GetBannerDTO struct {
	Id    int    `json:"Id"`
	Image string `json:"imagePath"`
}

type AddBannerDTO struct {
	Image string `json:"imagePath"`
}
