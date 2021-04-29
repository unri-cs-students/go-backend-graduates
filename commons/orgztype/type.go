package orgztype

import (
	"errors"
	"strings"
)

type OrgzType int

const (
	Kongres_KM_ITB OrgzType = iota
	MWA_WM_ITB
	Kabinet_KM_ITB
	HMJ
	BSO
	Komunitas
	UKM_Olahraga_Dan_Kesehatan
	UKM_Media
	UKM_Seni_Budaya
	UKM_Agama_Pendidikan_Dan_Kajian
	TPB
)

func (s OrgzType) String() string {
	return [...]string{"KONGRES_KM_ITB", "MWA_WM_ITB", "KABINET_KM_ITB", "HMJ", "BSO", "KOMUNITAS", "UKM_OLAHRAGA_DAN_KESEHATAN", "UKM_MEDIA", "UKM_SENI_BUDAYA", "UKM_AGAMA_PENDIDIKAN_DAN_KAJIAN", "TPB"}[s]
}

func GetEnum(any string) (string, error) {
	TYPES := [...]string{"KONGRES_KM_ITB", "MWA_WM_ITB", "KABINET_KM_ITB", "HMJ", "BSO", "KOMUNITAS", "UKM_OLAHRAGA_DAN_KESEHATAN", "UKM_MEDIA", "UKM_SENI_BUDAYA", "UKM_AGAMA_PENDIDIKAN_DAN_KAJIAN", "TPB"}

	anyConverted := strings.ReplaceAll(any, " ", "_")
	for i, x := range TYPES {
		if strings.EqualFold(x, anyConverted) {
			return TYPES[i], nil
		}
	}
	return "", errors.New("Unknown type")
}

