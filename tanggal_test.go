package tanggal

import (
	"reflect"
	"testing"
	"time"
)

func TestTanggal_Format(t *testing.T) {
	type fields struct {
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
	type args struct {
		separator string
		formats   []Format
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test Format Tertentu",
			fields: fields{
				Hari:               10,
				Bulan:              10,
				Lokasi:             "Jakarta",
				LokasiDenganKoma:   "Jakarta,",
				NamaBulan:          "Oktober",
				NamaHari:           "Senin",
				NamaHariDenganKoma: "Senin,",
				Pukul:              "22:22",
				PukulDenganDetik:   "22:22:22",
				Tahun:              1010,
				Timezone:           "WIB",
			},
			args: args{
				separator: " ",
				formats:   []Format{Hari, NamaBulan, Tahun},
			},
			want: "10 Oktober 1010",
		},
		{
			name: "Test Semua Format",
			fields: fields{
				Hari:               10,
				Bulan:              10,
				Lokasi:             "Jakarta",
				LokasiDenganKoma:   "Jakarta,",
				NamaBulan:          "Oktober",
				NamaHari:           "Senin",
				NamaHariDenganKoma: "Senin,",
				Pukul:              "22:22",
				PukulDenganDetik:   "22:22:22",
				Tahun:              1010,
				Timezone:           "WIB",
			},
			args: args{
				separator: " ",
				formats: []Format{
					Lokasi, LokasiDenganKoma,
					Hari, NamaHari, NamaHariDenganKoma,
					Bulan, NamaBulan, Tahun,
					Pukul, PukulDenganDetik, ZonaWaktu,
				},
			},
			want: "Jakarta Jakarta, 10 Senin Senin, 10 Oktober 1010 22:22 22:22:22 WIB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Tanggal{
				Hari:               tt.fields.Hari,
				NamaHari:           tt.fields.NamaHari,
				NamaHariDenganKoma: tt.fields.NamaHariDenganKoma,
				Bulan:              tt.fields.Bulan,
				NamaBulan:          tt.fields.NamaBulan,
				Tahun:              tt.fields.Tahun,
				Pukul:              tt.fields.Pukul,
				PukulDenganDetik:   tt.fields.PukulDenganDetik,
				Lokasi:             tt.fields.Lokasi,
				LokasiDenganKoma:   tt.fields.LokasiDenganKoma,
				Timezone:           tt.fields.Timezone,
			}
			if got := w.Format(tt.args.separator, tt.args.formats); got != tt.want {
				t.Errorf("Tanggal.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPapar(t *testing.T) {
	ttw, _ := time.Parse(time.RFC3339, "2009-11-10T23:00:00Z")
	type args struct {
		t      time.Time
		lokasi string
		tz     Timezone
	}
	tests := []struct {
		name    string
		args    args
		want    Tanggal
		wantErr bool
	}{
		{
			name: "Test Papar WIB",
			args: args{
				t:      ttw,
				lokasi: "Jakarta",
				tz:     WIB,
			},
			want: Tanggal{
				Hari:               11,
				NamaHari:           "Selasa",
				Bulan:              11,
				NamaBulan:          "November",
				Lokasi:             "Jakarta",
				LokasiDenganKoma:   "Jakarta,",
				NamaHariDenganKoma: "Selasa,",
				Pukul:              "06:00",
				PukulDenganDetik:   "06:00:00",
				Tahun:              2009,
				Timezone:           "WIB",
			},
			wantErr: false,
		},
		{
			name: "Test Papar WITA",
			args: args{
				t:      ttw,
				lokasi: "Mataram",
				tz:     WITA,
			},
			want: Tanggal{
				Hari:               11,
				NamaHari:           "Selasa",
				Bulan:              11,
				NamaBulan:          "November",
				Lokasi:             "Mataram",
				LokasiDenganKoma:   "Mataram,",
				NamaHariDenganKoma: "Selasa,",
				Pukul:              "07:00",
				PukulDenganDetik:   "07:00:00",
				Tahun:              2009,
				Timezone:           "WITA",
			},
			wantErr: false,
		},
		{
			name: "Test Papar WIT",
			args: args{
				t:      ttw,
				lokasi: "Papua",
				tz:     WIT,
			},
			want: Tanggal{
				Hari:               11,
				NamaHari:           "Selasa",
				Bulan:              11,
				NamaBulan:          "November",
				Lokasi:             "Papua",
				LokasiDenganKoma:   "Papua,",
				NamaHariDenganKoma: "Selasa,",
				Pukul:              "08:00",
				PukulDenganDetik:   "08:00:00",
				Tahun:              2009,
				Timezone:           "WIT",
			},
			wantErr: false,
		},
		{
			name: "Test Papar Gagal",
			args: args{
				t:      time.Now(),
				lokasi: "",
				tz:     Timezone(""),
			},
			want:    Tanggal{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Papar(tt.args.t, tt.args.lokasi, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("Papar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Papar() = %v, want %v", got, tt.want)
			}
		})
	}
}
