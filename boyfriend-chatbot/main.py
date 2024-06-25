from flask import Flask, request, jsonify
from load_env import get_environment, next_tokne
from collection import Collection
from inference import get_model, get_next_text

ip, port, user, pw, db = get_environment()
coll = Collection(ip, port, user, pw, db)
coll.open()
model = get_model()

app = Flask(__name__)

@app.route('/request/next_sentence', methods=['POST'])
def json_example():
    global coll, model, jubu_id
    if request.is_json:
        data = request.get_json()
        size = data.get('size', 3)
        jubu_id = data.get('jubu_id', -1)

        texts = coll.get_text()
        texts = reversed(texts)
        texts = next_tokne(texts)
        result = get_next_text(jubu_id, texts, model, 3)
        r = result.split("\n")
        r = list(filter(lambda x: x != "", r))
        r = list(map(lambda x: x[len('(00:00:00) 차차 :'):].strip(), r))
        re = "\n".join(r)
        return jsonify(message=f"{re}")
            
    else:
        return jsonify(message="Request was not JSON"), 400

if __name__ == '__main__':
    app.run(port=5000, host='0.0.0.0')
