# APIの概要
Wikipediaにおいてあるページからあるページに6ステップ以内で到達可能かどうかを判定するAPI
可能な場合はルートを返す。

# 環境
Go
DB: Mysql
Docker を利用
Framework: Echo

# 処理
出発点のページのID(number)と到着点のページのID(number)を受け取り、
出発点から到着点に到達可能かどうかを判定するし、ルートを返す。
探索は双方向 BFS で行い計算量を抑える。

# データの形式
pagelinks というテーブルに pl_from と pl_target_id というカラムがある。
pl_from はページのIDで、pl_target_id はリンク先のページのIDである。

