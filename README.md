# TankSurvival2

Defoldを使用して個人制作した、2D戦車シューティングゲームです。

プレイヤーは3種類の機体から1つを選択し、
ステージ内の敵を倒しながら全10ステージのクリアを目指します。

ブラウザ版を公開しているため、インストール不要でプレイできます。

---

## 目次

- [概要](#概要)
- [紹介動画](#紹介動画)
- [操作方法](#操作方法)
- [工夫した点](#工夫した点)
- [遊び方](#遊び方)
- [開発者・コンタクト](#開発者コンタクト)

---

## 概要

| 項目 | 内容 |
| --- | --- |
| タイトル | TankSurvival2 |
| ジャンル | 2D戦車シューティング |
| プレイ人数 | 1人 |
| プラットフォーム | Web / Windows |
| 使用ゲームエンジン | Defold |
| 使用言語 | Lua |
| 制作形態 | 個人制作 |
| 制作期間 | 2019/10～2020/3 |
| ステージ数 | 全10ステージ |

---

## 紹介動画

ゲームの内容を約1分で紹介しています。

[![TankSurvival2 紹介動画](https://img.youtube.com/vi/vZmOgyaHwis/maxresdefault.jpg)](https://youtu.be/vZmOgyaHwis)

[YouTubeで紹介動画を見る](https://youtu.be/vZmOgyaHwis)

---

## 操作方法

| 操作 | キー |
| --- | --- |
| 移動 | W / A / S / D |
| 照準 | マウス |
| 攻撃 | 左クリック |
| フルスクリーン | F11 |
| ゲーム終了 / タイトルへ戻る | ESC |

※機体タイプによって使用できる能力が異なります。

---

## 工夫した点

### 3種類のプレイスタイル

本作では、性能の異なる3種類の機体を用意しました。

#### スピードタイプ

高速な移動を活かして敵の攻撃を避けながら戦うタイプです。

<img width="356" alt="スピードタイプ" src="https://github.com/user-attachments/assets/d7ec7d7f-7ed3-4510-bb7f-8c2c2a65fd88" />

#### バランスタイプ

基本性能のバランスが良く、
ステージ中に取得できる強化要素を利用して戦うタイプです。

<img width="299" alt="バランスタイプ" src="https://github.com/user-attachments/assets/a682ff44-a3ac-41ea-8dfe-26e36f1ce7f8" />

#### パワータイプ

攻撃をチャージすることで、
通常より強力な攻撃を行えるタイプです。

<img width="315" alt="パワータイプ" src="https://github.com/user-attachments/assets/a9610524-95cf-4bd5-bded-731ba5505283" />

---

### 敵・ボスの行動制御

敵ごとに移動速度や耐久力、攻撃方法を変更し、
ステージが進むにつれて異なる戦い方が必要になるよう設計しました。

また、敵の移動方向や壁との衝突処理などもLuaで実装しています。

<img width="612" alt="敵との戦闘" src="https://github.com/user-attachments/assets/51d57207-3c78-41a6-8ec3-4bb427a54901" />

---

### HUDによる状態表示

HPだけでなく、選択した機体に応じて
チャージゲージや強化状態など必要な情報を表示するHUDを実装しました。

<img width="786" alt="HUD" src="https://github.com/user-attachments/assets/e17a5f19-3071-4585-a5b1-db347a29501c" />

<img width="787" alt="HUD" src="https://github.com/user-attachments/assets/e4ffc537-8630-4f88-82bc-a77d6f5a105f" />

---

## 遊び方

### ブラウザ版

itch.ioで公開しています。

▶ [ブラウザでTankSurvival2をプレイ](https://miyasin0422.itch.io/tank-survival2)

インストール不要でそのままプレイできます。

### Windows版

GitHub ReleasesからZIPファイルをダウンロードできます。

▶ [Windows版をダウンロード](https://github.com/miyasin0422/Tank-Survival2/releases)

1. ZIPファイルをダウンロード
2. ファイルを解凍
3. `TankSurvival2.exe` を起動

---

## 開発者・コンタクト

### GitHub

[GitHub - miyasin0422](https://github.com/miyasin0422)

### X

[X - @miyashin_lab](https://x.com/miyashin_lab)

### YouTube

[YouTube - miyashin_games](https://www.youtube.com/@miyashin_games)

---

## License

This project was created as a personal game development project.
