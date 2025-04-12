package cp_template

const HELPER_FUNCTIONS = `
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
`
