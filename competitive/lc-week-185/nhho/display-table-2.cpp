/*
 *
 * Given the array orders, which represents the orders that customers have done in a restaurant.
 * More specifically orders[i]=[customerNamei,tableNumberi,foodItemi] where customerNamei is the
 * name of the customer, tableNumberi is the table customer sit at, and foodItemi is the item customer orders.
 *
 * Return the restaurant's “display table”. The “display table” is a table whose row entries denote how many
 * of each food item each table ordered. The first column is the table number and the remaining columns correspond
 * to each food item in alphabetical order. The first row should be a header whose first column is “Table”, followed
 * by the names of the food items. Note that the customer names are not part of the table. Additionally, the rows should
 * be sorted in numerically increasing order.
 *
 * Example 1:
 *
 * Input: orders = [["Laura","2","Bean Burrito"],["John","2","Beef Burrito"],["Melissa","2","Soda"]]
 * Output: [["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]
 *
 *
 * Example 2:
 *
 * Input: orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],
 *                 ["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]]
 * Output: [["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]] 
 *
 * Explanation: 
 * For the table 1: Adam and Brianna order "Canadian Waffles".
 * For the table 12: James, Ratesh and Amadeus order "Fried Chicken".
 *
 *
 * Constraints:
 *
 *     1 <= orders.length <= 5 * 10^4
 *     orders[i].length == 3
 *     1 <= customerName_i.length, foodItem_i.length <= 20
 *     customerName_i and foodItem_i consist of lowercase and uppercase English letters and the space character.
 *     tableNumber_i is a valid integer between 1 and 500.
 *
 */

/* Submission by user `nhho` (1st place) */
#include<iostream>
#include<string>
#include<set>
#include<vector>
#include<map>
using namespace std;

// receives *[][]string, returns [][]string
vector<vector<string>> displayTable(vector<vector<string>>& orders) {
	// sorted set of unique objects of type string
	set<string> fd;
	// key int, value: map[string]int
	map<int, map<string, int>> m;

	for (auto& i: orders) {
		// insert the item ordered in the set
		fd.insert(i[2]);
		// add the table number as a key and increment the item ordered in the val mapping
		m[stoi(i[1])][i[2]]++;
	}

	vector<vector<string>> ans(1);
	ans[0].push_back("Table");
	ans[0].insert(ans[0].end(), fd.begin(), fd.end());
	for (auto& i: m) {
		// TODO: understand the difference between `push_back` and `emplace_back`
		// TODO: understand what calling this without args mean
		ans.emplace_back();
		// TODO: understand this operation
		ans.back().push_back(to_string(i.first));
		for (auto& j: fd)
			ans.back().push_back(to_string(i.second[j]));
	}
	return ans;
}

using namespace std;
int main() {
	vector<vector<string>> order{ {"Laura","1","Burrito"}, {"John","2","Burrito"}, {"Melissa","2","Soda"} };
	vector<vector<string>> result;

	result = displayTable(order);
	
	for (int i = 0; i < result.size(); i++) {
		for (int j = 0; j < result[0].size(); j++) {
			cout << result[i][j] << " ";
		}
		cout << endl;
	}

	return 0;
}



