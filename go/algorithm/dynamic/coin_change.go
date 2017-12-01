package dynamic

func MakeChange(money int, coinKinds []int) int {
    /**
     * 硬币找零问题：给定待换面额，换成给定零钱面额，使得硬币数目最小。
     * 硬币找零问题，对于部分给定零钱面额是可以使用贪心的，但不是任意面额，
     * 所以此类问题最好解决方法还是动态规划
     * Created by Administrator on 2015/4/21.
     */
    F := make([]int, money + 1)
    for i := 1; i <= money; i ++ {
        F[i] = i
    }

    for i := 1; i <= money; i ++ {
        F[i] = F[i-1] + 1
        for _, v := range coinKinds {
            if v <= i {
                F[i] = min(F[i], F[i-v] + 1)
            }
        }
    }
    return F[money]
}
