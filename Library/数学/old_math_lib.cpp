#include <ctime>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <deque>
#include <list>
#include <queue>
#include <stack>
#include <set>
#include <map>
#include <numeric>
#include <algorithm>
#include <functional>
using namespace std;

typedef int       int32;
typedef long long int64;

const int32  mod  = 10007;
const int32  maxn = 10010;
const int32  inf  = 0x3f3f3f3f;
const int64  INF  = 0x3f3f3f3f3f3f3f3fLL;
const double pi   = 3.1415926;
const double eps  = 1e-8;

int32 gcd(int32 a, int32 b) {
    return b == 0 ? a : gcd(b, a % b);
}

int64 gcd(int64 a, int64 b) {
    return b == 0 ? a : gcd(b, a % b);
}

int32 lcm(int32 a, int32 b) {
    return a / gcd(a, b) * b;
}

int64 lcm(int64 a, int64 b) {
    return a / gcd(a, b) * b;
}

int32 extend_gcd(int32 a, int32 b, int32 &x, int32 &y) {
    if(b == 0) {
        x = 1;
        y = 0;
        return a;
    }
    int32 d = extend_gcd(b, a % b, y, x);
    y -= x * (a / b);
    return d;
}

int64 extend_gcd(int64 a, int64 b, int64 &x, int64 &y) {
    if(b == 0) {
        x = 1;
        y = 0;
        return a;
    }
    int64 d = extend_gcd(b, a % b, y, x);
    y -= x * (a / b);
    return d;
}

// O(nlogn) 素数筛法
bool isp[maxn];
int  prime[maxn], nump;

void get_prime_table() {
    int i, j;
    nump = 0;
    prime[nump++] = 2;
    memset(isp, true, sizeof(isp));
    for(i = 4; i < maxn; i += 2) isp[i] = false;
    for(i = 3; i * i <= maxn; i += 2) {
        if(isp[i]) {
            prime[nump++] = i;
            for(j = i * i; j < maxn; j += i) {
                isp[j] = false;
            }
        }
    }
    for( ; i < maxn; i += 2) if(isp[i]) prime[nump++] = i;
    return ;
}

// O(n) 线性素数筛法
void get_prime_table2() {
    nump = 0;
    memset(isp, true, sizeof(isp));
    for(int i = 2; i < maxn; ++i) {
        if(isp[i]) {
            prime[nump++] = i;
        }
        for(int j = 0; j < nump && i * prime[j] < maxn; ++j) {
            isp[i*prime[j]] = false;
            if(i % prime[j] == 0) break;
        }
    }
    return ;
}

// ax = b(mod n)
void mod_equation(int a, int b, int n) {
    int e, i, d, x, y;
    d = extend_gcd(a, n, x, y);
    if(b % d) {
        printf("No Answer!\n");
        return ;
    } else {
        e = (x * b / d) % n;
        for(i = 0; i < d; ++i) {
            printf("%d\n", e + i * (n / d));
        }
    }
    return ;
}

// 模逆元
int inv(int a, int n) {
    int x, y;
    int d = extend_gcd(a, n, x, y);
    if(d != 1) return -1;
    return (x % n + n) % n;
}

// 中国剩余定理
// a = b[i] (mod w[i]), gcd(w[i], w[j]) == 1
int china_reminder1(int b[], int w[], int k) {
    int i, d, x, y, m, a = 0, n = 1;
    for(i = 0; i < k; ++i) n *= w[i];
    for(i = 0; i < k; ++i) {
        m = n / w[i];
        d = extend_gcd(w[i], m, x, y);
        printf("gcd is: %d\n", d);
        a = (a + y * m * b[i]) % n;
    }
    if(a < 0) a += n;
    return a;
}

bool merge_equation(int64 a1, int64 n1, int64 a2, int64 n2, int64 &a3, int64 &n3) {
    int64 d = gcd(n1, n2);
    int64 c = a2 - a1;
    if(c % d) return false;
    c = (c % n2 + n2) % n2;
    c /= d;
    n1 /= d;
    n2 /= d;
    c *= inv(n1, n2);
    c %= n2;
    c *= (n1 * d);
    c %= n2;
    c *= (n1 * d);
    c += a1;
    n3 = n1 * n2 * d;
    a3 = (c % n3 + n3) % n3;
    return true;
}

// 非互质情况下中国剩余定理
// x = a[i] (mod n[i])
int64 china_reminder2(int len, int64 a[], int64 n[]) {
    int64 a1 = a[0], n1 = n[0];
    int64 a2, n2;
    for(int i = 1; i < len; ++i) {
        int64 aa, nn;
        a2 = a[i], n2 = n[i];
        if(!merge_equation(a1, n1, a2, n2, aa, nn)) return -1;
        a1 = aa;
        n1 = nn;
    }
    return (a1 % n1 + n1) % n1;
}

// 米勒拉宾素数测试
// a ^ (p-1) = 1 (mod p) + ����̽��
int witness(int a, int n) {
    int x, d = 1, i = ceil(log(n*1.0-1.0)/log(2.0)) - 1;
    for( ; i >= 0; --i) {
        x = d;
        d = (d * d) % n;
        if(d == 1 && x != 1 && x != n - 1) return 1;
        if(((n-1)&(1<<i)) > 0) d = (d * a) % n;
    }
    return d == 1 ? 0 : 1;
}

bool miller(int n, int s = 50) {
    if(n == 2) return true;
    if(n % 2 == 0) return false;
    int i, a;
    srand((unsigned)time(NULL));
    for(i = 0; i < s; ++i) {
        a = rand() * (n - 2) / RAND_MAX + 1;
        if(witness(a, n)) return false;
    }
    return true;
}

// Pollard-Rho���ӷֽ��㷨
// find(n, C)
// result�е������Ӵ����ظ�
// pcnt���ֽ�������ӵĸ���
// res���洢�ֽ���������
const int C = 201;
int pcnt;
int res[maxn];

int64 random(int64 n) {
    return (int64)((double)rand() / RAND_MAX * n + 0.5);
}

int64 mul_mod(int64 a, int64 b, int64 c) {
    int64 ans = 0;
    while(b) {
        if(b & 1) ans = (ans + a) % c;
        a = 2 * a % c;
        b >>= 1;
    }
    return ans;
}

int64 pow_mod(int64 a, int64 b, int64 c) {
    int64 ans = 1;
    while(b) {
        if(b & 1) ans = mul_mod(ans, a, c);
        a = mul_mod(a, a, c);
        b >>= 1;
    }
    return ans;
}

int64 pollard_rho(int64 n, int64 c) {
    int64 x, y, d, i = 1, k = 2;
    x = random(n - 1) + 1;
    y = x;
    while(1) {
        ++i;
        x = (mul_mod(x, x, n) + c) % n;
        d = gcd(y - x + n, n);
        if(d > 1 && d < n) return d;
        if(y == x) return n;
        if(i == k) {
            y = x;
            k *= 2;
        }
    }
    return n;
}

void find_prime(int64 n, int k) {
    if(n == 1) return ;
    if(miller(n)) {
        res[pcnt++] = n;
        return ;
    }
    int64 p = n;
    while(p >= n) {
        p = pollard_rho(p, k - 1);
    }
    find_prime(p, k);
    find_prime(n / p, k);
    return ;
}

// 欧拉函数
int get_euler(int x) {
    int i, ret = x;
    for(i = 2; i * i <= x; ++i) {
        if(x % i == 0) {
            ret = ret / i * (i - 1);
            while(x % i == 0) x /= i;
        }
    }
    if(x > 1) ret = ret / x * (x - 1);
    return ret;
}

// 递推求解欧拉函数
int phi[maxn+10];
void get_euler_table() {
    for(int i = 1; i < maxn; ++i) phi[i] = i;
    for(int i = 2; i < maxn; i += 2) phi[i] /= 2;
    for(int i = 3; i < maxn; i += 2) {
        if(phi[i] == i) {
            for(int j = i; j < maxn; j += i) {
                phi[j] = phi[j] / i * (i - 1);
            }
        }
    }
    return ;
}

// ���sigma(gcd(i,j)) 1 <= i, j <= n
// f(n) = sigma(gcd(i,n))
// s(n) = sigma(f(i))
int f[maxn], s[maxn];
void solve() {
    for(int i = 1; i < maxn; ++i) f[i] = phi[i];
    for(int i = 2; i * i <= maxn; ++i) {
        f[i*i] += i * phi[i];
        for(int j = i * i + i, k = i + 1; j < maxn; j += i, ++k) {
            f[j] += i * phi[k] + k * phi[i];
        }
    }
    s[1] = f[1];
    for(int i = 2; i < maxn; ++i) {
        s[i] = s[i-1] + f[i];
    }
    return ;
}

// 素因子计算
int factor[maxn];
int cal_factor(int n) {
    int cnt = 0;
    for(int i = 2; i * i <= n; i += 2) {
        while(n % i == 0) {
            factor[cnt++] = i;
            n /= i;
        }
        if(i == 2) --i;
    }
    if(n > 1) factor[cnt++] = n;
    return cnt;
}

int cal_factor_num(int n) {
    int ans = 1, cnt;
    for(int i = 2; i * i <= n; i += 2) {
        if(n % i == 0) {
            cnt = 1;
            while(n % i == 0) n /= i, cnt++;
            ans *= cnt;
        }
        if(i == 2) --i;
    }
    if(n > 1) ans *= 2;
    return ans;
}

int cal_factor_sum(int n) {
    int ans = 0, tmp;
    for(int i = 2; i * i <= n; i += 2) {
        if(n % i == 0) {
            tmp = i;
            while(n % i == 0) n /= i, tmp *= i;
            ans = ans * (tmp - 1) / (i - 1);
        }
        if(i == 2) --i;
    }
    if(n > 1) ans = ans * (n + 1);
    return ans;
}

// ԭ�����
// ԭ���ķֲ��ܹ㷺����С��ԭ��ͨ��Ҳ��С��
// �ʿ��Բ�ȡ��С����ö�ٵķ������ٵ�Ѱ��һ��ԭ��
// �洢p-1��ÿһ��������
vector<int64> fact;
bool g_test(int64 g, int64 p) {
    int sz = fact.size();
    for(int i = 0; i < sz; ++i) {
        if(pow_mod(g, (p-1)/fact[i], p) == 1) {
            return false;
        }
    }
    return true;
}

int64 primitive_root(int64 p) {
    int64 tmp = p - 1;
    for(int i = 2; (int64)i * i <= tmp; i += 2) {
        if(tmp % i == 0) {
            fact.push_back(i);
            while(tmp % i == 0) tmp /= i;
        }
        if(i == 2) --i;
    }
    if(tmp > 1) fact.push_back(tmp);
    int64 g = 1;
    while(true) {
        if(g_test(g, p)) return g;
        ++g;
    }
    return -1;
}

// ƽ��ʣ��
// x^2 = a (mod n)
int modsqr(int a, int n) {
    int b, k, i, x;
    if(n == 2) return a % n;
    if(pow_mod(a, (n-1)/2, n) == 1) {
        if(n % 4 == 3) x = pow_mod(a, (n+1)/4, n);
        else {
            for(b = 1; pow_mod(b, (n-1)/2, n) == 1; ++b);
            i = (n - 1) / 2;
            k = 0;
            do {
                i /= 2;
                k /= 2;
                if((pow_mod(a, i, n)*pow_mod(b, k, n)+1) % n == 0) {
                    k += (n - 1) / 2;
                }
            } while(i % 2 == 0);
            x = (pow_mod(a, (i+1)/2, n)) * pow_mod(b, k/2, n) % n;
        }
        if(x * 2 > n) x = n - x;
        return x;
    }
    return -1;
}

// 离散对数
// ��ɢ����
// ����x, n, m�� x^y = n (mod m),����mΪ����
int64 discrete_log(int x, int n, int m) {
    map<int64, int> rec;
    int s = (int)(sqrt(m*1.0));
    for( ; (int64)s*s <= m; ) ++s;
    int64 cur = 1;
    for(int i = 0; i < s; ++i) {
        rec[cur] = i;
        cur = cur * x % m;
    }
    int64 mul = cur;
    cur = 1;
    for(int i = 0; i < s; ++i) {
        int64 more = (int64)n*pow_mod(cur, m - 2, m) % m;
        if(rec.count(more)) {
            return i * s + rec[more];
        }
        cur = cur * mul % m;
    }
    return -1;
}

// ����x, n, m����x^y = n(mod m)������m��һ��������
struct Hash {
    int a, b, next;
}hhash[maxn<<1];
int flag[maxn];
int top, idx;

int inv_hash(int a, int b, int n) {
    int x, y, e;
    extend_gcd(a, n, x, y);
    e = (int64)x*b%n;
    return e < 0 ? e + n : e;
}

void insert_hash(int a, int b) {
    int k = b & maxn;
    if(flag[k] != idx) {
        flag[k] = idx;
        hhash[k].next = -1;
        hhash[k].a = a;
        hhash[k].b = b;
        return ;
    }
    while(hhash[k].next != -1) {
        if(hhash[k].b == b) return ;
        k = hhash[k].next;
    }
    hhash[k].next = ++top;
    hhash[top].next = -1;
    hhash[top].a = a;
    hhash[top].b = b;
    return ;
}

int find_hash(int b) {
    int k = b & maxn;
    if(flag[k] != idx) return -1;
    while(k != -1) {
        if(hhash[k].b == b) return hhash[k].a;
        k = hhash[k].next;
    }
    return -1;
}

int baby(int A, int B, int C) {
    top = maxn;
    ++idx;
    int64 buf = 1 % C, D = buf, k;
    int i, d = 0, tmp;
    for(i = 0; i <= 100; buf = buf * A % C, ++i) if(buf == B) return i;
    while((tmp=gcd(A, C)) != 1) {
        if(B % tmp) return -1;
        ++d;
        C /= tmp;
        B /= tmp;
        D = D * A / tmp % C;
    }
    int M = (int)ceil(sqrt(C*1.0));
    for(buf = 1 % C, i = 0; i <= M; buf = buf * A % C, ++i) insert_hash(i, buf);
    for(i = 0, k = pow_mod(A, M, C); i <= M; D = D * k % C, ++i) {
        tmp = inv_hash(D, B, C);
        int w;
        if(tmp > 0 && (w=find_hash(tmp)) != -1) return i * M + w + d;
    }
    return -1;
}

// N��ʣ��
// ����N, a, p, ���x^N = a (mod p)��ģp�����µ����н⣬����pΪ����
vector<int> residue(int p, int N, int a) {
    int g = primitive_root(p);
    int64 m = discrete_log(g, a, p);
    vector<int> ret;
    if(a == 0) {
        ret.push_back(0);
        return ret;
    }
    if(m == -1) return ret;
    int64 A = N, B = p - 1, C = m, x, y;
    int64 d = extend_gcd(A, B, x, y);
    if(C % d != 0) return ret;
    x = x * (C / d) % B;
    int64 delta = B / d;
    for(int i = 0; i < d; ++i) {
        x = ((x + delta) % B + B) % B;
        ret.push_back(pow_mod(g, x, p));
    }
    sort(ret.begin(), ret.end());
    ret.erase(unique(ret.begin(), ret.end()), ret.end());
    return ret;
}

// 目函数
// Ī����˹�������
// ���Ʒ������Ī����˹������O(nlogn)
int mu[maxn+10];
void getMu() {
    for(int i = 1; i <= maxn; ++i) {
        int target = (i == 1 ? 1 : 0);
        int delta = target - mu[i];
        mu[i] = delta;
        for(int j = i + i; j <= maxn; j += i) {
            mu[j] += delta;
        }
    }
    return ;
}

// ����ɸ�����Ī����˹����
int check[maxn];
//int prime[maxn];
//int mu[maxn];
void moblus() {
    memset(check, 0, sizeof(check));
    mu[1] = 1;
    int tot = 0;
    for(int i = 2; i <= 1000; ++i) {
        if(!check[i]) {
            prime[tot++] = i;
            mu[i] = -1;
        }
        for(int j = 0; j < tot; ++j) {
            if(i * prime[j] > maxn) break;
            check[i*prime[j]] = 1;
            if(i % prime[j] == 0) {
                mu[i*prime[j]] = 0;
                break;
            } else {
                mu[i*prime[j]] = -mu[i];
            }
        }
    }
    return ;
}

// ������: dfs + ��֦ ��ⷴ����
// ��ͨĸ����
// (1+x^1+x^2+...)(1+x^2+...)(1+x^k+x^(2k)+....)
int c1[maxn];
int c2[maxn];

void cal_mu_func(int n) {
    for(int i = 0; i <= n; ++i) {
        c1[i] = 1;
        c2[i] = 0;
    }
    for(int i = 2; i <= n; ++i) {
        for(int j = 0; j <= n; ++j) {
            for(int k = 0; k <= n; k += j) {
                if(k + j > n) break;
                c2[j+k] += c1[j];
            }
        }
        for(int j = 0; j <= n; ++j) {
            c1[j] = c2[j];
            c2[j] = 0;
        }
    }
    return ;
}

// ָ����ĸ����
/*
double fac[maxn];
double ans[maxn];
double tmp[maxn];
double num[maxn];

void init_fac() {
    fac[0] = 1;
    for(int i = 1; i <= 50; ++i) {
        fac[i] = fac[i-1] * i;
    }
}

void cal_zs_mu_func(int n, int m) {
    for(int i = 0; i < n; ++i) cin >> num[i];
    memset(ans, 0, sizeof(ans));
    memset(tmp, 0, sizeof(tmp));
    for(int i = 0; i <= num[0]; ++i) {
        ans[i] = 1.0 / fac[i];
    }
    for(int i = 1; i < n; ++i) {
        for(int j = 0; j < 50; ++j) {
            for(int k = 0; k <= num[i] && k + j < 50; ++k) {
                tmp[j+k] += ans[j] / fac[k];
            }
        }
        for(int j = 0; j < 50; ++j) {
            ans[j] = tmp[j];
            tmp[j] = 0;
        }
    }
    for(int i = 1; i < 50; ++i) {
        printf("%d\n", (int)(ans[m]*fac[m]));
    }
    return ;
}
*/

// Lucas��ģ
// C(n,m)%p, n��m�ǷǸ�������pΪ����
// Lucas(n, m, p) = C(n%p, m%p)*Lucas(n/p,m/p,p)
// ��nת����p����a[0],a[1],....��ͬʱ��mת����p����b[0],b[1],...
// ��C(n,m)��C(a[0],b[0])*C(a[1],b[1])...ͬ��
// ����Lucas�����������ģ�����õݹ�ʽ��⼴��

// p��һ��Ϊ����ʱ�Ŀ����������ģ����
// C(n,m) = n! / ((n-m)! * m!)
// ö�����п��ܵ������ӣ��ֱ����n!,m!,(n-m)!�а��������ӵĴ�����Ȼ��˵����յĽ����

// �����ѧ�������ݳ�ԭ�������������ʱ,�����������ַ���:
// 1. ֱ��dfsö��������ϵ�������ּ��ϴ�С����ż�Լ��㼴��
// 2. ������ö�����е��Ӽ���Ȼ��ֱ������Ӧ���Ӽ��ϼ���
// 0--->(1<<n)-1


// ������
struct Matrix {
    int mat[10][10];
    Matrix() {
        memset(mat, 0, sizeof(mat));
    }
    void init() {
        memset(mat, 0, sizeof(mat));
        for(int i = 0; i < 10; ++i) mat[i][i] = 1;
    }
    friend Matrix operator + (const Matrix &a, const Matrix &b);
    friend Matrix operator * (const Matrix &a, const Matrix &b);
    friend Matrix operator ^ (Matrix a, int x);
    friend Matrix sum(Matrix a, int n);
};

Matrix operator + (const Matrix &a, const Matrix &b) {
    Matrix c;
    for(int i = 0; i < 10; ++i) {
        for(int j = 0; j < 10; ++j) {
            c.mat[i][j] = (a.mat[i][j] + b.mat[i][j]) % mod;
        }
    }
    return c;
}

Matrix operator * (const Matrix &a, const Matrix &b) {
    Matrix c;
    for(int i = 0; i < 10; ++i) {
        for(int j = 0; j < 10; ++j) {
            for(int k = 0; k < 10; ++k) {
                c.mat[i][j] = (c.mat[i][j] + a.mat[i][k] * b.mat[k][j]) % mod;
            }
        }
    }
    return c;
}

Matrix operator ^ (Matrix a, int x) {
    Matrix c; c.init();
    while(x) {
        if(x & 1) c = c * a;
        a = a * a;
        x >>= 1;
    }
    return c;
}

// ��λ����
// E.init()
Matrix E;

// ans = a + a^1 + a^2 + ... + a^n
Matrix sum(Matrix a, int n) {
    if(n == 1) return a;
    if(n % 2 == 0) return (E + (a^(n/2))) * sum(a, n/2);
    else return (a^n) + sum(a, n - 1);
}

int main() {

    cout << ~(1+2+3) << endl;
    cout << ~1 + ~2 + ~3 << endl;

    return 0;
}


class Solution {
public:
    vector<int> t;
    vector<vector<int>> ans;

    vector<vector<int>> subsetsWithDup(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        for (int mask = 0; mask < (1 << n); ++mask) {
            t.clear();
            bool flag = true;
            for (int i = 0; i < n; ++i) {
                if (mask & (1 << i)) {
                    if (i > 0 && (mask >> (i - 1) & 1) == 0 && nums[i] == nums[i - 1]) {
                        flag = false;
                        break;
                    }
                    t.push_back(nums[i]);
                }
            }
            if (flag) {
                ans.push_back(t);
            }
        }
        return ans;
    }
};

/*
class Solution {
public:
    vector<int> t;
    vector<vector<int>> ans;

    void dfs(bool choosePre, int cur, vector<int> &nums) {
        if (cur == nums.size()) {
            ans.push_back(t);
            return;
        }
        dfs(false, cur + 1, nums);
        if (!choosePre && cur > 0 && nums[cur - 1] == nums[cur]) {
            return;
        }
        t.push_back(nums[cur]);
        dfs(true, cur + 1, nums);
        t.pop_back();
    }

    vector<vector<int>> subsetsWithDup(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        dfs(false, 0, nums);
        return ans;
    }
};
*/