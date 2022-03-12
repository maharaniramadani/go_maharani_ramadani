package main

type kendaraan struct {
	totalRoda int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan (){
	m.tambahKecepatan (10)
}

func (m *mobil) tambahKecepatan (kecepatanBaru int) {
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main(){
	mobilcepat := mobil{}
	Count := 3
	for i := 0; i < Count; i++ {
		mobilcepat.berjalan()
	}
	 
	mobillamban := mobil{}
	mobillamban.berjalan()
}