package main

import "fmt"

func main() {
	// fmt.Println("---ex0---")
	// ex0()
	// fmt.Println("---ex1---")
	// ex1()
	// fmt.Println("---ex1.2---")
	ex1_2(20,2)




	fmt.Println("--ex2---")
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	fmt.Println("ค่าที่น้อยที่สุดคือ ",ex2min(x))
	fmt.Println("ค่าที่มากที่สุดคือ ",ex2max(x))







}

//0.จากหน้า 19 จงเปลี่ยนให้เป็น if
func ex0() {
	i := 2
	fmt.Println("Example-: if condition")
	if i == 0 {
		fmt.Println("zero")
	} else if i == 1 {
		fmt.Println("one")
	} else if i == 2 {
		fmt.Println("two")
	} else if i == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("Your i not in case")
	}
}

//1.ระหว่าง 1-100 มีเลขที่หาร3ลงตัวกี่ตัว อะไรบ้าง (for if)
func ex1() {
	count := 0

	fmt.Println("เลขที่หารด้วย 3 ลงตัว :")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d, ", i)
			count++
		}
	}
	fmt.Println()
	fmt.Printf("เลขที่หารด้วย 3 ลงตัวมีทั้งหมด: %d", count)
}

//1.2 สร้างฟังชั่นคำนวณเลขยกกำลัง เช่น num(20,2) (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า27-30)ไม่เอาfuncสำเร็จรูป math.pow()
func ex1_2(x int,y int) int{
	result := 1 
	for i := 0; i < y; i++ { 
		result *= x 
	}
    fmt.Printf("%d ยกกำลัง %d มีค่าเท่ากับ: %d\n", x, y, result) 
	return result
}

//2.เขียนโปรแกรมหาเลขที่น้อยสุดและมากสุด (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า27-30) ไม่เอาfuncสำเร็จรูป math.min() math.max()
func ex2min(nums []int)int{
	min := nums[0]
	for _, v := range nums {
		if v < min{
			min = v
		}
	}
	return min
}

func ex2max(nums []int)int{
	max := nums[0]
	for _, v:= range nums{
		if v>max {
			max =v
		}
	}
	return max
}















//2.เขียนโปรแกรมหาเลขที่น้อยสุดและมากสุด (เรียกใช้เป็นฟังชั่น) ฝึกการใช้ฟังชั่น (หน้า27-30) ไม่เอาfuncสำเร็จรูป math.min() math.max()
