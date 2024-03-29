import os, requests
import qrcode_terminal

server_url = "http://192.168.88.88:8080"
dev = False

def  query_create_meeting():
    if dev :
        session_id = "mock_session_id"
        session_key = "mock_sessino_key"
        print("query_create_meeting", session_id, session_key)
        return session_id, session_key
    r = requests.post(server_url + "/v1/meeting/")
    print("query_create_meeting", r.json()['session_id'], r.json()['session_key'])
    return r.json()['session_id'], r.json()['session_key']

def enc_recn(session_key):
    if dev:
        print("enc_recn")
        return
    file = "../upload/d06c56b5-0622-466b-b9db-3c18bfc5e3ed/recndec"
    os.system("../aes " + session_key + " " + file)
    print("enc_recn")

def upload_recn(session_id, session_key):
    if dev:
        print("upload_recn")
        return
    file = session_key
    r = requests.post(
        server_url + "/v1/meeting/"+session_id+"/rec/recn",
        files={"file": open(file, 'rb')}
    )
    print("upload_recn", r.text)

def upload_recj(session_id):
    if dev:
        print("upload_recj")
        return
    file = "../upload/d06c56b5-0622-466b-b9db-3c18bfc5e3ed/recj"
    r = requests.post(
        server_url + "/v1/meeting/"+session_id+"/rec/recj",
        files={"file": open(file, 'rb')}
    )
    print("upload_recj", r.text)

def query_end_reg(session_id):
    if dev:
        print("query_end_reg")
        return
    r = requests.post(server_url + "/v1/meeting/"+session_id+"/end")
    print("query_end_reg", r.text)

def start_rec():
    # TODO
    print("start_rec")

def end_rec():
    # TODO
    print("end_rec")

def new_session():
    session_id, session_key = query_create_meeting()
    enc_recn(session_key)
    upload_recn(session_id, session_key)
    os.system("rm " + session_key)
    qrcode_terminal.draw(server_url + "/app/" +session_id) # TODO
    return session_id

def start_session(session_id):
    query_end_reg(session_id)
    start_rec()

def end_session(session_id):
    end_rec()
    upload_recj(session_id)

if __name__ == '__main__':
    if dev :
        print("DEV MODE")

    input("\nCreate Meeting Session, Press key...") # TODO
    session_id = new_session()

    input("\nStart Meeting Session, Press key...") # TODO
    start_session(session_id)

    input("\nEnd Meeting Session, Press key...") # TODO
    end_session(session_id)
