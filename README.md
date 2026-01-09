# Pomidor programming language🍅

Pomidor is an **experimental programming language** where you grow and harvest your code like a garden.  
Everything you write in Pomidor **works exactly as written**, with no hidden magic.  

---

## Philosophy

1. **Explicit and self-explanatory code**  
   Every line of Pomidor code should **explain itself**.  

2. **Code works exactly as written**
   What you type is what the program does — no surprises.  

---

## File Extension

All Pomidor programs use the `.pmd` extension.

---

## Running Pomidor Programs 🍅

You need the **Pomidor Interpreter** to run `.pmd` files.

### Using Python directly

```bash
python3 pomidor_interpreter.py <file.pmd>
```


## Using the pomidor alias (recommended):

To avoid typing python3 `pomidor_interpreter.py` every time, you can create an alias in your shell:


```bash
alias pomidor="python3 /full/path/to/pomidor_interpreter.py"
```
### Then run any .pmd file like this:

```bash
pomidor examples/hello_v1.pmd
```

### Output:

```bash
🍅 Pomidor Language Interpreter v0.1
Running file: examples/hello_v1.pmd

Hello, Pomidor!

```

## Pomidor interpreter now in GO

You can try out new Pomidor interpreter written in Go using the following commands:

1. Build
    ```bash
    cd interpreter && go build -o ../pomidor && cd ..
    ```
2. Run
    ```bash
    pomidor examples/hello_v2.pmd
    ```
