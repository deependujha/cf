package cp_template

const PREPROCESSING_HEADER = `
using namespace std;

#define F                   first
#define S                   second
#define pb                  push_back
#define int					long long
#define double         		long double
#define pq_max        		priority_queue <int>
#define pq_min       		priority_queue <int, vi, greater<int>>
#define print(a)       		for(auto x : a) cout << x << " "; cout << endl
#define print1(a)      		for(auto x : a) cout << x.F << " " << x.S << endl
#define print2(a,x,y)  		for(int i = x; i < y; i++) cout<< a[i]<< " "; cout << endl
#define lcm(a,b) 			(a*(b/__gcd(a,b)))
#define ull 				unsigned long long
#define mm 					1000000007
#define PRIZE_FIGHTER 		ios_base::sync_with_stdio(false);cin.tie(NULL);  
#define rep(n) 				for(int i=0;i<n;i++)
#define lop(a,n) 			for(int a=0;a<n;a++)
#define REP(a,i,n)			for(int a=i;a<n;a++)

const int N = 200005;
`
