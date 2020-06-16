package tanggal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Format string
type Timezone string

const (
	Hari               Format = "hari"
	NamaHari           Format = "namaHari"
	NamaHariDenganKoma Format = "namaHariDenganKoma"
	Minggu             Format = "minggu"
	NamaMinggu         Format = "namaMinggu"
	Bulan              Format = "bulan"
	NamaBulan          Format = "namaBulan"
	Tahun              Format = "tahun"
	Pukul              Format = "pukul"
	PukulDenganDetik   Format = "pukulDenganDetik"
	Lokasi             Format = "lokasi"
	LokasiDenganKoma   Format = "lokasiDenganKoma"
	ZonaWaktu          Format = "zonaWaktu"
)

const WIB, WITA, WIT Timezone = "WIB", "WITA", "WIT"

type Tanggal struct {
	Hari               uint8
	NamaHari           string
	NamaHariDenganKoma string
	Bulan              uint8
	NamaBulan          string
	Tahun              uint64
	Pukul              string
	PukulDenganDetik   string
	Lokasi             string
	LokasiDenganKoma   string
	Timezone           string
}

func (t Tanggal) Format(separator string, formats []Format) string {
	var f []string
	for _, v := range formats {
		switch v {
		case Hari:
			f = append(f, strconv.FormatUint(uint64(t.Hari), 10))
		case NamaHari:
			f = append(f, t.NamaHari)
		case NamaHariDenganKoma:
			f = append(f, t.NamaHariDenganKoma)
		case Bulan:
			f = append(f, strconv.FormatUint(uint64(t.Bulan), 10))
		case NamaBulan:
			f = append(f, t.NamaBulan)
		case Tahun:
			f = append(f, strconv.FormatUint(t.Tahun, 10))
		case Pukul:
			f = append(f, t.Pukul)
		case PukulDenganDetik:
			f = append(f, t.PukulDenganDetik)
		case Lokasi:
			if t.Lokasi != "" {
				f = append(f, t.Lokasi)
			}
		case LokasiDenganKoma:
			if t.Lokasi != "" {
				f = append(f, t.LokasiDenganKoma)
			}
		case ZonaWaktu:
			f = append(f, t.Timezone)
		}
	}
	return strings.Join(f, separator)
}

// Papar merubah dari Time menjadi struct yang diisi dengan memaparkan Time yang dioper.
//
// Kalau lokasi diisi dengan string kosong, ketika memapar jadi string, akan mengabaikan format Lokasi dan LokasiDenganKoma.
func Papar(t time.Time, lokasi string, tz Timezone) (Tanggal, error) {
	add, err := parseTimeZone(tz)
	if err != nil {
		return Tanggal{}, err
	}
	now := t.UTC().Add(time.Hour * add).String()
	year, _ := strconv.ParseUint(now[0:4], 10, 0)
	month, _ := strconv.ParseUint(now[5:7], 10, 0)
	day, _ := strconv.ParseUint(now[8:10], 10, 0)
	clock := now[11:19]
	shortClock := now[11:16]
	return Tanggal{
		Hari:               uint8(day),
		NamaHari:           cariNamaHari(t),
		Bulan:              uint8(month),
		NamaBulan:          cariNamaBulan(uint8(month)),
		Tahun:              year,
		PukulDenganDetik:   clock,
		Pukul:              shortClock,
		Lokasi:             lokasi,
		LokasiDenganKoma:   lokasi + ",",
		NamaHariDenganKoma: cariNamaHari(t) + ",",
		Timezone:           string(tz),
	}, nil
}

func cariNamaHari(t time.Time) string {
	switch int(t.Weekday()) {
	case 0:
		return "Minggu"
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jumat"
	case 6:
		return "Sabtu"
	}
	return ""
}

func cariNamaBulan(m uint8) string {
	switch m {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	}
	return ""
}

func parseTimeZone(tz Timezone) (time.Duration, error) {
	switch tz {
	case WIB:
		return time.Duration(7), nil
	case WITA:
		return time.Duration(8), nil
	case WIT:
		return time.Duration(9), nil
	default:
		return time.Duration(0), errors.New(fmt.Sprintf("Failed to parse timezone. %s is not supported", tz))
	}
}
