import os, requests, subprocess, signal
import qrcode
import RPi.GPIO as GPIO
from PIL import Image,ImageDraw,ImageFont
import epd2in7

font = ImageFont.truetype(os.path.join("pi", 'Font.ttc'), 24)

server_url = "http://192.168.88.88:8080"
dev = False
btn_pin = 23
light_pin = 27
jammer_pin = 16



def  query_create_meeting():
    if dev :
        session_id = "mockuuid-uuid-uuid-uuid-uuidmockuuid"
        session_key = "mocksessinokeymocksessinokeymock"
        print("query_create_meeting", session_id, session_key)
        return session_id, session_key
    r = requests.post(server_url + "/v1/meeting/")
    print("query_create_meeting", r.json()['session_id'], r.json()['session_key'])
    return r.json()['session_id'], r.json()['session_key']

def enc_recn(session_key):
    if dev:
        print("enc_recn")
        return
    file = "recn.wav"
    os.system("./aes " + session_key + " " + file)
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

def remove_recn(session_key):
    if dev:
        print("remove_recn")
        return
    os.system("rm " + session_key)

def upload_recj(session_id):
    if dev:
        print("upload_recj")
        return
    file = "recj.wav"
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
    # arecord -r 44100 -f FLOAT_LE -c 2 -t wav -v recj.wav
    proc_args = ['arecord','-r', '44100', '-f', 'FLOAT_LE', '-c', '2', '-t', 'wav', 'recj.wav']
    rec_proc = subprocess.Popen(proc_args, shell=False, preexec_fn=os.setsid)
    print("start_rec", rec_proc.pid)
    return rec_proc

def end_rec(rec_proc):
    os.killpg(rec_proc.pid, signal.SIGTERM)
    rec_proc.terminate()
    print("end_rec", rec_proc.pid)
    rec_proc = None

def start_jammer():
    GPIO.output(jammer_pin, GPIO.HIGH)
    print("start_jammer")

def end_jammer():
    GPIO.output(jammer_pin, GPIO.LOW)
    print("end_jammer")

def genqrcode(url):
    qr = qrcode.QRCode(
        version=None,
        error_correction=qrcode.constants.ERROR_CORRECT_M,
        box_size=4,
        border=0,
    )
    qr.add_data(url)
    qr.make(fit=True)
    img = qr.make_image(fill_color="black", back_color="white")
    qrfile = "qr.bmp"
    img.save(qrfile)
    return qrfile

def gen_frame(show_text):
    page = Image.new('1', (epd.width, epd.height), 255)
    draw = ImageDraw.Draw(page)
    w, h = draw.textsize(show_text, font=font)
    draw.text(((epd.width-w)/2, (epd.height-epd.width-h)/2), show_text, font=font, fill = 0)
    return page

def screen_text(show_text):
    frame = gen_frame(show_text)
    epd.display(epd.getbuffer(frame))

def screen_qr(show_text, qrfile):
    page = gen_frame(show_text)
    bmp = Image.open(qrfile)
    bmp = bmp.resize((epd.width-10, epd.width-10))
    page.paste(bmp, (5, epd.height - epd.width + 5))
    epd.display(epd.getbuffer(page))

def screen_clear():
    page = Image.new('1', (epd.width, epd.height), 255)
    epd.display(epd.getbuffer(page))

def new_session():
    screen_text("Init  Session")
    session_id, session_key = query_create_meeting()
    enc_recn(session_key)
    upload_recn(session_id, session_key)
    remove_recn(session_key)
    screen_qr("Reg  Owner", genqrcode(server_url + "/app/" +session_id))
    return session_id

def start_session(session_id):
    screen_text("Start  Session")
    start_jammer()
    query_end_reg(session_id)
    rec_proc = start_rec()
    GPIO.output(light_pin, GPIO.HIGH)
    return rec_proc

def end_session(session_id, rec_proc):
    screen_text("End  Session")
    GPIO.output(light_pin, GPIO.LOW)
    end_rec(rec_proc)
    end_jammer()
    upload_recj(session_id)
    screen_clear()

def setup_gpio():
    GPIO.cleanup(light_pin)
    GPIO.setup(light_pin, GPIO.OUT)
    GPIO.output(light_pin, GPIO.LOW)

    GPIO.cleanup(jammer_pin)
    GPIO.setup(jammer_pin, GPIO.OUT)
    GPIO.output(jammer_pin, GPIO.LOW)

    GPIO.cleanup(btn_pin)
    GPIO.setup(btn_pin, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.add_event_detect(btn_pin, GPIO.FALLING, bouncetime=200)

if __name__ == '__main__':
    if dev :
        print("DEV MODE")

    epd = epd2in7.EPD()
    epd.init()
    setup_gpio()
    
    while True:
        screen_text("Meeting  Box")

        while not GPIO.event_detected(btn_pin):
            pass
        print("\n--- Init Meeting Session ---")
        session_id = new_session()
    
        while not GPIO.event_detected(btn_pin):
            pass
        print("\n--- Start Meeting Session ---")
        rec_proc = start_session(session_id)

        while not GPIO.event_detected(btn_pin):
            pass
        print("\n--- End Meeting Session ---")
        end_session(session_id, rec_proc)