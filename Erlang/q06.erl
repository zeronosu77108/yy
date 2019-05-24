%% コラッツの予想（改良版）
%%
%% 自然数n について
%% * n が偶数の場合 n を 2 で割る
%% * n が奇数の場合 n に 3 をかけて 1 を足す
%% 初回のみ n に 3 をかけて 1 足す． 
%% 10000以下の偶数のうち「最初の数に戻る数」がいくつあるか．
-module(q06).
-export([main/0]).

main() -> collatz(10000).

collatz(0) -> 0;
collatz(X) -> collatz_(X*3+1, X) + collatz(X-2).

collatz_(1, _) -> 0;
collatz_(S, S) -> 1;
collatz_(N, S) when N rem 2 =:= 0 -> collatz_(N div 2, S);
collatz_(N, S) when N rem 2 =/= 0 -> collatz_(N*3+1, S).