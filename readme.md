# Tanggal (GO)

Tanggal (GO) merupakan Golang package untuk memaparkan golang time menjadi format tanggal bahasa Indonesia.

Install

```shell
go get github.com/TigorLazuardi/tanggal
```

Usage Example:

```go
tgl, err := tanggal.Papar(time.Now(), "Jakarta", tanggal.WIB)
if err != nil {
    log.Fatal(err)
}
aa, _ := json.MarshalIndent(tgl, "", "  ")
fmt.Println(string(aa))

// Formatting akan mengikuti urutan dari slice, dan separator untuk customize pemisah antar elemen
format := []tanggal.Format{
    tanggal.LokasiDenganKoma, tanggal.Hari, tanggal.NamaBulan, tanggal.Tahun, tanggal.PukulDenganDetik, tanggal.ZonaWaktu,
}

ss := tgl.Format(" ", format)
fmt.Println(ss)
```

```
Available Formats:

Hari               
NamaHari           
NamaHariDenganKoma 
Minggu             
NamaMinggu         
Bulan              
NamaBulan          
Tahun              
Pukul              
PukulDenganDetik   
Lokasi             
LokasiDenganKoma   
ZonaWaktu          
```
