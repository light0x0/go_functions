package piscine

// import "github.com/01-edu/z01"

// func PrintNbrBase(nbr int, base string){
//     baseRune := []rune(base)
//     baselen := len(baseRune)
//     if baselen < 2 {
//         return
//     }
//     for index1:= 0; index1 < baselen; index1++{
//         value := baseRune[index1]
//         if value == '+' || value == '-' || value <= ' ' || value == 127 {
//             return
//         }
//         for index2:= index1 + 1; index2 < baselen; index2++{
//             if baseRune[index2] == value {
//                 return
//             }
//         }

//     }

// 	if nbr == 0 {
// 		z01.PrintRune(baseRune[0])
// 		return
// 	}
// 	if nbr < 0 {
//         z01.PrintRune('-')
// 		PrintNbrBase(-nbr,base)
// 		return
// 	}

// 	if nbr >= baselen {
// 		PrintNbrBase(nbr / baselen,base)
// 	}

//     z01.PrintRune(baseRune[nbr%baselen])

// }

// func main() {
// 	PrintNbrBase(125, "0123456789")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "0123456789ABCDEF")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "choumi")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "aa")
// 	z01.PrintRune('\n')
// }
