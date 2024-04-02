package butterfly

type Storer struct {

}

//			  1   2  3  4  5   6
//1 string = key:key:-:key:-:key
//             3     5
//2 string = trash:trash
//[3; infinitive] = value
// 					value
//					  -
//					value
//					  -
//					value