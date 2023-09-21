package pipeline

type PackageId uint16

type Package interface {
	Id() PackageId
}
