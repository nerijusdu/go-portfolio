# Repurpose an old landline phone with a Raspberry Pi

## Prelude

Once upon a time in land far far away there was a small startup selling UGC videos. They had a platform and an good old boring release pipeline. Amidst the monotony of their reliable release pipeline, a spark of innovation was kindled. They made a [shinny big red button](/blog/overengineer-a-button) to deploy to production, bringing a touch of excitement to their routine.

And everything was well in the startup land. However, as the team decided to shift from mass releases to a more strategic, one-by-one approach for their microservices, the button was not enough anymore. Suddenly the team was again stuck with the boring old software buttons and the release process was not as exciting as it used to be.

Frustrated with the limitations of their once-beloved big red button, the team found themselves yearning for a more dynamic and engaging release process. In an ordinary 1-on-1 meeting, a whimsical idea emerged: what if they used an old phone to trigger code releases? Initially dismissed as a joke, the notion lingered. A week later, the team stumbled upon an old Winnie the Pooh phone while browsing online. Recognizing the potential, they haggled for the nostalgic relic, ultimately paying half the asking price. Little did they know that this decision would not only transform their release process but also infuse an unexpected dose of childhood whimsy into their high-tech startup environment.

## The idea

Enough with the fairy tale, let's talk about the project. The main idea is to rework the [old release button](/blog/overengineer-a-button) to a phone. The phone will trigger an Argo CD release the same way the button did, but it will also have an input keypad to select the service to release and an output speaker to announce the release status.

## The speaker

Let's start with the speaker as it was the easiest to implement. To wire up the speaker I used an generic USB to AUX adapter that had two inputs: one for the microphone and one for the speaker.

<div class="img-lg">
![Phone receiver with USB adapter](/data/images/phone/speaker.jpg)
</div><!---->

There were 4 wires coming from the phones receiver, I just had to figure out which two were for the speaker and which two were for the microphone (you can disassemble the receiver to see which wires go where). And then solder them to the adapter contacts (I just tested all the combinations until the speaker started making sounds lol).

<div class="img-sm">
![USB AUX adapter soldering](/data/images/phone/speaker-adapter.jpg)
</div><!---->

Then it was as easy as plugging in the USB adapter to the Raspberry Pi. I used AWS Polly for text-to-speech and pyaudio to play the audio. I'm not going to get into the code but if you are interested you can check it out on my [GitHub](https://github.com/nerijusdu/release-button).

## The keypad

The keypad was a bit more complicated. There are 14 buttons and 12 wires coming from the keypad. What I did was disassemble the keyboard and traced the connections from each button to the wires. There are 2 connections going to each button, sometimes they go through other buttons, but they eventually connect to the output wires.

<div class="img-md">
![Phone keypad](/data/images/phone/keypad.jpg)
</div><!---->

The diagram of the keypad connections looks like this (excuse my drawing skills):

<div class="img-md">
![Phone keypad diagram](/data/images/phone/keypad-diagram.jpg)
</div><!---->

Then I connected all the wires to Raspberry Pi GPIO pins. And wrote some python code to control it. The code is pretty simple, I have a map of all the buttons and their corresponding 2 connections. I run an infinite loop, through all the buttons sending a pulse to one of the connections, if I receive the puls on the other connection I know the button was pressed.

```python
GPIO.setmode(GPIO.BCM)

# pins1 = [40,38,36,32,26,24] # Board
pins1 = [21,20,16,12,7,8] #BCM
# pins2 = [23,29,31,33,35,37] # Board
pins2 = [11,5,6,13,19,26] # BCM

buttons = {
  1: [pins1[2], pins2[4]],
  4: [pins1[3], pins2[3]],
  9: [pins1[4], pins2[4]],
  2: [pins1[0], pins1[2]],
  5: [pins1[3], pins2[4]],
  3: [pins1[1], pins1[2]],
  8: [pins1[4], pins2[3]],
  6: [pins1[3], pins2[5]],
  7: [pins1[1], pins1[3]],
  0: [pins1[4], pins2[2]],
  '#': [pins1[5], pins2[1]],
  '*': [pins1[1], pins1[4]],
  'R': [pins1[5], pins2[4]],
  'Redial': [pins2[0], pins2[1]],
}

buttonKeys = list(buttons.keys())

def listen_to_keypad():
  while True:
    for key in buttonKeys:
      GPIO.setup(buttons[key][0], GPIO.OUT)
      GPIO.setup(buttons[key][1], 
        GPIO.IN, 
        pull_up_down=GPIO.PUD_DOWN)

      GPIO.output(buttons[key][0], GPIO.HIGH)

      if GPIO.input(buttons[key][1]) == GPIO.HIGH:
        print('Button pressed: ' + str(key))
        # avoid long press triggering multiple times
        sleep(0.7) 

      GPIO.cleanup(buttons[key])
    sleep(0.1)
```

I also connected the button on top of the phone where the receiver is and only run this `listen_to_keypad()` function when the phone is picked up.

## Final result

I won't go into the details of how I integrated this with Argo CD, because it's pretty much the same as the old release button. For those interested in the technical details, the entire project is available on my [Github](https://github.com/nerijusdu/release-button). 

And everyone lived happily ever after in the startup land. For now...

<div class="img-lg">
![Phone](/data/images/phone/result.jpg)
</div><!---->
