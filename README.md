# 🚀 Custom Unix-Like Shell

A **Unix-like shell implemented in Go**, built from scratch as a coding challange to me.  
This project explores **system programming, processes, and concurrency** by implementing core shell features step by step.

---

## ✨ Features (Challenge Steps)

1. ✅ Run external commands (e.g., `ls`, `pwd`)  
2. ✅ Implement `exit` built-in command  
3. ✅ Implement `cd` built-in command  
4. ✅ Implement `echo` built-in command  
5. ✅ Implement `export` and `unset` environment variable commands  
6. 🚧 **Pipes (`|`) support** *(Work in Progress)*  
7. 🚧 **Input/Output redirection (`>`, `<`)** *(Work in Progress)*  
8. 🚧 **Command history** *(Work in Progress)*  

---

## 🛠️ Tech Stack

- **Language:** Go  
- **Concepts:** Concurrency, System Programming, Process Management, Environment Handling  

---

## 🚀 Getting Started

### Prerequisites
- Install [Go 1.20+](https://go.dev/dl/)  
- Install [Git](https://git-scm.com/)  
- ⚠️ For best results, run this project in a **Linux environment**.  
  - External commands (like `ls`, `pwd`, etc.) may not work correctly on Windows.  

### Clone the repository
```bash
git clone https://github.com/Habiba-Adel/Building_My_Own_Shell.git
cd Building_My_Own_Shell
```

## Run the shell 
``` bash
go run main.go
```
Now you will see a prompt like this :
```bash
🚀<Habiba_Shell>:/home/user $
```
## 🖥️ Usage
You can now type commands directly into your shell:
## Built-in commands:
```bash
cd Documents                # change directory
export NAME=Habiba          # set environment variable
echo $NAME                  # print variable
unset NAME                  # remove variable
exit                        # close the shell
```
## External commands:
``` bash
ls                          # list files
pwd                         # show current directory
cat file.txt                # print file contents
```
---

💡 This project is being developed with dedication and curiosity — a valuable journey to explore system programming concepts, gain practical experience, and enjoy building something from scratch ❤️.

