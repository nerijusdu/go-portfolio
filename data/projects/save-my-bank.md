# Save My Bank

This is my most finished project, that I'm really proud of and learned a lot while building it.
This is a budgeting platform where you can easily track your spending, create a budget and learn to save money.

## Tech stack

Backend is built using .NET 5. It uses Entity Framework, PostgreSQL database, IdentityServer and all the other good stuff that I wanted to learn.


Frontend is build using React, TypeScript, ChakraUI, ReactQuery and React Hook Form.

## The journey

### The beginning

This started as a small project when my brother suggest an idea to creat a budgeting platform that would be suited for Lithuanian market, including syncing with your bank to make tracking your expenses effortless. However, I soon found out that connecting to bank APIs in Lithuania requires having a bussines, some certifications and other legal trouble that I was not prepared to go through. So for now I postponed this feature (to never) and settled with a CSV import of transactions that you can export from your bank. This can be configured for different type of delimiters which is probably confusing for the user but two biggest banks in Lithuania uses different default delimiters for CSV exports 🙃.

![Import modal screenshot](/data/images/save-my-bank-import.png)

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

TODO: additional features, debt plan, investment, tutorial, support requests, landing page, blog.

### Mobile

### The endgame