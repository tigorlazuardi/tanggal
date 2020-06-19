# Tanggal (GO)

Tanggal (GO) merupakan Golang package untuk memaparkan golang time menjadi format tanggal bahasa Indonesia.

Install

```shell
go get github.com/TigorLazuardi/tanggal
```

Usage Example:

```go
// tgl merupakan struct dan dapat digunakan untuk mengakses berbagai bagian dari tanggal tersetut.
// lokasi dapat dikosongkan, dan kalau dikosongkan, format lokasi akan diabaikan sepenuhnya.
// timezone mempengaruhi hasil waktu yang dipapar. WITA akan maju 1 jam didepan WIB,
// dan demikian pula WIT dengan WITA. timezone NONE akan mengabaikan formatting.
tgl, err := tanggal.Papar(time.Now(), "Jakarta", tanggal.WIB)
if err != nil {
    log.Fatal(err)
}

// contoh hasil struct. Note: Struct tidak memiliki json tags
aa, _ := json.MarshalIndent(tgl, "", "  ")
fmt.Println(string(aa))
// {
//     "Hari": 16,
//     "NamaHari": "Selasa",
//     "NamaHariDenganKoma": "Selasa,",
//     "Bulan": 6,
//     "NamaBulan": "Juni",
//     "Tahun": 2020,
//     "Pukul": "21:19",
//     "PukulDenganDetik": "21:19:16",
//     "Lokasi": "Jakarta",
//     "LokasiDenganKoma": "Jakarta,",
//     "Timezone": "WIB"
// }

// Dibawah ini cara custom formatting string
// Formatting akan mengikuti urutan dari slice, dan separator untuk customize pemisah antar elemen
format := []tanggal.Format{
    tanggal.LokasiDenganKoma, 
    tanggal.Hari, 
    tanggal.NamaBulan, 
    tanggal.Tahun, 
    tanggal.PukulDenganDetik, 
    tanggal.ZonaWaktu,
}
ss := tgl.Format(" ", format)
fmt.Println(ss)
// Jakarta, 16 Juni 2020 21:19:16 WIB
```

```
Available Formats:

tanggal.Hari               
tanggal.NamaHari           
tanggal.NamaHariDenganKoma 
tanggal.Minggu             
tanggal.NamaMinggu         
tanggal.Bulan              
tanggal.NamaBulan          
tanggal.Tahun              
tanggal.Pukul              
tanggal.PukulDenganDetik   
tanggal.Lokasi             
tanggal.LokasiDenganKoma   
tanggal.ZonaWaktu          
```

```
Available Timezones:
tanggal.WIB  // UTC+7
tanggal.WITA // UTC+8
tanggal.WIT  // UTC+9
tanggal.NONE // UTC
```
