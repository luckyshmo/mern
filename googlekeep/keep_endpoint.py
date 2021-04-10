from flask import Flask, jsonify, request
from flask_json import FlaskJSON, JsonError, json_response, as_json

import gkeepapi

# # Find by string
# gnotes = keep.find(query='Title')

# # Find by filter function
# gnotes = keep.find(func=lambda x: x.deleted and x.title == 'Title')

# # Find by labels
# gnotes = keep.find(labels=[keep.findLabel('todo')])

# # Find by colors
# gnotes = keep.find(colors=[gkeepapi.node.ColorValue.White])

# # Find by pinned/archived/trashed state
# gnotes = keep.find(pinned=True, archived=False, trashed=False)

def getFromKeep(email,token,name):
    keep = gkeepapi.Keep()
    success = keep.login(email, token)

    gnotes = keep.find(query=name)
    for note in gnotes:
        gnote = note
    # noda = nodes['177b5438cd8.82b17d320d61d816']

    keep.sync()
    return gnote.text.split('\n')

app = Flask(__name__)
FlaskJSON(app)

# curl -X GET --data '{"email":"mishka2017@gmail.com","token":"iopnilguhgigbbht","name":"English words"}' http://localhost:5001/get_words
# curl -X GET --data '{"email":"mishka2017@gmail.com","token":"iopnilguhgigbbht","name":"English words"}' http://localhost:3000/api/keep

@app.route('/get_words', methods=['GET'])
@as_json
def get_value():
    data = request.get_json(force=True)

    email = (data['email'])     #mishka2017@gmail.com
    token = (data['token'])     #iopnilguhgigbbht
    note_name = (data['name'])  #english words

    noda = getFromKeep(email, token, note_name)
    return dict(value=noda)

if __name__ == '__main__':
    app.run(port=5001) #TODO config
