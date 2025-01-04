# 啦八機 | LaBaG
Go 語言版本 | Golang version

  ```plaintext
  ├── LaBaG.go          # LaBaG 遊戲 (Go 語言版本) 只能在終端機遊玩
  |                       LaBaG game (Go language version) Only be played in the terminal
  ├── TargetJson.go     # 產生用於 LaBaG(Python版本) 中模擬一局的 .json 檔案
  |                       Generate a .json file for simulating a game in LaBaG (Python version)
  └── README.md         # 說明文件 
                          Documentation
  ```

## 使用 | Usage

### LaBaG.go
遊玩一局啦八機遊戲
Play a round of LaBaG game

- 直接運行 | Run

  ```bash
  go run LaBaG.go  
  ```

- 建置後運行 | Build

  ```bash
  go build LaBaG.go  
  .\LaBaG
  ```

### TargetJson.go
產生用於 LaBaG(Python版本) 中模擬一局的 .json 檔案 
Generate a .json file for simulating a game in LaBaG (Python version)

1. 建置後運行 | Build

  ```bash
  go build TargetJson.go  
  .\TargetJson
   ```

2. 輸入目標 | Input the target

3. 等待達到目標後將產生 `.json` 檔案 | Wait until the target is reached to generate a `.json` file

4. 在 LaBaG(Python版本) 將 `.json` 檔案的路徑輸入 | Input the path of the `.json` file in LaBaG (Python version)

- [LaBaG-PythonProject](https://github.com/Yorn90104/LaBaG-PythonProject.git)

5. 根據指引進行遊戲 | Follow the prompts to play the game.