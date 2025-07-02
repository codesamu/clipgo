# clipgo 

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)  

> Simple CLI clipboard manager for storing named snippets, listing them, and copying any item to your system clipboard.

## 📋 Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)

## ✨ Features

- Add named clipboard snippets with `add` command
- List stored snippets in a neat table
- Copy snippet value to clipboard by ID
- Remove snippets by ID
- Persistent JSON storage

## 🛠️ Technologies Used

- Go programming language
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard) for cross-platform clipboard support

## 🚀Installation

```bash
git clone https://github.com/codesamu/clipgo.git
cd clipgo
go build -o clipgo
```

## 🗃️Usage 
```bash
clipgo add "title" "Value"
clipgo list
clipgo copy <num>
cligo remove <num>
```
