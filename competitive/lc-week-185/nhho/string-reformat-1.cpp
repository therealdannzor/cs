/*
 * Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).
 *
 * You have to find a permutation of the string where no letter is followed by another
 * letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.
 *
 * Return the reformatted string or return an empty string if it is impossible to reformat the string.
 *
 */

/* Submission by user `nhho` (1st) place */
#include<iostream>
#include<string>
using namespace std;

string reformat(string s) {
	int ints = 0, chars = 0;
	string sa = "", sb = "";
	for (char i: s) {
		if (i >= '0' && i <= '9')
			ints++, sa.push_back(i);
		else 
			chars++, sb.push_back(i);
	}


	if (abs(ints - chars) > 1)
		return "";
	if (ints > chars) 
		sa.swap(sb);
	string ans = "";
	for (int i = 0; i < s.size(); i++) {
		if (i % 2)
			ans.push_back(sa[i/2]);
		else
			ans.push_back(sb[i/2]);
	}
	return ans;
}


using namespace std;
int main() {
	string res;
	res = reformat("a0b1c2");
	cout << res << endl;
}
