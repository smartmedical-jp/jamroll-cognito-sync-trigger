package setting

var (
	region string
)

func GetRegion() string {
	return region
}

func setRegion(r string) {
	region = r
}
