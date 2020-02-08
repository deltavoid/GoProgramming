package main

import "fmt"


func dfs(n int, k int, a []int, vis []bool) {
	// fmt.Println(n, k, a[1], vis[1]);

	if  (k > n) {
	    for i := 1; i <= n; i++ {
			fmt.Printf("%d ", a[i]);
		}
		fmt.Println();
	} else {
        for i := 1; i <= n; i++ {
			if  !vis[i] {
				a[k] = i;
				vis[i] = true;
				dfs(n, k + 1, a[:], vis[:]);
				vis[i] = false;
				a[k] = 0;
			}
		}
	}


}

func main() {

	const N = 10;
	var n = 3;
	var a [N]int;
	var vis [N]bool;

	// for i := 0; i < N; i++ {
	// 	a[i] = 0;
	// 	vis[i] = false;
	// }

	// for i := 0; i < N; i++ {
	// 	fmt.Println(a[i], vis[i]);
	// }

	dfs(n, 1, a[:], vis[:]);



}