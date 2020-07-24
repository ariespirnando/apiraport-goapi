package model

type Raport struct {
	Namakelas      string `json:"namakelas"`
	Kodejuruan     string `json:"kodejurusan"`
	Niksiswa       string `json:"niksiswa"`
	Kodekelas      string `json:"kodekelas"`
	Guidraport     string `json:"guidraport"`
	Statuskenaikan string `json:"statuskenaikan"`
	Statusspp      string `json:"statusspp"`
	Semester       string `json:"semester"`
	Tahunajaran    string `json:"tahunajaran"`
}

type Kelompok struct {
	Kodekelompok string      `json:"kodekelompok"`
	Namakelompok string      `json:"namakelompok"`
	Pelajaran    []Pelajaran `json:"matapelajaran,omitempty"`
}
type Pelajaran struct {
	Kodepelajaran string  `json:"kodepelajaran"`
	Namapelajaran string  `json:"namapelajaran"`
	Kkm           string  `json:"kkm"`
	Pengetahuan  string `json:"pengetahuan"`
	Keterampilan string `json:"keterampilan"`
	Sikap        string `json:"sikap"`
	Keterangan   string `json:"keterangan"` 
}

 
