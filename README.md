# steganography
主題：防止未公開海報外流

角色：小明、印刷廠、製片公司、電影導演

隱寫情境：小明是電影宣傳海報負責人。小明因為擔心海報原始檔再發布日前被外洩，所以小明在海報原始圖檔中埋下了隱藏訊息，再分別傳給印刷廠、製片公司跟電影導演...

隱寫方式：在要分別寄給印刷廠、製片公司和電影導演的海報圖檔中，分別隱寫了三種不同的qrcode，qrcode的內容對應到了(印刷廠、製片公司、電影導演)，這樣一旦外流的話，就可以知道是哪一方外流的。

解密情境/方式：在海報發布日之前，海報圖檔竟然被外流了！！小明下載被外流的圖檔(marked.png)並解密，找到底是誰外流的。

encrypt.go：
輸入 go run encrypt.go
輸出 三張被隱寫的海報

decrypt.go：
輸入 go run decrypt.go
輸出 解開隱寫後出現的盲水印
