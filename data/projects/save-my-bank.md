# Save My Bank

This is my most finished project, that I'm really proud of and learned a lot while building it.
This is a budgeting platform where you can easily track your spending, create a budget and learn to save money.

<div class="img-lg">
![Home page](/data/images/save-my-bank-home.png)
</div><!---->

## Tech stack

Backend is built using .NET 5. It uses Entity Framework, PostgreSQL database, IdentityServer and all the other good stuff that I wanted to learn.


Frontend is build using React, TypeScript, ChakraUI, ReactQuery and React Hook Form.

## The journey

### The beginning

This started as a small project when my brother suggest an idea to creat a budgeting platform that would be suited for Lithuanian market, including syncing with your bank to make tracking your expenses effortless. However, I soon found out that connecting to bank APIs in Lithuania requires having a bussines, some certifications and other legal trouble that I was not prepared to go through. So for now I postponed this feature (to never) and settled with a CSV import of transactions that you can export from your bank. This can be configured for different type of delimiters which is probably confusing for the user but two biggest banks in Lithuania uses different default delimiters for CSV exports 🙃.

<div class="img-lg">
![Import modal screenshot](/data/images/save-my-bank-import.png)
</div><!---->

After that came all the other basic features of a budgeting platform:

- Setting a budget
- Creating recurring transactions
- Fancy dashboard
- Dark theme (that's a must)
- etc.

### Deployment

When Save My Bank became an actually usable platform the time came for deployments. I thought about cloud hosting it but database hosting is quite expensive (a month of DB hosting costs as much as a year of a cheap VPS hosting), so I got a cheap VPS instead. And then I had to learn how to configure servers 😱.


It was challenging at first but I learned quite a lot of new things:

- Setting up and configuring Nginx
- Configuring a firewall on linux
- Setting up SSL with LetsEncrypt
- Creating a private docker registry
- Github Actions

The end result was a docker-compose file that had all the configuration to run the docker image of Save My Bank.
Docker image was built by Github Actions and pushed to my private registry. After that I had a manually triggered Github Action that would SSH into my server, pull the latest image and rerun the docker-compose file.

### Polishing

This is one of my most polished projects. It's deployed to production environment, it has a landing page, step-by-step tutorials on how to use different features.

<div class="img-lg">
![Tutorial preview](/data/images/save-my-bank-tutorial.png)
</div><!---->

Save My Bank also has quite a lot of self hosted supporting services:

- Analytics - [Plausible](https://plausible.io/)
- Monitoring - [Graphana](https://grafana.com/)
- Blog - [Ghost](https://ghost.org) (didn't write anything in it  though)

### Mobile

I tried making a mobile app for this project 3 times.


First time was with React Native, because I thought that having a React frontend will make everything really easy and I could reuse most of the logic. Boy was I wrong! Started by creating a shared logic package that could be reused between web and mobile, it took a week or two until I got authentication and one component working, I was using a lot of packages that don't work well with React Native. And then React Native turned out to be a really bad dev experience with random crashes, hard to install packages, hot reloading not working and hard to write styles. So I dropped it after implementing a few pages.


Second time was with .NET MAUI and Blazor web view. It was a little easier because it's just a website, but I tried to make it feel like a mobile app so it was quit hard to implement smooth transitions and etc. Also .NET MAUI was still in preview when I started this, so dev experience was not that great. So I decided to wait for .NET MAUI to get out of preview and never came back to this app.


And the third time was with Flutter. This was by far the best experience while making an app. Out of the box widgets cover most of what I needed and everything just worked. It took me about 6 hours to make the same progress as I did with React Native or Blazor in a week. However, when I started making this app with flutter I had no intention to finnish it because I had already lost any motivation to continue this project. 

### The endgame

I worked on this project for about a year until I lost all motivation to continue working on it. It was getting boring to implement new features when no one was actually using them, so I stopped. In summary, I learned a shit ton of new things while working on it and it was definately worth it, 10/10 would recommend.
