import os, requests, subprocess, signal
import qrcode
import RPi.GPIO as GPIO
from PIL import Image,ImageDraw,ImageFont
import epd2in7

font = ImageFont.truetype(os.path.join("pi", 'Font.ttc'), 24)

dev = False
btn_pin = 23
light_pin = 27
jammer_pin = 16

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

def start_session():
    screen_text("Start  Session")
    start_jammer()
    rec_proc = start_rec()
    GPIO.output(light_pin, GPIO.HIGH)
    return rec_proc

def end_session(rec_proc):
    screen_text("End  Session")
    GPIO.output(light_pin, GPIO.LOW)
    end_rec(rec_proc)
    end_jammer()
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
        print("\n--- Start Meeting Session ---")
        rec_proc = start_session()

        while not GPIO.event_detected(btn_pin):
            pass
        print("\n--- End Meeting Session ---")
        end_session(rec_proc)