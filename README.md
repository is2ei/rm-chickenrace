rmコマンドチキンレース!
===

チキチキ！ rmコマンドチキンレースの時間です！

# 注意！
**これは危険なゲームです！！！！！！！！！！**   
**よくわかっていない人は実施しないでください！！！！！！！！**

# 概要
* 2人で行う対戦型ゲーム！
* 仮想マシンの上で root ユーザーになり、交互にファイルを削除していきます！
* 削除には、特製の「競技用rmコマンド」を利用します
  * 使用例: `/rm-user1 /opt`
  * ※オプションなしで `rm -fr` 相当の挙動を示します
* 削除した内容は Twitter 及び Typetalk へ通知されます
* 削除に失敗した場合 or 削除できるものがなくなった場合、負けになります！

## ルール詳細
### ゲームの流れ
* root ユーザーで仮想マシンへログインします
* ユーザーのターン開始
  * マシンの状態などを調査します
  * 自分用の競技用rmコマンドを使用し、ファイルを削除します
* ターンエンド、操作者変更

競技用rmコマンドを実行するとターンエンドです。  
それまでの間、調査は好きなだけ行って構いません。

### 勝敗条件
以下のケースで負け判定となります！
* 削除に失敗した場合(rootユーザーでも削除できない場合)
* 削除してはいけないファイル(後述)を削除した場合
* ファイル削除後、SNS連携できなくなった場合

### 削除してはいけないファイル
以下のファイルを削除してはいけません。
* `/rm*` (競技用rmコマンド)
* `/rm.twitter.conf`, `/rm.typetalk.conf` (SNS連携用設定ファイル) 

## 準備
実施に先立ち、下記の準備が必要です
* 競技用rmコマンドの設置
  * 競技用rmコマンドを、`/rm-user1`, `/rm-user1` のように実行ユーザーがわかるようにして `/` 配下にコピーします
* SNS連携用設定ファイルの設置
  * `/rm.twitter.conf` (Twitter 連携用)
  * `/rm.typetalk.conf` (Typetalk 連携用)

## 注意点！
* 実施環境は、壊しても良い、完全仮想化環境(Vagrant や VMWare)で実施することを強く推奨します。
    * コンテナなどを使うのは、ホストに影響を及ぼす可能性があるため **禁止!!**