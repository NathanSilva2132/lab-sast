# LABORATÓRIO DE SAST: código vulnerável de propósito. Não usar em produção.
import hashlib
import pickle
import sqlite3
import subprocess

from flask import Flask, request

app = Flask(__name__)

# P6: chave secreta fixa no código
app.secret_key = "chave-fixa-no-codigo"


# P1: injeção de SQL
@app.route("/usuario")
def usuario():
    nome = request.args.get("nome", "")
    conn = sqlite3.connect("lab.db")
    cur = conn.execute(f"SELECT id FROM usuarios WHERE nome = '{nome}'")
    return str(cur.fetchall())


# P-ok: consulta parametrizada, NÃO é falha (serve para medir falso positivo)
@app.route("/usuario-seguro")
def usuario_seguro():
    nome = request.args.get("nome", "")
    conn = sqlite3.connect("lab.db")
    cur = conn.execute("SELECT id FROM usuarios WHERE nome = ?", (nome,))
    return str(cur.fetchall())


# P2: injeção de comando
@app.route("/ping")
def ping():
    host = request.args.get("host", "")
    return subprocess.check_output("ping -c 1 " + host, shell=True)


# P3: execução de código arbitrário
@app.route("/calc")
def calc():
    return str(eval(request.args.get("expr", "1+1")))


# P4: desserialização insegura
@app.route("/carregar", methods=["POST"])
def carregar():
    obj = pickle.loads(request.data)
    return str(obj)


# P5: MD5 para senha
def hash_senha(senha):
    return hashlib.md5(senha.encode()).hexdigest()


# P7: debug ligado e escutando em todas as interfaces
if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)
