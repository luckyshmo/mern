const express = require('express')
const config = require("config")
const mongoose = require("mongoose")

const app = express()

//now endpoints can parse data properly
app.use(express.json({ extended:true }))

app.use('/api/auth', require('./routes/auth.routes'))

const PORT = config.get("port") || 5000 //from config or 5000

//npm run server - starts server
//npx create-react-app client - creates react app (client folder)
//npm run dev - starts our server+client config from package.json
async function start() {
    try {
        await mongoose.connect(config.get("mongoUri"), {
            useNewUrlParser: true,
            useCreateIndex: true,
            useUnifiedTopology: true //TODO what is this?
        })
        app.listen(PORT, 'localhost', () => console.log(`App started on port: ${PORT}`)) //todo address?
        console.log("started node")
    } catch (e) {
        console.log("Server Error", e.message)
        process.exit(1)
    }
}

start()