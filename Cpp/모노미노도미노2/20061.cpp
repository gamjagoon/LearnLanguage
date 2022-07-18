/*
Date : 07/18/2020
problem : 모노미노도미노2
summary : 구현
*/
#include <iostream>
#include <vector>
#include <cstring>

#define endl '\n'
using namespace std;

//Global
int B_arr [4][6];
int G_arr [4][6];
void Array_view(int(*arr)[6], int r, int c) {
    cout<<endl;
    for (int i = 0; i < r; i++) { 
        for (int j = 0; j < c; j++) {
            cout << arr[i][j] << " ";
        }cout<<endl;
    }
}

// t : 1 1x1
//   : 2 1x2 
//   : 3 2x1
void push_block(int(*arr)[6], int sr, int t){
    int ec = 5;
    for (int c = 2; c < 6; c++) {
        if (arr[sr][c] > 0) {
            ec = c - 1;
            break;
        }
        if (t == 3 && arr[sr+1][c] > 0) {
            ec = c - 1;
            break;
        }
    }
    switch (t)
    {
    case 1:
        arr[sr][ec] = 1;
        break;
    case 2:
        arr[sr][ec] = 1;
        arr[sr][ec-1] = 1;
        break;
    default:
        arr[sr][ec] = 1;
        arr[sr+1][ec] = 1;
        break;
    }
}

int remove_column(int(*arr)[6]){
    for (size_t c = 2; c < 6; c++) {
        bool flag = true;
        for (size_t i = 0; i < 4; i++) {
            if(arr[i][c] == 0){
                flag = false;
                break;
            }
        }
        if(flag){
            for (int n = c; n > 0; n--) {
                arr[0][n] = arr[0][n-1];
                arr[1][n] = arr[1][n-1];
                arr[2][n] = arr[2][n-1];
                arr[3][n] = arr[3][n-1];
            }
            return 1;
        }
    }
    return 0;
}

int arr_cnt(int(*arr)[6]){
    int ret = 0;
    for (int i = 0; i < 4; i++) {
        for (int j = 0; j < 6; j++) {
            if (arr[i][j]) {
                ret++;
            }
        }
    }
    return ret;
}


void remove_overflow(int(*arr)[6]){
    int flag = 0;
    for (int c = 0; c < 2; c++) {
        for (int r = 0; r < 4; r++) {
            if (arr[r][c]) {
                flag = 1;
                break;
            }
        }
        if(flag > 0){
            if(c == 0)flag = 2;
            else flag = 1;
            for (int i = 0; i < flag; i++)
            {
                for (int n = 5; n > 0; n--) {
                    arr[0][n] = arr[0][n-1];
                    arr[1][n] = arr[1][n-1];
                    arr[2][n] = arr[2][n-1];
                    arr[3][n] = arr[3][n-1];
                }
                arr[0][0] = 0; 
                arr[1][0] = 0; 
                arr[2][0] = 0; 
                arr[3][0] = 0; 
            }
            break;
        }
    }
}

int main() {
  ios::sync_with_stdio(false);cin.tie(0);
  memset(B_arr,0, sizeof(B_arr));
  memset(G_arr,0, sizeof(G_arr));
  int q, t, x, y;
  int score = 0; 
  cin>>q;
  for (int i = 0; i < q; i++)
  {
    cin>>t>>x>>y;
    push_block(B_arr, x, t);
    score += remove_column(B_arr);
    if (t == 2) {
        score += remove_column(B_arr);
    }
    remove_overflow(B_arr);
    
    if (t == 2) t = 3;
    else if (t == 3) t = 2;
    push_block(G_arr, y, t);
    score += remove_column(G_arr);
    if (t == 2) {
        score += remove_column(G_arr);
    }
    remove_overflow(G_arr);
  }
  cout<<score<<endl;
  cout<<arr_cnt(B_arr) + arr_cnt(G_arr);
}