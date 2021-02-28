#install npm + react
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -

sudo apt install nodejs
# node --version
# npm --version

#Express & mongoose
npm install express mongoose

#TODO
npm install -D nodemon concurrently

#Config for constants
npm i config
mkdir config
vi config/default.json

#Mongo
docker pull mongo-express
docker pull mongo

#Encryption js lib
npm i bcryptjs

#Express validator
npm i express-validator

#generate tokens for auth
npm i jsonwebtoken