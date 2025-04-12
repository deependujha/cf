
/*

██████╗░███████╗███████╗██████╗░███████╗███╗░░██╗██████╗░██╗░░░██╗
██╔══██╗██╔════╝██╔════╝██╔══██╗██╔════╝████╗░██║██╔══██╗██║░░░██║
██║░░██║█████╗░░█████╗░░██████╔╝█████╗░░██╔██╗██║██║░░██║██║░░░██║
██║░░██║██╔══╝░░██╔══╝░░██╔═══╝░██╔══╝░░██║╚████║██║░░██║██║░░░██║
██████╔╝███████╗███████╗██║░░░░░███████╗██║░╚███║██████╔╝╚██████╔╝
╚═════╝░╚══════╝╚══════╝╚═╝░░░░░╚══════╝╚═╝░░╚══╝╚═════╝░░╚═════╝░


*/



#include <algorithm>
#include <bitset>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <queue>
#include <set>
#include <set>
#include <stack>
#include <string>
#include <unordered_map>
#include <vector>


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


ull power(ull x, ull y){
   if(y==0){
       return 1;
    }
   ull temp=power(x,y/2);
    temp=temp*temp;
    if(y%2!=0){
   temp=x*temp;
   }
   return temp;
}

bool isPrime(int val)
{
    if (val < 2)
        return false;
    for (int i = 2; i * i <= val; i++)
    {
        if (val % i == 0)
        {
            return false;
        }
    }
    return true;
}

int factorial(int n) {
    int factorial = 1;
    for (int i = 2; i <= n; i++)
        factorial = factorial * i;
    return factorial;
}


int nCr(int n, int r) {
    int denominator=factorial(r);
    int numerator=1;
    for(int i=n;i>=n-r+1;i--){
        numerator*=i;
    }
    return numerator/denominator;
}


void solve(){

    

}

int32_t main(){

   PRIZE_FIGHTER
/*
    #ifndef ONLINE_JUDGE
        freopen("input.txt","r",stdin);
        freopen("output.txt","w",stdout);
    #endif
*/

    int testcases=1;
    //cin>>testcases;
    for(int testcase_no=1;testcase_no<=testcases;testcase_no++){
        solve();
    // cout<<"Case #"<<testcase_no<<": "<<"OUTPUT"<<endl;
    }
    return 0;
}