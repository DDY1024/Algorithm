package 线段树

import "fmt"

// 比较经典的一道利用 priority_queue + set 的题目，
// 具体解题思路可以参考：https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/solution/cseter-fen-cha-zhao-priority_queue-by-russell_kean/
// 1. priority_queue 维护服务器下一次可以处理任务的时间，
// 2. set 维护在每次任务到达前，可以处理该任务的服务器
/*
class Solution {
public:
    vector<int> busiestServers(int k, vector<int>& arrival, vector<int>& load) {
        int n=arrival.size();
        vector<int> cnt(k,0);
        set<int,less<int>> s;
        priority_queue<pair<int,int>,vector<pair<int,int>>,greater<pair<int,int>>> lq;
        for(int i=0;i<k;i++){
            s.insert(i);
        }
        for(int i=0;i<n;i++){
            if(!s.empty()){
                int u;
                int til=arrival[i]+load[i];
                auto it=s.lower_bound(i%k);
                if(it==s.end()){
                    u=*(s.begin());
                    s.erase(s.begin());
                }else{
                    u=*it;
                    s.erase(it);
                }
                cnt[u]++;
                lq.push({til,u});
            }
            if(i+1<n){
                while(!lq.empty() && lq.top().first<=arrival[i+1]){
                    auto t=lq.top();
                    lq.pop();
                    s.insert(t.second);
                }
            }
        }
        int mx=0;
        for(int i=0;i<k;i++){
            mx=max(mx,cnt[i]);
        }
        vector<int> ans;
        for(int i=0;i<k;i++){
            if(cnt[i]==mx){
                ans.push_back(i);
            }
        }
        return ans;
    }
};
*/

//func main() {
//	var once sync.Once
//	once.Do()
//}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func push_up(arr []int, idx int) {
	arr[idx] = minInt(arr[idx<<1], arr[(idx<<1)|1])
}

func build(l, r, idx int, arr []int) {
	if l == r {
		arr[idx] = 0
		return
	}
	mid := (l + r) >> 1
	build(l, mid, idx<<1, arr)
	build(mid+1, r, (idx<<1)|1, arr)
	push_up(arr, idx)
	return
}

// 注意这种类型的区间查询不能只查左边，左边查不到的情况同时需要查右边
func query(l, r, val, s, t, idx int, arr []int) int {
	if l <= s && t <= r && s == t {
		return s
	}
	mid := (s + t) >> 1
	idx1, idx2 := -1, -1
	// 注意：[s, mid] 区间范围可能大于 [l, mid] 区间，但是结果落在 [s, l-1] 区间，不符合条件，因此我们在查完左边不存在以后，需要再查一遍右边
	// 不能直接根据左边得出结果。
	if l <= mid && arr[idx<<1] <= val {
		idx1 = query(l, r, val, s, mid, idx<<1, arr)
	}
	if idx1 != -1 {
		return idx1
	}
	if r > mid && arr[(idx<<1)|1] <= val {
		idx2 = query(l, r, val, mid+1, t, (idx<<1)|1, arr)
	}
	return idx2
}

func update(pos, val, s, t, idx int, arr []int) {
	if s == t {
		arr[idx] = val
		return
	}
	mid := (s + t) >> 1
	if pos <= mid {
		update(pos, val, s, mid, idx<<1, arr)
	} else {
		update(pos, val, mid+1, t, (idx<<1)|1, arr)
	}
	push_up(arr, idx)
	return
}

func busiestServers(k int, arrival []int, load []int) []int {
	minV, n := make([]int, (k+10)<<2), len(arrival)
	build(0, k-1, 1, minV)
	cnt := make([]int, k)
	for i := 0; i < n; i++ {
		idx1 := query(i%k, k-1, arrival[i], 0, k-1, 1, minV)
		if idx1 != -1 {
			cnt[idx1]++
			update(idx1, arrival[i]+load[i], 0, k-1, 1, minV)
			continue
		}
		idx2 := query(0, i%k, arrival[i], 0, k-1, 1, minV)
		if idx2 != -1 {
			cnt[idx2]++
			update(idx2, arrival[i]+load[i], 0, k-1, 1, minV)
			continue
		}
	}
	maxC, ans := 0, make([]int, 0)
	for i := 0; i < k; i++ {
		maxC = maxInt(maxC, cnt[i])
	}
	for i := 0; i < k; i++ {
		if cnt[i] == maxC {
			ans = append(ans, i)
		}
	}
	return ans
}

// k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3]
// 13
// [1,3,6,7,8,9,10,14,16,20,21,24,25,28,29,30,33,34]
// [20,27,27,14,14,9,15,8,23,1,34,2,28,25,7,6,24,15]

// 21,30,33,21,22,18,25,
func main() {
	fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}))
	fmt.Println(busiestServers(13, []int{1, 3, 6, 7, 8, 9, 10, 14, 16, 20, 21, 24, 25, 28, 29, 30, 33, 34},
		[]int{20, 27, 27, 14, 14, 9, 15, 8, 23, 1, 34, 2, 28, 25, 7, 6, 24, 15}))
}
