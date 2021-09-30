
int main(){

  string fn, sn;

  int score = 0;

  cin >> fn;
  cin >> sn;

  int fnPosTrack[100];
  int snPosTrack[100];

  for (int i = 0; i < fn.length(); i++){
    fnPosTrack[i] = 1;
  }

  for (int i = 0; i < sn.length(); i++){
    snPosTrack[i] = 1;
  }

  for (int i = 0; i < fn.length();i++){
    for (int j = 0; j < sn.length();j++){

      if(fn[i] == sn[j]){
        if(fnPosTrack[i] != -1 && snPosTrack[j] != -1){
          if(i == j){
            score += 2;
          }else{
            score += 1;
          }

          fnPosTrack[i] = -1;
          snPosTrack[j] = -1;
        }
      }

    }
  }

  cout << score;
}