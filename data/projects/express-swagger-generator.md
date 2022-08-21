# Swagger generator

This is a CLI program to generate Swagger documentation files for an existing API made with Express and TypeScript.

## Challenges

It was quite challenging to implement because there are almost no documentation on how typescript syntax trees looks like.
However, there's nothing that a bunch `console.log`s can't solve, am I right?

## Limitations

For now this only generates request and response types along with all the endpoints, but you can't add any descriptions and etc.


I did this only as a proof of concept, the code is all messy and there's a bunch of `@ts-ignore`, so don't judge me please.