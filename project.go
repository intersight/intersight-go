package intersight

type ProjectMode string
type ProjectRegion string

type Project struct {
	ID            string
	Name          string
	Mode          ProjectMode
	PrimaryRegion ProjectRegion
}
