package controller

import (
	"log" 
	"net/http"

	"github.com/apiraport/config"
	"github.com/apiraport/model"
	"github.com/gin-gonic/gin"
)

func Getraport(c *gin.Context) {
	niksiswa := c.MustGet("jwt_niksiswa") 
	var data []model.Raport
	var count int

	config.DB.Table("tbl_raport").
		Select("tbl_kelas.namakelas, tbl_kelas.kodekelas, tbl_kelas.kodejuruan, tbl_semester.tahunajaran, tbl_semester.semester, tbl_raport.statusspp, tbl_raport.statuskenaikan, tbl_raport.niksiswa, tbl_raport.guidraport").
		Joins("INNER JOIN tbl_kelas ON tbl_raport.guidkelas = tbl_kelas.guidkelas").
		Joins("INNER JOIN tbl_semester ON tbl_semester.guidsemester = tbl_raport.guidsemester").
		Where("tbl_raport.niksiswa = ?", niksiswa).
		Scan(&data).
		Count(&count)

	if len(data) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"error_code": "000000",
			"status":     "Data ditemukan",
			"data":       data,
			"total":      count,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error_code": "000001",
			"status":     "Data tidak ditemukan",
		})
	}

}

func Getraportdetail(c *gin.Context) {

	var json model.Requestraportdet

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "000002",
			"status":     "Bad Request",
		})
	} else {
		var kelompok model.Kelompok
		var datakelompok []model.Kelompok
		var count int

		rows, _ := config.DB.Table("tbl_pengaturankelompok").
			Select("tbl_pengaturankelompok.kodekelompok, tbl_pengaturankelompok.namakelompok").
			Order("tbl_pengaturankelompok.namakelompok asc").
			Rows()

		for rows.Next() {
			if err := rows.Scan(&kelompok.Kodekelompok, &kelompok.Namakelompok); err != nil {
				count++
				log.Fatal(err.Error())
			} else { 
				var datapelajaran []model.Pelajaran 
				//config.DB.LogMode(true)
				config.DB.Table("tb_jurusanxpelajaran").
					Select("tb_jurusanxpelajaran.kodepelajaran, tbl_pengaturanpelajaran.namapelajaran, tbl_pengaturanpelajaran.kkm,(SELECT t.nilai FROM tbl_raportdetail t WHERE t.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t.kodenilai ='N' AND t.guidraport= ? ) as 'pengetahuan', (SELECT t1.nilai FROM tbl_raportdetail t1 WHERE t1.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t1.kodenilai ='P' AND t1.guidraport= ? ) as 'keterampilan', (SELECT t2.nilai FROM tbl_raportdetail t2 WHERE t2.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t2.kodenilai ='S' AND t2.guidraport= ? ) as 'sikap', IF((SELECT t3.nilai FROM tbl_raportdetail t3 WHERE t3.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t3.kodenilai ='NK' AND t3.guidraport = ? ) = 'Tuntas', 'T','TT') as 'keterangan'", json.Guidraport,json.Guidraport,json.Guidraport,json.Guidraport).
					Joins("INNER JOIN tbl_pengaturanpelajaran ON tbl_pengaturanpelajaran.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran").
					Where("tb_jurusanxpelajaran.kodekelas = ? AND  tb_jurusanxpelajaran.kodekelompok = ? AND tb_jurusanxpelajaran.Kodejuruan = ?", json.Kodekelas, kelompok.Kodekelompok,json.Kodejuruan). 
					Scan(&datapelajaran)
				kelompok.Pelajaran = datapelajaran
				datakelompok = append(datakelompok, kelompok)
				// config.DB.Raw("CALL `getraportdetail`('X', 'KELA', '5ef02923c6d4626', 'IPS')", json.Kodekelas, kelompok.Kodekelompok, json.Guidraport, json.Kodejuruan).Scan(&products)
				// 
				// rows1, _ := config.DB.Table("tb_jurusanxpelajaran").
				// 	Select("tb_jurusanxpelajaran.kodepelajaran, tbl_pengaturanpelajaran.namapelajaran, tbl_pengaturanpelajaran.kkm").
				// 	Joins("INNER JOIN tbl_pengaturanpelajaran ON tbl_pengaturanpelajaran.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran").
				// 	Where("tb_jurusanxpelajaran.kodekelas = ? AND tb_jurusanxpelajaran.Kodejuruan = ? AND tb_jurusanxpelajaran.kodekelompok = ?", json.Kodekelas, json.Kodejuruan, kelompok.Kodekelompok).
				// 	Order("tb_jurusanxpelajaran.iurutan asc").
				// 	Rows()
				// for rows1.Next() {
				// 	if err1 := rows1.Scan(&pelajaran.Kodepelajaran, &pelajaran.Namapelajaran, &pelajaran.Kkm); err1 != nil {
				// 		count++
				// 		log.Fatal(err1.Error())
				// 	} else {
				// 		var nilai []model.Nilai
				// 		config.DB.Table("tbl_raportdetail").
				// 			Select("(SELECT t.nilai FROM tbl_raportdetail t WHERE t.kodepelajaran = ? AND t.kodenilai ='N' AND t.guidraport=?) as 'pengetahuan',(SELECT t1.nilai FROM tbl_raportdetail t1 WHERE t1.kodepelajaran = ? AND t1.kodenilai ='P' AND t1.guidraport=?) as 'keterampilan',(SELECT t2.nilai FROM tbl_raportdetail t2 WHERE t2.kodepelajaran = ? AND t2.kodenilai ='S' AND t2.guidraport=?) as 'sikap',IF((SELECT t3.nilai FROM tbl_raportdetail t3 WHERE t3.kodepelajaran = ? AND t3.kodenilai ='NK' AND t3.guidraport = ?) = 'Tuntas', 'T','TT') as 'keterangan'", pelajaran.Kodepelajaran, json.Guidraport, pelajaran.Kodepelajaran, json.Guidraport, pelajaran.Kodepelajaran, json.Guidraport, pelajaran.Kodepelajaran, json.Guidraport).
				// 			Where("guidraport=? AND kodepelajaran = ?", json.Guidraport, pelajaran.Kodepelajaran).
				// 			Limit(1).
				// 			Scan(&nilai)
				// 		pelajaran.Nilai = nilai
				// 		datapelajaran = append(datapelajaran, pelajaran)
				// 	}
				// }
				// kelompok.Pelajaran = datapelajaran
				// 
			}
		}

		if count > 0 {
			c.JSON(http.StatusOK, gin.H{
				"error_code": "000001",
				"status":     "Data tidak ditemukan",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code": "000000",
				"status":     "Data ditemukan",
				"data":       datakelompok,
			})
		}

	}

}



// SELECT tb_jurusanxpelajaran.kodepelajaran, 
// 		 tbl_pengaturanpelajaran.namapelajaran, 
// 		 tbl_pengaturanpelajaran.kkm,
// 		 (SELECT t.nilai FROM tbl_raportdetail t WHERE t.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t.kodenilai ='N' AND t.guidraport=?) as 'pengetahuan',
// 		 (SELECT t1.nilai FROM tbl_raportdetail t1 WHERE t1.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t1.kodenilai ='P' AND t1.guidraport=?) as 'keterampilan',
// 		 (SELECT t2.nilai FROM tbl_raportdetail t2 WHERE t2.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t2.kodenilai ='S' AND t2.guidraport=?) as 'sikap',
// 		 IF((SELECT t3.nilai FROM tbl_raportdetail t3 WHERE t3.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran AND t3.kodenilai ='NK' AND t3.guidraport = ?) = 'Tuntas', 'T','TT') as 'keterangan'
// FROM tb_jurusanxpelajaran INNER JOIN tbl_pengaturanpelajaran ON tbl_pengaturanpelajaran.kodepelajaran = tb_jurusanxpelajaran.kodepelajaran
// WHERE tb_jurusanxpelajaran.kodekelas = "X" AND tb_jurusanxpelajaran.Kodejuruan = "IPS" AND tb_jurusanxpelajaran.kodekelompok = "KELA"
 