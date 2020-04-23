/*
 * Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).
 *
 * You have to find a permutation of the string where no letter is followed by another
 * letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.
 *
 * Return the reformatted string or return an empty string if it is impossible to reformat the string.
 *
 */

// Submission by originally user `nhho` (1st) place
//
// Minor modifications done by author to understand how it works
#include<iostream>
#include<string>
using namespace std;

string reformat(string s) {
	int ta = 0, tb = 0;
	string sa = "", sb = "";
	// put characters and numbers in separate strings
	for (char i: s) {
		if (i >= '0' && i <= '9')
			ta++, sa.push_back(i);
		else 
			tb++, sb.push_back(i);
	}

	// if there are more than 1 number or character
	// offset, at least one will be adjacent to some each other
	if (abs(ta - tb) > 1) // shown in test case 6
		return "";
	if (ta > tb) // shown in test case 5
		sa.swap(sb);
	string ans = "";
	for (int i = 0; i < s.size(); i++) 
		if (i % 2) {
			ans.push_back(sa[i/2]);
		}
		else {
			ans.push_back(sb[i/2]);
		}
	
	return ans;
}


void outputResult(string input, string output) {
	cout << "Input: " << input << "\n" << "Output: " << output << endl;
}

using namespace std;
int main() {
	string case1 = "a0b1c2", res1 = "";
	res1 = reformat(case1); // same as case 1 but valid in the contest for some reason
	outputResult(case1, res1);

	string case2 = "leetcode", res2 = "";
	res2 = reformat(case2);
	outputResult(case2, res2);

	string case3 = "1229857369", res3 = "";
	res3 = reformat(case3);
	outputResult(case3, res3);

	string case4 = "covid2019", res4 = "";
	res4 = reformat(case4);
	outputResult(case4, res4);

	string case5 = "ab123", res5 = "";
	res5 = reformat(case5);
	outputResult(case5, res5);

	string case6 = "aaaa12", res6 = "";
	res6 = reformat(case6);
	outputResult(case6, res6);
}
