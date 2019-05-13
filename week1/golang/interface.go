package main

import (
	"fmt"
	"math"
)

type hinh interface {
	dien_tich() float64
}

type hinh_vuong struct {
	canh float64
}

type hinh_tron struct {
	r float64
}

type hinh_chu_nhat struct {
	chieu_dai, chieu_rong float64
}

func (hv hinh_vuong) dien_tich() float64{
	return hv.canh*hv.canh
}

func (ht hinh_tron) dien_tich() float64 {
	return ht.r*ht.r*math.Pi
}

func (hcn hinh_chu_nhat) dien_tich() float64 {
	return hcn.chieu_dai*hcn.chieu_rong
}

func tinh_dien_tich(h hinh) {
	fmt.Printf("Dien tich = %0.2f\n", h.dien_tich())
}

func main() {
	hv := hinh_vuong{4}
	hcn := hinh_chu_nhat{4, 5}
	ht := hinh_tron{5}

	fmt.Println("Hinh vuong:")
	tinh_dien_tich(hv)
	fmt.Println("Hinh tron:")
	tinh_dien_tich(ht)
	fmt.Println("Hinh chu nhat:")
	tinh_dien_tich(hcn)
}