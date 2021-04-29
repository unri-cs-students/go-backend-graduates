package content_type

import (
	"errors"
	"strings"
)

type ContentType int

const (
	Kontribusi ContentType = iota
	Prestasi
	Karya
	Tips
	Keanggotaan
	Funfact
)

func (s ContentType) String() string {
	return [...]string{"KONTRIBUSI", "PRESTASI", "KARYA", "TIPS_SUKSES", "KEANGGOTAAN", "FUNFACT"}[s]
}

func GetEnum(any string) (string, error) {
	TYPES := [...]string{"KONTRIBUSI", "PRESTASI", "KARYA", "TIPS_SUKSES", "KEANGGOTAAN", "FUNFACT"}

	anyConverted := strings.ReplaceAll(any, " ", "_")
	for i, x := range TYPES {
		if strings.EqualFold(x, anyConverted) {
			return TYPES[i], nil
		}
	}
	return "", errors.New("Unknown type")
}

