package do

type TPackage struct {
	PackageId   int    `xorm:"not null pk autoincr INT"`     // 专业ID
	PackageName string `xorm:"comment('专业名称') VARCHAR(255)"` // 专业名称
}
