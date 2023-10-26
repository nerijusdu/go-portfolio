# How to over-engineer a button

Do you think that releases to production are too mainstream? Are you looking for a more exciting way to release your crappy code to the users? Boy do I have a solution for you! A big red button! With LEDs! 

<div class="img-md">
![Release button](/data/images/release-button.jpg)
</div><!---->

If you're still reading this, I will tell you of how I built a release button. It's a physical button connected to a Raspberry Pi that makes releases on Argo CD platform.


## Stack

For this project's hardware I used Raspberry Pi Zero W and a Red push button with an integrated LED. 

The software is written in Go and Python.

## Over-engineering

When I started this project I was learning Go, so my first thought was to write everything in Go. My first steps were to find a package that would handle the GPIO for me, so I could just `if button.pressed() { doThings() }`, unfortunately things are never that easy.

I tried a few packages but for some reason none of them handled my button presses consistently, for some packages I always had to press the button a few times for it to register the click, and for some packages, the button clicks were registered when I was not clicking it (and that's not an option when I'm trying to make a button do releases to production).

Then I remembered that Raspberry Pi and Python are good friends.

<div class="img-md">
![Raspberry and Python being good friends](/data/images/raspberry-python.jpg)
</div><!---->

I tried to make a test app using Python, just to check if I wired everything correctly. And surprise! Everything works perfectly, every button click counts and no accidental clicks are registered. But I don't really like Python, I didn't want to write this whole project in Python so I kept testing different Go packages. Then the most brilliant idea came to my head, what if I over-engineered this and make it a microservice architecture. 

<div class="img-md">
![Over-engineering meme](/data/images/over-engineering.jpg)
</div><!---->

And that's exactly what I did. I wrote a microservice in Python that would control the GPIO pins and all the hardware and I put all the business logic in a separate service written in Go. And I also wrote a web UI with React for configuring the button just to make it even more over-engineered.

## IO Service

Python makes it very easy to work with GPIO pins so this is a very simple service that consists of two parts.

The first part listens for button clicks and then sends an HTTP request to the main service. And the seconds part is an HTTP server that exposes controls of the LEDs.

## Releaser service

This service listens for HTTP requests from the IO service and if the button is pressed it makes a request to Argo API to sync the applications.

It also periodically checks the status of the applications and lights up the LED on a button if some of the applications are out of sync.

## Conclusion

I guess I could've written everything in Python but I don't really enjoy writing Python and I wanted to improve my Go skills, so it became a fun over-engineering example. I'm very happy with the result, it was a lot of fun to build and hella fun to use. 

All the code is available on [GitHub](https://github.com/nerijusdu/release-button).

And here's how it looks in the end:

<div class="img-md">
![Release button picture](/data/images/release-button-pic.jpg)
</div><!---->
